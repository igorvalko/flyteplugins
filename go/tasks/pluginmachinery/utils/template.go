package utils

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/lyft/flytestdlib/logger"
	"github.com/pkg/errors"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"
)

var inputFileRegex = regexp.MustCompile(`(?i){{\s*[\.$]Input\s*}}`)
var inputPrefixRegex = regexp.MustCompile(`(?i){{\s*[\.$]InputPrefix\s*}}`)
var outputRegex = regexp.MustCompile(`(?i){{\s*[\.$]OutputPrefix\s*}}`)
var inputVarRegex = regexp.MustCompile(`(?i){{\s*[\.$]Inputs\.(?P<input_name>[^}\s]+)\s*}}`)

// Evaluates templates in each command with the equivalent value from passed args. Templates are case-insensitive
// Supported templates are:
// - {{ .InputFile }} to receive the input file path. The protocol used will depend on the underlying system
// 		configuration. E.g. s3://bucket/key/to/file.pb or /var/run/local.pb are both valid.
// - {{ .OutputPrefix }} to receive the path prefix for where to store the outputs.
// - {{ .Inputs.myInput }} to receive the actual value of the input passed. See docs on LiteralMapToTemplateArgs for how
// 		what to expect each literal type to be serialized as.
// If a command isn't a valid template or failed to evaluate, it'll be returned as is.
// NOTE: I wanted to do in-place replacement, until I realized that in-place replacement will alter the definition of the
// graph. This is not desirable, as we may have to retry and in that case the replacement will not work and we want
// to create a new location for outputs
func ReplaceTemplateCommandArgs(ctx context.Context, command []string, in io.InputReader, out io.OutputFilePaths) ([]string, error) {
	if len(command) == 0 {
		return []string{}, nil
	}
	if in == nil || out == nil {
		return nil, fmt.Errorf("input reader and output path cannot be nil")
	}
	res := make([]string, 0, len(command))
	for _, commandTemplate := range command {
		updated, err := replaceTemplateCommandArgs(ctx, commandTemplate, in, out)
		if err != nil {
			return res, err
		}

		res = append(res, updated)
	}

	return res, nil
}

func replaceTemplateCommandArgs(ctx context.Context, commandTemplate string, in io.InputReader, out io.OutputFilePaths) (string, error) {
	val := inputFileRegex.ReplaceAllString(commandTemplate, in.GetInputPath().String())
	val = outputRegex.ReplaceAllString(val, out.GetOutputPrefixPath().String())
	val = inputPrefixRegex.ReplaceAllString(val, in.GetInputPrefixPath().String())
	groupMatches := inputVarRegex.FindAllStringSubmatchIndex(val, -1)
	if len(groupMatches) == 0 {
		return val, nil
	} else if len(groupMatches) > 1 {
		return val, fmt.Errorf("only one level of inputs nesting is supported. Syntax in [%v] is invalid", commandTemplate)
	} else if len(groupMatches[0]) > 4 {
		return val, fmt.Errorf("longer submatches not supported. Syntax in [%v] is invalid", commandTemplate)
	}
	startIdx := groupMatches[0][0]
	endIdx := groupMatches[0][1]
	inputStartIdx := groupMatches[0][2]
	inputEndIdx := groupMatches[0][3]
	inputName := val[inputStartIdx:inputEndIdx]

	inputs, err := in.Get(ctx)
	if err != nil {
		return val, errors.Wrapf(err, "unable to read inputs for [%s]", inputName)
	}
	if inputs == nil || inputs.Literals == nil {
		return val, fmt.Errorf("no inputs provided, cannot bind input name [%s]", inputName)
	}
	inputVal, exists := inputs.Literals[inputName]
	if !exists {
		return val, fmt.Errorf("requested input is not found [%v] while processing template [%v]",
			inputName, commandTemplate)
	}

	v, err := serializeLiteral(ctx, inputVal)
	if err != nil {
		return val, errors.Wrapf(err, "failed to bind a value to inputName [%s]", inputName)
	}
	if endIdx >= len(val) {
		return val[:startIdx] + v, nil
	}

	return val[:startIdx] + v + val[endIdx:], nil

}

func serializePrimitive(p *core.Primitive) (string, error) {
	switch o := p.Value.(type) {
	case *core.Primitive_Integer:
		return fmt.Sprintf("%v", o.Integer), nil
	case *core.Primitive_Boolean:
		return fmt.Sprintf("%v", o.Boolean), nil
	case *core.Primitive_Datetime:
		return ptypes.TimestampString(o.Datetime), nil
	case *core.Primitive_Duration:
		return o.Duration.String(), nil
	case *core.Primitive_FloatValue:
		return fmt.Sprintf("%v", o.FloatValue), nil
	case *core.Primitive_StringValue:
		return o.StringValue, nil
	default:
		return "", fmt.Errorf("received an unexpected primitive type [%v]", reflect.TypeOf(p.Value))
	}
}

func serializeLiteralScalar(l *core.Scalar) (string, error) {
	switch o := l.Value.(type) {
	case *core.Scalar_Primitive:
		return serializePrimitive(o.Primitive)
	case *core.Scalar_Blob:
		return o.Blob.Uri, nil
	default:
		return "", fmt.Errorf("received an unexpected scalar type [%v]", reflect.TypeOf(l.Value))
	}
}

func serializeLiteral(ctx context.Context, l *core.Literal) (string, error) {
	switch o := l.Value.(type) {
	case *core.Literal_Collection:
		res := make([]string, 0, len(o.Collection.Literals))
		for _, sub := range o.Collection.Literals {
			s, err := serializeLiteral(ctx, sub)
			if err != nil {
				return "", err
			}

			res = append(res, s)
		}

		return fmt.Sprintf("[%v]", strings.Join(res, ",")), nil
	case *core.Literal_Scalar:
		return serializeLiteralScalar(o.Scalar)
	default:
		logger.Debugf(ctx, "received unexpected primitive type")
		return "", fmt.Errorf("received an unexpected primitive type [%v]", reflect.TypeOf(l.Value))
	}
}

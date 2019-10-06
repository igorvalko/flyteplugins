// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PluginStateWriter is an autogenerated mock type for the PluginStateWriter type
type PluginStateWriter struct {
	mock.Mock
}

type PluginStateWriter_Put struct {
	*mock.Call
}

func (_m PluginStateWriter_Put) Return(_a0 error) *PluginStateWriter_Put {
	return &PluginStateWriter_Put{Call: _m.Call.Return(_a0)}
}

func (_m *PluginStateWriter) OnPut(stateVersion uint8, v interface{}) *PluginStateWriter_Put {
	c := _m.On("Put")
	return &PluginStateWriter_Put{Call: c}
}
func (_m *PluginStateWriter) OnPutMatch(matchers ...interface{}) *PluginStateWriter_Put {
	c := _m.On("Put", matchers...)
	return &PluginStateWriter_Put{Call: c}
}

// Put provides a mock function with given fields: stateVersion, v
func (_m *PluginStateWriter) Put(stateVersion uint8, v interface{}) error {
	ret := _m.Called(stateVersion, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, interface{}) error); ok {
		r0 = rf(stateVersion, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type PluginStateWriter_Reset struct {
	*mock.Call
}

func (_m PluginStateWriter_Reset) Return(_a0 error) *PluginStateWriter_Reset {
	return &PluginStateWriter_Reset{Call: _m.Call.Return(_a0)}
}

func (_m *PluginStateWriter) OnReset() *PluginStateWriter_Reset {
	c := _m.On("Reset")
	return &PluginStateWriter_Reset{Call: c}
}
func (_m *PluginStateWriter) OnResetMatch(matchers ...interface{}) *PluginStateWriter_Reset {
	c := _m.On("Reset", matchers...)
	return &PluginStateWriter_Reset{Call: c}
}

// Reset provides a mock function with given fields:
func (_m *PluginStateWriter) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
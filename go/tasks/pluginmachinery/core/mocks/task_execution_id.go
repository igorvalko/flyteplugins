// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	flyteidlcore "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"
)

// TaskExecutionID is an autogenerated mock type for the TaskExecutionID type
type TaskExecutionID struct {
	mock.Mock
}

type TaskExecutionID_GetGeneratedName struct {
	*mock.Call
}

func (_m TaskExecutionID_GetGeneratedName) Return(_a0 string) *TaskExecutionID_GetGeneratedName {
	return &TaskExecutionID_GetGeneratedName{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionID) OnGetGeneratedName() *TaskExecutionID_GetGeneratedName {
	c := _m.On("GetGeneratedName")
	return &TaskExecutionID_GetGeneratedName{Call: c}
}

func (_m *TaskExecutionID) OnGetGeneratedNameMatch(matchers ...interface{}) *TaskExecutionID_GetGeneratedName {
	c := _m.On("GetGeneratedName", matchers...)
	return &TaskExecutionID_GetGeneratedName{Call: c}
}

// GetGeneratedName provides a mock function with given fields:
func (_m *TaskExecutionID) GetGeneratedName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type TaskExecutionID_GetID struct {
	*mock.Call
}

func (_m TaskExecutionID_GetID) Return(_a0 flyteidlcore.TaskExecutionIdentifier) *TaskExecutionID_GetID {
	return &TaskExecutionID_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionID) OnGetID() *TaskExecutionID_GetID {
	c := _m.On("GetID")
	return &TaskExecutionID_GetID{Call: c}
}

func (_m *TaskExecutionID) OnGetIDMatch(matchers ...interface{}) *TaskExecutionID_GetID {
	c := _m.On("GetID", matchers...)
	return &TaskExecutionID_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *TaskExecutionID) GetID() flyteidlcore.TaskExecutionIdentifier {
	ret := _m.Called()

	var r0 flyteidlcore.TaskExecutionIdentifier
	if rf, ok := ret.Get(0).(func() flyteidlcore.TaskExecutionIdentifier); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flyteidlcore.TaskExecutionIdentifier)
	}

	return r0
}

// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	calculator "github.com/halfnibble/learn-tdd-go/calculator"

	mock "github.com/stretchr/testify/mock"
)

// OperationProcessor is an autogenerated mock type for the OperationProcessor type
type OperationProcessor struct {
	mock.Mock
}

// ProcessOperation provides a mock function with given fields: operation
func (_m *OperationProcessor) ProcessOperation(operation *calculator.Operation) (*string, error) {
	ret := _m.Called(operation)

	if len(ret) == 0 {
		panic("no return value specified for ProcessOperation")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(*calculator.Operation) (*string, error)); ok {
		return rf(operation)
	}
	if rf, ok := ret.Get(0).(func(*calculator.Operation) *string); ok {
		r0 = rf(operation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(*calculator.Operation) error); ok {
		r1 = rf(operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOperationProcessor creates a new instance of OperationProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOperationProcessor(t interface {
	mock.TestingT
	Cleanup(func())
}) *OperationProcessor {
	mock := &OperationProcessor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

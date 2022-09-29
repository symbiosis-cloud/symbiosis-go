// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	resty "github.com/go-resty/resty/v2"
	mock "github.com/stretchr/testify/mock"
)

// CallOption is an autogenerated mock type for the CallOption type
type CallOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: req
func (_m *CallOption) Execute(req *resty.Request) {
	_m.Called(req)
}

type mockConstructorTestingTNewCallOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewCallOption creates a new instance of CallOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCallOption(t mockConstructorTestingTNewCallOption) *CallOption {
	mock := &CallOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	resty "github.com/go-resty/resty/v2"
	mock "github.com/stretchr/testify/mock"
)

// ClientOption is an autogenerated mock type for the ClientOption type
type ClientOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: c
func (_m *ClientOption) Execute(c *resty.Client) {
	_m.Called(c)
}

type mockConstructorTestingTNewClientOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientOption creates a new instance of ClientOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientOption(t mockConstructorTestingTNewClientOption) *ClientOption {
	mock := &ClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

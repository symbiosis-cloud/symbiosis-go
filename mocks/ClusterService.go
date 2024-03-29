// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	symbiosis "github.com/symbiosis-cloud/symbiosis-go"
)

// ClusterService is an autogenerated mock type for the ClusterService type
type ClusterService struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *ClusterService) Create(input *symbiosis.ClusterInput) (*symbiosis.Cluster, error) {
	ret := _m.Called(input)

	var r0 *symbiosis.Cluster
	if rf, ok := ret.Get(0).(func(*symbiosis.ClusterInput) *symbiosis.Cluster); ok {
		r0 = rf(input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*symbiosis.ClusterInput) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServiceAccount provides a mock function with given fields: clusterName, subjectId
func (_m *ClusterService) CreateServiceAccount(clusterName string, subjectId string) (*symbiosis.ServiceAccount, error) {
	ret := _m.Called(clusterName, subjectId)

	var r0 *symbiosis.ServiceAccount
	if rf, ok := ret.Get(0).(func(string, string) *symbiosis.ServiceAccount); ok {
		r0 = rf(clusterName, subjectId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(clusterName, subjectId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServiceAccountForSelf provides a mock function with given fields: clusterName
func (_m *ClusterService) CreateServiceAccountForSelf(clusterName string) (*symbiosis.ServiceAccount, error) {
	ret := _m.Called(clusterName)

	var r0 *symbiosis.ServiceAccount
	if rf, ok := ret.Get(0).(func(string) *symbiosis.ServiceAccount); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: clusterName
func (_m *ClusterService) Delete(clusterName string) error {
	ret := _m.Called(clusterName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(clusterName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteServiceAccount provides a mock function with given fields: clusterName, serviceAccountId
func (_m *ClusterService) DeleteServiceAccount(clusterName string, serviceAccountId string) error {
	ret := _m.Called(clusterName, serviceAccountId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(clusterName, serviceAccountId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Describe provides a mock function with given fields: clusterName
func (_m *ClusterService) Describe(clusterName string) (*symbiosis.Cluster, error) {
	ret := _m.Called(clusterName)

	var r0 *symbiosis.Cluster
	if rf, ok := ret.Get(0).(func(string) *symbiosis.Cluster); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeById provides a mock function with given fields: id
func (_m *ClusterService) DescribeById(id string) (*symbiosis.Cluster, error) {
	ret := _m.Called(id)

	var r0 *symbiosis.Cluster
	if rf, ok := ret.Get(0).(func(string) *symbiosis.Cluster); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIdentity provides a mock function with given fields: clusterName
func (_m *ClusterService) GetIdentity(clusterName string) (*symbiosis.ClusterIdentity, error) {
	ret := _m.Called(clusterName)

	var r0 *symbiosis.ClusterIdentity
	if rf, ok := ret.Get(0).(func(string) *symbiosis.ClusterIdentity); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.ClusterIdentity)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceAccount provides a mock function with given fields: clusterName, serviceAccountId
func (_m *ClusterService) GetServiceAccount(clusterName string, serviceAccountId string) (*symbiosis.ServiceAccount, error) {
	ret := _m.Called(clusterName, serviceAccountId)

	var r0 *symbiosis.ServiceAccount
	if rf, ok := ret.Get(0).(func(string, string) *symbiosis.ServiceAccount); ok {
		r0 = rf(clusterName, serviceAccountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.ServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(clusterName, serviceAccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: maxSize, page
func (_m *ClusterService) List(maxSize int, page int) (*symbiosis.ClusterList, error) {
	ret := _m.Called(maxSize, page)

	var r0 *symbiosis.ClusterList
	if rf, ok := ret.Get(0).(func(int, int) *symbiosis.ClusterList); ok {
		r0 = rf(maxSize, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.ClusterList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(maxSize, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodes provides a mock function with given fields: clusterName
func (_m *ClusterService) ListNodes(clusterName string) (*symbiosis.NodeList, error) {
	ret := _m.Called(clusterName)

	var r0 *symbiosis.NodeList
	if rf, ok := ret.Get(0).(func(string) *symbiosis.NodeList); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*symbiosis.NodeList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserServiceAccounts provides a mock function with given fields: clusterName
func (_m *ClusterService) ListUserServiceAccounts(clusterName string) ([]*symbiosis.UserServiceAccount, error) {
	ret := _m.Called(clusterName)

	var r0 []*symbiosis.UserServiceAccount
	if rf, ok := ret.Get(0).(func(string) []*symbiosis.UserServiceAccount); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*symbiosis.UserServiceAccount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClusterService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterService creates a new instance of ClusterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterService(t mockConstructorTestingTNewClusterService) *ClusterService {
	mock := &ClusterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

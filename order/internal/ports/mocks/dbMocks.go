// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/az58740/gRPC-Go-Microservice-Hexagolanal/order/internal/application/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// DBPort is an autogenerated mock type for the DBPort type
type DBPort struct {
	mock.Mock
}

// Get provides a mock function with given fields: id
func (_m *DBPort) Get(id string) (domain.Order, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 domain.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Order); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Order)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *DBPort) Save(_a0 *domain.Order) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Order) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDBPort creates a new instance of DBPort. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDBPort(t interface {
	mock.TestingT
	Cleanup(func())
}) *DBPort {
	mock := &DBPort{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

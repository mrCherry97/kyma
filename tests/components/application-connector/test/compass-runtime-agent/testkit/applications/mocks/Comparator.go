// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// Comparator is an autogenerated mock type for the Comparator type
type Comparator struct {
	mock.Mock
}

// Compare provides a mock function with given fields: test, expected, actual
func (_m *Comparator) Compare(test *testing.T, expected string, actual string) error {
	ret := _m.Called(test, expected, actual)

	var r0 error
	if rf, ok := ret.Get(0).(func(*testing.T, string, string) error); ok {
		r0 = rf(test, expected, actual)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewComparator interface {
	mock.TestingT
	Cleanup(func())
}

// NewComparator creates a new instance of Comparator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComparator(t mockConstructorTestingTNewComparator) *Comparator {
	mock := &Comparator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
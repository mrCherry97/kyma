// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	handlers "github.com/kyma-project/kyma/components/eventing-controller/pkg/handlers"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/kyma-project/kyma/components/eventing-controller/api/v1alpha1"
)

// NatsBackend is an autogenerated mock type for the NatsBackend type
type NatsBackend struct {
	mock.Mock
}

// DeleteSubscription provides a mock function with given fields: subscription
func (_m *NatsBackend) DeleteSubscription(subscription *v1alpha1.Subscription) error {
	ret := _m.Called(subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Subscription) error); ok {
		r0 = rf(subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initialize provides a mock function with given fields: connCloseHandler
func (_m *NatsBackend) Initialize(connCloseHandler handlers.ConnClosedHandler) error {
	ret := _m.Called(connCloseHandler)

	var r0 error
	if rf, ok := ret.Get(0).(func(handlers.ConnClosedHandler) error); ok {
		r0 = rf(connCloseHandler)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SyncSubscription provides a mock function with given fields: subscription
func (_m *NatsBackend) SyncSubscription(subscription *v1alpha1.Subscription) error {
	ret := _m.Called(subscription)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.Subscription) error); ok {
		r0 = rf(subscription)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
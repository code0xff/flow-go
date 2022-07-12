// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// EpochLookup is an autogenerated mock type for the EpochLookup type
type EpochLookup struct {
	mock.Mock
}

// EpochForView provides a mock function with given fields: view
func (_m *EpochLookup) EpochForView(view uint64) (uint64, error) {
	ret := _m.Called(view)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(view)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(view)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpochForViewWithFallback provides a mock function with given fields: view
func (_m *EpochLookup) EpochForViewWithFallback(view uint64) (uint64, error) {
	ret := _m.Called(view)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint64) uint64); ok {
		r0 = rf(view)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(view)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewEpochLookupT interface {
	mock.TestingT
	Cleanup(func())
}

// NewEpochLookup creates a new instance of EpochLookup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEpochLookup(t NewEpochLookupT) *EpochLookup {
	mock := &EpochLookup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

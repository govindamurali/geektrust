// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	enum "geektrust/enum"
	portfolio "geektrust/portfolio"

	mock "github.com/stretchr/testify/mock"
)

// Portfolio is an autogenerated mock type for the Portfolio type
type Portfolio struct {
	mock.Mock
}

// Allocate provides a mock function with given fields: allocation
func (_m *Portfolio) Allocate(allocation portfolio.ClasswiseAllocationMap) error {
	ret := _m.Called(allocation)

	var r0 error
	if rf, ok := ret.Get(0).(func(portfolio.ClasswiseAllocationMap) error); ok {
		r0 = rf(allocation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Change provides a mock function with given fields: month, change
func (_m *Portfolio) Change(month enum.Month, change portfolio.Change) error {
	ret := _m.Called(month, change)

	var r0 error
	if rf, ok := ret.Get(0).(func(enum.Month, portfolio.Change) error); ok {
		r0 = rf(month, change)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalance provides a mock function with given fields: month
func (_m *Portfolio) GetBalance(month enum.Month) (portfolio.ClasswiseAllocationMap, error) {
	ret := _m.Called(month)

	var r0 portfolio.ClasswiseAllocationMap
	if rf, ok := ret.Get(0).(func(enum.Month) portfolio.ClasswiseAllocationMap); ok {
		r0 = rf(month)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(portfolio.ClasswiseAllocationMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(enum.Month) error); ok {
		r1 = rf(month)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastRebalance provides a mock function with given fields:
func (_m *Portfolio) GetLastRebalance() (portfolio.ClasswiseAllocationMap, error) {
	ret := _m.Called()

	var r0 portfolio.ClasswiseAllocationMap
	if rf, ok := ret.Get(0).(func() portfolio.ClasswiseAllocationMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(portfolio.ClasswiseAllocationMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsRebalanced provides a mock function with given fields:
func (_m *Portfolio) IsRebalanced() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// StartSip provides a mock function with given fields: sip
func (_m *Portfolio) StartSip(sip portfolio.ClasswiseAllocationMap) {
	_m.Called(sip)
}

type NewPortfolioT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPortfolio creates a new instance of Portfolio. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPortfolio(t NewPortfolioT) *Portfolio {
	mock := &Portfolio{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

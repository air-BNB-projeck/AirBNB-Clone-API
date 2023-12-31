// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	users "alta/air-bnb/features/users"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// DeleteUserById provides a mock function with given fields: userId
func (_m *UserServiceInterface) DeleteUserById(userId uint) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditUserById provides a mock function with given fields: userId, userData
func (_m *UserServiceInterface) EditUserById(userId uint, userData users.CoreUserRequest) error {
	ret := _m.Called(userId, userData)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, users.CoreUserRequest) error); ok {
		r0 = rf(userId, userData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserById provides a mock function with given fields: userId
func (_m *UserServiceInterface) GetUserById(userId uint) (users.Core, error) {
	ret := _m.Called(userId)

	var r0 users.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (users.Core, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) users.Core); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: loginPayload
func (_m *UserServiceInterface) LoginUser(loginPayload users.CoreLoginUserRequest) (uint, error) {
	ret := _m.Called(loginPayload)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(users.CoreLoginUserRequest) (uint, error)); ok {
		return rf(loginPayload)
	}
	if rf, ok := ret.Get(0).(func(users.CoreLoginUserRequest) uint); ok {
		r0 = rf(loginPayload)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(users.CoreLoginUserRequest) error); ok {
		r1 = rf(loginPayload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: userData
func (_m *UserServiceInterface) RegisterUser(userData users.CoreUserRequest) (uint, error) {
	ret := _m.Called(userData)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(users.CoreUserRequest) (uint, error)); ok {
		return rf(userData)
	}
	if rf, ok := ret.Get(0).(func(users.CoreUserRequest) uint); ok {
		r0 = rf(userData)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(users.CoreUserRequest) error); ok {
		r1 = rf(userData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserServiceInterface(t mockConstructorTestingTNewUserServiceInterface) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

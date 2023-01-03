// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt/v4"
	mock "github.com/stretchr/testify/mock"
)

// AuthUtil is an autogenerated mock type for the AuthUtil type
type AuthUtil struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: hashedPwd, inputPwd
func (_m *AuthUtil) ComparePassword(hashedPwd string, inputPwd string) bool {
	ret := _m.Called(hashedPwd, inputPwd)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(hashedPwd, inputPwd)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GenerateAccessToken provides a mock function with given fields: username, role
func (_m *AuthUtil) GenerateAccessToken(username string, role string) (string, error) {
	ret := _m.Called(username, role)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateRefreshToken provides a mock function with given fields:
func (_m *AuthUtil) GenerateRefreshToken() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: encodedToken, signSecret
func (_m *AuthUtil) ValidateToken(encodedToken string, signSecret string) (*jwt.Token, error) {
	ret := _m.Called(encodedToken, signSecret)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(string, string) *jwt.Token); ok {
		r0 = rf(encodedToken, signSecret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(encodedToken, signSecret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthUtil interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUtil creates a new instance of AuthUtil. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUtil(t mockConstructorTestingTNewAuthUtil) *AuthUtil {
	mock := &AuthUtil{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

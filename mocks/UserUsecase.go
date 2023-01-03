// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetDetailRole provides a mock function with given fields: roleId
func (_m *UserUsecase) GetDetailRole(roleId int) (*entity.Role, error) {
	ret := _m.Called(roleId)

	var r0 *entity.Role
	if rf, ok := ret.Get(0).(func(int) *entity.Role); ok {
		r0 = rf(roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Role)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDetailUserByUsername provides a mock function with given fields: accessToken
func (_m *UserUsecase) GetDetailUserByUsername(accessToken string) (*entity.UserContext, error) {
	ret := _m.Called(accessToken)

	var r0 *entity.UserContext
	if rf, ok := ret.Get(0).(func(string) *entity.UserContext); ok {
		r0 = rf(accessToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accessToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *UserUsecase) Login(_a0 string, _a1 string) (*entity.UserLoginResBody, string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.UserLoginResBody
	if rf, ok := ret.Get(0).(func(string, string) *entity.UserLoginResBody); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserLoginResBody)
		}
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(_a0, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Refresh provides a mock function with given fields: _a0
func (_m *UserUsecase) Refresh(_a0 string) (*entity.UserLoginResBody, error) {
	ret := _m.Called(_a0)

	var r0 *entity.UserLoginResBody
	if rf, ok := ret.Get(0).(func(string) *entity.UserLoginResBody); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserLoginResBody)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: _a0
func (_m *UserUsecase) Register(_a0 *entity.UserRegisterReqBody) (*entity.UserRegisterResBody, error) {
	ret := _m.Called(_a0)

	var r0 *entity.UserRegisterResBody
	if rf, ok := ret.Get(0).(func(*entity.UserRegisterReqBody) *entity.UserRegisterResBody); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserRegisterResBody)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.UserRegisterReqBody) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserDetailsByUsername provides a mock function with given fields: username, updatePremises
func (_m *UserUsecase) UpdateUserDetailsByUsername(username string, updatePremises entity.UserEditDetailsReqBody) (*entity.UserContext, error) {
	ret := _m.Called(username, updatePremises)

	var r0 *entity.UserContext
	if rf, ok := ret.Get(0).(func(string, entity.UserEditDetailsReqBody) *entity.UserContext); ok {
		r0 = rf(username, updatePremises)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, entity.UserEditDetailsReqBody) error); ok {
		r1 = rf(username, updatePremises)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

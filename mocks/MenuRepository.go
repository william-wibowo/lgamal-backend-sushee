// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// MenuRepository is an autogenerated mock type for the MenuRepository type
type MenuRepository struct {
	mock.Mock
}

// AddMenu provides a mock function with given fields: newMenu
func (_m *MenuRepository) AddMenu(newMenu *entity.Menu) (*entity.Menu, error) {
	ret := _m.Called(newMenu)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(*entity.Menu) *entity.Menu); ok {
		r0 = rf(newMenu)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Menu) error); ok {
		r1 = rf(newMenu)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMenuByMenuId provides a mock function with given fields: menuId
func (_m *MenuRepository) DeleteMenuByMenuId(menuId int) error {
	ret := _m.Called(menuId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(menuId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAndValidatePromoMenu provides a mock function with given fields: menuId, promoId
func (_m *MenuRepository) GetAndValidatePromoMenu(menuId int, promoId int) (*entity.PromoMenu, error) {
	ret := _m.Called(menuId, promoId)

	var r0 *entity.PromoMenu
	if rf, ok := ret.Get(0).(func(int, int) *entity.PromoMenu); ok {
		r0 = rf(menuId, promoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.PromoMenu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(menuId, promoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenu provides a mock function with given fields: _a0
func (_m *MenuRepository) GetMenu(_a0 entity.MenuQuery) (*[]entity.Menu, error) {
	ret := _m.Called(_a0)

	var r0 *[]entity.Menu
	if rf, ok := ret.Get(0).(func(entity.MenuQuery) *[]entity.Menu); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.MenuQuery) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenuByMenuId provides a mock function with given fields: menuId
func (_m *MenuRepository) GetMenuByMenuId(menuId int) (*entity.Menu, error) {
	ret := _m.Called(menuId)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(int) *entity.Menu); ok {
		r0 = rf(menuId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(menuId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenuCount provides a mock function with given fields: q
func (_m *MenuRepository) GetMenuCount(q entity.MenuQuery) (int, error) {
	ret := _m.Called(q)

	var r0 int
	if rf, ok := ret.Get(0).(func(entity.MenuQuery) int); ok {
		r0 = rf(q)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.MenuQuery) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMenuDetailByMenuId provides a mock function with given fields: menuId
func (_m *MenuRepository) GetMenuDetailByMenuId(menuId int) (*entity.Menu, error) {
	ret := _m.Called(menuId)

	var r0 *entity.Menu
	if rf, ok := ret.Get(0).(func(int) *entity.Menu); ok {
		r0 = rf(menuId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Menu)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(menuId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPromotionMenu provides a mock function with given fields:
func (_m *MenuRepository) GetPromotionMenu() (*[]entity.Promotion, error) {
	ret := _m.Called()

	var r0 *[]entity.Promotion
	if rf, ok := ret.Get(0).(func() *[]entity.Promotion); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Promotion)
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

// UpdateMenuByMenuId provides a mock function with given fields: menuId, newMenu
func (_m *MenuRepository) UpdateMenuByMenuId(menuId int, newMenu *entity.Menu) error {
	ret := _m.Called(menuId, newMenu)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *entity.Menu) error); ok {
		r0 = rf(menuId, newMenu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMenuRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenuRepository creates a new instance of MenuRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenuRepository(t mockConstructorTestingTNewMenuRepository) *MenuRepository {
	mock := &MenuRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
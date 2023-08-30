// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	domain "funding-rate/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: chatID
func (_m *UserRepository) CreateUser(chatID int64) error {
	ret := _m.Called(chatID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(chatID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveUsers provides a mock function with given fields:
func (_m *UserRepository) RetrieveUsers() ([]domain.User, error) {
	ret := _m.Called()

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

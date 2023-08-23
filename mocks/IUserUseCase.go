// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	domain "funding-rate/domain"

	mock "github.com/stretchr/testify/mock"
)

// IUserUseCase is an autogenerated mock type for the IUserUseCase type
type IUserUseCase struct {
	mock.Mock
}

// GetUsersNotification provides a mock function with given fields:
func (_m *IUserUseCase) GetUsersNotification() []domain.Notification {
	ret := _m.Called()

	var r0 []domain.Notification
	if rf, ok := ret.Get(0).(func() []domain.Notification); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Notification)
		}
	}

	return r0
}

// NewUser provides a mock function with given fields: chatID
func (_m *IUserUseCase) NewUser(chatID int64) string {
	ret := _m.Called(chatID)

	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(chatID)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewIUserUseCase creates a new instance of IUserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserUseCase {
	mock := &IUserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
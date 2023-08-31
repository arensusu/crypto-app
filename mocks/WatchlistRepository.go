// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	domain "funding-rate/domain"

	mock "github.com/stretchr/testify/mock"
)

// WatchlistRepository is an autogenerated mock type for the WatchlistRepository type
type WatchlistRepository struct {
	mock.Mock
}

// CreateWatchlist provides a mock function with given fields: chatID, pair
func (_m *WatchlistRepository) CreateWatchlist(chatID int64, pair domain.Pair) error {
	ret := _m.Called(chatID, pair)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, domain.Pair) error); ok {
		r0 = rf(chatID, pair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWatchlist provides a mock function with given fields: chatID, pair
func (_m *WatchlistRepository) DeleteWatchlist(chatID int64, pair domain.Pair) error {
	ret := _m.Called(chatID, pair)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, domain.Pair) error); ok {
		r0 = rf(chatID, pair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveWatchlists provides a mock function with given fields: chatID
func (_m *WatchlistRepository) RetrieveWatchlists(chatID int64) ([]domain.Pair, error) {
	ret := _m.Called(chatID)

	var r0 []domain.Pair
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) ([]domain.Pair, error)); ok {
		return rf(chatID)
	}
	if rf, ok := ret.Get(0).(func(int64) []domain.Pair); ok {
		r0 = rf(chatID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Pair)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(chatID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWatchlistRepository creates a new instance of WatchlistRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWatchlistRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *WatchlistRepository {
	mock := &WatchlistRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	coinglass "funding-rate/coinglass"
	domain "funding-rate/domain"

	mock "github.com/stretchr/testify/mock"
)

// PairRepository is an autogenerated mock type for the PairRepository type
type PairRepository struct {
	mock.Mock
}

// CreateFundingSearched provides a mock function with given fields: chatID, pair
func (_m *PairRepository) CreateFundingSearched(chatID int64, pair domain.Pair) error {
	ret := _m.Called(chatID, pair)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, domain.Pair) error); ok {
		r0 = rf(chatID, pair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFundingHistory provides a mock function with given fields: exchange, symbol
func (_m *PairRepository) GetFundingHistory(exchange string, symbol string) ([]float64, error) {
	ret := _m.Called(exchange, symbol)

	var r0 []float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]float64, error)); ok {
		return rf(exchange, symbol)
	}
	if rf, ok := ret.Get(0).(func(string, string) []float64); ok {
		r0 = rf(exchange, symbol)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(exchange, symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPerpetualMarket provides a mock function with given fields: exchange, symbol
func (_m *PairRepository) GetPerpetualMarket(exchange string, symbol string) (coinglass.PerpetualMarket, error) {
	ret := _m.Called(exchange, symbol)

	var r0 coinglass.PerpetualMarket
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (coinglass.PerpetualMarket, error)); ok {
		return rf(exchange, symbol)
	}
	if rf, ok := ret.Get(0).(func(string, string) coinglass.PerpetualMarket); ok {
		r0 = rf(exchange, symbol)
	} else {
		r0 = ret.Get(0).(coinglass.PerpetualMarket)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(exchange, symbol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveFundingSearched provides a mock function with given fields: chatID
func (_m *PairRepository) RetrieveFundingSearched(chatID int64) ([]domain.Pair, error) {
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

// NewPairRepository creates a new instance of PairRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPairRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *PairRepository {
	mock := &PairRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
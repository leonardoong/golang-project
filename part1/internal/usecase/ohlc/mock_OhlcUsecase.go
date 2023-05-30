// Code generated by mockery v2.20.0. DO NOT EDIT.

package ohlc

import (
	context "context"
	model "part1/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// MockOhlcUsecase is an autogenerated mock type for the OhlcUsecase type
type MockOhlcUsecase struct {
	mock.Mock
}

// GetSummary provides a mock function with given fields: ctx, req
func (_m *MockOhlcUsecase) GetSummary(ctx context.Context, req model.GetSummaryRequest) (model.GetSummaryResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 model.GetSummaryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetSummaryRequest) (model.GetSummaryResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetSummaryRequest) model.GetSummaryResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.GetSummaryResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetSummaryRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitData provides a mock function with given fields: ctx
func (_m *MockOhlcUsecase) InitData(ctx context.Context) (model.InitDataResponse, error) {
	ret := _m.Called(ctx)

	var r0 model.InitDataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (model.InitDataResponse, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) model.InitDataResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(model.InitDataResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessOhlc provides a mock function with given fields: ctx, req
func (_m *MockOhlcUsecase) ProcessOhlc(ctx context.Context, req model.Transaction) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Transaction) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockOhlcUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOhlcUsecase creates a new instance of MockOhlcUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOhlcUsecase(t mockConstructorTestingTNewMockOhlcUsecase) *MockOhlcUsecase {
	mock := &MockOhlcUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

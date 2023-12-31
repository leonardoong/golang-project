// Code generated by mockery v2.20.0. DO NOT EDIT.

package ohlc

import (
	context "context"
	__ "part1/internal/model/proto"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	mock "github.com/stretchr/testify/mock"
)

// MockOhlcHandler is an autogenerated mock type for the OhlcHandler type
type MockOhlcHandler struct {
	mock.Mock
}

// GetSummary provides a mock function with given fields: ctx, req
func (_m *MockOhlcHandler) GetSummary(ctx context.Context, req *__.GetSummaryRequest) (*__.GetSummaryResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 *__.GetSummaryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *__.GetSummaryRequest) (*__.GetSummaryResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *__.GetSummaryRequest) *__.GetSummaryResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.GetSummaryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *__.GetSummaryRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitData provides a mock function with given fields: ctx, void
func (_m *MockOhlcHandler) InitData(ctx context.Context, void *emptypb.Empty) (*__.InitDataResponse, error) {
	ret := _m.Called(ctx, void)

	var r0 *__.InitDataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) (*__.InitDataResponse, error)); ok {
		return rf(ctx, void)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *emptypb.Empty) *__.InitDataResponse); ok {
		r0 = rf(ctx, void)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.InitDataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *emptypb.Empty) error); ok {
		r1 = rf(ctx, void)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockOhlcHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockOhlcHandler creates a new instance of MockOhlcHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockOhlcHandler(t mockConstructorTestingTNewMockOhlcHandler) *MockOhlcHandler {
	mock := &MockOhlcHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

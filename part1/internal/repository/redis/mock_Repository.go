// Code generated by mockery v2.20.0. DO NOT EDIT.

package redis

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// HGetAll provides a mock function with given fields: ctx, key
func (_m *MockRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ret := _m.Called(ctx, key)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (map[string]string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HSet provides a mock function with given fields: ctx, key, values
func (_m *MockRepository) HSet(ctx context.Context, key string, values ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (int64, error)); ok {
		return rf(ctx, key, values...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) int64); ok {
		r0 = rf(ctx, key, values...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, key, values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
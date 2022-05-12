// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/tendermint/spn/x/profile/types"
	"google.golang.org/grpc"
)

// ProfileClient is an autogenerated mock type for the ProfileClient type
type ProfileClient struct {
	mock.Mock
}

// Coordinator provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) Coordinator(ctx context.Context, in *types.QueryGetCoordinatorRequest, opts ...grpc.CallOption) (*types.QueryGetCoordinatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCoordinatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCoordinatorRequest, ...grpc.CallOption) *types.QueryGetCoordinatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCoordinatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCoordinatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CoordinatorAll provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) CoordinatorAll(ctx context.Context, in *types.QueryAllCoordinatorRequest, opts ...grpc.CallOption) (*types.QueryAllCoordinatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllCoordinatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllCoordinatorRequest, ...grpc.CallOption) *types.QueryAllCoordinatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllCoordinatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllCoordinatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CoordinatorByAddress provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) CoordinatorByAddress(ctx context.Context, in *types.QueryGetCoordinatorByAddressRequest, opts ...grpc.CallOption) (*types.QueryGetCoordinatorByAddressResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetCoordinatorByAddressResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetCoordinatorByAddressRequest, ...grpc.CallOption) *types.QueryGetCoordinatorByAddressResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetCoordinatorByAddressResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetCoordinatorByAddressRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validator provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) Validator(ctx context.Context, in *types.QueryGetValidatorRequest, opts ...grpc.CallOption) (*types.QueryGetValidatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetValidatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetValidatorRequest, ...grpc.CallOption) *types.QueryGetValidatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetValidatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetValidatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorAll provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) ValidatorAll(ctx context.Context, in *types.QueryAllValidatorRequest, opts ...grpc.CallOption) (*types.QueryAllValidatorResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryAllValidatorResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryAllValidatorRequest, ...grpc.CallOption) *types.QueryAllValidatorResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryAllValidatorResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryAllValidatorRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidatorByOperatorAddress provides a mock function with given fields: ctx, in, opts
func (_m *ProfileClient) ValidatorByOperatorAddress(ctx context.Context, in *types.QueryGetValidatorByOperatorAddressRequest, opts ...grpc.CallOption) (*types.QueryGetValidatorByOperatorAddressResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *types.QueryGetValidatorByOperatorAddressResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryGetValidatorByOperatorAddressRequest, ...grpc.CallOption) *types.QueryGetValidatorByOperatorAddressResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryGetValidatorByOperatorAddressResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryGetValidatorByOperatorAddressRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProfileClient creates a new instance of ProfileClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewProfileClient(t testing.TB) *ProfileClient {
	mock := &ProfileClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
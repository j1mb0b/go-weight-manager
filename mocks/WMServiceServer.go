// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	proto "github.com/j1mb0b/go-weight-manager/proto"
	mock "github.com/stretchr/testify/mock"
)

// WMServiceServer is an autogenerated mock type for the WMServiceServer type
type WMServiceServer struct {
	mock.Mock
}

// AddEntry provides a mock function with given fields: _a0, _a1
func (_m *WMServiceServer) AddEntry(_a0 context.Context, _a1 *proto.WeightEntry) (*proto.EntryResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddEntry")
	}

	var r0 *proto.EntryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WeightEntry) (*proto.EntryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WeightEntry) *proto.EntryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.EntryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.WeightEntry) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AnalyzeWeight provides a mock function with given fields: _a0, _a1
func (_m *WMServiceServer) AnalyzeWeight(_a0 context.Context, _a1 *proto.Empty) (*proto.EntryResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AnalyzeWeight")
	}

	var r0 *proto.EntryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Empty) (*proto.EntryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.Empty) *proto.EntryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.EntryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteEntry provides a mock function with given fields: _a0, _a1
func (_m *WMServiceServer) DeleteEntry(_a0 context.Context, _a1 *proto.EntryID) (*proto.EntryResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEntry")
	}

	var r0 *proto.EntryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.EntryID) (*proto.EntryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.EntryID) *proto.EntryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.EntryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.EntryID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEntries provides a mock function with given fields: _a0, _a1
func (_m *WMServiceServer) GetEntries(_a0 context.Context, _a1 *proto.UserID) (*proto.WeightEntries, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetEntries")
	}

	var r0 *proto.WeightEntries
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UserID) (*proto.WeightEntries, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UserID) *proto.WeightEntries); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.WeightEntries)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.UserID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEntry provides a mock function with given fields: _a0, _a1
func (_m *WMServiceServer) UpdateEntry(_a0 context.Context, _a1 *proto.WeightEntry) (*proto.EntryResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEntry")
	}

	var r0 *proto.EntryResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WeightEntry) (*proto.EntryResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *proto.WeightEntry) *proto.EntryResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.EntryResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *proto.WeightEntry) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewWMServiceServer creates a new instance of WMServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWMServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *WMServiceServer {
	mock := &WMServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
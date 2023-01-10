// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cloudtrace "github.com/observiq/cloud-trace-grafana-ds/pkg/plugin/cloudtrace"

	cloudtracepb "google.golang.org/genproto/googleapis/devtools/cloudtrace/v1"

	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *API) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListTraces provides a mock function with given fields: _a0, _a1
func (_m *API) ListTraces(_a0 context.Context, _a1 *cloudtrace.Query) ([]*cloudtracepb.Trace, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*cloudtracepb.Trace
	if rf, ok := ret.Get(0).(func(context.Context, *cloudtrace.Query) []*cloudtracepb.Trace); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cloudtracepb.Trace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cloudtrace.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjects provides a mock function with given fields: _a0
func (_m *API) ListProjects(_a0 context.Context) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestConnection provides a mock function with given fields: ctx, projectID
func (_m *API) TestConnection(ctx context.Context, projectID string) error {
	ret := _m.Called(ctx, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewAPI creates a new instance of API. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAPI(t mockConstructorTestingTNewAPI) *API {
	mock := &API{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

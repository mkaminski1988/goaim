// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	state "github.com/mk6i/retro-aim-server/state"
)

// mockBOSRouter is an autogenerated mock type for the BOSRouter type
type mockBOSRouter struct {
	mock.Mock
}

type mockBOSRouter_Expecter struct {
	mock *mock.Mock
}

func (_m *mockBOSRouter) EXPECT() *mockBOSRouter_Expecter {
	return &mockBOSRouter_Expecter{mock: &_m.Mock}
}

// Route provides a mock function with given fields: ctx, sess, r, w, sequence
func (_m *mockBOSRouter) Route(ctx context.Context, sess *state.Session, r io.Reader, w io.Writer, sequence *uint32) error {
	ret := _m.Called(ctx, sess, r, w, sequence)

	if len(ret) == 0 {
		panic("no return value specified for Route")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, io.Reader, io.Writer, *uint32) error); ok {
		r0 = rf(ctx, sess, r, w, sequence)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockBOSRouter_Route_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Route'
type mockBOSRouter_Route_Call struct {
	*mock.Call
}

// Route is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - r io.Reader
//   - w io.Writer
//   - sequence *uint32
func (_e *mockBOSRouter_Expecter) Route(ctx interface{}, sess interface{}, r interface{}, w interface{}, sequence interface{}) *mockBOSRouter_Route_Call {
	return &mockBOSRouter_Route_Call{Call: _e.mock.On("Route", ctx, sess, r, w, sequence)}
}

func (_c *mockBOSRouter_Route_Call) Run(run func(ctx context.Context, sess *state.Session, r io.Reader, w io.Writer, sequence *uint32)) *mockBOSRouter_Route_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(io.Reader), args[3].(io.Writer), args[4].(*uint32))
	})
	return _c
}

func (_c *mockBOSRouter_Route_Call) Return(_a0 error) *mockBOSRouter_Route_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockBOSRouter_Route_Call) RunAndReturn(run func(context.Context, *state.Session, io.Reader, io.Writer, *uint32) error) *mockBOSRouter_Route_Call {
	_c.Call.Return(run)
	return _c
}

// newMockBOSRouter creates a new instance of mockBOSRouter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockBOSRouter(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockBOSRouter {
	mock := &mockBOSRouter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

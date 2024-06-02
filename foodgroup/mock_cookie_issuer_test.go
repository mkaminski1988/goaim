// Code generated by mockery v2.40.1. DO NOT EDIT.

package foodgroup

import mock "github.com/stretchr/testify/mock"

// mockCookieIssuer is an autogenerated mock type for the CookieIssuer type
type mockCookieIssuer struct {
	mock.Mock
}

type mockCookieIssuer_Expecter struct {
	mock *mock.Mock
}

func (_m *mockCookieIssuer) EXPECT() *mockCookieIssuer_Expecter {
	return &mockCookieIssuer_Expecter{mock: &_m.Mock}
}

// Issue provides a mock function with given fields: data
func (_m *mockCookieIssuer) Issue(data []byte) ([]byte, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Issue")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockCookieIssuer_Issue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Issue'
type mockCookieIssuer_Issue_Call struct {
	*mock.Call
}

// Issue is a helper method to define mock.On call
//   - data []byte
func (_e *mockCookieIssuer_Expecter) Issue(data interface{}) *mockCookieIssuer_Issue_Call {
	return &mockCookieIssuer_Issue_Call{Call: _e.mock.On("Issue", data)}
}

func (_c *mockCookieIssuer_Issue_Call) Run(run func(data []byte)) *mockCookieIssuer_Issue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *mockCookieIssuer_Issue_Call) Return(_a0 []byte, _a1 error) *mockCookieIssuer_Issue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockCookieIssuer_Issue_Call) RunAndReturn(run func([]byte) ([]byte, error)) *mockCookieIssuer_Issue_Call {
	_c.Call.Return(run)
	return _c
}

// newMockCookieIssuer creates a new instance of mockCookieIssuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockCookieIssuer(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockCookieIssuer {
	mock := &mockCookieIssuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

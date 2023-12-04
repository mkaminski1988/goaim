// Code generated by mockery v2.38.0. DO NOT EDIT.

package server

import (
	context "context"

	oscar "github.com/mkaminski/goaim/oscar"
	mock "github.com/stretchr/testify/mock"

	state "github.com/mkaminski/goaim/state"
)

// mockChatHandler is an autogenerated mock type for the ChatHandler type
type mockChatHandler struct {
	mock.Mock
}

type mockChatHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *mockChatHandler) EXPECT() *mockChatHandler_Expecter {
	return &mockChatHandler_Expecter{mock: &_m.Mock}
}

// ChannelMsgToHostHandler provides a mock function with given fields: ctx, sess, chatID, inFrame, inBody
func (_m *mockChatHandler) ChannelMsgToHostHandler(ctx context.Context, sess *state.Session, chatID string, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*oscar.SNACMessage, error) {
	ret := _m.Called(ctx, sess, chatID, inFrame, inBody)

	if len(ret) == 0 {
		panic("no return value specified for ChannelMsgToHostHandler")
	}

	var r0 *oscar.SNACMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, string, oscar.SNACFrame, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*oscar.SNACMessage, error)); ok {
		return rf(ctx, sess, chatID, inFrame, inBody)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *state.Session, string, oscar.SNACFrame, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) *oscar.SNACMessage); ok {
		r0 = rf(ctx, sess, chatID, inFrame, inBody)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oscar.SNACMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *state.Session, string, oscar.SNACFrame, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) error); ok {
		r1 = rf(ctx, sess, chatID, inFrame, inBody)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockChatHandler_ChannelMsgToHostHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChannelMsgToHostHandler'
type mockChatHandler_ChannelMsgToHostHandler_Call struct {
	*mock.Call
}

// ChannelMsgToHostHandler is a helper method to define mock.On call
//   - ctx context.Context
//   - sess *state.Session
//   - chatID string
//   - inFrame oscar.SNACFrame
//   - inBody oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost
func (_e *mockChatHandler_Expecter) ChannelMsgToHostHandler(ctx interface{}, sess interface{}, chatID interface{}, inFrame interface{}, inBody interface{}) *mockChatHandler_ChannelMsgToHostHandler_Call {
	return &mockChatHandler_ChannelMsgToHostHandler_Call{Call: _e.mock.On("ChannelMsgToHostHandler", ctx, sess, chatID, inFrame, inBody)}
}

func (_c *mockChatHandler_ChannelMsgToHostHandler_Call) Run(run func(ctx context.Context, sess *state.Session, chatID string, inFrame oscar.SNACFrame, inBody oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost)) *mockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*state.Session), args[2].(string), args[3].(oscar.SNACFrame), args[4].(oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost))
	})
	return _c
}

func (_c *mockChatHandler_ChannelMsgToHostHandler_Call) Return(_a0 *oscar.SNACMessage, _a1 error) *mockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockChatHandler_ChannelMsgToHostHandler_Call) RunAndReturn(run func(context.Context, *state.Session, string, oscar.SNACFrame, oscar.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*oscar.SNACMessage, error)) *mockChatHandler_ChannelMsgToHostHandler_Call {
	_c.Call.Return(run)
	return _c
}

// newMockChatHandler creates a new instance of mockChatHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockChatHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockChatHandler {
	mock := &mockChatHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

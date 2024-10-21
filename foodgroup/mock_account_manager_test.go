// Code generated by mockery v2.46.3. DO NOT EDIT.

package foodgroup

import (
	mail "net/mail"

	state "github.com/mk6i/retro-aim-server/state"
	mock "github.com/stretchr/testify/mock"
)

// mockAccountManager is an autogenerated mock type for the AccountManager type
type mockAccountManager struct {
	mock.Mock
}

type mockAccountManager_Expecter struct {
	mock *mock.Mock
}

func (_m *mockAccountManager) EXPECT() *mockAccountManager_Expecter {
	return &mockAccountManager_Expecter{mock: &_m.Mock}
}

// ConfirmStatusByName provides a mock function with given fields: screnName
func (_m *mockAccountManager) ConfirmStatusByName(screnName state.IdentScreenName) (bool, error) {
	ret := _m.Called(screnName)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmStatusByName")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (bool, error)); ok {
		return rf(screnName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) bool); ok {
		r0 = rf(screnName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screnName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAccountManager_ConfirmStatusByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConfirmStatusByName'
type mockAccountManager_ConfirmStatusByName_Call struct {
	*mock.Call
}

// ConfirmStatusByName is a helper method to define mock.On call
//   - screnName state.IdentScreenName
func (_e *mockAccountManager_Expecter) ConfirmStatusByName(screnName interface{}) *mockAccountManager_ConfirmStatusByName_Call {
	return &mockAccountManager_ConfirmStatusByName_Call{Call: _e.mock.On("ConfirmStatusByName", screnName)}
}

func (_c *mockAccountManager_ConfirmStatusByName_Call) Run(run func(screnName state.IdentScreenName)) *mockAccountManager_ConfirmStatusByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_ConfirmStatusByName_Call) Return(_a0 bool, _a1 error) *mockAccountManager_ConfirmStatusByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAccountManager_ConfirmStatusByName_Call) RunAndReturn(run func(state.IdentScreenName) (bool, error)) *mockAccountManager_ConfirmStatusByName_Call {
	_c.Call.Return(run)
	return _c
}

// EmailAddressByName provides a mock function with given fields: screenName
func (_m *mockAccountManager) EmailAddressByName(screenName state.IdentScreenName) (*mail.Address, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for EmailAddressByName")
	}

	var r0 *mail.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (*mail.Address, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) *mail.Address); ok {
		r0 = rf(screenName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*mail.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAccountManager_EmailAddressByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EmailAddressByName'
type mockAccountManager_EmailAddressByName_Call struct {
	*mock.Call
}

// EmailAddressByName is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockAccountManager_Expecter) EmailAddressByName(screenName interface{}) *mockAccountManager_EmailAddressByName_Call {
	return &mockAccountManager_EmailAddressByName_Call{Call: _e.mock.On("EmailAddressByName", screenName)}
}

func (_c *mockAccountManager_EmailAddressByName_Call) Run(run func(screenName state.IdentScreenName)) *mockAccountManager_EmailAddressByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_EmailAddressByName_Call) Return(_a0 *mail.Address, _a1 error) *mockAccountManager_EmailAddressByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAccountManager_EmailAddressByName_Call) RunAndReturn(run func(state.IdentScreenName) (*mail.Address, error)) *mockAccountManager_EmailAddressByName_Call {
	_c.Call.Return(run)
	return _c
}

// RegStatusByName provides a mock function with given fields: screenName
func (_m *mockAccountManager) RegStatusByName(screenName state.IdentScreenName) (uint16, error) {
	ret := _m.Called(screenName)

	if len(ret) == 0 {
		panic("no return value specified for RegStatusByName")
	}

	var r0 uint16
	var r1 error
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) (uint16, error)); ok {
		return rf(screenName)
	}
	if rf, ok := ret.Get(0).(func(state.IdentScreenName) uint16); ok {
		r0 = rf(screenName)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	if rf, ok := ret.Get(1).(func(state.IdentScreenName) error); ok {
		r1 = rf(screenName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockAccountManager_RegStatusByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegStatusByName'
type mockAccountManager_RegStatusByName_Call struct {
	*mock.Call
}

// RegStatusByName is a helper method to define mock.On call
//   - screenName state.IdentScreenName
func (_e *mockAccountManager_Expecter) RegStatusByName(screenName interface{}) *mockAccountManager_RegStatusByName_Call {
	return &mockAccountManager_RegStatusByName_Call{Call: _e.mock.On("RegStatusByName", screenName)}
}

func (_c *mockAccountManager_RegStatusByName_Call) Run(run func(screenName state.IdentScreenName)) *mockAccountManager_RegStatusByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_RegStatusByName_Call) Return(_a0 uint16, _a1 error) *mockAccountManager_RegStatusByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockAccountManager_RegStatusByName_Call) RunAndReturn(run func(state.IdentScreenName) (uint16, error)) *mockAccountManager_RegStatusByName_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateConfirmStatus provides a mock function with given fields: confirmStatus, screenName
func (_m *mockAccountManager) UpdateConfirmStatus(confirmStatus bool, screenName state.IdentScreenName) error {
	ret := _m.Called(confirmStatus, screenName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConfirmStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, state.IdentScreenName) error); ok {
		r0 = rf(confirmStatus, screenName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAccountManager_UpdateConfirmStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateConfirmStatus'
type mockAccountManager_UpdateConfirmStatus_Call struct {
	*mock.Call
}

// UpdateConfirmStatus is a helper method to define mock.On call
//   - confirmStatus bool
//   - screenName state.IdentScreenName
func (_e *mockAccountManager_Expecter) UpdateConfirmStatus(confirmStatus interface{}, screenName interface{}) *mockAccountManager_UpdateConfirmStatus_Call {
	return &mockAccountManager_UpdateConfirmStatus_Call{Call: _e.mock.On("UpdateConfirmStatus", confirmStatus, screenName)}
}

func (_c *mockAccountManager_UpdateConfirmStatus_Call) Run(run func(confirmStatus bool, screenName state.IdentScreenName)) *mockAccountManager_UpdateConfirmStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_UpdateConfirmStatus_Call) Return(_a0 error) *mockAccountManager_UpdateConfirmStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAccountManager_UpdateConfirmStatus_Call) RunAndReturn(run func(bool, state.IdentScreenName) error) *mockAccountManager_UpdateConfirmStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDisplayScreenName provides a mock function with given fields: displayScreenName
func (_m *mockAccountManager) UpdateDisplayScreenName(displayScreenName state.DisplayScreenName) error {
	ret := _m.Called(displayScreenName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDisplayScreenName")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(state.DisplayScreenName) error); ok {
		r0 = rf(displayScreenName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAccountManager_UpdateDisplayScreenName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDisplayScreenName'
type mockAccountManager_UpdateDisplayScreenName_Call struct {
	*mock.Call
}

// UpdateDisplayScreenName is a helper method to define mock.On call
//   - displayScreenName state.DisplayScreenName
func (_e *mockAccountManager_Expecter) UpdateDisplayScreenName(displayScreenName interface{}) *mockAccountManager_UpdateDisplayScreenName_Call {
	return &mockAccountManager_UpdateDisplayScreenName_Call{Call: _e.mock.On("UpdateDisplayScreenName", displayScreenName)}
}

func (_c *mockAccountManager_UpdateDisplayScreenName_Call) Run(run func(displayScreenName state.DisplayScreenName)) *mockAccountManager_UpdateDisplayScreenName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.DisplayScreenName))
	})
	return _c
}

func (_c *mockAccountManager_UpdateDisplayScreenName_Call) Return(_a0 error) *mockAccountManager_UpdateDisplayScreenName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAccountManager_UpdateDisplayScreenName_Call) RunAndReturn(run func(state.DisplayScreenName) error) *mockAccountManager_UpdateDisplayScreenName_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEmailAddress provides a mock function with given fields: emailAddress, screenName
func (_m *mockAccountManager) UpdateEmailAddress(emailAddress *mail.Address, screenName state.IdentScreenName) error {
	ret := _m.Called(emailAddress, screenName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEmailAddress")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*mail.Address, state.IdentScreenName) error); ok {
		r0 = rf(emailAddress, screenName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAccountManager_UpdateEmailAddress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEmailAddress'
type mockAccountManager_UpdateEmailAddress_Call struct {
	*mock.Call
}

// UpdateEmailAddress is a helper method to define mock.On call
//   - emailAddress *mail.Address
//   - screenName state.IdentScreenName
func (_e *mockAccountManager_Expecter) UpdateEmailAddress(emailAddress interface{}, screenName interface{}) *mockAccountManager_UpdateEmailAddress_Call {
	return &mockAccountManager_UpdateEmailAddress_Call{Call: _e.mock.On("UpdateEmailAddress", emailAddress, screenName)}
}

func (_c *mockAccountManager_UpdateEmailAddress_Call) Run(run func(emailAddress *mail.Address, screenName state.IdentScreenName)) *mockAccountManager_UpdateEmailAddress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*mail.Address), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_UpdateEmailAddress_Call) Return(_a0 error) *mockAccountManager_UpdateEmailAddress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAccountManager_UpdateEmailAddress_Call) RunAndReturn(run func(*mail.Address, state.IdentScreenName) error) *mockAccountManager_UpdateEmailAddress_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRegStatus provides a mock function with given fields: regStatus, screenName
func (_m *mockAccountManager) UpdateRegStatus(regStatus uint16, screenName state.IdentScreenName) error {
	ret := _m.Called(regStatus, screenName)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRegStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint16, state.IdentScreenName) error); ok {
		r0 = rf(regStatus, screenName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockAccountManager_UpdateRegStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRegStatus'
type mockAccountManager_UpdateRegStatus_Call struct {
	*mock.Call
}

// UpdateRegStatus is a helper method to define mock.On call
//   - regStatus uint16
//   - screenName state.IdentScreenName
func (_e *mockAccountManager_Expecter) UpdateRegStatus(regStatus interface{}, screenName interface{}) *mockAccountManager_UpdateRegStatus_Call {
	return &mockAccountManager_UpdateRegStatus_Call{Call: _e.mock.On("UpdateRegStatus", regStatus, screenName)}
}

func (_c *mockAccountManager_UpdateRegStatus_Call) Run(run func(regStatus uint16, screenName state.IdentScreenName)) *mockAccountManager_UpdateRegStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uint16), args[1].(state.IdentScreenName))
	})
	return _c
}

func (_c *mockAccountManager_UpdateRegStatus_Call) Return(_a0 error) *mockAccountManager_UpdateRegStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockAccountManager_UpdateRegStatus_Call) RunAndReturn(run func(uint16, state.IdentScreenName) error) *mockAccountManager_UpdateRegStatus_Call {
	_c.Call.Return(run)
	return _c
}

// newMockAccountManager creates a new instance of mockAccountManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockAccountManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockAccountManager {
	mock := &mockAccountManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

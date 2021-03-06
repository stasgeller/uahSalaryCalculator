// Code generated by MockGen. DO NOT EDIT.
// Source: ./external/telegram/client.go

// Package mock_telegram is a generated GoMock package.
package mock_telegram

import (
	reflect "reflect"

	v5 "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	gomock "github.com/golang/mock/gomock"
)

// MockTgClient is a mock of TgClient interface.
type MockTgClient struct {
	ctrl     *gomock.Controller
	recorder *MockTgClientMockRecorder
}

// MockTgClientMockRecorder is the mock recorder for MockTgClient.
type MockTgClientMockRecorder struct {
	mock *MockTgClient
}

// NewMockTgClient creates a new mock instance.
func NewMockTgClient(ctrl *gomock.Controller) *MockTgClient {
	mock := &MockTgClient{ctrl: ctrl}
	mock.recorder = &MockTgClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTgClient) EXPECT() *MockTgClientMockRecorder {
	return m.recorder
}

// GetUpdatesChan mocks base method.
func (m *MockTgClient) GetUpdatesChan(arg0 v5.UpdateConfig) v5.UpdatesChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdatesChan", arg0)
	ret0, _ := ret[0].(v5.UpdatesChannel)
	return ret0
}

// GetUpdatesChan indicates an expected call of GetUpdatesChan.
func (mr *MockTgClientMockRecorder) GetUpdatesChan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdatesChan", reflect.TypeOf((*MockTgClient)(nil).GetUpdatesChan), arg0)
}

// Send mocks base method.
func (m *MockTgClient) Send(arg0 v5.Chattable) (v5.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(v5.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockTgClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTgClient)(nil).Send), arg0)
}

// StopReceivingUpdates mocks base method.
func (m *MockTgClient) StopReceivingUpdates() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopReceivingUpdates")
}

// StopReceivingUpdates indicates an expected call of StopReceivingUpdates.
func (mr *MockTgClientMockRecorder) StopReceivingUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopReceivingUpdates", reflect.TypeOf((*MockTgClient)(nil).StopReceivingUpdates))
}

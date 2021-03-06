// The MIT License (MIT)
//
// Copyright (c) 2017-2020 Uber Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: resetor.go

// Package reset is a generated GoMock package.
package reset

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	history "github.com/uber/cadence/.gen/go/history"
	shared "github.com/uber/cadence/.gen/go/shared"
	execution "github.com/uber/cadence/service/history/execution"
)

// MockWorkflowResetor is a mock of WorkflowResetor interface
type MockWorkflowResetor struct {
	ctrl     *gomock.Controller
	recorder *MockWorkflowResetorMockRecorder
}

// MockWorkflowResetorMockRecorder is the mock recorder for MockWorkflowResetor
type MockWorkflowResetorMockRecorder struct {
	mock *MockWorkflowResetor
}

// NewMockWorkflowResetor creates a new mock instance
func NewMockWorkflowResetor(ctrl *gomock.Controller) *MockWorkflowResetor {
	mock := &MockWorkflowResetor{ctrl: ctrl}
	mock.recorder = &MockWorkflowResetorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWorkflowResetor) EXPECT() *MockWorkflowResetorMockRecorder {
	return m.recorder
}

// ResetWorkflowExecution mocks base method
func (m *MockWorkflowResetor) ResetWorkflowExecution(ctx context.Context, resetRequest *shared.ResetWorkflowExecutionRequest, baseContext execution.Context, baseMutableState execution.MutableState, currContext execution.Context, currMutableState execution.MutableState) (*shared.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", ctx, resetRequest, baseContext, baseMutableState, currContext, currMutableState)
	ret0, _ := ret[0].(*shared.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution
func (mr *MockWorkflowResetorMockRecorder) ResetWorkflowExecution(ctx, resetRequest, baseContext, baseMutableState, currContext, currMutableState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockWorkflowResetor)(nil).ResetWorkflowExecution), ctx, resetRequest, baseContext, baseMutableState, currContext, currMutableState)
}

// ApplyResetEvent mocks base method
func (m *MockWorkflowResetor) ApplyResetEvent(ctx context.Context, request *history.ReplicateEventsRequest, domainID, workflowID, currentRunID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyResetEvent", ctx, request, domainID, workflowID, currentRunID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyResetEvent indicates an expected call of ApplyResetEvent
func (mr *MockWorkflowResetorMockRecorder) ApplyResetEvent(ctx, request, domainID, workflowID, currentRunID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyResetEvent", reflect.TypeOf((*MockWorkflowResetor)(nil).ApplyResetEvent), ctx, request, domainID, workflowID, currentRunID)
}

// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/tmds/handlers/v4/state (interfaces: AgentState)

// Package mock_state is a generated GoMock package.
package mock_state

import (
	reflect "reflect"

	state "github.com/aws/amazon-ecs-agent/ecs-agent/tmds/handlers/v4/state"
	gomock "github.com/golang/mock/gomock"
)

// MockAgentState is a mock of AgentState interface.
type MockAgentState struct {
	ctrl     *gomock.Controller
	recorder *MockAgentStateMockRecorder
}

// MockAgentStateMockRecorder is the mock recorder for MockAgentState.
type MockAgentStateMockRecorder struct {
	mock *MockAgentState
}

// NewMockAgentState creates a new mock instance.
func NewMockAgentState(ctrl *gomock.Controller) *MockAgentState {
	mock := &MockAgentState{ctrl: ctrl}
	mock.recorder = &MockAgentStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentState) EXPECT() *MockAgentStateMockRecorder {
	return m.recorder
}

// GetContainerMetadata mocks base method.
func (m *MockAgentState) GetContainerMetadata(arg0 string) (state.ContainerResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerMetadata", arg0)
	ret0, _ := ret[0].(state.ContainerResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerMetadata indicates an expected call of GetContainerMetadata.
func (mr *MockAgentStateMockRecorder) GetContainerMetadata(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerMetadata", reflect.TypeOf((*MockAgentState)(nil).GetContainerMetadata), arg0)
}

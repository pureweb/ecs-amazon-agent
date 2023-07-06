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
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/tmds/handlers/taskprotection/v1/handlers (interfaces: TaskProtectionClientFactoryInterface)

// Package handlers is a generated GoMock package.
package handlers

import (
	reflect "reflect"

	api "github.com/aws/amazon-ecs-agent/ecs-agent/api"
	credentials "github.com/aws/amazon-ecs-agent/ecs-agent/credentials"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskProtectionClientFactoryInterface is a mock of TaskProtectionClientFactoryInterface interface.
type MockTaskProtectionClientFactoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTaskProtectionClientFactoryInterfaceMockRecorder
}

// MockTaskProtectionClientFactoryInterfaceMockRecorder is the mock recorder for MockTaskProtectionClientFactoryInterface.
type MockTaskProtectionClientFactoryInterfaceMockRecorder struct {
	mock *MockTaskProtectionClientFactoryInterface
}

// NewMockTaskProtectionClientFactoryInterface creates a new mock instance.
func NewMockTaskProtectionClientFactoryInterface(ctrl *gomock.Controller) *MockTaskProtectionClientFactoryInterface {
	mock := &MockTaskProtectionClientFactoryInterface{ctrl: ctrl}
	mock.recorder = &MockTaskProtectionClientFactoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskProtectionClientFactoryInterface) EXPECT() *MockTaskProtectionClientFactoryInterfaceMockRecorder {
	return m.recorder
}

// NewTaskProtectionClient mocks base method.
func (m *MockTaskProtectionClientFactoryInterface) NewTaskProtectionClient(arg0 credentials.TaskIAMRoleCredentials) api.ECSTaskProtectionSDK {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTaskProtectionClient", arg0)
	ret0, _ := ret[0].(api.ECSTaskProtectionSDK)
	return ret0
}

// NewTaskProtectionClient indicates an expected call of NewTaskProtectionClient.
func (mr *MockTaskProtectionClientFactoryInterfaceMockRecorder) NewTaskProtectionClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTaskProtectionClient", reflect.TypeOf((*MockTaskProtectionClientFactoryInterface)(nil).NewTaskProtectionClient), arg0)
}

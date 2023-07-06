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
// Source: github.com/aws/amazon-ecs-agent/agent/ecscni (interfaces: NamespaceHelper)

// Package mock_ecscni is a generated GoMock package.
package mock_ecscni

import (
	context "context"
	reflect "reflect"

	ecscni "github.com/aws/amazon-ecs-agent/agent/ecscni"
	eni "github.com/aws/amazon-ecs-agent/ecs-agent/api/eni"
	types100 "github.com/containernetworking/cni/pkg/types/100"
	gomock "github.com/golang/mock/gomock"
)

// MockNamespaceHelper is a mock of NamespaceHelper interface.
type MockNamespaceHelper struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceHelperMockRecorder
}

// MockNamespaceHelperMockRecorder is the mock recorder for MockNamespaceHelper.
type MockNamespaceHelperMockRecorder struct {
	mock *MockNamespaceHelper
}

// NewMockNamespaceHelper creates a new mock instance.
func NewMockNamespaceHelper(ctrl *gomock.Controller) *MockNamespaceHelper {
	mock := &MockNamespaceHelper{ctrl: ctrl}
	mock.recorder = &MockNamespaceHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceHelper) EXPECT() *MockNamespaceHelperMockRecorder {
	return m.recorder
}

// ConfigureTaskNamespaceRouting mocks base method.
func (m *MockNamespaceHelper) ConfigureTaskNamespaceRouting(arg0 context.Context, arg1 *eni.ENI, arg2 *ecscni.Config, arg3 *types100.Result) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureTaskNamespaceRouting", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfigureTaskNamespaceRouting indicates an expected call of ConfigureTaskNamespaceRouting.
func (mr *MockNamespaceHelperMockRecorder) ConfigureTaskNamespaceRouting(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureTaskNamespaceRouting", reflect.TypeOf((*MockNamespaceHelper)(nil).ConfigureTaskNamespaceRouting), arg0, arg1, arg2, arg3)
}

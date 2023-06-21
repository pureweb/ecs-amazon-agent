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
// Source: github.com/containernetworking/cni/pkg/types (interfaces: Result)

// Package mock_types is a generated GoMock package.
package mock_types

import (
	io "io"
	reflect "reflect"

	types "github.com/containernetworking/cni/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockResult is a mock of Result interface.
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult.
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance.
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// GetAsVersion mocks base method.
func (m *MockResult) GetAsVersion(arg0 string) (types.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAsVersion", arg0)
	ret0, _ := ret[0].(types.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAsVersion indicates an expected call of GetAsVersion.
func (mr *MockResultMockRecorder) GetAsVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAsVersion", reflect.TypeOf((*MockResult)(nil).GetAsVersion), arg0)
}

// Print mocks base method.
func (m *MockResult) Print() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print")
	ret0, _ := ret[0].(error)
	return ret0
}

// Print indicates an expected call of Print.
func (mr *MockResultMockRecorder) Print() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockResult)(nil).Print))
}

// PrintTo mocks base method.
func (m *MockResult) PrintTo(arg0 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintTo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrintTo indicates an expected call of PrintTo.
func (mr *MockResultMockRecorder) PrintTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTo", reflect.TypeOf((*MockResult)(nil).PrintTo), arg0)
}

// Version mocks base method.
func (m *MockResult) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockResultMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockResult)(nil).Version))
}

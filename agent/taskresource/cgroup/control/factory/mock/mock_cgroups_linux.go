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
// Source: github.com/containerd/cgroups (interfaces: Cgroup)

// Package mock_cgroups is a generated GoMock package.
package mock_cgroups

import (
	reflect "reflect"

	cgroups "github.com/containerd/cgroups"
	v1 "github.com/containerd/cgroups/stats/v1"
	gomock "github.com/golang/mock/gomock"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// MockCgroup is a mock of Cgroup interface.
type MockCgroup struct {
	ctrl     *gomock.Controller
	recorder *MockCgroupMockRecorder
}

// MockCgroupMockRecorder is the mock recorder for MockCgroup.
type MockCgroupMockRecorder struct {
	mock *MockCgroup
}

// NewMockCgroup creates a new mock instance.
func NewMockCgroup(ctrl *gomock.Controller) *MockCgroup {
	mock := &MockCgroup{ctrl: ctrl}
	mock.recorder = &MockCgroupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCgroup) EXPECT() *MockCgroupMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCgroup) Add(arg0 cgroups.Process, arg1 ...cgroups.Name) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockCgroupMockRecorder) Add(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCgroup)(nil).Add), varargs...)
}

// AddProc mocks base method.
func (m *MockCgroup) AddProc(arg0 uint64, arg1 ...cgroups.Name) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProc", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProc indicates an expected call of AddProc.
func (mr *MockCgroupMockRecorder) AddProc(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProc", reflect.TypeOf((*MockCgroup)(nil).AddProc), varargs...)
}

// AddTask mocks base method.
func (m *MockCgroup) AddTask(arg0 cgroups.Process, arg1 ...cgroups.Name) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask.
func (mr *MockCgroupMockRecorder) AddTask(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockCgroup)(nil).AddTask), varargs...)
}

// Delete mocks base method.
func (m *MockCgroup) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCgroupMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCgroup)(nil).Delete))
}

// Freeze mocks base method.
func (m *MockCgroup) Freeze() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Freeze")
	ret0, _ := ret[0].(error)
	return ret0
}

// Freeze indicates an expected call of Freeze.
func (mr *MockCgroupMockRecorder) Freeze() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Freeze", reflect.TypeOf((*MockCgroup)(nil).Freeze))
}

// MoveTo mocks base method.
func (m *MockCgroup) MoveTo(arg0 cgroups.Cgroup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveTo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveTo indicates an expected call of MoveTo.
func (mr *MockCgroupMockRecorder) MoveTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveTo", reflect.TypeOf((*MockCgroup)(nil).MoveTo), arg0)
}

// New mocks base method.
func (m *MockCgroup) New(arg0 string, arg1 *specs.LinuxResources) (cgroups.Cgroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "New", arg0, arg1)
	ret0, _ := ret[0].(cgroups.Cgroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// New indicates an expected call of New.
func (mr *MockCgroupMockRecorder) New(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "New", reflect.TypeOf((*MockCgroup)(nil).New), arg0, arg1)
}

// OOMEventFD mocks base method.
func (m *MockCgroup) OOMEventFD() (uintptr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OOMEventFD")
	ret0, _ := ret[0].(uintptr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OOMEventFD indicates an expected call of OOMEventFD.
func (mr *MockCgroupMockRecorder) OOMEventFD() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OOMEventFD", reflect.TypeOf((*MockCgroup)(nil).OOMEventFD))
}

// Processes mocks base method.
func (m *MockCgroup) Processes(arg0 cgroups.Name, arg1 bool) ([]cgroups.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Processes", arg0, arg1)
	ret0, _ := ret[0].([]cgroups.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Processes indicates an expected call of Processes.
func (mr *MockCgroupMockRecorder) Processes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Processes", reflect.TypeOf((*MockCgroup)(nil).Processes), arg0, arg1)
}

// RegisterMemoryEvent mocks base method.
func (m *MockCgroup) RegisterMemoryEvent(arg0 cgroups.MemoryEvent) (uintptr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterMemoryEvent", arg0)
	ret0, _ := ret[0].(uintptr)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterMemoryEvent indicates an expected call of RegisterMemoryEvent.
func (mr *MockCgroupMockRecorder) RegisterMemoryEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterMemoryEvent", reflect.TypeOf((*MockCgroup)(nil).RegisterMemoryEvent), arg0)
}

// Stat mocks base method.
func (m *MockCgroup) Stat(arg0 ...cgroups.ErrorHandler) (*v1.Metrics, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stat", varargs...)
	ret0, _ := ret[0].(*v1.Metrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockCgroupMockRecorder) Stat(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockCgroup)(nil).Stat), arg0...)
}

// State mocks base method.
func (m *MockCgroup) State() cgroups.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(cgroups.State)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockCgroupMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockCgroup)(nil).State))
}

// Subsystems mocks base method.
func (m *MockCgroup) Subsystems() []cgroups.Subsystem {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subsystems")
	ret0, _ := ret[0].([]cgroups.Subsystem)
	return ret0
}

// Subsystems indicates an expected call of Subsystems.
func (mr *MockCgroupMockRecorder) Subsystems() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subsystems", reflect.TypeOf((*MockCgroup)(nil).Subsystems))
}

// Tasks mocks base method.
func (m *MockCgroup) Tasks(arg0 cgroups.Name, arg1 bool) ([]cgroups.Process, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tasks", arg0, arg1)
	ret0, _ := ret[0].([]cgroups.Process)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tasks indicates an expected call of Tasks.
func (mr *MockCgroupMockRecorder) Tasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tasks", reflect.TypeOf((*MockCgroup)(nil).Tasks), arg0, arg1)
}

// Thaw mocks base method.
func (m *MockCgroup) Thaw() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Thaw")
	ret0, _ := ret[0].(error)
	return ret0
}

// Thaw indicates an expected call of Thaw.
func (mr *MockCgroupMockRecorder) Thaw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Thaw", reflect.TypeOf((*MockCgroup)(nil).Thaw))
}

// Update mocks base method.
func (m *MockCgroup) Update(arg0 *specs.LinuxResources) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCgroupMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCgroup)(nil).Update), arg0)
}

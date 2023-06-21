// Copyright 2015-2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: dependencies.go in package docker
// Code generated by MockGen. DO NOT EDIT.

// Package docker is a generated GoMock package.
package docker

import (
	reflect "reflect"

	docker "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
)

// Mockdockerclient is a mock of dockerclient interface.
type Mockdockerclient struct {
	ctrl     *gomock.Controller
	recorder *MockdockerclientMockRecorder
}

// MockdockerclientMockRecorder is the mock recorder for Mockdockerclient.
type MockdockerclientMockRecorder struct {
	mock *Mockdockerclient
}

// NewMockdockerclient creates a new mock instance.
func NewMockdockerclient(ctrl *gomock.Controller) *Mockdockerclient {
	mock := &Mockdockerclient{ctrl: ctrl}
	mock.recorder = &MockdockerclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockdockerclient) EXPECT() *MockdockerclientMockRecorder {
	return m.recorder
}

// CreateContainer mocks base method.
func (m *Mockdockerclient) CreateContainer(opts docker.CreateContainerOptions) (*docker.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContainer", opts)
	ret0, _ := ret[0].(*docker.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer.
func (mr *MockdockerclientMockRecorder) CreateContainer(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*Mockdockerclient)(nil).CreateContainer), opts)
}

// ListContainers mocks base method.
func (m *Mockdockerclient) ListContainers(opts docker.ListContainersOptions) ([]docker.APIContainers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContainers", opts)
	ret0, _ := ret[0].([]docker.APIContainers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers.
func (mr *MockdockerclientMockRecorder) ListContainers(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*Mockdockerclient)(nil).ListContainers), opts)
}

// ListImages mocks base method.
func (m *Mockdockerclient) ListImages(opts docker.ListImagesOptions) ([]docker.APIImages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", opts)
	ret0, _ := ret[0].([]docker.APIImages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockdockerclientMockRecorder) ListImages(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*Mockdockerclient)(nil).ListImages), opts)
}

// LoadImage mocks base method.
func (m *Mockdockerclient) LoadImage(opts docker.LoadImageOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadImage", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage.
func (mr *MockdockerclientMockRecorder) LoadImage(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*Mockdockerclient)(nil).LoadImage), opts)
}

// Logs mocks base method.
func (m *Mockdockerclient) Logs(opts docker.LogsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logs", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logs indicates an expected call of Logs.
func (mr *MockdockerclientMockRecorder) Logs(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logs", reflect.TypeOf((*Mockdockerclient)(nil).Logs), opts)
}

// Ping mocks base method.
func (m *Mockdockerclient) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockdockerclientMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*Mockdockerclient)(nil).Ping))
}

// RemoveContainer mocks base method.
func (m *Mockdockerclient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveContainer", opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer.
func (mr *MockdockerclientMockRecorder) RemoveContainer(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*Mockdockerclient)(nil).RemoveContainer), opts)
}

// StartContainer mocks base method.
func (m *Mockdockerclient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartContainer", id, hostConfig)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainer indicates an expected call of StartContainer.
func (mr *MockdockerclientMockRecorder) StartContainer(id, hostConfig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*Mockdockerclient)(nil).StartContainer), id, hostConfig)
}

// StopContainer mocks base method.
func (m *Mockdockerclient) StopContainer(id string, timeout uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopContainer", id, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer.
func (mr *MockdockerclientMockRecorder) StopContainer(id, timeout interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*Mockdockerclient)(nil).StopContainer), id, timeout)
}

// WaitContainer mocks base method.
func (m *Mockdockerclient) WaitContainer(id string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitContainer", id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitContainer indicates an expected call of WaitContainer.
func (mr *MockdockerclientMockRecorder) WaitContainer(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitContainer", reflect.TypeOf((*Mockdockerclient)(nil).WaitContainer), id)
}

// MockdockerClientFactory is a mock of dockerClientFactory interface.
type MockdockerClientFactory struct {
	ctrl     *gomock.Controller
	recorder *MockdockerClientFactoryMockRecorder
}

// MockdockerClientFactoryMockRecorder is the mock recorder for MockdockerClientFactory.
type MockdockerClientFactoryMockRecorder struct {
	mock *MockdockerClientFactory
}

// NewMockdockerClientFactory creates a new mock instance.
func NewMockdockerClientFactory(ctrl *gomock.Controller) *MockdockerClientFactory {
	mock := &MockdockerClientFactory{ctrl: ctrl}
	mock.recorder = &MockdockerClientFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdockerClientFactory) EXPECT() *MockdockerClientFactoryMockRecorder {
	return m.recorder
}

// NewVersionedClient mocks base method.
func (m *MockdockerClientFactory) NewVersionedClient(endpoint, apiVersionString string) (dockerclient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVersionedClient", endpoint, apiVersionString)
	ret0, _ := ret[0].(dockerclient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVersionedClient indicates an expected call of NewVersionedClient.
func (mr *MockdockerClientFactoryMockRecorder) NewVersionedClient(endpoint, apiVersionString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVersionedClient", reflect.TypeOf((*MockdockerClientFactory)(nil).NewVersionedClient), endpoint, apiVersionString)
}

// MockfileSystem is a mock of fileSystem interface.
type MockfileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockfileSystemMockRecorder
}

// MockfileSystemMockRecorder is the mock recorder for MockfileSystem.
type MockfileSystemMockRecorder struct {
	mock *MockfileSystem
}

// NewMockfileSystem creates a new mock instance.
func NewMockfileSystem(ctrl *gomock.Controller) *MockfileSystem {
	mock := &MockfileSystem{ctrl: ctrl}
	mock.recorder = &MockfileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileSystem) EXPECT() *MockfileSystemMockRecorder {
	return m.recorder
}

// ReadFile mocks base method.
func (m *MockfileSystem) ReadFile(filename string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", filename)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockfileSystemMockRecorder) ReadFile(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockfileSystem)(nil).ReadFile), filename)
}

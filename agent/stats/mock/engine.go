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
// Source: github.com/aws/amazon-ecs-agent/agent/stats (interfaces: Engine)

// Package mock_stats is a generated GoMock package.
package mock_stats

import (
	reflect "reflect"
	time "time"

	stats "github.com/aws/amazon-ecs-agent/agent/stats"
	ecstcs "github.com/aws/amazon-ecs-agent/agent/tcs/model/ecstcs"
	types "github.com/docker/docker/api/types"
	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// ContainerDockerStats mocks base method.
func (m *MockEngine) ContainerDockerStats(arg0, arg1 string) (*types.StatsJSON, *stats.NetworkStatsPerSec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerDockerStats", arg0, arg1)
	ret0, _ := ret[0].(*types.StatsJSON)
	ret1, _ := ret[1].(*stats.NetworkStatsPerSec)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ContainerDockerStats indicates an expected call of ContainerDockerStats.
func (mr *MockEngineMockRecorder) ContainerDockerStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerDockerStats", reflect.TypeOf((*MockEngine)(nil).ContainerDockerStats), arg0, arg1)
}

// GetInstanceMetrics mocks base method.
func (m *MockEngine) GetInstanceMetrics(arg0 bool) (*ecstcs.MetricsMetadata, []*ecstcs.TaskMetric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceMetrics", arg0)
	ret0, _ := ret[0].(*ecstcs.MetricsMetadata)
	ret1, _ := ret[1].([]*ecstcs.TaskMetric)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetInstanceMetrics indicates an expected call of GetInstanceMetrics.
func (mr *MockEngineMockRecorder) GetInstanceMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceMetrics", reflect.TypeOf((*MockEngine)(nil).GetInstanceMetrics), arg0)
}

// GetPublishMetricsTicker mocks base method.
func (m *MockEngine) GetPublishMetricsTicker() *time.Ticker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishMetricsTicker")
	ret0, _ := ret[0].(*time.Ticker)
	return ret0
}

// GetPublishMetricsTicker indicates an expected call of GetPublishMetricsTicker.
func (mr *MockEngineMockRecorder) GetPublishMetricsTicker() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishMetricsTicker", reflect.TypeOf((*MockEngine)(nil).GetPublishMetricsTicker))
}

// GetPublishServiceConnectTickerInterval mocks base method.
func (m *MockEngine) GetPublishServiceConnectTickerInterval() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishServiceConnectTickerInterval")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetPublishServiceConnectTickerInterval indicates an expected call of GetPublishServiceConnectTickerInterval.
func (mr *MockEngineMockRecorder) GetPublishServiceConnectTickerInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishServiceConnectTickerInterval", reflect.TypeOf((*MockEngine)(nil).GetPublishServiceConnectTickerInterval))
}

// GetTaskHealthMetrics mocks base method.
func (m *MockEngine) GetTaskHealthMetrics() (*ecstcs.HealthMetadata, []*ecstcs.TaskHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskHealthMetrics")
	ret0, _ := ret[0].(*ecstcs.HealthMetadata)
	ret1, _ := ret[1].([]*ecstcs.TaskHealth)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTaskHealthMetrics indicates an expected call of GetTaskHealthMetrics.
func (mr *MockEngineMockRecorder) GetTaskHealthMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskHealthMetrics", reflect.TypeOf((*MockEngine)(nil).GetTaskHealthMetrics))
}

// SetPublishServiceConnectTickerInterval mocks base method.
func (m *MockEngine) SetPublishServiceConnectTickerInterval(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPublishServiceConnectTickerInterval", arg0)
}

// SetPublishServiceConnectTickerInterval indicates an expected call of SetPublishServiceConnectTickerInterval.
func (mr *MockEngineMockRecorder) SetPublishServiceConnectTickerInterval(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPublishServiceConnectTickerInterval", reflect.TypeOf((*MockEngine)(nil).SetPublishServiceConnectTickerInterval), arg0)
}

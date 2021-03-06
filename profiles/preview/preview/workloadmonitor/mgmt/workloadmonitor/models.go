// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package workloadmonitor

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/workloadmonitor/mgmt/2020-01-13-preview/workloadmonitor"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type HealthState = original.HealthState

const (
	Critical HealthState = original.Critical
	Healthy  HealthState = original.Healthy
	Unknown  HealthState = original.Unknown
	Warning  HealthState = original.Warning
)

type BaseClient = original.BaseClient
type DefaultError = original.DefaultError
type DefaultErrorError = original.DefaultErrorError
type ErrorDetails = original.ErrorDetails
type Monitor = original.Monitor
type MonitorList = original.MonitorList
type MonitorListIterator = original.MonitorListIterator
type MonitorListPage = original.MonitorListPage
type MonitorProperties = original.MonitorProperties
type MonitorStateChange = original.MonitorStateChange
type MonitorStateChangeList = original.MonitorStateChangeList
type MonitorStateChangeListIterator = original.MonitorStateChangeListIterator
type MonitorStateChangeListPage = original.MonitorStateChangeListPage
type MonitorStateChangeProperties = original.MonitorStateChangeProperties
type MonitorsClient = original.MonitorsClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource

func New() BaseClient {
	return original.New()
}
func NewMonitorListIterator(page MonitorListPage) MonitorListIterator {
	return original.NewMonitorListIterator(page)
}
func NewMonitorListPage(getNextPage func(context.Context, MonitorList) (MonitorList, error)) MonitorListPage {
	return original.NewMonitorListPage(getNextPage)
}
func NewMonitorStateChangeListIterator(page MonitorStateChangeListPage) MonitorStateChangeListIterator {
	return original.NewMonitorStateChangeListIterator(page)
}
func NewMonitorStateChangeListPage(getNextPage func(context.Context, MonitorStateChangeList) (MonitorStateChangeList, error)) MonitorStateChangeListPage {
	return original.NewMonitorStateChangeListPage(getNextPage)
}
func NewMonitorsClient() MonitorsClient {
	return original.NewMonitorsClient()
}
func NewMonitorsClientWithBaseURI(baseURI string) MonitorsClient {
	return original.NewMonitorsClientWithBaseURI(baseURI)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleHealthStateValues() []HealthState {
	return original.PossibleHealthStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}

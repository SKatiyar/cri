/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package performance provides commands and events for Performance domain.
package performance

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Performance domain
const (
	Enable     = "Performance.enable"
	Disable    = "Performance.disable"
	GetMetrics = "Performance.getMetrics"
)

// List of events in Performance domain
const (
	Metrics = "Performance.metrics"
)

type Performance struct {
	conn cri.Connector
}

// New creates a Performance instance
func New(conn cri.Connector) *Performance {
	return &Performance{conn}
}

// Enable collecting and reporting metrics.
func (obj *Performance) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disable collecting and reporting metrics.
func (obj *Performance) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type GetMetricsResponse struct {
	// Current values for run-time metrics.
	Metrics []types.Performance_Metric `json:"metrics"`
}

// Retrieve current values of run-time metrics.
func (obj *Performance) GetMetrics() (response GetMetricsResponse, err error) {
	err = obj.conn.Send(GetMetrics, nil, &response)
	return
}

type MetricsParams struct {
	// Current values of the metrics.
	Metrics []types.Performance_Metric `json:"metrics"`
	// Timestamp title.
	Title string `json:"title"`
}

// Current values of the metrics.
func (obj *Performance) Metrics(fn func(params *MetricsParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(Metrics, closeChn)
	go func() {
		for {
			params := MetricsParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

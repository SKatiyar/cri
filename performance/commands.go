package performance

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Performance struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Performance {
	return &Performance{conn}
}

// Enable collecting and reporting metrics.
func (obj *Performance) Enable() (err error) {
	err = obj.conn.Send("Performance.enable", nil, nil)
	return
}

// Disable collecting and reporting metrics.
func (obj *Performance) Disable() (err error) {
	err = obj.conn.Send("Performance.disable", nil, nil)
	return
}

type GetMetricsResponse struct {
	// Current values for run-time metrics.
	Metrics []types.Performance_Metric `json:"metrics"`
}

// Retrieve current values of run-time metrics.
func (obj *Performance) GetMetrics() (response GetMetricsResponse, err error) {
	err = obj.conn.Send("Performance.getMetrics", nil, &response)
	return
}

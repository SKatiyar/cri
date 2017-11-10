package profiler

import types "github.com/SKatiyar/cri/types"
import "github.com/SKatiyar/cri"

type Profiler struct {
	conn cri.Connector
}

func New(conn cri.Connector) *Profiler {
	return &Profiler{conn}
}
func (obj *Profiler) Enable() (err error) {
	err = obj.conn.Send("Profiler.enable", nil, nil)
	return
}
func (obj *Profiler) Disable() (err error) {
	err = obj.conn.Send("Profiler.disable", nil, nil)
	return
}

type SetSamplingIntervalRequest struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (obj *Profiler) SetSamplingInterval(request *SetSamplingIntervalRequest) (err error) {
	err = obj.conn.Send("Profiler.setSamplingInterval", request, nil)
	return
}
func (obj *Profiler) Start() (err error) {
	err = obj.conn.Send("Profiler.start", nil, nil)
	return
}

type StopResponse struct {
	// Recorded profile.
	Profile types.Profiler_Profile `json:"profile"`
}

func (obj *Profiler) Stop() (response StopResponse, err error) {
	err = obj.conn.Send("Profiler.stop", nil, &response)
	return
}

type StartPreciseCoverageRequest struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount *bool `json:"callCount,omitempty"`
	// Collect block-based coverage.
	Detailed *bool `json:"detailed,omitempty"`
}

// Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
func (obj *Profiler) StartPreciseCoverage(request *StartPreciseCoverageRequest) (err error) {
	err = obj.conn.Send("Profiler.startPreciseCoverage", request, nil)
	return
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (obj *Profiler) StopPreciseCoverage() (err error) {
	err = obj.conn.Send("Profiler.stopPreciseCoverage", nil, nil)
	return
}

type TakePreciseCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
func (obj *Profiler) TakePreciseCoverage() (response TakePreciseCoverageResponse, err error) {
	err = obj.conn.Send("Profiler.takePreciseCoverage", nil, &response)
	return
}

type GetBestEffortCoverageResponse struct {
	// Coverage data for the current isolate.
	Result []types.Profiler_ScriptCoverage `json:"result"`
}

// Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
func (obj *Profiler) GetBestEffortCoverage() (response GetBestEffortCoverageResponse, err error) {
	err = obj.conn.Send("Profiler.getBestEffortCoverage", nil, &response)
	return
}

// Enable type profile.
func (obj *Profiler) StartTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.startTypeProfile", nil, nil)
	return
}

// Disable type profile. Disabling releases type profile data collected so far.
func (obj *Profiler) StopTypeProfile() (err error) {
	err = obj.conn.Send("Profiler.stopTypeProfile", nil, nil)
	return
}

type TakeTypeProfileResponse struct {
	// Type profile for all scripts since startTypeProfile() was turned on.
	Result []types.Profiler_ScriptTypeProfile `json:"result"`
}

// Collect type profile.
func (obj *Profiler) TakeTypeProfile() (response TakeTypeProfileResponse, err error) {
	err = obj.conn.Send("Profiler.takeTypeProfile", nil, &response)
	return
}
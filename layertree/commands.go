/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package layertree provides commands and events for LayerTree domain.
package layertree

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in LayerTree domain
const (
	Enable             = "LayerTree.enable"
	Disable            = "LayerTree.disable"
	CompositingReasons = "LayerTree.compositingReasons"
	MakeSnapshot       = "LayerTree.makeSnapshot"
	LoadSnapshot       = "LayerTree.loadSnapshot"
	ReleaseSnapshot    = "LayerTree.releaseSnapshot"
	ProfileSnapshot    = "LayerTree.profileSnapshot"
	ReplaySnapshot     = "LayerTree.replaySnapshot"
	SnapshotCommandLog = "LayerTree.snapshotCommandLog"
)

// List of events in LayerTree domain
const (
	LayerTreeDidChange = "LayerTree.layerTreeDidChange"
	LayerPainted       = "LayerTree.layerPainted"
)

type LayerTree struct {
	conn cri.Connector
}

// New creates a LayerTree instance
func New(conn cri.Connector) *LayerTree {
	return &LayerTree{conn}
}

// Enables compositing tree inspection.
func (obj *LayerTree) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disables compositing tree inspection.
func (obj *LayerTree) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type CompositingReasonsRequest struct {
	// The id of the layer for which we want to get the reasons it was composited.
	LayerId types.LayerTree_LayerId `json:"layerId"`
}

type CompositingReasonsResponse struct {
	// A list of strings specifying reasons for the given layer to become composited.
	CompositingReasons []string `json:"compositingReasons"`
}

// Provides the reasons why the given layer was composited.
func (obj *LayerTree) CompositingReasons(request *CompositingReasonsRequest) (response CompositingReasonsResponse, err error) {
	err = obj.conn.Send(CompositingReasons, request, &response)
	return
}

type MakeSnapshotRequest struct {
	// The id of the layer.
	LayerId types.LayerTree_LayerId `json:"layerId"`
}

type MakeSnapshotResponse struct {
	// The id of the layer snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

// Returns the layer snapshot identifier.
func (obj *LayerTree) MakeSnapshot(request *MakeSnapshotRequest) (response MakeSnapshotResponse, err error) {
	err = obj.conn.Send(MakeSnapshot, request, &response)
	return
}

type LoadSnapshotRequest struct {
	// An array of tiles composing the snapshot.
	Tiles []types.LayerTree_PictureTile `json:"tiles"`
}

type LoadSnapshotResponse struct {
	// The id of the snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

// Returns the snapshot identifier.
func (obj *LayerTree) LoadSnapshot(request *LoadSnapshotRequest) (response LoadSnapshotResponse, err error) {
	err = obj.conn.Send(LoadSnapshot, request, &response)
	return
}

type ReleaseSnapshotRequest struct {
	// The id of the layer snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

// Releases layer snapshot captured by the back-end.
func (obj *LayerTree) ReleaseSnapshot(request *ReleaseSnapshotRequest) (err error) {
	err = obj.conn.Send(ReleaseSnapshot, request, nil)
	return
}

type ProfileSnapshotRequest struct {
	// The id of the layer snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
	// The maximum number of times to replay the snapshot (1, if not specified).
	MinRepeatCount *int `json:"minRepeatCount,omitempty"`
	// The minimum duration (in seconds) to replay the snapshot.
	MinDuration *float32 `json:"minDuration,omitempty"`
	// The clip rectangle to apply when replaying the snapshot.
	ClipRect *types.DOM_Rect `json:"clipRect,omitempty"`
}

type ProfileSnapshotResponse struct {
	// The array of paint profiles, one per run.
	Timings []types.LayerTree_PaintProfile `json:"timings"`
}

func (obj *LayerTree) ProfileSnapshot(request *ProfileSnapshotRequest) (response ProfileSnapshotResponse, err error) {
	err = obj.conn.Send(ProfileSnapshot, request, &response)
	return
}

type ReplaySnapshotRequest struct {
	// The id of the layer snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
	// The first step to replay from (replay from the very start if not specified).
	FromStep *int `json:"fromStep,omitempty"`
	// The last step to replay to (replay till the end if not specified).
	ToStep *int `json:"toStep,omitempty"`
	// The scale to apply while replaying (defaults to 1).
	Scale *float32 `json:"scale,omitempty"`
}

type ReplaySnapshotResponse struct {
	// A data: URL for resulting image.
	DataURL string `json:"dataURL"`
}

// Replays the layer snapshot and returns the resulting bitmap.
func (obj *LayerTree) ReplaySnapshot(request *ReplaySnapshotRequest) (response ReplaySnapshotResponse, err error) {
	err = obj.conn.Send(ReplaySnapshot, request, &response)
	return
}

type SnapshotCommandLogRequest struct {
	// The id of the layer snapshot.
	SnapshotId types.LayerTree_SnapshotId `json:"snapshotId"`
}

type SnapshotCommandLogResponse struct {
	// The array of canvas function calls.
	CommandLog []map[string]interface{} `json:"commandLog"`
}

// Replays the layer snapshot and returns canvas log.
func (obj *LayerTree) SnapshotCommandLog(request *SnapshotCommandLogRequest) (response SnapshotCommandLogResponse, err error) {
	err = obj.conn.Send(SnapshotCommandLog, request, &response)
	return
}

type LayerTreeDidChangeParams struct {
	// Layer tree, absent if not in the comspositing mode.
	Layers []types.LayerTree_Layer `json:"layers,omitempty"`
}

func (obj *LayerTree) LayerTreeDidChange(fn func(params *LayerTreeDidChangeParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(LayerTreeDidChange, closeChn)
	go func() {
		for {
			params := LayerTreeDidChangeParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type LayerPaintedParams struct {
	// The id of the painted layer.
	LayerId types.LayerTree_LayerId `json:"layerId"`
	// Clip rectangle.
	Clip types.DOM_Rect `json:"clip"`
}

func (obj *LayerTree) LayerPainted(fn func(params *LayerPaintedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(LayerPainted, closeChn)
	go func() {
		for {
			params := LayerPaintedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

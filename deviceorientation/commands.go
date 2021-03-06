/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package deviceorientation provides commands and events for DeviceOrientation domain.
package deviceorientation

import (
	"github.com/SKatiyar/cri"
)

// List of commands in DeviceOrientation domain
const (
	SetDeviceOrientationOverride   = "DeviceOrientation.setDeviceOrientationOverride"
	ClearDeviceOrientationOverride = "DeviceOrientation.clearDeviceOrientationOverride"
)

type DeviceOrientation struct {
	conn cri.Connector
}

// New creates a DeviceOrientation instance
func New(conn cri.Connector) *DeviceOrientation {
	return &DeviceOrientation{conn}
}

type SetDeviceOrientationOverrideRequest struct {
	// Mock alpha
	Alpha float32 `json:"alpha"`
	// Mock beta
	Beta float32 `json:"beta"`
	// Mock gamma
	Gamma float32 `json:"gamma"`
}

// Overrides the Device Orientation.
func (obj *DeviceOrientation) SetDeviceOrientationOverride(request *SetDeviceOrientationOverrideRequest) (err error) {
	err = obj.conn.Send(SetDeviceOrientationOverride, request, nil)
	return
}

// Clears the overridden Device Orientation.
func (obj *DeviceOrientation) ClearDeviceOrientationOverride() (err error) {
	err = obj.conn.Send(ClearDeviceOrientationOverride, nil, nil)
	return
}

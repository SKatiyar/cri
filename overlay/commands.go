/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package overlay provides commands and events for Overlay domain.
package overlay

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Overlay domain
const (
	Enable                       = "Overlay.enable"
	Disable                      = "Overlay.disable"
	SetShowPaintRects            = "Overlay.setShowPaintRects"
	SetShowDebugBorders          = "Overlay.setShowDebugBorders"
	SetShowFPSCounter            = "Overlay.setShowFPSCounter"
	SetShowScrollBottleneckRects = "Overlay.setShowScrollBottleneckRects"
	SetShowViewportSizeOnResize  = "Overlay.setShowViewportSizeOnResize"
	SetPausedInDebuggerMessage   = "Overlay.setPausedInDebuggerMessage"
	SetSuspended                 = "Overlay.setSuspended"
	SetInspectMode               = "Overlay.setInspectMode"
	HighlightRect                = "Overlay.highlightRect"
	HighlightQuad                = "Overlay.highlightQuad"
	HighlightNode                = "Overlay.highlightNode"
	HighlightFrame               = "Overlay.highlightFrame"
	HideHighlight                = "Overlay.hideHighlight"
	GetHighlightObjectForTest    = "Overlay.getHighlightObjectForTest"
)

// List of events in Overlay domain
const (
	NodeHighlightRequested = "Overlay.nodeHighlightRequested"
	InspectNodeRequested   = "Overlay.inspectNodeRequested"
	ScreenshotRequested    = "Overlay.screenshotRequested"
)

// This domain provides various functionality related to drawing atop the inspected page.
type Overlay struct {
	conn cri.Connector
}

// New creates a Overlay instance
func New(conn cri.Connector) *Overlay {
	return &Overlay{conn}
}

// Enables domain notifications.
func (obj *Overlay) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disables domain notifications.
func (obj *Overlay) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type SetShowPaintRectsRequest struct {
	// True for showing paint rectangles
	Result bool `json:"result"`
}

// Requests that backend shows paint rectangles
func (obj *Overlay) SetShowPaintRects(request *SetShowPaintRectsRequest) (err error) {
	err = obj.conn.Send(SetShowPaintRects, request, nil)
	return
}

type SetShowDebugBordersRequest struct {
	// True for showing debug borders
	Show bool `json:"show"`
}

// Requests that backend shows debug borders on layers
func (obj *Overlay) SetShowDebugBorders(request *SetShowDebugBordersRequest) (err error) {
	err = obj.conn.Send(SetShowDebugBorders, request, nil)
	return
}

type SetShowFPSCounterRequest struct {
	// True for showing the FPS counter
	Show bool `json:"show"`
}

// Requests that backend shows the FPS counter
func (obj *Overlay) SetShowFPSCounter(request *SetShowFPSCounterRequest) (err error) {
	err = obj.conn.Send(SetShowFPSCounter, request, nil)
	return
}

type SetShowScrollBottleneckRectsRequest struct {
	// True for showing scroll bottleneck rects
	Show bool `json:"show"`
}

// Requests that backend shows scroll bottleneck rects
func (obj *Overlay) SetShowScrollBottleneckRects(request *SetShowScrollBottleneckRectsRequest) (err error) {
	err = obj.conn.Send(SetShowScrollBottleneckRects, request, nil)
	return
}

type SetShowViewportSizeOnResizeRequest struct {
	// Whether to paint size or not.
	Show bool `json:"show"`
}

// Paints viewport size upon main frame resize.
func (obj *Overlay) SetShowViewportSizeOnResize(request *SetShowViewportSizeOnResizeRequest) (err error) {
	err = obj.conn.Send(SetShowViewportSizeOnResize, request, nil)
	return
}

type SetPausedInDebuggerMessageRequest struct {
	// The message to display, also triggers resume and step over controls.
	Message *string `json:"message,omitempty"`
}

func (obj *Overlay) SetPausedInDebuggerMessage(request *SetPausedInDebuggerMessageRequest) (err error) {
	err = obj.conn.Send(SetPausedInDebuggerMessage, request, nil)
	return
}

type SetSuspendedRequest struct {
	// Whether overlay should be suspended and not consume any resources until resumed.
	Suspended bool `json:"suspended"`
}

func (obj *Overlay) SetSuspended(request *SetSuspendedRequest) (err error) {
	err = obj.conn.Send(SetSuspended, request, nil)
	return
}

type SetInspectModeRequest struct {
	// Set an inspection mode.
	Mode types.Overlay_InspectMode `json:"mode"`
	// A descriptor for the highlight appearance of hovered-over nodes. May be omitted if <code>enabled == false</code>.
	HighlightConfig *types.Overlay_HighlightConfig `json:"highlightConfig,omitempty"`
}

// Enters the 'inspect' mode. In this mode, elements that user is hovering over are highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.
func (obj *Overlay) SetInspectMode(request *SetInspectModeRequest) (err error) {
	err = obj.conn.Send(SetInspectMode, request, nil)
	return
}

type HighlightRectRequest struct {
	// X coordinate
	X int `json:"x"`
	// Y coordinate
	Y int `json:"y"`
	// Rectangle width
	Width int `json:"width"`
	// Rectangle height
	Height int `json:"height"`
	// The highlight fill color (default: transparent).
	Color *types.DOM_RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

// Highlights given rectangle. Coordinates are absolute with respect to the main frame viewport.
func (obj *Overlay) HighlightRect(request *HighlightRectRequest) (err error) {
	err = obj.conn.Send(HighlightRect, request, nil)
	return
}

type HighlightQuadRequest struct {
	// Quad to highlight
	Quad types.DOM_Quad `json:"quad"`
	// The highlight fill color (default: transparent).
	Color *types.DOM_RGBA `json:"color,omitempty"`
	// The highlight outline color (default: transparent).
	OutlineColor *types.DOM_RGBA `json:"outlineColor,omitempty"`
}

// Highlights given quad. Coordinates are absolute with respect to the main frame viewport.
func (obj *Overlay) HighlightQuad(request *HighlightQuadRequest) (err error) {
	err = obj.conn.Send(HighlightQuad, request, nil)
	return
}

type HighlightNodeRequest struct {
	// A descriptor for the highlight appearance.
	HighlightConfig types.Overlay_HighlightConfig `json:"highlightConfig"`
	// Identifier of the node to highlight.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node to highlight.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node to be highlighted.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

// Highlights DOM node with given id or with the given JavaScript object wrapper. Either nodeId or objectId must be specified.
func (obj *Overlay) HighlightNode(request *HighlightNodeRequest) (err error) {
	err = obj.conn.Send(HighlightNode, request, nil)
	return
}

type HighlightFrameRequest struct {
	// Identifier of the frame to highlight.
	FrameId types.Page_FrameId `json:"frameId"`
	// The content box highlight fill color (default: transparent).
	ContentColor *types.DOM_RGBA `json:"contentColor,omitempty"`
	// The content box highlight outline color (default: transparent).
	ContentOutlineColor *types.DOM_RGBA `json:"contentOutlineColor,omitempty"`
}

// Highlights owner element of the frame with given id.
func (obj *Overlay) HighlightFrame(request *HighlightFrameRequest) (err error) {
	err = obj.conn.Send(HighlightFrame, request, nil)
	return
}

// Hides any highlight.
func (obj *Overlay) HideHighlight() (err error) {
	err = obj.conn.Send(HideHighlight, nil, nil)
	return
}

type GetHighlightObjectForTestRequest struct {
	// Id of the node to get highlight object for.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetHighlightObjectForTestResponse struct {
	// Highlight data for the node.
	Highlight map[string]interface{} `json:"highlight"`
}

// For testing.
func (obj *Overlay) GetHighlightObjectForTest(request *GetHighlightObjectForTestRequest) (response GetHighlightObjectForTestResponse, err error) {
	err = obj.conn.Send(GetHighlightObjectForTest, request, &response)
	return
}

type NodeHighlightRequestedParams struct {
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Fired when the node should be highlighted. This happens after call to <code>setInspectMode</code>.
func (obj *Overlay) NodeHighlightRequested(fn func(params *NodeHighlightRequestedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(NodeHighlightRequested, closeChn)
	go func() {
		for {
			params := NodeHighlightRequestedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type InspectNodeRequestedParams struct {
	// Id of the node to inspect.
	BackendNodeId types.DOM_BackendNodeId `json:"backendNodeId"`
}

// Fired when the node should be inspected. This happens after call to <code>setInspectMode</code> or when user manually inspects an element.
func (obj *Overlay) InspectNodeRequested(fn func(params *InspectNodeRequestedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(InspectNodeRequested, closeChn)
	go func() {
		for {
			params := InspectNodeRequestedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ScreenshotRequestedParams struct {
	// Viewport to capture, in CSS.
	Viewport types.Page_Viewport `json:"viewport"`
}

// Fired when user asks to capture screenshot of some area on the page.
func (obj *Overlay) ScreenshotRequested(fn func(params *ScreenshotRequestedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ScreenshotRequested, closeChn)
	go func() {
		for {
			params := ScreenshotRequestedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

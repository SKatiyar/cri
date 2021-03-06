/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package domsnapshot provides commands and events for DOMSnapshot domain.
package domsnapshot

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in DOMSnapshot domain
const (
	GetSnapshot = "DOMSnapshot.getSnapshot"
)

// This domain facilitates obtaining document snapshots with DOM, layout, and style information.
type DOMSnapshot struct {
	conn cri.Connector
}

// New creates a DOMSnapshot instance
func New(conn cri.Connector) *DOMSnapshot {
	return &DOMSnapshot{conn}
}

type GetSnapshotRequest struct {
	// Whitelist of computed styles to return.
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`
}

type GetSnapshotResponse struct {
	// The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	DomNodes []types.DOMSnapshot_DOMNode `json:"domNodes"`
	// The nodes in the layout tree.
	LayoutTreeNodes []types.DOMSnapshot_LayoutTreeNode `json:"layoutTreeNodes"`
	// Whitelisted ComputedStyle properties for each node in the layout tree.
	ComputedStyles []types.DOMSnapshot_ComputedStyle `json:"computedStyles"`
}

// Returns a document snapshot, including the full DOM tree of the root node (including iframes, template contents, and imported documents) in a flattened array, as well as layout and white-listed computed style information for the nodes. Shadow DOM in the returned DOM tree is flattened.
func (obj *DOMSnapshot) GetSnapshot(request *GetSnapshotRequest) (response GetSnapshotResponse, err error) {
	err = obj.conn.Send(GetSnapshot, request, &response)
	return
}

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/skatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package schema provides commands and events for Schema domain.
package schema

import (
	"github.com/skatiyar/cri"
	types "github.com/skatiyar/cri/types"
)

// List of commands in Schema domain
const (
	GetDomains = "Schema.getDomains"
)

// This domain is deprecated.
type Schema struct {
	conn cri.Connector
}

// New creates a Schema instance
func New(conn cri.Connector) *Schema {
	return &Schema{conn}
}

type GetDomainsResponse struct {
	// List of supported domains.
	Domains []types.Schema_Domain `json:"domains"`
}

// Returns supported domains.
func (obj *Schema) GetDomains() (response GetDomainsResponse, err error) {
	err = obj.conn.Send(GetDomains, nil, &response)
	return
}

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package storage provides commands and events for Storage domain.
package storage

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in Storage domain
const (
	ClearDataForOrigin           = "Storage.clearDataForOrigin"
	GetUsageAndQuota             = "Storage.getUsageAndQuota"
	TrackCacheStorageForOrigin   = "Storage.trackCacheStorageForOrigin"
	TrackIndexedDBForOrigin      = "Storage.trackIndexedDBForOrigin"
	UntrackCacheStorageForOrigin = "Storage.untrackCacheStorageForOrigin"
	UntrackIndexedDBForOrigin    = "Storage.untrackIndexedDBForOrigin"
)

// List of events in Storage domain
const (
	CacheStorageContentUpdated = "Storage.cacheStorageContentUpdated"
	CacheStorageListUpdated    = "Storage.cacheStorageListUpdated"
	IndexedDBContentUpdated    = "Storage.indexedDBContentUpdated"
	IndexedDBListUpdated       = "Storage.indexedDBListUpdated"
)

type Storage struct {
	conn cri.Connector
}

// New creates a Storage instance
func New(conn cri.Connector) *Storage {
	return &Storage{conn}
}

type ClearDataForOriginRequest struct {
	// Security origin.
	Origin string `json:"origin"`
	// Comma separated origin names.
	StorageTypes string `json:"storageTypes"`
}

// Clears storage for origin.
func (obj *Storage) ClearDataForOrigin(request *ClearDataForOriginRequest) (err error) {
	err = obj.conn.Send(ClearDataForOrigin, request, nil)
	return
}

type GetUsageAndQuotaRequest struct {
	// Security origin.
	Origin string `json:"origin"`
}

type GetUsageAndQuotaResponse struct {
	// Storage usage (bytes).
	Usage float32 `json:"usage"`
	// Storage quota (bytes).
	Quota float32 `json:"quota"`
	// Storage usage per type (bytes).
	UsageBreakdown []types.Storage_UsageForType `json:"usageBreakdown"`
}

// Returns usage and quota in bytes.
func (obj *Storage) GetUsageAndQuota(request *GetUsageAndQuotaRequest) (response GetUsageAndQuotaResponse, err error) {
	err = obj.conn.Send(GetUsageAndQuota, request, &response)
	return
}

type TrackCacheStorageForOriginRequest struct {
	// Security origin.
	Origin string `json:"origin"`
}

// Registers origin to be notified when an update occurs to its cache storage list.
func (obj *Storage) TrackCacheStorageForOrigin(request *TrackCacheStorageForOriginRequest) (err error) {
	err = obj.conn.Send(TrackCacheStorageForOrigin, request, nil)
	return
}

type TrackIndexedDBForOriginRequest struct {
	// Security origin.
	Origin string `json:"origin"`
}

// Registers origin to be notified when an update occurs to its IndexedDB.
func (obj *Storage) TrackIndexedDBForOrigin(request *TrackIndexedDBForOriginRequest) (err error) {
	err = obj.conn.Send(TrackIndexedDBForOrigin, request, nil)
	return
}

type UntrackCacheStorageForOriginRequest struct {
	// Security origin.
	Origin string `json:"origin"`
}

// Unregisters origin from receiving notifications for cache storage.
func (obj *Storage) UntrackCacheStorageForOrigin(request *UntrackCacheStorageForOriginRequest) (err error) {
	err = obj.conn.Send(UntrackCacheStorageForOrigin, request, nil)
	return
}

type UntrackIndexedDBForOriginRequest struct {
	// Security origin.
	Origin string `json:"origin"`
}

// Unregisters origin from receiving notifications for IndexedDB.
func (obj *Storage) UntrackIndexedDBForOrigin(request *UntrackIndexedDBForOriginRequest) (err error) {
	err = obj.conn.Send(UntrackIndexedDBForOrigin, request, nil)
	return
}

type CacheStorageContentUpdatedParams struct {
	// Origin to update.
	Origin string `json:"origin"`
	// Name of cache in origin.
	CacheName string `json:"cacheName"`
}

// A cache's contents have been modified.
func (obj *Storage) CacheStorageContentUpdated(fn func(event string, params CacheStorageContentUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(CacheStorageContentUpdated)
	go func() {
		defer closer()
		for {
			var params CacheStorageContentUpdatedParams
			if !fn(CacheStorageContentUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type CacheStorageListUpdatedParams struct {
	// Origin to update.
	Origin string `json:"origin"`
}

// A cache has been added/deleted.
func (obj *Storage) CacheStorageListUpdated(fn func(event string, params CacheStorageListUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(CacheStorageListUpdated)
	go func() {
		defer closer()
		for {
			var params CacheStorageListUpdatedParams
			if !fn(CacheStorageListUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type IndexedDBContentUpdatedParams struct {
	// Origin to update.
	Origin string `json:"origin"`
	// Database to update.
	DatabaseName string `json:"databaseName"`
	// ObjectStore to update.
	ObjectStoreName string `json:"objectStoreName"`
}

// The origin's IndexedDB object store has been modified.
func (obj *Storage) IndexedDBContentUpdated(fn func(event string, params IndexedDBContentUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(IndexedDBContentUpdated)
	go func() {
		defer closer()
		for {
			var params IndexedDBContentUpdatedParams
			if !fn(IndexedDBContentUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

type IndexedDBListUpdatedParams struct {
	// Origin to update.
	Origin string `json:"origin"`
}

// The origin's IndexedDB database list has been modified.
func (obj *Storage) IndexedDBListUpdated(fn func(event string, params IndexedDBListUpdatedParams, err error) bool) {
	listen, closer := obj.conn.On(IndexedDBListUpdated)
	go func() {
		defer closer()
		for {
			var params IndexedDBListUpdatedParams
			if !fn(IndexedDBListUpdated, params, listen(&params)) {
				return
			}
		}
	}()
}

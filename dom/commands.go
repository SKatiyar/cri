/*
* CODE GENERATED AUTOMATICALLY WITH github.com/SKatiyar/cri/cmd/cri-gen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

// Package dom provides commands and events for DOM domain.
package dom

import (
	"github.com/SKatiyar/cri"
	types "github.com/SKatiyar/cri/types"
)

// List of commands in DOM domain
const (
	Enable                          = "DOM.enable"
	Disable                         = "DOM.disable"
	GetDocument                     = "DOM.getDocument"
	GetFlattenedDocument            = "DOM.getFlattenedDocument"
	CollectClassNamesFromSubtree    = "DOM.collectClassNamesFromSubtree"
	RequestChildNodes               = "DOM.requestChildNodes"
	QuerySelector                   = "DOM.querySelector"
	QuerySelectorAll                = "DOM.querySelectorAll"
	SetNodeName                     = "DOM.setNodeName"
	SetNodeValue                    = "DOM.setNodeValue"
	RemoveNode                      = "DOM.removeNode"
	SetAttributeValue               = "DOM.setAttributeValue"
	SetAttributesAsText             = "DOM.setAttributesAsText"
	RemoveAttribute                 = "DOM.removeAttribute"
	GetOuterHTML                    = "DOM.getOuterHTML"
	SetOuterHTML                    = "DOM.setOuterHTML"
	PerformSearch                   = "DOM.performSearch"
	GetSearchResults                = "DOM.getSearchResults"
	DiscardSearchResults            = "DOM.discardSearchResults"
	RequestNode                     = "DOM.requestNode"
	HighlightRect                   = "DOM.highlightRect"
	HighlightNode                   = "DOM.highlightNode"
	HideHighlight                   = "DOM.hideHighlight"
	PushNodeByPathToFrontend        = "DOM.pushNodeByPathToFrontend"
	PushNodesByBackendIdsToFrontend = "DOM.pushNodesByBackendIdsToFrontend"
	SetInspectedNode                = "DOM.setInspectedNode"
	ResolveNode                     = "DOM.resolveNode"
	GetAttributes                   = "DOM.getAttributes"
	CopyTo                          = "DOM.copyTo"
	MoveTo                          = "DOM.moveTo"
	Undo                            = "DOM.undo"
	Redo                            = "DOM.redo"
	MarkUndoableState               = "DOM.markUndoableState"
	Focus                           = "DOM.focus"
	SetFileInputFiles               = "DOM.setFileInputFiles"
	GetBoxModel                     = "DOM.getBoxModel"
	GetNodeForLocation              = "DOM.getNodeForLocation"
	GetRelayoutBoundary             = "DOM.getRelayoutBoundary"
	DescribeNode                    = "DOM.describeNode"
)

// List of events in DOM domain
const (
	DocumentUpdated         = "DOM.documentUpdated"
	SetChildNodes           = "DOM.setChildNodes"
	AttributeModified       = "DOM.attributeModified"
	AttributeRemoved        = "DOM.attributeRemoved"
	InlineStyleInvalidated  = "DOM.inlineStyleInvalidated"
	CharacterDataModified   = "DOM.characterDataModified"
	ChildNodeCountUpdated   = "DOM.childNodeCountUpdated"
	ChildNodeInserted       = "DOM.childNodeInserted"
	ChildNodeRemoved        = "DOM.childNodeRemoved"
	ShadowRootPushed        = "DOM.shadowRootPushed"
	ShadowRootPopped        = "DOM.shadowRootPopped"
	PseudoElementAdded      = "DOM.pseudoElementAdded"
	PseudoElementRemoved    = "DOM.pseudoElementRemoved"
	DistributedNodesUpdated = "DOM.distributedNodesUpdated"
)

// This domain exposes DOM read/write operations. Each DOM Node is represented with its mirror object that has an <code>id</code>. This <code>id</code> can be used to get additional information on the Node, resolve it into the JavaScript object wrapper, etc. It is important that client receives DOM events only for the nodes that are known to the client. Backend keeps track of the nodes that were sent to the client and never sends the same node twice. It is client's responsibility to collect information about the nodes that were sent to the client.<p>Note that <code>iframe</code> owner elements will return corresponding document elements as their child nodes.</p>
type DOM struct {
	conn cri.Connector
}

// New creates a DOM instance
func New(conn cri.Connector) *DOM {
	return &DOM{conn}
}

// Enables DOM agent for the given page.
func (obj *DOM) Enable() (err error) {
	err = obj.conn.Send(Enable, nil, nil)
	return
}

// Disables DOM agent for the given page.
func (obj *DOM) Disable() (err error) {
	err = obj.conn.Send(Disable, nil, nil)
	return
}

type GetDocumentRequest struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth *int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce *bool `json:"pierce,omitempty"`
}

type GetDocumentResponse struct {
	// Resulting node.
	Root types.DOM_Node `json:"root"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
func (obj *DOM) GetDocument(request *GetDocumentRequest) (response GetDocumentResponse, err error) {
	err = obj.conn.Send(GetDocument, request, &response)
	return
}

type GetFlattenedDocumentRequest struct {
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth *int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce *bool `json:"pierce,omitempty"`
}

type GetFlattenedDocumentResponse struct {
	// Resulting node.
	Nodes []types.DOM_Node `json:"nodes"`
}

// Returns the root DOM node (and optionally the subtree) to the caller.
func (obj *DOM) GetFlattenedDocument(request *GetFlattenedDocumentRequest) (response GetFlattenedDocumentResponse, err error) {
	err = obj.conn.Send(GetFlattenedDocument, request, &response)
	return
}

type CollectClassNamesFromSubtreeRequest struct {
	// Id of the node to collect class names.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type CollectClassNamesFromSubtreeResponse struct {
	// Class name list.
	ClassNames []string `json:"classNames"`
}

// Collects class names for the node with given id and all of it's child nodes.
func (obj *DOM) CollectClassNamesFromSubtree(request *CollectClassNamesFromSubtreeRequest) (response CollectClassNamesFromSubtreeResponse, err error) {
	err = obj.conn.Send(CollectClassNamesFromSubtree, request, &response)
	return
}

type RequestChildNodesRequest struct {
	// Id of the node to get children for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth *int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the sub-tree (default is false).
	Pierce *bool `json:"pierce,omitempty"`
}

// Requests that children of the node with given id are returned to the caller in form of <code>setChildNodes</code> events where not only immediate children are retrieved, but all children down to the specified depth.
func (obj *DOM) RequestChildNodes(request *RequestChildNodesRequest) (err error) {
	err = obj.conn.Send(RequestChildNodes, request, nil)
	return
}

type QuerySelectorRequest struct {
	// Id of the node to query upon.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

type QuerySelectorResponse struct {
	// Query selector result.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Executes <code>querySelector</code> on a given node.
func (obj *DOM) QuerySelector(request *QuerySelectorRequest) (response QuerySelectorResponse, err error) {
	err = obj.conn.Send(QuerySelector, request, &response)
	return
}

type QuerySelectorAllRequest struct {
	// Id of the node to query upon.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Selector string.
	Selector string `json:"selector"`
}

type QuerySelectorAllResponse struct {
	// Query selector result.
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}

// Executes <code>querySelectorAll</code> on a given node.
func (obj *DOM) QuerySelectorAll(request *QuerySelectorAllRequest) (response QuerySelectorAllResponse, err error) {
	err = obj.conn.Send(QuerySelectorAll, request, &response)
	return
}

type SetNodeNameRequest struct {
	// Id of the node to set name for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// New node's name.
	Name string `json:"name"`
}

type SetNodeNameResponse struct {
	// New node's id.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Sets node name for a node with given id.
func (obj *DOM) SetNodeName(request *SetNodeNameRequest) (response SetNodeNameResponse, err error) {
	err = obj.conn.Send(SetNodeName, request, &response)
	return
}

type SetNodeValueRequest struct {
	// Id of the node to set value for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// New node's value.
	Value string `json:"value"`
}

// Sets node value for a node with given id.
func (obj *DOM) SetNodeValue(request *SetNodeValueRequest) (err error) {
	err = obj.conn.Send(SetNodeValue, request, nil)
	return
}

type RemoveNodeRequest struct {
	// Id of the node to remove.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Removes node with given id.
func (obj *DOM) RemoveNode(request *RemoveNodeRequest) (err error) {
	err = obj.conn.Send(RemoveNode, request, nil)
	return
}

type SetAttributeValueRequest struct {
	// Id of the element to set attribute for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Attribute name.
	Name string `json:"name"`
	// Attribute value.
	Value string `json:"value"`
}

// Sets attribute for an element with given id.
func (obj *DOM) SetAttributeValue(request *SetAttributeValueRequest) (err error) {
	err = obj.conn.Send(SetAttributeValue, request, nil)
	return
}

type SetAttributesAsTextRequest struct {
	// Id of the element to set attributes for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Text with a number of attributes. Will parse this text using HTML parser.
	Text string `json:"text"`
	// Attribute name to replace with new attributes derived from text in case text parsed successfully.
	Name *string `json:"name,omitempty"`
}

// Sets attributes on element with given id. This method is useful when user edits some existing attribute value and types in several attribute name/value pairs.
func (obj *DOM) SetAttributesAsText(request *SetAttributesAsTextRequest) (err error) {
	err = obj.conn.Send(SetAttributesAsText, request, nil)
	return
}

type RemoveAttributeRequest struct {
	// Id of the element to remove attribute from.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Name of the attribute to remove.
	Name string `json:"name"`
}

// Removes attribute with given name from an element with given id.
func (obj *DOM) RemoveAttribute(request *RemoveAttributeRequest) (err error) {
	err = obj.conn.Send(RemoveAttribute, request, nil)
	return
}

type GetOuterHTMLRequest struct {
	// Identifier of the node.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

type GetOuterHTMLResponse struct {
	// Outer HTML markup.
	OuterHTML string `json:"outerHTML"`
}

// Returns node's HTML markup.
func (obj *DOM) GetOuterHTML(request *GetOuterHTMLRequest) (response GetOuterHTMLResponse, err error) {
	err = obj.conn.Send(GetOuterHTML, request, &response)
	return
}

type SetOuterHTMLRequest struct {
	// Id of the node to set markup for.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Outer HTML markup to set.
	OuterHTML string `json:"outerHTML"`
}

// Sets node HTML markup, returns new node id.
func (obj *DOM) SetOuterHTML(request *SetOuterHTMLRequest) (err error) {
	err = obj.conn.Send(SetOuterHTML, request, nil)
	return
}

type PerformSearchRequest struct {
	// Plain text or query selector or XPath search query.
	Query string `json:"query"`
	// True to search in user agent shadow DOM.
	IncludeUserAgentShadowDOM *bool `json:"includeUserAgentShadowDOM,omitempty"`
}

type PerformSearchResponse struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
	// Number of search results.
	ResultCount int `json:"resultCount"`
}

// Searches for a given string in the DOM tree. Use <code>getSearchResults</code> to access search results or <code>cancelSearch</code> to end this search session.
func (obj *DOM) PerformSearch(request *PerformSearchRequest) (response PerformSearchResponse, err error) {
	err = obj.conn.Send(PerformSearch, request, &response)
	return
}

type GetSearchResultsRequest struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
	// Start index of the search result to be returned.
	FromIndex int `json:"fromIndex"`
	// End index of the search result to be returned.
	ToIndex int `json:"toIndex"`
}

type GetSearchResultsResponse struct {
	// Ids of the search result nodes.
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}

// Returns search results from given <code>fromIndex</code> to given <code>toIndex</code> from the search with the given identifier.
func (obj *DOM) GetSearchResults(request *GetSearchResultsRequest) (response GetSearchResultsResponse, err error) {
	err = obj.conn.Send(GetSearchResults, request, &response)
	return
}

type DiscardSearchResultsRequest struct {
	// Unique search session identifier.
	SearchId string `json:"searchId"`
}

// Discards search results from the session with the given id. <code>getSearchResults</code> should no longer be called for that search.
func (obj *DOM) DiscardSearchResults(request *DiscardSearchResultsRequest) (err error) {
	err = obj.conn.Send(DiscardSearchResults, request, nil)
	return
}

type RequestNodeRequest struct {
	// JavaScript object id to convert into node.
	ObjectId types.Runtime_RemoteObjectId `json:"objectId"`
}

type RequestNodeResponse struct {
	// Node id for given object.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Requests that the node is sent to the caller given the JavaScript node object reference. All nodes that form the path from the node to the root are also sent to the client as a series of <code>setChildNodes</code> notifications.
func (obj *DOM) RequestNode(request *RequestNodeRequest) (response RequestNodeResponse, err error) {
	err = obj.conn.Send(RequestNode, request, &response)
	return
}

// Highlights given rectangle.
func (obj *DOM) HighlightRect() (err error) {
	err = obj.conn.Send(HighlightRect, nil, nil)
	return
}

// Highlights DOM node.
func (obj *DOM) HighlightNode() (err error) {
	err = obj.conn.Send(HighlightNode, nil, nil)
	return
}

// Hides any highlight.
func (obj *DOM) HideHighlight() (err error) {
	err = obj.conn.Send(HideHighlight, nil, nil)
	return
}

type PushNodeByPathToFrontendRequest struct {
	// Path to node in the proprietary format.
	Path string `json:"path"`
}

type PushNodeByPathToFrontendResponse struct {
	// Id of the node for given path.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Requests that the node is sent to the caller given its path. // FIXME, use XPath
func (obj *DOM) PushNodeByPathToFrontend(request *PushNodeByPathToFrontendRequest) (response PushNodeByPathToFrontendResponse, err error) {
	err = obj.conn.Send(PushNodeByPathToFrontend, request, &response)
	return
}

type PushNodesByBackendIdsToFrontendRequest struct {
	// The array of backend node ids.
	BackendNodeIds []types.DOM_BackendNodeId `json:"backendNodeIds"`
}

type PushNodesByBackendIdsToFrontendResponse struct {
	// The array of ids of pushed nodes that correspond to the backend ids specified in backendNodeIds.
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}

// Requests that a batch of nodes is sent to the caller given their backend node ids.
func (obj *DOM) PushNodesByBackendIdsToFrontend(request *PushNodesByBackendIdsToFrontendRequest) (response PushNodesByBackendIdsToFrontendResponse, err error) {
	err = obj.conn.Send(PushNodesByBackendIdsToFrontend, request, &response)
	return
}

type SetInspectedNodeRequest struct {
	// DOM node id to be accessible by means of $x command line API.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (obj *DOM) SetInspectedNode(request *SetInspectedNodeRequest) (err error) {
	err = obj.conn.Send(SetInspectedNode, request, nil)
	return
}

type ResolveNodeRequest struct {
	// Id of the node to resolve.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Backend identifier of the node to resolve.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup *string `json:"objectGroup,omitempty"`
}

type ResolveNodeResponse struct {
	// JavaScript object wrapper for given node.
	Object types.Runtime_RemoteObject `json:"object"`
}

// Resolves the JavaScript node object for a given NodeId or BackendNodeId.
func (obj *DOM) ResolveNode(request *ResolveNodeRequest) (response ResolveNodeResponse, err error) {
	err = obj.conn.Send(ResolveNode, request, &response)
	return
}

type GetAttributesRequest struct {
	// Id of the node to retrieve attibutes for.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetAttributesResponse struct {
	// An interleaved array of node attribute names and values.
	Attributes []string `json:"attributes"`
}

// Returns attributes for the specified node.
func (obj *DOM) GetAttributes(request *GetAttributesRequest) (response GetAttributesResponse, err error) {
	err = obj.conn.Send(GetAttributes, request, &response)
	return
}

type CopyToRequest struct {
	// Id of the node to copy.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Id of the element to drop the copy into.
	TargetNodeId types.DOM_NodeId `json:"targetNodeId"`
	// Drop the copy before this node (if absent, the copy becomes the last child of <code>targetNodeId</code>).
	InsertBeforeNodeId *types.DOM_NodeId `json:"insertBeforeNodeId,omitempty"`
}

type CopyToResponse struct {
	// Id of the node clone.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Creates a deep copy of the specified node and places it into the target container before the given anchor.
func (obj *DOM) CopyTo(request *CopyToRequest) (response CopyToResponse, err error) {
	err = obj.conn.Send(CopyTo, request, &response)
	return
}

type MoveToRequest struct {
	// Id of the node to move.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Id of the element to drop the moved node into.
	TargetNodeId types.DOM_NodeId `json:"targetNodeId"`
	// Drop node before this one (if absent, the moved node becomes the last child of <code>targetNodeId</code>).
	InsertBeforeNodeId *types.DOM_NodeId `json:"insertBeforeNodeId,omitempty"`
}

type MoveToResponse struct {
	// New id of the moved node.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Moves node into the new container, places it before the given anchor.
func (obj *DOM) MoveTo(request *MoveToRequest) (response MoveToResponse, err error) {
	err = obj.conn.Send(MoveTo, request, &response)
	return
}

// Undoes the last performed action.
func (obj *DOM) Undo() (err error) {
	err = obj.conn.Send(Undo, nil, nil)
	return
}

// Re-does the last undone action.
func (obj *DOM) Redo() (err error) {
	err = obj.conn.Send(Redo, nil, nil)
	return
}

// Marks last undoable state.
func (obj *DOM) MarkUndoableState() (err error) {
	err = obj.conn.Send(MarkUndoableState, nil, nil)
	return
}

type FocusRequest struct {
	// Identifier of the node.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

// Focuses the given element.
func (obj *DOM) Focus(request *FocusRequest) (err error) {
	err = obj.conn.Send(Focus, request, nil)
	return
}

type SetFileInputFilesRequest struct {
	// Array of file paths to set.
	Files []string `json:"files"`
	// Identifier of the node.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

// Sets files for the given file input element.
func (obj *DOM) SetFileInputFiles(request *SetFileInputFilesRequest) (err error) {
	err = obj.conn.Send(SetFileInputFiles, request, nil)
	return
}

type GetBoxModelRequest struct {
	// Identifier of the node.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
}

type GetBoxModelResponse struct {
	// Box model for the node.
	Model types.DOM_BoxModel `json:"model"`
}

// Returns boxes for the given node.
func (obj *DOM) GetBoxModel(request *GetBoxModelRequest) (response GetBoxModelResponse, err error) {
	err = obj.conn.Send(GetBoxModel, request, &response)
	return
}

type GetNodeForLocationRequest struct {
	// X coordinate.
	X int `json:"x"`
	// Y coordinate.
	Y int `json:"y"`
	// False to skip to the nearest non-UA shadow root ancestor (default: false).
	IncludeUserAgentShadowDOM *bool `json:"includeUserAgentShadowDOM,omitempty"`
}

type GetNodeForLocationResponse struct {
	// Id of the node at given coordinates.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Returns node id at given location.
func (obj *DOM) GetNodeForLocation(request *GetNodeForLocationRequest) (response GetNodeForLocationResponse, err error) {
	err = obj.conn.Send(GetNodeForLocation, request, &response)
	return
}

type GetRelayoutBoundaryRequest struct {
	// Id of the node.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

type GetRelayoutBoundaryResponse struct {
	// Relayout boundary node id for the given node.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Returns the id of the nearest ancestor that is a relayout boundary.
func (obj *DOM) GetRelayoutBoundary(request *GetRelayoutBoundaryRequest) (response GetRelayoutBoundaryResponse, err error) {
	err = obj.conn.Send(GetRelayoutBoundary, request, &response)
	return
}

type DescribeNodeRequest struct {
	// Identifier of the node.
	NodeId *types.DOM_NodeId `json:"nodeId,omitempty"`
	// Identifier of the backend node.
	BackendNodeId *types.DOM_BackendNodeId `json:"backendNodeId,omitempty"`
	// JavaScript object id of the node wrapper.
	ObjectId *types.Runtime_RemoteObjectId `json:"objectId,omitempty"`
	// The maximum depth at which children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth *int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false).
	Pierce *bool `json:"pierce,omitempty"`
}

type DescribeNodeResponse struct {
	// Node description.
	Node types.DOM_Node `json:"node"`
}

// Describes node given its id, does not require domain to be enabled. Does not start tracking any objects, can be used for automation.
func (obj *DOM) DescribeNode(request *DescribeNodeRequest) (response DescribeNodeResponse, err error) {
	err = obj.conn.Send(DescribeNode, request, &response)
	return
}

// Fired when <code>Document</code> has been totally updated. Node ids are no longer valid.
func (obj *DOM) DocumentUpdated(fn func(err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DocumentUpdated, closeChn)
	go func() {
		for {

			readErr := decoder(nil)
			if !fn(readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type SetChildNodesParams struct {
	// Parent node id to populate with children.
	ParentId types.DOM_NodeId `json:"parentId"`
	// Child nodes array.
	Nodes []types.DOM_Node `json:"nodes"`
}

// Fired when backend wants to provide client with the missing DOM structure. This happens upon most of the calls requesting node ids.
func (obj *DOM) SetChildNodes(fn func(params *SetChildNodesParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(SetChildNodes, closeChn)
	go func() {
		for {
			params := SetChildNodesParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type AttributeModifiedParams struct {
	// Id of the node that has changed.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// Attribute name.
	Name string `json:"name"`
	// Attribute value.
	Value string `json:"value"`
}

// Fired when <code>Element</code>'s attribute is modified.
func (obj *DOM) AttributeModified(fn func(params *AttributeModifiedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AttributeModified, closeChn)
	go func() {
		for {
			params := AttributeModifiedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type AttributeRemovedParams struct {
	// Id of the node that has changed.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// A ttribute name.
	Name string `json:"name"`
}

// Fired when <code>Element</code>'s attribute is removed.
func (obj *DOM) AttributeRemoved(fn func(params *AttributeRemovedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(AttributeRemoved, closeChn)
	go func() {
		for {
			params := AttributeRemovedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type InlineStyleInvalidatedParams struct {
	// Ids of the nodes for which the inline styles have been invalidated.
	NodeIds []types.DOM_NodeId `json:"nodeIds"`
}

// Fired when <code>Element</code>'s inline style is modified via a CSS property modification.
// NOTE Experimental
func (obj *DOM) InlineStyleInvalidated(fn func(params *InlineStyleInvalidatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(InlineStyleInvalidated, closeChn)
	go func() {
		for {
			params := InlineStyleInvalidatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type CharacterDataModifiedParams struct {
	// Id of the node that has changed.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// New text value.
	CharacterData string `json:"characterData"`
}

// Mirrors <code>DOMCharacterDataModified</code> event.
func (obj *DOM) CharacterDataModified(fn func(params *CharacterDataModifiedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(CharacterDataModified, closeChn)
	go func() {
		for {
			params := CharacterDataModifiedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ChildNodeCountUpdatedParams struct {
	// Id of the node that has changed.
	NodeId types.DOM_NodeId `json:"nodeId"`
	// New node count.
	ChildNodeCount int `json:"childNodeCount"`
}

// Fired when <code>Container</code>'s child node count has changed.
func (obj *DOM) ChildNodeCountUpdated(fn func(params *ChildNodeCountUpdatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ChildNodeCountUpdated, closeChn)
	go func() {
		for {
			params := ChildNodeCountUpdatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ChildNodeInsertedParams struct {
	// Id of the node that has changed.
	ParentNodeId types.DOM_NodeId `json:"parentNodeId"`
	// If of the previous siblint.
	PreviousNodeId types.DOM_NodeId `json:"previousNodeId"`
	// Inserted node data.
	Node types.DOM_Node `json:"node"`
}

// Mirrors <code>DOMNodeInserted</code> event.
func (obj *DOM) ChildNodeInserted(fn func(params *ChildNodeInsertedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ChildNodeInserted, closeChn)
	go func() {
		for {
			params := ChildNodeInsertedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ChildNodeRemovedParams struct {
	// Parent id.
	ParentNodeId types.DOM_NodeId `json:"parentNodeId"`
	// Id of the node that has been removed.
	NodeId types.DOM_NodeId `json:"nodeId"`
}

// Mirrors <code>DOMNodeRemoved</code> event.
func (obj *DOM) ChildNodeRemoved(fn func(params *ChildNodeRemovedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ChildNodeRemoved, closeChn)
	go func() {
		for {
			params := ChildNodeRemovedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ShadowRootPushedParams struct {
	// Host element id.
	HostId types.DOM_NodeId `json:"hostId"`
	// Shadow root.
	Root types.DOM_Node `json:"root"`
}

// Called when shadow root is pushed into the element.
// NOTE Experimental
func (obj *DOM) ShadowRootPushed(fn func(params *ShadowRootPushedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ShadowRootPushed, closeChn)
	go func() {
		for {
			params := ShadowRootPushedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type ShadowRootPoppedParams struct {
	// Host element id.
	HostId types.DOM_NodeId `json:"hostId"`
	// Shadow root id.
	RootId types.DOM_NodeId `json:"rootId"`
}

// Called when shadow root is popped from the element.
// NOTE Experimental
func (obj *DOM) ShadowRootPopped(fn func(params *ShadowRootPoppedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(ShadowRootPopped, closeChn)
	go func() {
		for {
			params := ShadowRootPoppedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type PseudoElementAddedParams struct {
	// Pseudo element's parent element id.
	ParentId types.DOM_NodeId `json:"parentId"`
	// The added pseudo element.
	PseudoElement types.DOM_Node `json:"pseudoElement"`
}

// Called when a pseudo element is added to an element.
// NOTE Experimental
func (obj *DOM) PseudoElementAdded(fn func(params *PseudoElementAddedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(PseudoElementAdded, closeChn)
	go func() {
		for {
			params := PseudoElementAddedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type PseudoElementRemovedParams struct {
	// Pseudo element's parent element id.
	ParentId types.DOM_NodeId `json:"parentId"`
	// The removed pseudo element id.
	PseudoElementId types.DOM_NodeId `json:"pseudoElementId"`
}

// Called when a pseudo element is removed from an element.
// NOTE Experimental
func (obj *DOM) PseudoElementRemoved(fn func(params *PseudoElementRemovedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(PseudoElementRemoved, closeChn)
	go func() {
		for {
			params := PseudoElementRemovedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

type DistributedNodesUpdatedParams struct {
	// Insertion point where distrubuted nodes were updated.
	InsertionPointId types.DOM_NodeId `json:"insertionPointId"`
	// Distributed nodes for given insertion point.
	DistributedNodes []types.DOM_BackendNode `json:"distributedNodes"`
}

// Called when distrubution is changed.
// NOTE Experimental
func (obj *DOM) DistributedNodesUpdated(fn func(params *DistributedNodesUpdatedParams, err error) bool) {
	closeChn := make(chan struct{})
	decoder := obj.conn.On(DistributedNodesUpdated, closeChn)
	go func() {
		for {
			params := DistributedNodesUpdatedParams{}
			readErr := decoder(&params)
			if !fn(&params, readErr) {
				close(closeChn)
				break
			}
		}
	}()
}

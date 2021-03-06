// Code generated by goa v3.2.6, DO NOT EDIT.
//
// goasnmp HTTP client types
//
// Command:
// $ goa gen github.com/rshade/goasnmp/design

package client

import (
	goasnmpviews "github.com/rshade/goasnmp/gen/goasnmp/views"
)

// ListResponseBody is the type of the "goasnmp" service "list" endpoint HTTP
// response body.
type ListResponseBody []*HostResponse

// AddResponseBody is the type of the "goasnmp" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// Whether or not to walk public tree
	Public *bool `form:"public,omitempty" json:"public,omitempty" xml:"public,omitempty"`
	// Whether or not Ondemand polling is supported
	OnDemand *bool `form:"on_demand,omitempty" json:"on_demand,omitempty" xml:"on_demand,omitempty"`
}

// HostResponse is used to define fields on response body types.
type HostResponse struct {
	// Whether or not to walk public tree
	Public *bool `form:"public,omitempty" json:"public,omitempty" xml:"public,omitempty"`
	// Whether or not Ondemand polling is supported
	OnDemand *bool `form:"on_demand,omitempty" json:"on_demand,omitempty" xml:"on_demand,omitempty"`
}

// NewListHostCollectionOK builds a "goasnmp" service "list" endpoint result
// from a HTTP "OK" response.
func NewListHostCollectionOK(body ListResponseBody) goasnmpviews.HostCollectionView {
	v := make([]*goasnmpviews.HostView, len(body))
	for i, val := range body {
		v[i] = unmarshalHostResponseToGoasnmpviewsHostView(val)
	}
	return v
}

// NewAddHostOK builds a "goasnmp" service "add" endpoint result from a HTTP
// "OK" response.
func NewAddHostOK(body *AddResponseBody) *goasnmpviews.HostView {
	v := &goasnmpviews.HostView{
		Public:   body.Public,
		OnDemand: body.OnDemand,
	}

	return v
}

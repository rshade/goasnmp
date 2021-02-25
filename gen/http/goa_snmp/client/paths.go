// Code generated by goa v3.2.6, DO NOT EDIT.
//
// HTTP request path constructors for the goa-snmp service.
//
// Command:
// $ goa gen github.com/rshade/goasnmp/design

package client

import (
	"fmt"
)

// ListGoaSnmpPath returns the URL path to the goa-snmp service list HTTP endpoint.
func ListGoaSnmpPath() string {
	return "/hosts"
}

// AddGoaSnmpPath returns the URL path to the goa-snmp service add HTTP endpoint.
func AddGoaSnmpPath(hostname string) string {
	return fmt.Sprintf("/hosts/%v", hostname)
}

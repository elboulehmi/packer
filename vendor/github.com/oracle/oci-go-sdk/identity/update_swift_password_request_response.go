// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateSwiftPasswordRequest wrapper for the UpdateSwiftPassword operation
type UpdateSwiftPasswordRequest struct {

	// The OCID of the user.
	UserId *string `mandatory:"true" contributesTo:"path" name:"userId"`

	// The OCID of the Swift password.
	SwiftPasswordId *string `mandatory:"true" contributesTo:"path" name:"swiftPasswordId"`

	// Request object for updating a Swift password.
	UpdateSwiftPasswordDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateSwiftPasswordRequest) String() string {
	return common.PointerString(request)
}

// UpdateSwiftPasswordResponse wrapper for the UpdateSwiftPassword operation
type UpdateSwiftPasswordResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SwiftPassword instance
	SwiftPassword `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateSwiftPasswordResponse) String() string {
	return common.PointerString(response)
}

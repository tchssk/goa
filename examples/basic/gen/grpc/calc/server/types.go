// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen goa.design/goa/examples/basic/design -o
// $(GOPATH)/src/goa.design/goa/examples/basic

package server

import (
	calcsvc "goa.design/goa/examples/basic/gen/calc"
	calcpb "goa.design/goa/examples/basic/gen/grpc/calc/pb"
)

// NewAddPayload builds the payload of the "add" endpoint of the "calc" service
// from the gRPC request type.
func NewAddPayload(message *calcpb.AddRequest) *calcsvc.AddPayload {
	v := &calcsvc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "calc" service.
func NewAddResponse(result int) *calcpb.AddResponse {
	message := &calcpb.AddResponse{}
	message.Field = int32(result)
	return message
}
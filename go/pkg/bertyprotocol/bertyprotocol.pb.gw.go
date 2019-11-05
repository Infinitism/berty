// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: bertyprotocol.proto

/*
Package bertyprotocol is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package bertyprotocol

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_Instance_InstanceExportData_0(ctx context.Context, marshaler runtime.Marshaler, client InstanceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InstanceExportDataRequest
	var metadata runtime.ServerMetadata

	msg, err := client.InstanceExportData(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Instance_InstanceExportData_0(ctx context.Context, marshaler runtime.Marshaler, server InstanceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InstanceExportDataRequest
	var metadata runtime.ServerMetadata

	msg, err := server.InstanceExportData(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterInstanceHandlerServer registers the http handlers for service Instance to "mux".
// UnaryRPC     :call InstanceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterInstanceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InstanceServer) error {

	mux.Handle("POST", pattern_Instance_InstanceExportData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Instance_InstanceExportData_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Instance_InstanceExportData_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterInstanceHandlerFromEndpoint is same as RegisterInstanceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInstanceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInstanceHandler(ctx, mux, conn)
}

// RegisterInstanceHandler registers the http handlers for service Instance to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInstanceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInstanceHandlerClient(ctx, mux, NewInstanceClient(conn))
}

// RegisterInstanceHandlerClient registers the http handlers for service Instance
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InstanceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InstanceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InstanceClient" to call the correct interceptors.
func RegisterInstanceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InstanceClient) error {

	mux.Handle("POST", pattern_Instance_InstanceExportData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Instance_InstanceExportData_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Instance_InstanceExportData_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Instance_InstanceExportData_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"instance", "export-data"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Instance_InstanceExportData_0 = runtime.ForwardResponseMessage
)

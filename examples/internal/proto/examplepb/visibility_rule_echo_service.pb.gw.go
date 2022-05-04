// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: examples/internal/proto/examplepb/visibility_rule_echo_service.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"context"
	"io"
	"net/http"

	"github.com/sezabis/grpc-gateway/v2/runtime"
	"github.com/sezabis/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_VisibilityRuleEchoService_Echo_0 = &utilities.DoubleArray{Encoding: map[string]int{"id": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_VisibilityRuleEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, client VisibilityRuleEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_Echo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Echo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VisibilityRuleEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, server VisibilityRuleEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_Echo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Echo(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_VisibilityRuleEchoService_EchoInternal_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_VisibilityRuleEchoService_EchoInternal_0(ctx context.Context, marshaler runtime.Marshaler, client VisibilityRuleEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoInternal_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.EchoInternal(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VisibilityRuleEchoService_EchoInternal_0(ctx context.Context, marshaler runtime.Marshaler, server VisibilityRuleEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoInternal_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.EchoInternal(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_VisibilityRuleEchoService_EchoPreview_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_VisibilityRuleEchoService_EchoPreview_0(ctx context.Context, marshaler runtime.Marshaler, client VisibilityRuleEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoPreview_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.EchoPreview(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VisibilityRuleEchoService_EchoPreview_0(ctx context.Context, marshaler runtime.Marshaler, server VisibilityRuleEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoPreview_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.EchoPreview(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_VisibilityRuleEchoService_EchoInternalAndPreview_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx context.Context, marshaler runtime.Marshaler, client VisibilityRuleEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoInternalAndPreview_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.EchoInternalAndPreview(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx context.Context, marshaler runtime.Marshaler, server VisibilityRuleEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleEchoService_EchoInternalAndPreview_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.EchoInternalAndPreview(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_VisibilityRuleInternalEchoService_Echo_0 = &utilities.DoubleArray{Encoding: map[string]int{"id": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_VisibilityRuleInternalEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, client VisibilityRuleInternalEchoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleInternalEchoService_Echo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Echo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_VisibilityRuleInternalEchoService_Echo_0(ctx context.Context, marshaler runtime.Marshaler, server VisibilityRuleInternalEchoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq VisibilityRuleSimpleMessage
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_VisibilityRuleInternalEchoService_Echo_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Echo(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterVisibilityRuleEchoServiceHandlerServer registers the http handlers for service VisibilityRuleEchoService to "mux".
// UnaryRPC     :call VisibilityRuleEchoServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterVisibilityRuleEchoServiceHandlerFromEndpoint instead.
func RegisterVisibilityRuleEchoServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server VisibilityRuleEchoServiceServer) error {

	mux.Handle("POST", pattern_VisibilityRuleEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/Echo", runtime.WithHTTPPathPattern("/v1/example/echo/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VisibilityRuleEchoService_Echo_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_Echo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoInternal_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoInternal", runtime.WithHTTPPathPattern("/v1/example/echo_internal"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VisibilityRuleEchoService_EchoInternal_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoInternal_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoPreview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoPreview", runtime.WithHTTPPathPattern("/v1/example/echo_preview"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VisibilityRuleEchoService_EchoPreview_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoPreview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoInternalAndPreview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoInternalAndPreview", runtime.WithHTTPPathPattern("/v1/example/echo_internal_and_preview"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterVisibilityRuleInternalEchoServiceHandlerServer registers the http handlers for service VisibilityRuleInternalEchoService to "mux".
// UnaryRPC     :call VisibilityRuleInternalEchoServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterVisibilityRuleInternalEchoServiceHandlerFromEndpoint instead.
func RegisterVisibilityRuleInternalEchoServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server VisibilityRuleInternalEchoServiceServer) error {

	mux.Handle("POST", pattern_VisibilityRuleInternalEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleInternalEchoService/Echo", runtime.WithHTTPPathPattern("/v1/example/internal/echo/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_VisibilityRuleInternalEchoService_Echo_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleInternalEchoService_Echo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterVisibilityRuleEchoServiceHandlerFromEndpoint is same as RegisterVisibilityRuleEchoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterVisibilityRuleEchoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterVisibilityRuleEchoServiceHandler(ctx, mux, conn)
}

// RegisterVisibilityRuleEchoServiceHandler registers the http handlers for service VisibilityRuleEchoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterVisibilityRuleEchoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterVisibilityRuleEchoServiceHandlerClient(ctx, mux, NewVisibilityRuleEchoServiceClient(conn))
}

// RegisterVisibilityRuleEchoServiceHandlerClient registers the http handlers for service VisibilityRuleEchoService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "VisibilityRuleEchoServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "VisibilityRuleEchoServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "VisibilityRuleEchoServiceClient" to call the correct interceptors.
func RegisterVisibilityRuleEchoServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client VisibilityRuleEchoServiceClient) error {

	mux.Handle("POST", pattern_VisibilityRuleEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/Echo", runtime.WithHTTPPathPattern("/v1/example/echo/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VisibilityRuleEchoService_Echo_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_Echo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoInternal_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoInternal", runtime.WithHTTPPathPattern("/v1/example/echo_internal"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VisibilityRuleEchoService_EchoInternal_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoInternal_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoPreview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoPreview", runtime.WithHTTPPathPattern("/v1/example/echo_preview"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VisibilityRuleEchoService_EchoPreview_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoPreview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_VisibilityRuleEchoService_EchoInternalAndPreview_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleEchoService/EchoInternalAndPreview", runtime.WithHTTPPathPattern("/v1/example/echo_internal_and_preview"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleEchoService_EchoInternalAndPreview_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_VisibilityRuleEchoService_Echo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"v1", "example", "echo", "id"}, ""))

	pattern_VisibilityRuleEchoService_EchoInternal_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "echo_internal"}, ""))

	pattern_VisibilityRuleEchoService_EchoPreview_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "echo_preview"}, ""))

	pattern_VisibilityRuleEchoService_EchoInternalAndPreview_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "example", "echo_internal_and_preview"}, ""))
)

var (
	forward_VisibilityRuleEchoService_Echo_0 = runtime.ForwardResponseMessage

	forward_VisibilityRuleEchoService_EchoInternal_0 = runtime.ForwardResponseMessage

	forward_VisibilityRuleEchoService_EchoPreview_0 = runtime.ForwardResponseMessage

	forward_VisibilityRuleEchoService_EchoInternalAndPreview_0 = runtime.ForwardResponseMessage
)

// RegisterVisibilityRuleInternalEchoServiceHandlerFromEndpoint is same as RegisterVisibilityRuleInternalEchoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterVisibilityRuleInternalEchoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterVisibilityRuleInternalEchoServiceHandler(ctx, mux, conn)
}

// RegisterVisibilityRuleInternalEchoServiceHandler registers the http handlers for service VisibilityRuleInternalEchoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterVisibilityRuleInternalEchoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterVisibilityRuleInternalEchoServiceHandlerClient(ctx, mux, NewVisibilityRuleInternalEchoServiceClient(conn))
}

// RegisterVisibilityRuleInternalEchoServiceHandlerClient registers the http handlers for service VisibilityRuleInternalEchoService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "VisibilityRuleInternalEchoServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "VisibilityRuleInternalEchoServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "VisibilityRuleInternalEchoServiceClient" to call the correct interceptors.
func RegisterVisibilityRuleInternalEchoServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client VisibilityRuleInternalEchoServiceClient) error {

	mux.Handle("POST", pattern_VisibilityRuleInternalEchoService_Echo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.VisibilityRuleInternalEchoService/Echo", runtime.WithHTTPPathPattern("/v1/example/internal/echo/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_VisibilityRuleInternalEchoService_Echo_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_VisibilityRuleInternalEchoService_Echo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_VisibilityRuleInternalEchoService_Echo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"v1", "example", "internal", "echo", "id"}, ""))
)

var (
	forward_VisibilityRuleInternalEchoService_Echo_0 = runtime.ForwardResponseMessage
)

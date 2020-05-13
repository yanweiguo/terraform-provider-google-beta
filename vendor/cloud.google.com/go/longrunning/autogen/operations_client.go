// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package longrunning

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// OperationsCallOptions contains the retry settings for each method of OperationsClient.
type OperationsCallOptions struct {
	ListOperations  []gax.CallOption
	GetOperation    []gax.CallOption
	DeleteOperation []gax.CallOption
	CancelOperation []gax.CallOption
	WaitOperation   []gax.CallOption
}

func defaultOperationsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("longrunning.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultOperationsCallOptions() *OperationsCallOptions {
	return &OperationsCallOptions{
		ListOperations: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    500 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		GetOperation: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    500 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		DeleteOperation: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    500 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		CancelOperation: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    500 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 2.00,
				})
			}),
		},
		WaitOperation: []gax.CallOption{},
	}
}

// OperationsClient is a client for interacting with Long Running Operations API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type OperationsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// The gRPC API client.
	operationsClient longrunningpb.OperationsClient

	// The call options for this service.
	CallOptions *OperationsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewOperationsClient creates a new operations client.
//
// Manages long-running operations with an API service.
//
// When an API method normally takes long time to complete, it can be designed
// to return Operation to the client, and the client can use this
// interface to receive the real response asynchronously by polling the
// operation resource, or pass the operation resource to another API (such as
// Google Cloud Pub/Sub API) to receive the response.  Any API service that
// returns long-running operations should implement the Operations interface
// so developers can have a consistent client experience.
func NewOperationsClient(ctx context.Context, opts ...option.ClientOption) (*OperationsClient, error) {
	connPool, err := gtransport.DialPool(ctx, append(defaultOperationsClientOptions(), opts...)...)
	if err != nil {
		return nil, err
	}
	c := &OperationsClient{
		connPool:    connPool,
		CallOptions: defaultOperationsCallOptions(),

		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.SetGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *OperationsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *OperationsClient) Close() error {
	return c.connPool.Close()
}

// SetGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *OperationsClient) SetGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListOperations lists operations that match the specified filter in the request. If the
// server doesn’t support this method, it returns UNIMPLEMENTED.
//
// NOTE: the name binding allows API services to override the binding
// to use different resource name schemes, such as users/*/operations. To
// override the binding, API services can add a binding such as
// "/v1/{name=users/*}/operations" to their service configuration.
// For backwards compatibility, the default name includes the operations
// collection id, however overriding users must ensure the name binding
// is the parent resource, without the operations collection id.
func (c *OperationsClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListOperations[0:len(c.CallOptions.ListOperations):len(c.CallOptions.ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		var resp *longrunningpb.ListOperationsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.operationsClient.ListOperations(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.Operations, resp.NextPageToken, nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.PageSize)
	it.pageInfo.Token = req.PageToken
	return it
}

// GetOperation gets the latest state of a long-running operation.  Clients can use this
// method to poll the operation result at intervals as recommended by the API
// service.
func (c *OperationsClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetOperation[0:len(c.CallOptions.GetOperation):len(c.CallOptions.GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteOperation deletes a long-running operation. This method indicates that the client is
// no longer interested in the operation result. It does not cancel the
// operation. If the server doesn’t support this method, it returns
// google.rpc.Code.UNIMPLEMENTED.
func (c *OperationsClient) DeleteOperation(ctx context.Context, req *longrunningpb.DeleteOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteOperation[0:len(c.CallOptions.DeleteOperation):len(c.CallOptions.DeleteOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.DeleteOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CancelOperation starts asynchronous cancellation on a long-running operation.  The server
// makes a best effort to cancel the operation, but success is not
// guaranteed.  If the server doesn’t support this method, it returns
// google.rpc.Code.UNIMPLEMENTED.  Clients can use
// Operations.GetOperation or
// other methods to check whether the cancellation succeeded or whether the
// operation completed despite cancellation. On successful cancellation,
// the operation is not deleted; instead, it becomes an operation with
// an Operation.error value with a google.rpc.Status.code of 1,
// corresponding to Code.CANCELLED.
func (c *OperationsClient) CancelOperation(ctx context.Context, req *longrunningpb.CancelOperationRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CancelOperation[0:len(c.CallOptions.CancelOperation):len(c.CallOptions.CancelOperation)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.operationsClient.CancelOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// WaitOperation waits for the specified long-running operation until it is done or reaches
// at most a specified timeout, returning the latest state.  If the operation
// is already done, the latest state is immediately returned.  If the timeout
// specified is greater than the default HTTP/RPC timeout, the HTTP/RPC
// timeout is used.  If the server does not support this method, it returns
// google.rpc.Code.UNIMPLEMENTED.
// Note that this method is on a best-effort basis.  It may return the latest
// state before the specified timeout (including immediately), meaning even an
// immediate response is no guarantee that the operation is done.
func (c *OperationsClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.WaitOperation[0:len(c.CallOptions.WaitOperation):len(c.CallOptions.WaitOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.WaitOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// OperationIterator manages a stream of *longrunningpb.Operation.
type OperationIterator struct {
	items    []*longrunningpb.Operation
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*longrunningpb.Operation, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *OperationIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *OperationIterator) Next() (*longrunningpb.Operation, error) {
	var item *longrunningpb.Operation
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *OperationIterator) bufLen() int {
	return len(it.items)
}

func (it *OperationIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

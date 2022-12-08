// Copyright 2022 Google LLC
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

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newServiceMonitoringClientHook clientHook

// ServiceMonitoringCallOptions contains the retry settings for each method of ServiceMonitoringClient.
type ServiceMonitoringCallOptions struct {
	CreateService               []gax.CallOption
	GetService                  []gax.CallOption
	ListServices                []gax.CallOption
	UpdateService               []gax.CallOption
	DeleteService               []gax.CallOption
	CreateServiceLevelObjective []gax.CallOption
	GetServiceLevelObjective    []gax.CallOption
	ListServiceLevelObjectives  []gax.CallOption
	UpdateServiceLevelObjective []gax.CallOption
	DeleteServiceLevelObjective []gax.CallOption
}

func defaultServiceMonitoringGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServiceMonitoringCallOptions() *ServiceMonitoringCallOptions {
	return &ServiceMonitoringCallOptions{
		CreateService: []gax.CallOption{},
		GetService: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListServices: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateService: []gax.CallOption{},
		DeleteService: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateServiceLevelObjective: []gax.CallOption{},
		GetServiceLevelObjective: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListServiceLevelObjectives: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateServiceLevelObjective: []gax.CallOption{},
		DeleteServiceLevelObjective: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalServiceMonitoringClient is an interface that defines the methods available from Cloud Monitoring API.
type internalServiceMonitoringClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateService(context.Context, *monitoringpb.CreateServiceRequest, ...gax.CallOption) (*monitoringpb.Service, error)
	GetService(context.Context, *monitoringpb.GetServiceRequest, ...gax.CallOption) (*monitoringpb.Service, error)
	ListServices(context.Context, *monitoringpb.ListServicesRequest, ...gax.CallOption) *ServiceIterator
	UpdateService(context.Context, *monitoringpb.UpdateServiceRequest, ...gax.CallOption) (*monitoringpb.Service, error)
	DeleteService(context.Context, *monitoringpb.DeleteServiceRequest, ...gax.CallOption) error
	CreateServiceLevelObjective(context.Context, *monitoringpb.CreateServiceLevelObjectiveRequest, ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error)
	GetServiceLevelObjective(context.Context, *monitoringpb.GetServiceLevelObjectiveRequest, ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error)
	ListServiceLevelObjectives(context.Context, *monitoringpb.ListServiceLevelObjectivesRequest, ...gax.CallOption) *ServiceLevelObjectiveIterator
	UpdateServiceLevelObjective(context.Context, *monitoringpb.UpdateServiceLevelObjectiveRequest, ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error)
	DeleteServiceLevelObjective(context.Context, *monitoringpb.DeleteServiceLevelObjectiveRequest, ...gax.CallOption) error
}

// ServiceMonitoringClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The Cloud Monitoring Service-Oriented Monitoring API has endpoints for
// managing and querying aspects of a workspace’s services. These include the
// Service's monitored resources, its Service-Level Objectives, and a taxonomy
// of categorized Health Metrics.
type ServiceMonitoringClient struct {
	// The internal transport-dependent client.
	internalClient internalServiceMonitoringClient

	// The call options for this service.
	CallOptions *ServiceMonitoringCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServiceMonitoringClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServiceMonitoringClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ServiceMonitoringClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateService create a Service.
func (c *ServiceMonitoringClient) CreateService(ctx context.Context, req *monitoringpb.CreateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	return c.internalClient.CreateService(ctx, req, opts...)
}

// GetService get the named Service.
func (c *ServiceMonitoringClient) GetService(ctx context.Context, req *monitoringpb.GetServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	return c.internalClient.GetService(ctx, req, opts...)
}

// ListServices list Services for this workspace.
func (c *ServiceMonitoringClient) ListServices(ctx context.Context, req *monitoringpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	return c.internalClient.ListServices(ctx, req, opts...)
}

// UpdateService update this Service.
func (c *ServiceMonitoringClient) UpdateService(ctx context.Context, req *monitoringpb.UpdateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	return c.internalClient.UpdateService(ctx, req, opts...)
}

// DeleteService soft delete this Service.
func (c *ServiceMonitoringClient) DeleteService(ctx context.Context, req *monitoringpb.DeleteServiceRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteService(ctx, req, opts...)
}

// CreateServiceLevelObjective create a ServiceLevelObjective for the given Service.
func (c *ServiceMonitoringClient) CreateServiceLevelObjective(ctx context.Context, req *monitoringpb.CreateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	return c.internalClient.CreateServiceLevelObjective(ctx, req, opts...)
}

// GetServiceLevelObjective get a ServiceLevelObjective by name.
func (c *ServiceMonitoringClient) GetServiceLevelObjective(ctx context.Context, req *monitoringpb.GetServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	return c.internalClient.GetServiceLevelObjective(ctx, req, opts...)
}

// ListServiceLevelObjectives list the ServiceLevelObjectives for the given Service.
func (c *ServiceMonitoringClient) ListServiceLevelObjectives(ctx context.Context, req *monitoringpb.ListServiceLevelObjectivesRequest, opts ...gax.CallOption) *ServiceLevelObjectiveIterator {
	return c.internalClient.ListServiceLevelObjectives(ctx, req, opts...)
}

// UpdateServiceLevelObjective update the given ServiceLevelObjective.
func (c *ServiceMonitoringClient) UpdateServiceLevelObjective(ctx context.Context, req *monitoringpb.UpdateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	return c.internalClient.UpdateServiceLevelObjective(ctx, req, opts...)
}

// DeleteServiceLevelObjective delete the given ServiceLevelObjective.
func (c *ServiceMonitoringClient) DeleteServiceLevelObjective(ctx context.Context, req *monitoringpb.DeleteServiceLevelObjectiveRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteServiceLevelObjective(ctx, req, opts...)
}

// serviceMonitoringGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type serviceMonitoringGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ServiceMonitoringClient
	CallOptions **ServiceMonitoringCallOptions

	// The gRPC API client.
	serviceMonitoringClient monitoringpb.ServiceMonitoringServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServiceMonitoringClient creates a new service monitoring service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The Cloud Monitoring Service-Oriented Monitoring API has endpoints for
// managing and querying aspects of a workspace’s services. These include the
// Service's monitored resources, its Service-Level Objectives, and a taxonomy
// of categorized Health Metrics.
func NewServiceMonitoringClient(ctx context.Context, opts ...option.ClientOption) (*ServiceMonitoringClient, error) {
	clientOpts := defaultServiceMonitoringGRPCClientOptions()
	if newServiceMonitoringClientHook != nil {
		hookOpts, err := newServiceMonitoringClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ServiceMonitoringClient{CallOptions: defaultServiceMonitoringCallOptions()}

	c := &serviceMonitoringGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		serviceMonitoringClient: monitoringpb.NewServiceMonitoringServiceClient(connPool),
		CallOptions:             &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *serviceMonitoringGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *serviceMonitoringGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *serviceMonitoringGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *serviceMonitoringGRPCClient) CreateService(ctx context.Context, req *monitoringpb.CreateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateService[0:len((*c.CallOptions).CreateService):len((*c.CallOptions).CreateService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.CreateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) GetService(ctx context.Context, req *monitoringpb.GetServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetService[0:len((*c.CallOptions).GetService):len((*c.CallOptions).GetService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.GetService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) ListServices(ctx context.Context, req *monitoringpb.ListServicesRequest, opts ...gax.CallOption) *ServiceIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListServices[0:len((*c.CallOptions).ListServices):len((*c.CallOptions).ListServices)], opts...)
	it := &ServiceIterator{}
	req = proto.Clone(req).(*monitoringpb.ListServicesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.Service, string, error) {
		resp := &monitoringpb.ListServicesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.serviceMonitoringClient.ListServices(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServices(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *serviceMonitoringGRPCClient) UpdateService(ctx context.Context, req *monitoringpb.UpdateServiceRequest, opts ...gax.CallOption) (*monitoringpb.Service, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service.name", url.QueryEscape(req.GetService().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateService[0:len((*c.CallOptions).UpdateService):len((*c.CallOptions).UpdateService)], opts...)
	var resp *monitoringpb.Service
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.UpdateService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) DeleteService(ctx context.Context, req *monitoringpb.DeleteServiceRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteService[0:len((*c.CallOptions).DeleteService):len((*c.CallOptions).DeleteService)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.serviceMonitoringClient.DeleteService(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *serviceMonitoringGRPCClient) CreateServiceLevelObjective(ctx context.Context, req *monitoringpb.CreateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateServiceLevelObjective[0:len((*c.CallOptions).CreateServiceLevelObjective):len((*c.CallOptions).CreateServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.CreateServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) GetServiceLevelObjective(ctx context.Context, req *monitoringpb.GetServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetServiceLevelObjective[0:len((*c.CallOptions).GetServiceLevelObjective):len((*c.CallOptions).GetServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.GetServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) ListServiceLevelObjectives(ctx context.Context, req *monitoringpb.ListServiceLevelObjectivesRequest, opts ...gax.CallOption) *ServiceLevelObjectiveIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListServiceLevelObjectives[0:len((*c.CallOptions).ListServiceLevelObjectives):len((*c.CallOptions).ListServiceLevelObjectives)], opts...)
	it := &ServiceLevelObjectiveIterator{}
	req = proto.Clone(req).(*monitoringpb.ListServiceLevelObjectivesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.ServiceLevelObjective, string, error) {
		resp := &monitoringpb.ListServiceLevelObjectivesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.serviceMonitoringClient.ListServiceLevelObjectives(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetServiceLevelObjectives(), resp.GetNextPageToken(), nil
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
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *serviceMonitoringGRPCClient) UpdateServiceLevelObjective(ctx context.Context, req *monitoringpb.UpdateServiceLevelObjectiveRequest, opts ...gax.CallOption) (*monitoringpb.ServiceLevelObjective, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_level_objective.name", url.QueryEscape(req.GetServiceLevelObjective().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateServiceLevelObjective[0:len((*c.CallOptions).UpdateServiceLevelObjective):len((*c.CallOptions).UpdateServiceLevelObjective)], opts...)
	var resp *monitoringpb.ServiceLevelObjective
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceMonitoringClient.UpdateServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *serviceMonitoringGRPCClient) DeleteServiceLevelObjective(ctx context.Context, req *monitoringpb.DeleteServiceLevelObjectiveRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteServiceLevelObjective[0:len((*c.CallOptions).DeleteServiceLevelObjective):len((*c.CallOptions).DeleteServiceLevelObjective)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.serviceMonitoringClient.DeleteServiceLevelObjective(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ServiceIterator manages a stream of *monitoringpb.Service.
type ServiceIterator struct {
	items    []*monitoringpb.Service
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
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.Service, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceIterator) Next() (*monitoringpb.Service, error) {
	var item *monitoringpb.Service
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ServiceLevelObjectiveIterator manages a stream of *monitoringpb.ServiceLevelObjective.
type ServiceLevelObjectiveIterator struct {
	items    []*monitoringpb.ServiceLevelObjective
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
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.ServiceLevelObjective, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ServiceLevelObjectiveIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ServiceLevelObjectiveIterator) Next() (*monitoringpb.ServiceLevelObjective, error) {
	var item *monitoringpb.ServiceLevelObjective
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ServiceLevelObjectiveIterator) bufLen() int {
	return len(it.items)
}

func (it *ServiceLevelObjectiveIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

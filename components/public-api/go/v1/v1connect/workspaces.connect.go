// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: gitpod/v1/workspaces.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/gitpod-io/gitpod/public-api/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// WorkspacesServiceName is the fully-qualified name of the WorkspacesService service.
	WorkspacesServiceName = "gitpod.v1.WorkspacesService"
)

// WorkspacesServiceClient is a client for the gitpod.v1.WorkspacesService service.
type WorkspacesServiceClient interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// GetOwnerToken returns an owner token.
	GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StartWorkspace starts an existing workspace.
	StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error)
	// GetRunningWorkspaceInstance returns the currently active instance of a workspace.
	// Errors:
	//
	//	FAILED_PRECONDITION: if a workspace does not a currently active instance
	GetActiveWorkspaceInstance(context.Context, *connect_go.Request[v1.GetActiveWorkspaceInstanceRequest]) (*connect_go.Response[v1.GetActiveWorkspaceInstanceResponse], error)
	// GetWorkspaceInstanceOwnerToken returns the owner token of a workspace instance.
	// Note: the owner token is not part of the workspace instance status so that we can scope its access on the
	//
	//	API function level.
	GetWorkspaceInstanceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceInstanceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceInstanceOwnerTokenResponse], error)
	// ListenToWorkspaceInstance listens to workspace instance updates.
	ListenToWorkspaceInstance(context.Context, *connect_go.Request[v1.ListenToWorkspaceInstanceRequest]) (*connect_go.ServerStreamForClient[v1.ListenToWorkspaceInstanceResponse], error)
	// ListenToImageBuildLogs streams (currently or previously) running workspace image build logs
	ListenToImageBuildLogs(context.Context, *connect_go.Request[v1.ListenToImageBuildLogsRequest]) (*connect_go.ServerStreamForClient[v1.ListenToImageBuildLogsResponse], error)
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.ServerStreamForClient[v1.StopWorkspaceResponse], error)
}

// NewWorkspacesServiceClient constructs a client for the gitpod.v1.WorkspacesService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWorkspacesServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WorkspacesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &workspacesServiceClient{
		listWorkspaces: connect_go.NewClient[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/ListWorkspaces",
			opts...,
		),
		getWorkspace: connect_go.NewClient[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/GetWorkspace",
			opts...,
		),
		getOwnerToken: connect_go.NewClient[v1.GetOwnerTokenRequest, v1.GetOwnerTokenResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/GetOwnerToken",
			opts...,
		),
		createAndStartWorkspace: connect_go.NewClient[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/CreateAndStartWorkspace",
			opts...,
		),
		startWorkspace: connect_go.NewClient[v1.StartWorkspaceRequest, v1.StartWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/StartWorkspace",
			opts...,
		),
		getActiveWorkspaceInstance: connect_go.NewClient[v1.GetActiveWorkspaceInstanceRequest, v1.GetActiveWorkspaceInstanceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/GetActiveWorkspaceInstance",
			opts...,
		),
		getWorkspaceInstanceOwnerToken: connect_go.NewClient[v1.GetWorkspaceInstanceOwnerTokenRequest, v1.GetWorkspaceInstanceOwnerTokenResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/GetWorkspaceInstanceOwnerToken",
			opts...,
		),
		listenToWorkspaceInstance: connect_go.NewClient[v1.ListenToWorkspaceInstanceRequest, v1.ListenToWorkspaceInstanceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/ListenToWorkspaceInstance",
			opts...,
		),
		listenToImageBuildLogs: connect_go.NewClient[v1.ListenToImageBuildLogsRequest, v1.ListenToImageBuildLogsResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/ListenToImageBuildLogs",
			opts...,
		),
		stopWorkspace: connect_go.NewClient[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse](
			httpClient,
			baseURL+"/gitpod.v1.WorkspacesService/StopWorkspace",
			opts...,
		),
	}
}

// workspacesServiceClient implements WorkspacesServiceClient.
type workspacesServiceClient struct {
	listWorkspaces                 *connect_go.Client[v1.ListWorkspacesRequest, v1.ListWorkspacesResponse]
	getWorkspace                   *connect_go.Client[v1.GetWorkspaceRequest, v1.GetWorkspaceResponse]
	getOwnerToken                  *connect_go.Client[v1.GetOwnerTokenRequest, v1.GetOwnerTokenResponse]
	createAndStartWorkspace        *connect_go.Client[v1.CreateAndStartWorkspaceRequest, v1.CreateAndStartWorkspaceResponse]
	startWorkspace                 *connect_go.Client[v1.StartWorkspaceRequest, v1.StartWorkspaceResponse]
	getActiveWorkspaceInstance     *connect_go.Client[v1.GetActiveWorkspaceInstanceRequest, v1.GetActiveWorkspaceInstanceResponse]
	getWorkspaceInstanceOwnerToken *connect_go.Client[v1.GetWorkspaceInstanceOwnerTokenRequest, v1.GetWorkspaceInstanceOwnerTokenResponse]
	listenToWorkspaceInstance      *connect_go.Client[v1.ListenToWorkspaceInstanceRequest, v1.ListenToWorkspaceInstanceResponse]
	listenToImageBuildLogs         *connect_go.Client[v1.ListenToImageBuildLogsRequest, v1.ListenToImageBuildLogsResponse]
	stopWorkspace                  *connect_go.Client[v1.StopWorkspaceRequest, v1.StopWorkspaceResponse]
}

// ListWorkspaces calls gitpod.v1.WorkspacesService.ListWorkspaces.
func (c *workspacesServiceClient) ListWorkspaces(ctx context.Context, req *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return c.listWorkspaces.CallUnary(ctx, req)
}

// GetWorkspace calls gitpod.v1.WorkspacesService.GetWorkspace.
func (c *workspacesServiceClient) GetWorkspace(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return c.getWorkspace.CallUnary(ctx, req)
}

// GetOwnerToken calls gitpod.v1.WorkspacesService.GetOwnerToken.
func (c *workspacesServiceClient) GetOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error) {
	return c.getOwnerToken.CallUnary(ctx, req)
}

// CreateAndStartWorkspace calls gitpod.v1.WorkspacesService.CreateAndStartWorkspace.
func (c *workspacesServiceClient) CreateAndStartWorkspace(ctx context.Context, req *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return c.createAndStartWorkspace.CallUnary(ctx, req)
}

// StartWorkspace calls gitpod.v1.WorkspacesService.StartWorkspace.
func (c *workspacesServiceClient) StartWorkspace(ctx context.Context, req *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	return c.startWorkspace.CallUnary(ctx, req)
}

// GetActiveWorkspaceInstance calls gitpod.v1.WorkspacesService.GetActiveWorkspaceInstance.
func (c *workspacesServiceClient) GetActiveWorkspaceInstance(ctx context.Context, req *connect_go.Request[v1.GetActiveWorkspaceInstanceRequest]) (*connect_go.Response[v1.GetActiveWorkspaceInstanceResponse], error) {
	return c.getActiveWorkspaceInstance.CallUnary(ctx, req)
}

// GetWorkspaceInstanceOwnerToken calls gitpod.v1.WorkspacesService.GetWorkspaceInstanceOwnerToken.
func (c *workspacesServiceClient) GetWorkspaceInstanceOwnerToken(ctx context.Context, req *connect_go.Request[v1.GetWorkspaceInstanceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceInstanceOwnerTokenResponse], error) {
	return c.getWorkspaceInstanceOwnerToken.CallUnary(ctx, req)
}

// ListenToWorkspaceInstance calls gitpod.v1.WorkspacesService.ListenToWorkspaceInstance.
func (c *workspacesServiceClient) ListenToWorkspaceInstance(ctx context.Context, req *connect_go.Request[v1.ListenToWorkspaceInstanceRequest]) (*connect_go.ServerStreamForClient[v1.ListenToWorkspaceInstanceResponse], error) {
	return c.listenToWorkspaceInstance.CallServerStream(ctx, req)
}

// ListenToImageBuildLogs calls gitpod.v1.WorkspacesService.ListenToImageBuildLogs.
func (c *workspacesServiceClient) ListenToImageBuildLogs(ctx context.Context, req *connect_go.Request[v1.ListenToImageBuildLogsRequest]) (*connect_go.ServerStreamForClient[v1.ListenToImageBuildLogsResponse], error) {
	return c.listenToImageBuildLogs.CallServerStream(ctx, req)
}

// StopWorkspace calls gitpod.v1.WorkspacesService.StopWorkspace.
func (c *workspacesServiceClient) StopWorkspace(ctx context.Context, req *connect_go.Request[v1.StopWorkspaceRequest]) (*connect_go.ServerStreamForClient[v1.StopWorkspaceResponse], error) {
	return c.stopWorkspace.CallServerStream(ctx, req)
}

// WorkspacesServiceHandler is an implementation of the gitpod.v1.WorkspacesService service.
type WorkspacesServiceHandler interface {
	// ListWorkspaces enumerates all workspaces belonging to the authenticated user.
	ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error)
	// GetWorkspace returns a single workspace.
	GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error)
	// GetOwnerToken returns an owner token.
	GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error)
	// CreateAndStartWorkspace creates a new workspace and starts it.
	CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error)
	// StartWorkspace starts an existing workspace.
	StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error)
	// GetRunningWorkspaceInstance returns the currently active instance of a workspace.
	// Errors:
	//
	//	FAILED_PRECONDITION: if a workspace does not a currently active instance
	GetActiveWorkspaceInstance(context.Context, *connect_go.Request[v1.GetActiveWorkspaceInstanceRequest]) (*connect_go.Response[v1.GetActiveWorkspaceInstanceResponse], error)
	// GetWorkspaceInstanceOwnerToken returns the owner token of a workspace instance.
	// Note: the owner token is not part of the workspace instance status so that we can scope its access on the
	//
	//	API function level.
	GetWorkspaceInstanceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceInstanceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceInstanceOwnerTokenResponse], error)
	// ListenToWorkspaceInstance listens to workspace instance updates.
	ListenToWorkspaceInstance(context.Context, *connect_go.Request[v1.ListenToWorkspaceInstanceRequest], *connect_go.ServerStream[v1.ListenToWorkspaceInstanceResponse]) error
	// ListenToImageBuildLogs streams (currently or previously) running workspace image build logs
	ListenToImageBuildLogs(context.Context, *connect_go.Request[v1.ListenToImageBuildLogsRequest], *connect_go.ServerStream[v1.ListenToImageBuildLogsResponse]) error
	// StopWorkspace stops a running workspace (instance).
	// Errors:
	//
	//	NOT_FOUND:           the workspace_id is unkown
	//	FAILED_PRECONDITION: if there's no running instance
	StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest], *connect_go.ServerStream[v1.StopWorkspaceResponse]) error
}

// NewWorkspacesServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWorkspacesServiceHandler(svc WorkspacesServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/gitpod.v1.WorkspacesService/ListWorkspaces", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/ListWorkspaces",
		svc.ListWorkspaces,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/GetWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/GetWorkspace",
		svc.GetWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/GetOwnerToken", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/GetOwnerToken",
		svc.GetOwnerToken,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/CreateAndStartWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/CreateAndStartWorkspace",
		svc.CreateAndStartWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/StartWorkspace", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/StartWorkspace",
		svc.StartWorkspace,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/GetActiveWorkspaceInstance", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/GetActiveWorkspaceInstance",
		svc.GetActiveWorkspaceInstance,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/GetWorkspaceInstanceOwnerToken", connect_go.NewUnaryHandler(
		"/gitpod.v1.WorkspacesService/GetWorkspaceInstanceOwnerToken",
		svc.GetWorkspaceInstanceOwnerToken,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/ListenToWorkspaceInstance", connect_go.NewServerStreamHandler(
		"/gitpod.v1.WorkspacesService/ListenToWorkspaceInstance",
		svc.ListenToWorkspaceInstance,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/ListenToImageBuildLogs", connect_go.NewServerStreamHandler(
		"/gitpod.v1.WorkspacesService/ListenToImageBuildLogs",
		svc.ListenToImageBuildLogs,
		opts...,
	))
	mux.Handle("/gitpod.v1.WorkspacesService/StopWorkspace", connect_go.NewServerStreamHandler(
		"/gitpod.v1.WorkspacesService/StopWorkspace",
		svc.StopWorkspace,
		opts...,
	))
	return "/gitpod.v1.WorkspacesService/", mux
}

// UnimplementedWorkspacesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedWorkspacesServiceHandler struct{}

func (UnimplementedWorkspacesServiceHandler) ListWorkspaces(context.Context, *connect_go.Request[v1.ListWorkspacesRequest]) (*connect_go.Response[v1.ListWorkspacesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.ListWorkspaces is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetWorkspace(context.Context, *connect_go.Request[v1.GetWorkspaceRequest]) (*connect_go.Response[v1.GetWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.GetWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetOwnerToken(context.Context, *connect_go.Request[v1.GetOwnerTokenRequest]) (*connect_go.Response[v1.GetOwnerTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.GetOwnerToken is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) CreateAndStartWorkspace(context.Context, *connect_go.Request[v1.CreateAndStartWorkspaceRequest]) (*connect_go.Response[v1.CreateAndStartWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.CreateAndStartWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) StartWorkspace(context.Context, *connect_go.Request[v1.StartWorkspaceRequest]) (*connect_go.Response[v1.StartWorkspaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.StartWorkspace is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetActiveWorkspaceInstance(context.Context, *connect_go.Request[v1.GetActiveWorkspaceInstanceRequest]) (*connect_go.Response[v1.GetActiveWorkspaceInstanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.GetActiveWorkspaceInstance is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) GetWorkspaceInstanceOwnerToken(context.Context, *connect_go.Request[v1.GetWorkspaceInstanceOwnerTokenRequest]) (*connect_go.Response[v1.GetWorkspaceInstanceOwnerTokenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.GetWorkspaceInstanceOwnerToken is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) ListenToWorkspaceInstance(context.Context, *connect_go.Request[v1.ListenToWorkspaceInstanceRequest], *connect_go.ServerStream[v1.ListenToWorkspaceInstanceResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.ListenToWorkspaceInstance is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) ListenToImageBuildLogs(context.Context, *connect_go.Request[v1.ListenToImageBuildLogsRequest], *connect_go.ServerStream[v1.ListenToImageBuildLogsResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.ListenToImageBuildLogs is not implemented"))
}

func (UnimplementedWorkspacesServiceHandler) StopWorkspace(context.Context, *connect_go.Request[v1.StopWorkspaceRequest], *connect_go.ServerStream[v1.StopWorkspaceResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("gitpod.v1.WorkspacesService.StopWorkspace is not implemented"))
}
// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: google/cloud/discoveryengine/v1alpha/acl_config_service.proto

package discoveryenginepb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AclConfigService_UpdateAclConfig_FullMethodName = "/google.cloud.discoveryengine.v1alpha.AclConfigService/UpdateAclConfig"
	AclConfigService_GetAclConfig_FullMethodName    = "/google.cloud.discoveryengine.v1alpha.AclConfigService/GetAclConfig"
)

// AclConfigServiceClient is the client API for AclConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AclConfigServiceClient interface {
	// Default ACL configuration for use in a location of a customer's project.
	// Updates will only reflect to new data stores. Existing data stores will
	// still use the old value.
	UpdateAclConfig(ctx context.Context, in *UpdateAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error)
	// Gets the [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig].
	GetAclConfig(ctx context.Context, in *GetAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error)
}

type aclConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAclConfigServiceClient(cc grpc.ClientConnInterface) AclConfigServiceClient {
	return &aclConfigServiceClient{cc}
}

func (c *aclConfigServiceClient) UpdateAclConfig(ctx context.Context, in *UpdateAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error) {
	out := new(AclConfig)
	err := c.cc.Invoke(ctx, AclConfigService_UpdateAclConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aclConfigServiceClient) GetAclConfig(ctx context.Context, in *GetAclConfigRequest, opts ...grpc.CallOption) (*AclConfig, error) {
	out := new(AclConfig)
	err := c.cc.Invoke(ctx, AclConfigService_GetAclConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AclConfigServiceServer is the server API for AclConfigService service.
// All implementations should embed UnimplementedAclConfigServiceServer
// for forward compatibility
type AclConfigServiceServer interface {
	// Default ACL configuration for use in a location of a customer's project.
	// Updates will only reflect to new data stores. Existing data stores will
	// still use the old value.
	UpdateAclConfig(context.Context, *UpdateAclConfigRequest) (*AclConfig, error)
	// Gets the [AclConfig][google.cloud.discoveryengine.v1alpha.AclConfig].
	GetAclConfig(context.Context, *GetAclConfigRequest) (*AclConfig, error)
}

// UnimplementedAclConfigServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAclConfigServiceServer struct {
}

func (UnimplementedAclConfigServiceServer) UpdateAclConfig(context.Context, *UpdateAclConfigRequest) (*AclConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAclConfig not implemented")
}
func (UnimplementedAclConfigServiceServer) GetAclConfig(context.Context, *GetAclConfigRequest) (*AclConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAclConfig not implemented")
}

// UnsafeAclConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AclConfigServiceServer will
// result in compilation errors.
type UnsafeAclConfigServiceServer interface {
	mustEmbedUnimplementedAclConfigServiceServer()
}

func RegisterAclConfigServiceServer(s grpc.ServiceRegistrar, srv AclConfigServiceServer) {
	s.RegisterService(&AclConfigService_ServiceDesc, srv)
}

func _AclConfigService_UpdateAclConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAclConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclConfigServiceServer).UpdateAclConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AclConfigService_UpdateAclConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclConfigServiceServer).UpdateAclConfig(ctx, req.(*UpdateAclConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AclConfigService_GetAclConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAclConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AclConfigServiceServer).GetAclConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AclConfigService_GetAclConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AclConfigServiceServer).GetAclConfig(ctx, req.(*GetAclConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AclConfigService_ServiceDesc is the grpc.ServiceDesc for AclConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AclConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.AclConfigService",
	HandlerType: (*AclConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAclConfig",
			Handler:    _AclConfigService_UpdateAclConfig_Handler,
		},
		{
			MethodName: "GetAclConfig",
			Handler:    _AclConfigService_GetAclConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/acl_config_service.proto",
}

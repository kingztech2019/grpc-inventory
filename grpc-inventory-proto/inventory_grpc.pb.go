// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: inventory.proto

package grpc_inventory_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	InventoryService_GetProduct_FullMethodName    = "/inventory.InventoryService/GetProduct"
	InventoryService_UpdateProduct_FullMethodName = "/inventory.InventoryService/UpdateProduct"
	InventoryService_DeleteProduct_FullMethodName = "/inventory.InventoryService/DeleteProduct"
	InventoryService_CreateOrder_FullMethodName   = "/inventory.InventoryService/CreateOrder"
	InventoryService_GetOrder_FullMethodName      = "/inventory.InventoryService/GetOrder"
	InventoryService_UpdateOrder_FullMethodName   = "/inventory.InventoryService/UpdateOrder"
	InventoryService_DeleteOrder_FullMethodName   = "/inventory.InventoryService/DeleteOrder"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, InventoryService_UpdateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, InventoryService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_UpdateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_DeleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations must embed UnimplementedInventoryServiceServer
// for forward compatibility.
type InventoryServiceServer interface {
	GetProduct(context.Context, *GetProductRequest) (*ProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error)
	GetOrder(context.Context, *GetOrderRequest) (*OrderResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

// UnimplementedInventoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) GetProduct(context.Context, *GetProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedInventoryServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedInventoryServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedInventoryServiceServer) GetOrder(context.Context, *GetOrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedInventoryServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedInventoryServiceServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedInventoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _InventoryService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _InventoryService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _InventoryService_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _InventoryService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _InventoryService_GetOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _InventoryService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _InventoryService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}

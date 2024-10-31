package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/kingztech2019/proto_grpc/proto_inventory/grpc-inventory-proto" // Update this import path
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// In-memory storage for products and orders
var (
	productStore = make(map[string]*pb.Product)
	orderStore   = make(map[string]*pb.Order)
	mu           sync.Mutex
)

// Server struct implementing pb.InventoryServiceServer
type server struct {
	pb.UnimplementedInventoryServiceServer
}

// GetProduct fetches a product by its ID
func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	product, exists := productStore[req.ProductId]
	if !exists {
		return &pb.ProductResponse{Status: "error", Message: "Product not found"}, nil
	}
	return &pb.ProductResponse{Product: product, Status: "success"}, nil
}

// UpdateProduct updates an existing product
func (s *server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	productStore[req.Product.Id] = req.Product
	return &pb.ProductResponse{Product: req.Product, Status: "success", Message: "Product updated"}, nil
}

// DeleteProduct deletes a product by its ID
func (s *server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := productStore[req.ProductId]; exists {
		delete(productStore, req.ProductId)
		return &pb.DeleteProductResponse{Success: true}, nil
	}
	return &pb.DeleteProductResponse{Success: false}, nil
}

// CreateOrder creates a new order
func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	req.Order.OrderDate = timestamppb.Now()
	orderStore[req.Order.Id] = req.Order
	return &pb.OrderResponse{Order: req.Order, Status: "success", Message: "Order created"}, nil
}

// GetOrder fetches an order by its ID
func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	order, exists := orderStore[req.OrderId]
	if !exists {
		return &pb.OrderResponse{Status: "error", Message: "Order not found"}, nil
	}
	return &pb.OrderResponse{Order: order, Status: "success"}, nil
}

// UpdateOrder updates an existing order
func (s *server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	orderStore[req.Order.Id] = req.Order
	return &pb.UpdateOrderResponse{Order: req.Order, Status: "success", Message: "Order updated"}, nil
}

// DeleteOrder deletes an order by its ID
func (s *server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := orderStore[req.OrderId]; exists {
		delete(orderStore, req.OrderId)
		return &pb.DeleteOrderResponse{Success: true}, nil
	}
	return &pb.DeleteOrderResponse{Success: false, Message: "Order not found"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, &server{})

	fmt.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

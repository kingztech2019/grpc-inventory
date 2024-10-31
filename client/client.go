package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/kingztech2019/proto_grpc/proto_inventory/grpc-inventory-proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewInventoryServiceClient(conn)

	createProduct(client)
	getProduct(client, "product-123")
	updateProduct(client)
	deleteProduct(client, "product-123")

	createOrder(client)
	getOrder(client, "order-123")
	updateOrder(client)
	deleteOrder(client, "order-123")
}

// Helper function for formatted JSON output
func printFormattedResponse(prefix string, v interface{}) {
	formatted, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("Failed to format response: %v", err)
	}
	fmt.Printf("%s:\n%s\n", prefix, string(formatted))
}

func createProduct(client pb.InventoryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	product := &pb.Product{
		Id:             "product-123",
		Name:           "Sample Product",
		Price:          9.99,
		InventoryLevel: 100,
	}

	req := &pb.UpdateProductRequest{Product: product}
	res, err := client.UpdateProduct(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create product: %v", err)
	}
	printFormattedResponse("Create Product Response", res)
}

func getProduct(client pb.InventoryServiceClient, productId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetProductRequest{ProductId: productId}
	res, err := client.GetProduct(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get product: %v", err)
	}
	printFormattedResponse("Get Product Response", res)
}

func updateProduct(client pb.InventoryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	product := &pb.Product{
		Id:             "product-123",
		Name:           "Updated Product",
		Price:          12.99,
		InventoryLevel: 150,
	}

	req := &pb.UpdateProductRequest{Product: product}
	res, err := client.UpdateProduct(ctx, req)
	if err != nil {
		log.Fatalf("Failed to update product: %v", err)
	}
	printFormattedResponse("Update Product Response", res)
}

func deleteProduct(client pb.InventoryServiceClient, productId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteProductRequest{ProductId: productId}
	res, err := client.DeleteProduct(ctx, req)
	if err != nil {
		log.Fatalf("Failed to delete product: %v", err)
	}
	printFormattedResponse("Delete Product Response", res)
}

func createOrder(client pb.InventoryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order := &pb.Order{
		Id:        "order-123",
		ProductId: "product-123",
		Quantity:  5,
		OrderDate: timestamppb.Now(),
	}

	req := &pb.CreateOrderRequest{Order: order}
	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}
	printFormattedResponse("Create Order Response", res)
}

func getOrder(client pb.InventoryServiceClient, orderId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetOrderRequest{OrderId: orderId}
	res, err := client.GetOrder(ctx, req)
	if err != nil {
		log.Fatalf("Failed to get order: %v", err)
	}
	printFormattedResponse("Get Order Response", res)
}

func updateOrder(client pb.InventoryServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	order := &pb.Order{
		Id:        "order-123",
		ProductId: "product-123",
		Quantity:  10,
		OrderDate: timestamppb.Now(),
	}

	req := &pb.UpdateOrderRequest{Order: order}
	res, err := client.UpdateOrder(ctx, req)
	if err != nil {
		log.Fatalf("Failed to update order: %v", err)
	}
	printFormattedResponse("Update Order Response", res)
}

func deleteOrder(client pb.InventoryServiceClient, orderId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteOrderRequest{OrderId: orderId}
	res, err := client.DeleteOrder(ctx, req)
	if err != nil {
		log.Fatalf("Failed to delete order: %v", err)
	}
	printFormattedResponse("Delete Order Response", res)
}

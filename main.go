package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	ID     int
	status string
}

func main() {
	orders := generateOrders(20)

	processOrders(orders)

	updateOrderStatuses(orders)

	fmt.Println("All operations completed. Exiting...")
}

// Imagine this to send HTTP Request or some request to another microservice
// where it actually processes the orders
func processOrders(orders []*Order) {
	for _, order := range orders {
		// Generate random int between 0 and 500
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		status := []string{"processing", "shipped", "delivered"}[rand.Intn(3)]
		order.status = status
		fmt.Printf("Updating order %d status: %s\n", order.ID, order.status)
	}
}

// return pointer of orders because we want to manipulate or modify the orders
func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			status: "pending",
		}
	}
	return orders
}

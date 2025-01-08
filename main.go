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

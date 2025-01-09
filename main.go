package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	status string
}

func main() {
	// Imagine it is like Increment or counter
	var wg sync.WaitGroup

	// Initialize 3 size the Increment or counter
	wg.Add(3)

	orders := generateOrders(20)

	// Add increment by 1
	go func() {
		// remove increment by -1
		defer wg.Done()
		processOrders(orders)
	}()

	// Add increment by 1
	go func() {
		// remove increment by -1
		defer wg.Done()
		go updateOrderStatuses(orders)
	}()

	// Add increment by 1
	go func() {
		// remove increment by -1
		defer wg.Done()
		go reportOrderStatus(orders)
	}()

	// Blocks the whole execution until the counter (3) become zero again.
	wg.Wait()

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

// Utility function to print orders
func reportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n---Order Status Report---")
		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.status)
		}
		fmt.Println("------------------------")
	}
}

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
	mux    sync.Mutex
}

// Global variable
var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	// Imagine it is like Increment or counter
	var wg sync.WaitGroup

	// Initialize 3 size the Increment or counter
	wg.Add(3)

	orders := generateOrders(20)

	// Add increment by 1
	//go func() {
	//	// remove increment by -1
	//	defer wg.Done()
	//	processOrders(orders)
	//}()

	for i := 0; i < 3; i++ {
		// Add increment by 1
		go func() {
			// remove increment by -1
			defer wg.Done()
			for _, o := range orders {
				updateOrderStatus(o)
			}
		}()
	}

	// Blocks the whole execution until the counter (3) become zero again.
	wg.Wait()

	reportOrderStatus(orders)

	fmt.Println("All operations completed. Exiting...")
	fmt.Println("Total updates: ", totalUpdates)
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

func updateOrderStatus(order *Order) {
	// Lock to access or update status variable from another Goroutine
	order.mux.Lock()
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
	status := []string{"processing", "shipped", "delivered"}[rand.Intn(3)]
	order.status = status
	fmt.Printf("Updating order %d status: %s\n", order.ID, order.status)
	// Unlock, so the next Goroutine can access or update status variable
	order.mux.Unlock()

	// Lock to access or update totalUpdates variable from another Goroutine
	updateMutex.Lock()
	// Unlock, so the next Goroutine can access or update totalUpdates variable
	defer updateMutex.Unlock()
	time.Sleep(5 * time.Millisecond)
	totalUpdates = totalUpdates + 1
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
	fmt.Println("\n---Order Status Report---")
	for _, order := range orders {
		fmt.Printf("Order %d: %s\n", order.ID, order.status)
	}
	fmt.Println("------------------------")
}

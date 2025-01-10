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

func main() {
	// Imagine it is like Increment or counter
	var wg sync.WaitGroup

	// Initialize 2 size the Increment or counter
	wg.Add(2)

	orderChannel := make(chan *Order)

	// Generating orders in a Goroutine
	// And we want to receive these orders in another Goroutine
	go func() {
		defer wg.Done()
		defer close(orderChannel)
		for _, order := range generateOrders(20) {
			orderChannel <- order
		}
		fmt.Println("Done with generating orders")
	}()

	go processOrders(orderChannel, &wg)

	// Blocks the whole execution until the counter (3) become zero again.
	wg.Wait()

	fmt.Println("All operations completed. Exiting...")
}

// Imagine this to send HTTP Request or some request to another microservice
// where it actually processes the orders
func processOrders(orderChannel <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChannel {
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

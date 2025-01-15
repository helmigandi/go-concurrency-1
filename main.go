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

	// Initialize 3 size the Increment or counter
	wg.Add(3)

	orderChannel := make(chan *Order, 20)
	processChannel := make(chan *Order, 20)

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

	go processOrders(orderChannel, processChannel, &wg)

	go func() {
		defer wg.Done()

		// Infinite loop is there to continuously check the channels, if there are any updates
		for {
			select {
			// Waits until a received value comes from the processChannel
			case processOrder, ok := <-processChannel:
				if !ok {
					// if Channel closed
					fmt.Println("Processor channel closed")
					return
				}
				fmt.Printf("Processing order %d with status %s\n", processOrder.ID, processOrder.status)
			// Function for timeout
			case <-time.After(time.Second * 1):
				// If no other case ready after 10 seconds
				fmt.Println("Time out waiting for processing orders")
				return
			}
		}
	}()

	// Blocks the whole execution until the counter (3) become zero again.
	wg.Wait()

	fmt.Println("All operations completed. Exiting...")
}

// Imagine this to send HTTP Request or some request to another microservice
// where it actually processes the orders
// inChan: Channel that only receive orders
// outChan: Channel that only send the process order to the Channel
func processOrders(inChan <-chan *Order, outChan chan<- *Order, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		// Generate random int between 0 and 500
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		order.status = "processed"

		// Send the order into the out channel
		// Whenever data comes into the input channel (inChan),
		// then this order going to be sent into the out channel (outChan).
		outChan <- order
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

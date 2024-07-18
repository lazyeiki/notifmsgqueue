package main

import (
	"fmt"
	"time"

	"github.com/lazyeiki/notifmsgqueue/notifmsgqueue"
)

func main() {
	// Create a message queue in SingleGoroutine mode.
	singleQueue := notifmsgqueue.New(10, notifmsgqueue.Single, 0)

	// Define a function to handle messages.
	singleHandler := func(msg any) {
		fmt.Printf("Processed message (SingleGoroutine): %v\n", msg)
	}

	// Run the message queue in SingleGoroutine mode.
	singleQueue.Run(singleHandler)

	// Add messages to the queue.
	for i := 0; i < 5; i++ {
		err := singleQueue.Push(fmt.Sprintf("Message %d", i))
		if err != nil {
			fmt.Printf("Failed to add message (SingleGoroutine): %v\n", err)
		}
	}

	// Give some time for the queue to process the messages.
	time.Sleep(2 * time.Second)

	// Stop the message queue in SingleGoroutine mode.
	singleQueue.Stop()

	// Create a message queue in WorkerPool mode.
	workerQueue := notifmsgqueue.New(10, notifmsgqueue.WorkerPool, 3)

	// Define a function to handle messages.
	workerHandler := func(msg any) {
		fmt.Printf("Processed message (WorkerPool): %v\n", msg)
		// time.Sleep(500 * time.Millisecond) // Simulate worker processing time
	}

	// Run the message queue in WorkerPool mode.
	workerQueue.Run(workerHandler)

	// Add messages to the queue.
	for i := 0; i < 15; i++ {
		err := workerQueue.Push(fmt.Sprintf("Message %d", i))
		if err != nil {
			fmt.Printf("Failed to add message (WorkerPool): %v\n", err)
		}
	}

	// Give some time for the queue to process the messages.
	time.Sleep(5 * time.Second)

	// Stop the message queue in WorkerPool mode.
	workerQueue.Stop()
}

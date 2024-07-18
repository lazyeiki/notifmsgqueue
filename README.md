# NotifMsgQueue

## Table of Contents

- [About](#about)
- [for Developers](#for_developer)

## About <a name = "about"></a>
`NotifMsgQueue` is a simple and efficient message queue package for Go that supports both single goroutine and worker pool modes. This package allows you to process messages in a sequential or parallel manner based on your configuration.

## License
This project is licensed under the MIT License.

## for Developers :purple_heart:  <a name = "for_developer"> </a>
To install the package, use the following command:

```sh
go get github.com/lazyeiki/notifmsgqueue
```

### Usage
Here is how you can use the NotifMsgQueue package in your Go project.

### Single Goroutine Mode
In single goroutine mode, messages are processed sequentially by a single goroutine.

```
package main

import (
    "fmt"
    "time"

    "t github.com/lazyeiki/notifmsgqueue"
)

func main() {
    // Create a message queue in SingleGoroutine mode.
    singleQueue := notifmsgqueue.New(10, notifmsgqueue.SingleGoroutine, 0)

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
}
```

### Worker Pool Mode
In worker pool mode, messages are processed in parallel by multiple workers.

```
package main

import (
    "fmt"
    "time"

    "t github.com/lazyeiki/notifmsgqueue"
)

func main() {
    // Create a message queue in WorkerPool mode.
    workerQueue := notifmsgqueue.New(10, notifmsgqueue.WorkerPool, 3)

    // Define a function to handle messages.
    workerHandler := func(msg any) {
        fmt.Printf("Processed message (WorkerPool): %v\n", msg)
        time.Sleep(500 * time.Millisecond) // Simulate worker processing time
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
```

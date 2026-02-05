package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	cc     = flag.Bool("concurrent", false, "If true, run the concurrent function")
	chFlag = flag.Bool("channel", false, "If true, run the concurrent function with channels")
	ch     = make(chan string)
	items  = []string{"task1", "task2", "task3", "task4", "task5"}
)

func sequential() {
	for _, item := range items {
		// Simulate some work with sleep
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Printing sequential task: " + item)
	}
}

func concurrent() {
	for _, item := range items {
		go func(item string) {
			// Simulate some work with sleep
			time.Sleep(100 * time.Millisecond)

			if *chFlag {
				// send the message to the channel
				ch <- "Sent task to channel: " + item
				return
			}

			// print directly
			fmt.Println("Printing concurrent task: " + item)
		}(item)
	}
}

func main() {
	flag.Parse()

	// Log how long did it take.
	defer func(begin time.Time) {
		fmt.Println("Duration:", time.Since(begin))
	}(time.Now())

	if !*cc {
		sequential()
	} else {
		concurrent()
	}

	if !*chFlag {
		return
	}

	// receive the message from the channel
	for range items {
		msg := <-ch
		fmt.Println(msg)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// ============================================================
// MAIN
// ============================================================

func main() {

	// ========================================================
	// SECTION 1: NO CONCURRENCY
	// ========================================================
	//
	// Without the "go" keyword, each function must completely
	// finish before the next function can begin.
	//
	// greet("Nice to meet you!")
	// greet("How are you?")
	// slowGreet("How ... are ... you ...?")
	// greet("I hope you're liking the course!")

	// ========================================================
	// SECTION 2: GOROUTINES + ONE SHARED CHANNEL
	// ========================================================
	//
	// Adding "go" starts each function as a goroutine.
	// This allows the functions to execute concurrently.
	//
	// The done channel allows each goroutine to signal main()
	// when it has finished its work.
	//
	// done := make(chan bool)
	//
	// go greetWithChannel("Nice to meet you!", done)
	// go greetWithChannel("How are you?", done)
	// go slowGreetWithChannel("How ... are ... you ...?", done)
	// go greetWithChannel("I hope you're liking the course!", done)
	//
	// Each <-done waits for one goroutine to send a completion
	// signal through the channel.
	//
	// <-done
	// <-done
	// <-done
	// <-done

	// ========================================================
	// SECTION 3: GOROUTINES + INDIVIDUAL CHANNELS
	// ========================================================
	//
	// Instead of sharing one channel, each goroutine can have
	// its own completion channel.
	//
	// dones := make([]chan bool, 4)
	//
	// dones[0] = make(chan bool)
	// dones[1] = make(chan bool)
	// dones[2] = make(chan bool)
	// dones[3] = make(chan bool)
	//
	// go greetWithChannel("Nice to meet you!", dones[0])
	// go greetWithChannel("How are you?", dones[1])
	// go slowGreetWithChannel("How ... are ... you ...?", dones[2])
	// go greetWithChannel("I hope you're liking the course!", dones[3])
	//
	// Wait for one completion signal from each channel.
	//
	// for _, done := range dones {
	// 	<-done
	// }

	// ========================================================
	// SECTION 4: CLOSING + RANGING OVER A SHARED CHANNEL
	// ========================================================
	//
	// A channel can be ranged over just like a slice.
	//
	// The difference is that:
	//
	//     for range done
	//
	// keeps waiting for new values until the channel is closed.
	//
	// We use a WaitGroup here so that we know ALL goroutines
	// are finished before closing the shared channel.

	done := make(chan bool)

	var wg sync.WaitGroup

	// Tell the WaitGroup that four goroutines are about to run.
	wg.Add(4)

	go greet("Nice to meet you!", done, &wg)
	go greet("How are you?", done, &wg)
	go slowGreet("How ... are ... you ...?", done, &wg)
	go greet("I hope you're liking the course!", done, &wg)

	// Start another goroutine whose job is to wait until all
	// four worker goroutines finish.
	//
	// Once they are all finished, this goroutine closes the
	// channel so the for-range loop below knows when to stop.
	go func() {
		wg.Wait()
		close(done)
	}()

	// Receive values from done until the channel is closed.
	//
	// Every worker sends true when it completes.
	for range done {
	}
}

// ============================================================
// HELPERS — BASIC NON-CONCURRENT VERSIONS
// ============================================================

// basicGreet prints a greeting message. This version does not
// use a channel and uses normal sequential execution.
func basicGreet(phrase string) {
	fmt.Println("Hello!", phrase)
}

// basicSlowGreet waits three seconds before printing a greeting.
func basicSlowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

// ============================================================
// HELPERS — CHANNEL VERSIONS
// ============================================================

// greet prints a greeting and then signals that its goroutine has
// finished. doneChan is shared with main and receives a completion
// signal. wg tracks when this goroutine has completely finished
// executing.
func greet(phrase string, doneChan chan bool, wg *sync.WaitGroup) {
	// Tell the WaitGroup this goroutine is finished when the
	// function exits.
	defer wg.Done()

	fmt.Println("Hello!", phrase)

	// Send a completion signal through the channel.
	doneChan <- true
}

// slowGreet simulates a slower concurrent task before printing its
// greeting. After finishing, it sends a completion signal through
// doneChan. The channel is notT closed here because multiple goroutines
// share it The channel should only be closed after all senders have
// finished.
func slowGreet(phrase string, doneChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate a slow or long-running operation.
	time.Sleep(3 * time.Second)

	fmt.Println("Hello!", phrase)

	// Signal that this goroutine has completed.
	doneChan <- true
}

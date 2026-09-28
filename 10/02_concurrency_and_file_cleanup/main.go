package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

// ============================================================
// MAIN
// ============================================================

func main() {
	// Define the tax rates that we want to process.
	// Each tax rate will be handled as its own concurrent job.
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// Create one completion channel and one error channel
	// for each tax rate/job.
	//
	// doneChans will receive a signal when a job finishes successfully.
	// errorChans will receive an error if a job fails.
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	// ========================================================
	// CREATE AND START CONCURRENT JOBS
	// ========================================================

	for index, taxRate := range taxRates {
		// Create a separate completion channel for this job.
		doneChans[index] = make(chan bool)

		// Create a separate error channel for this job.
		errorChans[index] = make(chan error)

		// Create a file manager for this specific tax rate.
		//
		// Example output filenames:
		// 0%  -> result_0.json
		// 7%  -> result_7.json
		// 10% -> result_10.json
		// 15% -> result_15.json
		fm := filemanager.New(
			"prices.txt",
			fmt.Sprintf("result_%.0f.json", taxRate*100),
		)

		// Create a new price-processing job using:
		// - the file manager
		// - the current tax rate
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		// Start the job concurrently in its own goroutine.
		//
		// The job communicates back to main through:
		// - doneChans[index] if it completes successfully
		// - errorChans[index] if something goes wrong
		go priceJob.Process(doneChans[index], errorChans[index])

		// Old synchronous-style error handling:
		//
		// This is commented out because Process now runs in a goroutine
		// and communicates errors through a channel instead of returning
		// the error directly to main.
		//
		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	// ========================================================
	// WAIT FOR EACH JOB TO FINISH
	// ========================================================

	for index := range taxRates {
		// select waits until one of the listed channel operations
		// becomes ready.
		//
		// For each job, we are waiting for either:
		// 1. an error
		// 2. a successful completion signal
		select {
		case err := <-errorChans[index]:
			// This case runs if the job sends an error.
			if err != nil {
				fmt.Println("Job failed:", err)
			}

		case <-doneChans[index]:
			// This case runs if the job completes successfully.
			fmt.Println("Done.")
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func countReports(numSentCh chan int) int {
	total := 0
	for {
		numReports, ok := <-numSentCh
		if !ok {
			break
		}
		total += numReports
	}
	return total
}

// TEST SUITE - Don't touch below this line

func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(
		numBatches,
		numSentCh,
	)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf(
		"%v reports sent!\n",
		numReports,
	)
	fmt.Println("========================")
}

func main() {
	test(3)
	test(4)
	test(5)
	test(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf(
			"Sent batch of %v reports\n",
			numReports,
		)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

/*
ASSIGNMENT
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the emails it's given. It should then write the emails to the channel in order, and finally return the channel.
*/

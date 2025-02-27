package main

import (
	"fmt"
	"time"
)

func concurrrentFib(n int) {
	chInt := make(chan int)
	go func() {
		fibonacci(
			n,
			chInt,
		)
	}()
	for v := range chInt {
		fmt.Println(v)
	}
}

// TEST SUITE - Don't touch below this line

func test(n int) {
	fmt.Printf(
		"Printing %v numbers...\n",
		n,
	)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func main() {
	test(10)
	test(5)
	test(20)
	test(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

/*
ASSIGNMENT
It's that time again, Mailio is hiring and we've been assigned to do the interview. The Fibonacci sequence is Mailio's interview problem of choice. We've been tasked with building a small toy program we can use in the interview.

Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci concurrently
Use a range loop to read from the channel and append the values to a slice
Return the slice
*/

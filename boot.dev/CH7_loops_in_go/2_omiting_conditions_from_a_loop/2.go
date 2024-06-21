package main

import "fmt"

func maxMessages(thresh int) int {
	totalCost := 0
	for i := 0; ; i++ {
		totalCost += 100 + i
		if totalCost > thresh {
			return i
		}
	}
}
func test(thresh int) {
	fmt.Printf(
		"Max messages for %v pennies: %v\n",
		thresh,
		maxMessages(thresh),
	)
}
func main() {
	test(100)
	test(200)
	test(300)
}

/*
ASSIGNMENT
Complete the maxMessages function. Given a cost threshold, it should calculate the maximum number of messages that can be sent.

Each message costs 100 pennies, plus an additional fee. The fee structure is:

1st message: 100 + 0
2nd message: 100 + 1
3rd message: 100 + 2
4th message: 100 + 3
BROWSER FREEZE
If you lock up your browser by creating an infinite loop that isn't breaking, just click the cancel button.
*/

package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf(
		"SMS that costs $%.2f to be sent to '%s' can not be sent",
		cost,
		recipient,
	)
}
func test(cost float64, recipient string) {
	s := getSMSErrorString(
		cost,
		recipient,
	)
	fmt.Println(s)
	fmt.Println("----------------")
}
func main() {
	test(
		0.0,
		"1234567890",
	)
	test(
		1.0,
		"1234567890",
	)
	test(
		1.23,
		"1234567890",
	)
	test(
		1.234,
		"1234567890",
	)
	test(
		1.2345,
		"1234567890",
	)
	test(
		1.23456,
		"123",
	)
}

/*
We need better error logs for our backend developers to help them debug their code.

Complete the getSMSErrorString() function. It should return a string with this format:

SMS that costs $COST to be sent to 'RECIPIENT' can not be sent
Copy icon
COST is the cost of the SMS, always showing the price formatted to 2 decimal places.
RECIPIENT is the stringified representation of the recipient's phone number
Be sure to include the $ symbol and the single quotes
*/

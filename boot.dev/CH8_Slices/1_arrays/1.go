package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	return [3]string{primary, secondary, tertiary}, [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
}

func main() {
	// Test cases
	s, l := getMessageWithRetries(
		"hello",
		"world",
		"!",
	)
	fmt.Printf(
		"getMessageWithRetries(hello, world, !) = %v, %v \n",
		s,
		l,
	) //expected: getMessageWithRetries(hello, world, !) = [hello world !] [5 10 11]
}

/*
ASSIGNMENT
When our clients don't respond to a message, they can be reminded with up to 2 additional messages. getMessageWithRetries returns:

An array of 3 strings
An array of 3 integers
The strings are just the original messages structured as an array. The first is the primary message, the second is the first reminder, and the third is the last reminder. The integers represent the cost of sending each message.

The cost of a message is equal to the length of the message, plus the cost of any previous messages. For example:

"hello" costs 5
"world" costs 10
"!" costs 11
*/

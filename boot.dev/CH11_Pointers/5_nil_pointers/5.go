package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(
		messageVal,
		"fubb",
		"****",
	)
	messageVal = strings.ReplaceAll(
		messageVal,
		"shiz",
		"****",
	)
	messageVal = strings.ReplaceAll(
		messageVal,
		"witch",
		"*****",
	)
	*message = messageVal
}
func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

// test for getting nil pointer
func testNil(messages []string) {
	for _, _ = range messages {
		var messagePointer *string
		removeProfanity(messagePointer)
		fmt.Println(messagePointer)
	}
}
func main() {
	messages := []string{
		"Hello, fubb",
		"shiz, it's cold",
		"witch, please",
	}
	test(messages)
	testNil(messages)
}

/*
NIL POINTERS
Pointers can be very dangerous.

If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.

ASSIGNMENT
Let's make our profanity checker safe. Update the removeProfanity function. If message is nil, return early to avoid a panic. After all, there are no bad words to remove.
*/

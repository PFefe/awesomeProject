package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
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
func main() {
	messages := []string{
		"Hello, fubb",
		"shiz, it's cold",
		"witch, please",
	}
	test(messages)
}

/*
UST BECAUSE YOU CAN DOESN'T MEAN YOU SHOULD
We're doing this exercise to understand that pointers can be used in this way. That said, pointers can be very dangerous. It's generally a better idea to have your functions accept non-pointers and return new values rather than mutating pointer inputs.

ASSIGNMENT
Complete the removeProfanity function.

It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

"fubb" -> "****"
"shiz" -> "****"
"witch" -> "*****"
It should mutate the value in the pointer and return nothing.

Do not alter the function signature.
*/

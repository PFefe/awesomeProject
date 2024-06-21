package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

func test(status string) {
	err := validateStatus(status)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	println("status is valid")
}

// create a string repeat function to repeat a string which is coming before it
func repeat(s string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}
func main() {
	test("")
	test("hello")
	test("hello world")
	test("hello world, how are you doing today?, long message long message long message long message long message long message long message long message long message long message long message long message ")
	test(
		repeat(
			" long message ",
			10,
		),
	)
}

/*
USER INPUT
In Textio, users can set their profile status to communicate their current activity to those that choose to read it... However, there are some restrictions on what these statuses can contain. Your task is to implement a function that validates a user's status update. The status update cannot be empty and must not exceed 140 characters.

ASSIGNMENT
Complete the validateStatus function. It should return an error when the status update violates any of the rules:

If the status is empty, return an error that reads status cannot be empty.
If the status exceeds 140 characters, return an error that says status exceeds 140 characters.
*/

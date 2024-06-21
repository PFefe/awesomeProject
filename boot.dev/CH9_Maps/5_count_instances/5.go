package main

import "fmt"

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, exists := validUsers[user]; exists {
			validUsers[user]++
		}
	}
}

func test(messagedUsers []string, validUsers map[string]int) {
	getCounts(
		messagedUsers,
		validUsers,
	)
	for k, v := range validUsers {
		fmt.Println(
			k,
			v,
		)
	}
}

func main() {
	// Test cases
	test(
		[]string{"eddie", "eddie", "eddie", "jennie", "jennie", "jennie", "jennie", "jennie"},
		map[string]int{},
	)
	fmt.Println("---")
	test(
		[]string{"eddie", "eddie", "eddie", "jennie", "jennie", "jennie", "jennie", "jennie"},
		map[string]int{"eddie": 0, "jennie": 0},
	)
}

/*
ASSIGNMENT
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the getCounts function. It takes a slice of strings messagedUsers and a map of string -> int validUsers. It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/

/*
ASSIGNMENT
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many messages each user received.

Implement the getCounts function. It takes a slice of strings messagedUsers and a map of string -> int validUsers. It should update the validUsers map with the number of times each user has received a message. Each string in the slice is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/

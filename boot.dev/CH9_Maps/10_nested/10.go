package main

import "fmt"

// getNameCounts function
func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		first := rune(name[0])
		if _, ok := counts[first]; !ok {
			counts[first] = make(map[string]int)
		}
		counts[first][name]++
	}
	return counts
}

// test function
func test(names []string, expected map[rune]map[string]int) {
	counts := getNameCounts(names)
	for k, v := range expected {
		for name, count := range v {
			actualCount, ok := counts[k][name]
			if !ok {
				fmt.Printf(
					"Error: Key '%c' with name '%s' not found\n",
					k,
					name,
				)
			} else if actualCount != count {
				fmt.Printf(
					"Error: Key '%c' with name '%s' expected %d but got %d\n",
					k,
					name,
					count,
					actualCount,
				)
			} else {
				fmt.Printf(
					"Success: Key '%c' with name '%s' has correct count %d\n",
					k,
					name,
					count,
				)
			}
		}
	}
	fmt.Println("--------------------------------")
}

// main function
func main() {
	// Test cases
	test(
		[]string{"billy", "billy", "bob", "joe"},
		map[rune]map[string]int{
			'b': {
				"billy": 2,
				"bob":   1,
			},
			'j': {
				"joe": 1,
			},
		},
	)
	test(
		[]string{"alice", "alice", "bob", "bob", "bob", "charlie"},
		map[rune]map[string]int{
			'a': {
				"alice": 2,
			},
			'b': {
				"bob": 3,
			},
			'c': {
				"charlie": 1,
			},
		},
	)
}

/*
NESTED
Maps can contain maps, creating a nested structure. For example:

map[string]map[string]int
map[rune]map[string]int
map[int]map[string]map[string]int
Copy icon
ASSIGNMENT
Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

Complete the getNameCounts function. It takes a slice of strings (names) and returns a nested map where the first key is all the unique first characters (runes) of the names, the second key is all the names themselves, and the value is the count of each name.

For example:

billy
billy
bob
joe
Copy icon
Creates the following nested map:

b: {
    billy: 2,
    bob: 1
},
j: {
    joe: 1
}
Copy icon
Note: the test suite is not printing the map you're returning directly, but instead checking some specific keys.




*/

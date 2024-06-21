package main

func countDistinctWords(messages []string) int {
	// ?
}

/*
DISTINCT WORDS
Complete the countDistinctWords function. It should take a slice of strings and return the total count of distinct words across all the strings. Assume words are separated by spaces. Casing should not matter. (e.g., "Hello" and "hello" should be considered the same word).

For example:

messages := []string{"Hello world", "hello there", "General Kenobi"}
count := countDistinctWords(messages)
Copy icon
count should be 5 as the distinct words are "hello", "world", "there", "general" and "kenobi" irrespective of casing.

HINT
Go's strings package can be very helpful here. Specifically the Fields method and the ToLower method.




*/

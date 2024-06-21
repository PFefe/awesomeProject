package main

func getMessageCosts(messages []string) []float64 {
	costs := make(
		[]float64,
		len(messages),
	)
	for i := 0; i < len(messages); i++ {
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}

func testMessageCosts() {
	messages := []string{
		"hello",
		"world",
		"!",
		"👋",
		"🌍",
	}
	costs := getMessageCosts(messages)
	if len(costs) != 2 {
		panic("❌")
	}
	if costs[0] != 0.05 {
		panic("❌")
	}
	if costs[1] != 0.05 {
		panic("❌")
	}
	println("✅")
}
func main() {
	testMessageCosts()
}

/*
ASSIGNMENT
We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.
*/

package main

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}
	if actualCostInPennies > float64(maxCostInPennies) {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}
func test(costMultiplier float64, maxCostInPennies int) {
	maxMessagesToSend := getMaxMessagesToSend(
		costMultiplier,
		maxCostInPennies,
	)
	println(maxMessagesToSend)
}
func main() {
	test(
		2.0,
		100,
	)
	test(
		2.0,
		200,
	)
	test(
		2.0,
		300,
	)
}

/*
ASSIGNMENT
We have an interesting new cost structure from our SMS vendor. They charge exponentially more money for each consecutive text we send! Let's write a function that calculates how many messages we can send in a given batch given a costMultiplier and a maxCostInPennies.

In a nutshell, the first message costs a penny, and each message after that costs the same as the previous message multiplied by the costMultiplier.

There is a bug in the code! Add a condition to the for loop to fix the bug. The loop should stop when actualCostInPennies is greater than or equal to the maxCostInPennies.
*/

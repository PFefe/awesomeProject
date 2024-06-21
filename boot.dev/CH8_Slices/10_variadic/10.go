package main

import "fmt"

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]

	}
	return total
}

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf(
		"Summing %v costs...\n",
		len(nums),
	)
	fmt.Printf(
		"Bill for the month: %.2f\n",
		total,
	)
	fmt.Println("=END REPORT=")
}
func main() {
	test(
		1.0,
		2.0,
		3.0,
	)
	test(
		4.0,
		5.0,
		6.0,
		7.0,
	)
	test(
		8.0,
		9.0,
		10.0,
		11.0,
		12.0,
	)

}

/*
ASSIGNMENT
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.
*/

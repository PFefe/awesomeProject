package main

import "fmt"

func getMonthlyPrice2(tier string) int {
	switch tier {
	case "basic":
		return 10000 // $100.00 in pennies
	case "premium":
		return 15000 // $150.00 in pennies
	case "enterprise":
		return 50000 // $500.00 in pennies
	default:
		return 0
	}
}
func main() {
	test2("basic")
	test2("premium")
	test2("enterprise")
}

func test2(s1 string) {
	fmt.Println(getMonthlyPrice2(s1))
}

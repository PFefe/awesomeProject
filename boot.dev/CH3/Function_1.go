package main

import "fmt"

func concat(s1, s2 string) string {
	return s1 + s2
}

func getMonthlyPrice(tier string) int {
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

// don't touch below this line

func main() {
	test(
		"Lane,",
		" happy birthday!",
	)
	test(
		"Elon,",
		" hope that Tesla thing works out",
	)
	test(
		"Go",
		" is fantastic",
	)
	test1("basic")
	test1("premium")
	test1("enterprise")
}

func test(s1 string, s2 string) {
	fmt.Println(
		concat(
			s1,
			s2,
		),
	)
}
func test1(s1 string) {
	fmt.Println(getMonthlyPrice(s1))
}

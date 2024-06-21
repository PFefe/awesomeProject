package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for len(costByDay) <= cost.day {
			costByDay = append(
				costByDay,
				0.0,
			)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}
func test(costs []cost) {
	fmt.Printf(
		"Costs: %v\n",
		costs,
	)
	fmt.Printf(
		"Costs by day: %v\n",
		getCostsByDay(costs),
	)
}
func main() {
	test(
		[]cost{
			{0, 4.0},
			{1, 2.1},
			{5, 2.5},
			{1, 3.1},
		},
	)
}

/*
ASSIGNMENT
We've been asked to "bucket" costs per day, in a given period.

Complete the getCostsByDay() function using the append() function. It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the daily costs slice should be the number of days contained in the costs slice, up to and including the last day. There should only be one "bucket" of costs per day. Be sure to include entries for intermediate days, even if they don't have any costs.

Days are numbered and start at 0.

EXAMPLE
Given this input:

[]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
}
Copy icon
We expect this result:

[]float64{
    4.0, // first day
    5.2, // 2.1 + 3.1
    0.0, // intermediate days with no costs
    0.0, // ...
    0.0, // ...
    2.5, // last day
}
Copy icon



*/

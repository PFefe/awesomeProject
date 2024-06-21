package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	} else if plan == planFree {
		return messages[:2], nil
	} else {
		return nil, errors.New("unsupported plan")
	}
}

func main() {
	messages := [3]string{"retry 1", "retry 2", "retry 3"}

	plan := planFree
	fmt.Println(
		getMessageWithRetriesForPlan(
			plan,
			messages,
		),
	) // [retry 1 retry 2] <nil>

	plan = planPro
	fmt.Println(
		getMessageWithRetriesForPlan(
			plan,
			messages,
		),
	) // [retry 1 retry 2 retry 3] <nil>

	plan = "unsupported"
	fmt.Println(
		getMessageWithRetriesForPlan(
			plan,
			messages,
		),
	) // [] unsupported plan
}

/*
ASSIGNMENT
Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the getMessageWithRetriesForPlan function. It takes a plan variable as input as well as an array of 3 messages. You've been provided with constants representing the plan types at the top of the file.

If the plan is a pro plan, return all the strings from the messages input in a slice.
If the plan is a free plan, return the first 2 strings from the messages input in a slice.
If the plan isn't either of those, return a nil slice and an error that says unsupported plan.
*/

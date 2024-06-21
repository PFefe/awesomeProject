package main

import (
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance < 0.0 {
		return nil, 0.0, fmt.Errorf("insufficient funds")
	}
	newItems := append(
		oldItems,
		newItem,
	)
	return newItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf(
		"%s subscription",
		s.interval,
	)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf(
		"one time usage plan with %v emails",
		otup.numEmailsAllowed,
	)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	test(
		subscription{
			userEmail: "john@example.com",
			startDate: time.Now().UTC(),
			interval:  "yearly",
		},
		[]subscription{},
		1000.00,
	)
	test(
		subscription{
			userEmail: "jane@example.com",
			startDate: time.Now().UTC(),
			interval:  "monthly",
		},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	test(
		oneTimeUsagePlan{
			userEmail:        "dillon@example.com",
			numEmailsAllowed: 5000,
		},
		[]oneTimeUsagePlan{},
		756.20,
	)
	test(
		oneTimeUsagePlan{
			userEmail:        "dalton@example.com",
			numEmailsAllowed: 100000,
		},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)
}

func test[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf(
		"Charging customer for a '%s', current balance is %v...\n",
		newItem.GetName(),
		balance,
	)
	newItems, newBalance, err := chargeForLineItem(
		newItem,
		oldItems,
		balance,
	)
	if err != nil {
		fmt.Printf(
			"Got error: %v\n",
			err,
		)
		return
	}
	fmt.Printf(
		"New balance is: %v. Total number of line items is now %v\n",
		newBalance,
		len(newItems),
	)
}

/*
ASSIGNMENT
We have different kinds of "line items" that we charge our customer's credit cards for. Line items can be things like "subscriptions" or "one-time payments" for email usage.

Complete the chargeForLineItem function. First, it should check if the user has a balance with enough funds to be able to pay for the cost of the newItem. If they don't, then return an "insufficient funds" error and zero values for the other return values.

If they do have enough funds:

Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.
*/

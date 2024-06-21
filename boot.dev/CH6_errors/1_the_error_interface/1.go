package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costForCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, fmt.Errorf(
			"error sending message to customer: %v",
			err,
		)
	}
	costForSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, fmt.Errorf(
			"error sending message to spouse: %v",
			err,
		)
	}
	return costForSpouse + costForCustomer, nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf(
			"can't send texts over %v characters",
			maxTextLen,
		)
	}
	return costPerChar * len(message), nil
}

func main() {
	cost, err := sendSMSToCouple(
		"hello",
		"world",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(
		"total cost:",
		cost,
	)
	//send long sms to get error
	cost, err = sendSMSToCouple(
		"hello, this is a long message",
		"world",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(
		"total cost:",
		cost,
	)
}

package main

import "fmt"

func monthlyBillIncrease(costPerSend, lastMonthBill, thisMonthBill int) int {

	getBillForMonth(
		costPerSend,
		lastMonthBill,
	)
	getBillForMonth(
		costPerSend,
		thisMonthBill,
	)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

func main() {
	// Call monthlyBillIncrease with sample arguments
	result := monthlyBillIncrease(
		10,
		100,
		120,
	)
	fmt.Println(
		"Monthly bill increase:",
		result,
	)
}

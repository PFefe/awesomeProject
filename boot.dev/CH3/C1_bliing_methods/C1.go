package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(
		costPerMessage,
		numMessages,
	)
	discount := calculateDiscount(numMessages)
	return baseBill - (baseBill * discount)
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent >= 5000 {
		return 0.2
	}
	if messagesSent >= 1000 {
		return 0.1
	}
	return 0
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main() {
	// Call calculateFinalBill with sample arguments
	result := calculateFinalBill(
		0.1,
		100,
	)
	println(
		"Final bill:",
		result,
	)
}

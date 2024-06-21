package main

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	price := calcPrice(
		productID,
		quantity,
	)
	if price == 0 {
		return false, 0
	}
	if accountBalance < price {
		return false, 0
	}
	if amountInStock(productID) < quantity {
		return false, 0
	}
	return true, accountBalance - price
}

// Don't touch below this line

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}

func main() {
	// Call placeOrder with sample arguments
	success, newBalance := placeOrder(
		"8",
		2,
		150.00,
	)
	println(
		"Order successful:",
		success,
	)
	println(
		"New balance:",
		newBalance,
	)
}

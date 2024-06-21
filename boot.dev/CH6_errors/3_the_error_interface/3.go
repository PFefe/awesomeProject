package main

import "fmt"

type divideError struct {
	dividend float64
}

func (de divideError) Error() string {
	return fmt.Sprintf(
		"can not divide %v by zero",
		de.dividend,
	)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
func main() {
	result, err := divide(
		1,
		0,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

/*
ASSIGNMENT
Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the divideError type implements the error interface. Its Error() method should just return a string formatted in the following way:

can not divide DIVIDEND by zero
Copy icon
Where DIVIDEND is the actual dividend of the divideError. Use the %v verb to format the dividend as a float.
*/

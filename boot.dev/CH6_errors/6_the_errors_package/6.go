package main

import "errors"

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
}

func test(x, y float64) {
	result, err := divide(
		x,
		y,
	)
	if err != nil {
		println(err.Error())
		return
	}
	println(result)
}

func main() {
	test(
		1,
		0,
	)
	test(
		1,
		1,
	)
	test(
		1,
		2,
	)
	test(
		1,
		3,
	)
	test(
		1,
		4,
	)
	test(
		1,
		5,
	)
}

/*
ASSIGNMENT
Textio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string no dividing by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads "no dividing by 0".

HINT
Remember that it's conventional to return the "zero" values of all other return values when you return a non-nil error in Go.
*/

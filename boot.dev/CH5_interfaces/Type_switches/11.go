package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch e := e.(type) {
	case email:
		return e.toAddress, e.cost()
	case sms:
		return e.toPhoneNumber, e.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
func main() {
	// Call getExpenseReport with sample arguments
	e := email{
		isSubscribed: true,
		body:         "Hello, World!",
		toAddress:    "kgu",
	}
	s := sms{
		isSubscribed:  true,
		body:          "Hello, World!",
		toPhoneNumber: "123456789",
	}
	fmt.Println(getExpenseReport(e))
	fmt.Println(getExpenseReport(s))
	fmt.Println(getExpenseReport(invalid{}))
}

/*
fmt.Printf("%T\n", v) prints the type of a variable.

ASSIGNMENT
After submitting our last snippet of code for review, a more experienced gopher told us to use a type switch instead of successive assertions. Let's make that improvement!

Implement the getExpenseReport function using a type switch.

If the expense is an email then it should return the email's toAddress and the cost of the email.
If the expense is an sms then it should return the sms's toPhoneNumber and its cost.
If the expense has any other underlying type, just return an empty string and 0.0 for the cost.
*/

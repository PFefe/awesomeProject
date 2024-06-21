package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	if e.isSubscribed {
		return fmt.Sprintf(
			"'%s' | Subscribed",
			e.body,
		)
	}
	return fmt.Sprintf(
		"'%s' | Not Subscribed",
		e.body,
	)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "Hello, World!",
	}
	fmt.Println(e.cost())
	fmt.Println(e.format())
	println("---------------------")
	e = email{
		isSubscribed: false,
		body:         "Hello, World!",
	}
	fmt.Println(e.cost())
	fmt.Println(e.format())
}

/*
ASSIGNMENT
Complete the required methods so that the email type implements both the expense and formatter interfaces.

COST()
If the email is not "subscribed", then the cost is 5 cents for each character in the body. If it is, then the cost is 2 cents per character.

Return the total cost of the entire email in cents.

FORMAT()
The format method should return a string in this format:

'CONTENT' | Subscribed
Copy icon
If the email is not subscribed, change the second part to "Not Subscribed":

'CONTENT' | Not Subscribed
Copy icon
The single quotes are included in the string, and CONTENT is the email's body. For example:

'Hello, World!' | Subscribed
*/

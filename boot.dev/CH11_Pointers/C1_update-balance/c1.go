package main

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?

/*
UPDATE BALANCE
Textio's needs a new way to update users account balance.

ASSIGNMENT
Implement the updateBalance function it should take a pointer to a customer, and a transaction. Depending on the transactionType it should either add or subtract the amount from the customers balance. If the customer does not have enough money it should return the error insufficient funds, or if the transactionType that is passed in isn't a withdrawal or deposit it should return the error unknown transaction type. Otherwise it should process the transaction and return nil.

alice := customer{id: 1, balance: 100.0},
deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

updateBalance(&alice, deposit)
// id: 1 balance: 150
Copy icon
TIP
When your function receives a pointer to a struct, you might try to access a field like this and encounter an error:

balance := *customer.balance
Copy icon
Instead, use this simple approach using a selector expression.

balance := customer.balance
Copy icon
This approach is shorthand for (*customer).balance and is the recommended, simplest way to access struct fields in Go.
*/

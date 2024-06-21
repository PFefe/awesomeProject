package main

func SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= 100 {
		return message, true
	}
	return "", false
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

func (u User) SendMessage(s string, i int) (string, bool) {
	if i <= u.MessageCharLimit {
		return s, true
	}
	return "", false
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}
	if membershipType == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membershipType == TypePremium {
		membership.MessageCharLimit = 1000
	}
	return User{Name: name, Membership: membership}
}

func main() {
	// Call newUser with sample arguments
	result := newUser(
		"John",
		TypePremium,
	)
	println(
		"User name:",
		result.Name,
	)
	println(
		"Membership type:",
		result.Type,
	)
	println(
		"Message char limit:",
		result.MessageCharLimit,
	)
	println("---------------------")
	result2 := newUser(
		"Jane",
		TypeStandard,
	)
	println(
		"User name:",
		result2.Name,
	)
	println(
		"Membership type:",
		result2.Type,
	)
	println(
		"Message char limit:",
		result2.MessageCharLimit,
	)
	println("---------------------")
	// Call SendMessage with sample arguments
	message, canSend := result.SendMessage(
		"Hello, World!",
		1001,
	)
	println(
		"Message:",
		message,
	)
	println(
		"Can send:",
		canSend,
	)
}

/*
Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the messagelength is within the user's message character limit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.
*/

package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType membershipType) User {
	if membershipType == TypePremium {
		return User{
			Name: name,
			Membership: Membership{
				Type:             membershipType,
				MessageCharLimit: 1000,
			},
		}
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}
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
}

/*
Create a new struct called Membership, it should have:

A Type field with the pre-defined membershipType string type.
A MessageCharLimit integer field.
Update the User struct to embed a Membership.

Complete the newUser function. It should return a new User with all the fields set as you would expect based on the inputs. If the user is premium, the MessageCharLimit should be 1000, otherwise, it should only be 100.
*/

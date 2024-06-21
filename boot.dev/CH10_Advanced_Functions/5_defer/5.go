package main

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(
		users,
		name,
	)
	user, ok := users[name]
	if !ok {

		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

// tests for the logAndDelete function
func deleteTests() {
	tests := []struct {
		name     string
		expected string
	}{
		{"jane (not admin)", logDeleted},
		{"john (admin)", logAdmin},
		{"missing (not found)", logNotFound},
	}
	failed := false
	for _, test := range tests {
		users := map[string]user{
			"jane": user{name: "jane", number: 1234, admin: false},
			"john": user{name: "john", number: 5678, admin: true},
		}
		log := logAndDelete(
			users,
			test.name,
		)
		if log != test.expected {
			println(
				"Test failed for user:",
				test.name,
			)
			failed = true
		}
	}
	if failed {
		println("Some tests failed")
	} else {
		println("All tests passed")
	}
}
func main() {
	deleteTests()
}

/*
ASSIGNMENT
There is a bug in the logAndDelete function, fix it!

This function should always delete the user from the user's map, which is a map that stores the user's name as keys. It also returns a log string that indicates to the caller some information about the user's deletion.

To avoid bugs like this in the future, instead of calling delete before each return, just defer the delete once at the beginning of the function.
*/

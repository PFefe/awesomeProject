package main

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasUpper := false
	hasNumber := false
	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
		if c >= '0' && c <= '9' {
			hasNumber = true
		}
	}
	return hasUpper && hasNumber
}

func test(password string, expected bool) {
	if isValidPassword(password) == expected {
		println("PASS")
	} else {
		println("FAIL")
	}
}

func main() {
	test(
		"abcde",
		false,
	)
	test(
		"abcdef",
		false,
	)
	test(
		"abcde1",
		false,
	)
	test(
		"abcdeA",
		false,
	)
	test(
		"abcdeA1",
		true,
	)
	test(
		"abcdeA1fghijklmnop",
		false,
	)
	test(
		"abcdeA1fgh",
		true,
	)
	test(
		"abcdeA1fghi",
		true,
	)
}

/*
PASSWORD STRENGTH
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one number.
ASSIGNMENT
Implement the isValidPassword function. Use a loop to inspect each character in the password string to check for its length, and the presence of an uppercase letter and a number.
*/

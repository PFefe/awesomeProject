package main

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(
			formattedMessages,
			formatter(message),
		)
	}
	return formattedMessages
}

/*
ASSIGNMENT
Textio is launching a new email messaging product, "Mailio"!

Fix the compile-time bug in the getFormattedMessages function. The function body is correct, but the function signature is not.
*/

package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(
		`
To: %v
Message: %v
`,
		m.Recipient,
		m.Text,
	)
}
func test() {
	m := Message{
		Recipient: "John",
		Text:      "Hello",
	}
	fmt.Println(getMessageText(m))
}
func main() {
	test()
}

/*
ASSIGNMENT
Fix the bug in the getMessageText function. It's supposed to return a nicely formatted message to the console containing an SMS's recipient and message body. However, it's not working as expected. Run the code and see what happens, then fix the bug.
*/

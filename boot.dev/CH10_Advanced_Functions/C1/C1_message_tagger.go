package main

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
}

/*
MESSAGE TAGGER
Textio needs a way to tag messages based on specific criteria

ASSIGNMENT
Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
It should loop through each message and set the tags to the result of the passed in function.
Complete the tagger function. It should take a sms message and return a slice of strings.
For any message that contains "urgent" in the content the Urgent tag should be applied.
For any message that contains "sale" in the content the Promo tag should be applied.
Example usage:

messages := []sms{
	{id: "001", content: "Urgent! Last chance to see!"},
	{id: "002", content: "Big sale on all items!"},
	// Additional messages...
}
taggedMessages := tagMessages(messages, tagger)
// `taggedMessages` will now have tags based on the content.
// 001 = [Urgent]
// 002 = [Promo]
Copy icon
TIP
The go strings package, specifically the contains method can be very helpful here!




*/

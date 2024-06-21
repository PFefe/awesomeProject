package main

func processMessages(messages []string) []string {
	// ?
}

/*
PROCESS MESSAGES
Textio needs to speed up message processing using concurrency.

ASSIGNMENT
Implement the processMessages function, it takes a slice of messages. It should process each message concurrently within a goroutine. The processing for each message is simple: append -processed to the content. Use a channel to ensure that all messages are processed and collected correctly then return the slice of processed messages.

Example output:

messages: []string{"Here are some messsages", "Here is another", "and another"}
processedMessages := processMessages(messages)
fmt.Printf(processedMessages)
// prints ["Here are some messsages-processed", "Here is another-processed", "and another-processed"]
*/

package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var filteredMessages []Message
	for _, message := range messages {
		if message.Type() == filterType {
			filteredMessages = append(
				filteredMessages,
				message,
			)
		}
	}
	return filteredMessages
}

func testFilterMessages(messages []Message, filterType string) {
	filteredMessages := filterMessages(
		messages,
		filterType,
	)
	for _, message := range filteredMessages {
		switch message.(type) {
		case TextMessage:
			tm := message.(TextMessage)
			println(
				"TextMessage: ",
				tm.Content,
			)
		case MediaMessage:
			mm := message.(MediaMessage)
			println(
				"MediaMessage: ",
				mm.Content,
			)
		case LinkMessage:
			lm := message.(LinkMessage)
			println(
				"LinkMessage: ",
				lm.Content,
			)
		}
	}
}

func main() {
	messages := []Message{
		TextMessage{"John", "Hello, how are you?"},
		MediaMessage{"Jane", "image", "Check out this image"},
		LinkMessage{"John", "http://example.com", "Click here"},
		TextMessage{"Jane", "I'm good, thanks!"},
		MediaMessage{"John", "video", "Watch this video"},
		LinkMessage{"Jane", "http://example.com", "Click here"},
	}

	testFilterMessages(
		messages,
		"text",
	)
	testFilterMessages(
		messages,
		"media",
	)
	testFilterMessages(
		messages,
		"link",
	)

}

/*
MESSAGE FILTER
Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

ASSIGNMENT
Your task is to implement a function that filters a slice of messages based on the message type.

Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.
*/

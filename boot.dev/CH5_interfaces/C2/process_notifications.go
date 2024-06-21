package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}
func (gm groupMessage) importance() int {
	return gm.priorityLevel
}
func (sa systemAlert) importance() int {
	return 100
}
func processNotification(n notification) (string, int) {
	switch n := n.(type) {
	case directMessage:
		return n.senderUsername, n.importance()
	case groupMessage:
		return n.groupName, n.importance()
	case systemAlert:
		return n.alertCode, n.importance()
	}
	return "", 0
}
func main() {
	// Call processNotification with sample arguments
	dm := directMessage{
		senderUsername: "John",
		messageContent: "Hello, World!",
		priorityLevel:  1,
		isUrgent:       true,
	}
	gm := groupMessage{
		groupName:      "Team",
		messageContent: "Standup meeting at 3pm",
		priorityLevel:  2,
	}
	sa := systemAlert{
		alertCode:      "1001",
		messageContent: "Server down",
	}
	fmt.Println(processNotification(dm))
	fmt.Println(processNotification(gm))
	fmt.Println(processNotification(sa))
}

/*
ASSIGNMENT
Implement the importance methods for each message type. They should return the importance score for each message type.
For a directMessage the importance score is based on if the message isUrgent or not. If it is the importance score is 50 otherwise the importance score is equal to the DM's priorityLevel.
For a groupMessage the importance score is based on the messages priorityLevel
All systemAlert messages should return a 100 importance score.
Complete the processNotification function. It should identify the type and return different values for each type
For a directMessage, return the sender's username and importance score.
For a groupMessage, return the group's name and the importance score.
For an systemAlert, return the alert code and the importance score.
If the notification does not match any known type, return an empty string and a score of 0.
*/

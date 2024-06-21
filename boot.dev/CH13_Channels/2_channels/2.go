package main

import (
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [4]email) [4]bool {
	isOldChan := make(chan bool)
	go func() {
		sendIsOld(
			isOldChan,
			emails,
		)
	}()
	isOld := [4]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	isOld[3] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [4]email) {
	for _, e := range emails {
		if e.date.Before(
			time.Date(
				2020,
				0,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}
func main() {
	emails := [4]email{
		{
			body: "Hello there Kaladin!",
			date: time.Now(),
		},
		{
			body: "Hi there Shallan!",
			date: time.Now(),
		},
		{
			body: "Hey there Dalinar!",
			date: time.Now(),
		},
		// add old email
		{
			body: "Hey there Dalinar!",
			date: time.Date(
				2019,
				0,
				0,
				0,
				0,
				0,
				0,
				time.UTC,
			),
		},
	}

	isOld := checkEmailAge(emails)
	for i, old := range isOld {
		if old {
			println(
				"Email",
				i,
				"is old",
			)
		} else {
			println(
				"Email",
				i,
				"is not old",
			)
		}
	}
}

/*
ASSIGNMENT
Run the program. You'll see that it deadlocks and never exits. The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.
*/

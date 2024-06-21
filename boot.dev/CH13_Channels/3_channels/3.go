package main

import (
	"fmt"
	"time"
)

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

// don't touch below this line

func test(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf(
		"Waiting for %v databases...\n",
		numDBs,
	)
	waitForDbs(
		numDBs,
		dbChan,
	)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func main() {
	test(3)
	test(4)
	test(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf(
				"Database %v is online\n",
				i+1,
			)
		}
	}()
	return ch
}

/*
ASSIGNMENT
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to waitForDBs stays blocked.

Fix the waitForDBs function. It should block until it receives numDBs tokens on the dbChan channel. Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you.
*/

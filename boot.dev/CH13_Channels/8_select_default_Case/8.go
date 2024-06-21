package main

import (
	"fmt"
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// TEST SUITE - Don't touch below this line

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(
		snapshotTicker,
		saveAfter,
	)
	fmt.Println("===========================")
}

func main() {
	test()
}

/*
ASSIGNMENT
Like all good back-end engineers, we frequently save backup snapshots of the Mailio database.

Complete the saveBackups function.

It should read values from the snapshotTicker and saveAfter channels simultaneously and continuously.

If a value is received from snapshotTicker, call takeSnapshot()
If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds. After all, we want to show in the logs that the snapshot service is running.
*/

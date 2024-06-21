package main

import (
	"fmt"
)

func getPacketSize(message string) int {
	length := len(message)
	for i := length - 1; i >= 2; i-- {
		if length%i == 0 && isComposite(i) {
			return i
		}
	}
	return length
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isComposite(num int) bool {
	// A number is composite if it is greater than 1 and not prime
	return num > 1 && !isPrime(num)
}

func test(message string) {
	fmt.Println(getPacketSize(message))
}

func main() {
	test("You shall not pass!!")                 // expected: 5
	test("May the odds be ever in your favor!!") // expected: 9
}

/*
PACKET SIZE
Textio needs to divide messages up into data packets. Normally, packet sizes are uniform and divided into bytes. An intern had the bright idea to customize packet sizes per message. Help textio test this inefficient idea.

ASSIGNMENT
Implement the getPacketSize function to take a string message and return an integer. Find the packet size for which the message characters can be evenly divided into the highest composite number of packets. Use the provided isPrime function within getPacketSize.

getPacketSize("You shall not pass!!")
// returns 5

getPacketSize("May the odds be ever in your favor!!")
// returns 9
Copy icon



*/

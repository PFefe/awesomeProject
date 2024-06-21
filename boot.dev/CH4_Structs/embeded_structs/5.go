package main

type sender struct {
	rateLimit int
}

type user struct {
	name   string
	number int
	sender
}

//At Textio, a "user" struct represents an account holder, and a "sender" is just a "user" with some additional "sender" specific data. A "sender" is a user that has a rateLimit field that tells us how many messages they are allowed to send.
//
//Fix the bug by embedding the proper struct in the other.

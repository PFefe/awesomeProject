package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return "Authorization: Basic " + a.username + ":" + a.password
}

func main() {
	// create a new authenticationInfo struct
	auth := authenticationInfo{
		username: "username",
		password: "password",
	}
	// call the getBasicAuth method on the auth struct
	fmt.Println(auth.getBasicAuth())
}

// create the method below

//Let's clean up Textio's authentication logic. We store our user's authentication data inside an authenticationInfo struct. We need a method that can take that data and return a basic authorization string.
//
//The format of the string should be:
//
//Authorization: Basic USERNAME:PASSWORD
//Copy icon
//Create a method on the authenticationInfo struct called getBasicAuth that returns the formatted string.

// SAmple:
// type rect struct {
//  width int
//  height int
//}
//
//// area has a receiver of (r rect)
//func (r rect) area() int {
//  return r.width * r.height
//}
//
//var r = rect{
//  width: 5,
//  height: 10,
//}
//
//fmt.Println(r.area())
//// prints 50

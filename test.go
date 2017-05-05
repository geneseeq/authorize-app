package main

import "fmt"

import "github.com/hsluoyz/casbin"

var gys int 
func main() {
    // fmt.Println("Hello, World!")
	// fmt.Println(gys)
	e := casbin.NewEnforcer("casbin.conf")
	sub := "alice" // the user that wants to access a resource.
	obj := "www.geneseeq.com/sample" // the resource that is going to be accessed.
	act := "read" // the operation that the user performs on the resource.
	if e.Enforce(sub, obj, act) == true {
		fmt.Println("this is true testing.")
	} else {
		fmt.Println("this is false testing.")
	}
}


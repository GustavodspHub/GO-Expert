/*Declaração e atribuição*/
package main

const helloWorld = "Hello world"

var helloWorldAsVar string = "Hello world as a var"
var (
	test         string = "This message is a test"
	testMessage3 string = "We are learning how variables and constants work"
	testMessage4 string = "It's very interesting"
)
var testMessage5 string

func main() {
	testMessage5 = "There're several ways to do this"
	testMessage6 := "This is an abbreviation"

	println(helloWorld)
	println(helloWorldAsVar)
	println(test)
	println(testMessage3)
	println(testMessage4)
	println(testMessage5)
	println(testMessage6)
}

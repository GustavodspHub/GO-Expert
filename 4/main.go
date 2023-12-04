/*Importando fmt*/
package main

import "fmt"

const helloWorld = "Hello world"
const numberTest = 1
const booleanTest = true
const floatTest = 1.0

func main() {

	fmt.Printf("The type of helloWorld constant is %T", helloWorld)
	println("")
	fmt.Printf("The type of numberTest constant is %T", numberTest)
	println("")
	fmt.Printf("The type of booleanTest constant is %T", booleanTest)
	println("")
	fmt.Printf("The type of floatTest constant is %T", floatTest)
}

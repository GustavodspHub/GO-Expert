/*Percorrendo Arrays*/

package main

import "fmt"

var myArray [5]int

func main() {

	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50

	for index, value := range myArray {
		fmt.Printf("The index is %d and the value is %d\n", index, value)
	}

}

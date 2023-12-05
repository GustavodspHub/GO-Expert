/*Closures*/
package main

import "fmt"

func main() {
	total := func() int {
		return sum(20, 30) * 2
	}()

	fmt.Println(total)
}

func sum(numberOne, numberTwo int) int {
	return numberOne + numberTwo
}

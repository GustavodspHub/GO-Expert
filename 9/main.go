/*Funções variádicas*/

package main

import "fmt"

func main() {
	fmt.Println(sum(432432432, 423423, 432432, 4324324, 32432432, 4324432))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

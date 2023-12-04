/*Functions*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	value, error := sum(5, 50)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(value)
}

func sum(numberOne, numberTwo int) (int, error) {
	if numberOne+numberTwo >= 50 {
		return 0, errors.New("A soma ultrapassou de 50")
	}

	return numberOne + numberTwo, nil
}

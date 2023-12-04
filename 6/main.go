package main

import "fmt"

func main() {
	mySlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len = %d cap = %d value = %v\n", len(mySlice), cap(mySlice), mySlice)
	fmt.Printf("len = %d cap = %d value = %v\n", len(mySlice[:5]), cap(mySlice[:5]), mySlice[:5])
	fmt.Printf("len = %d cap = %d value = %v\n", len(mySlice[5:]), cap(mySlice[5:]), mySlice[5:])

	mySlice = append(mySlice, 110)

	fmt.Printf("len = %d cap = %d value = %v\n", len(mySlice), cap(mySlice), mySlice)

}

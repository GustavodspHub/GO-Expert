/*Structs*/

package main

import "fmt"

func main() {
	type Client struct {
		name   string
		age    int
		active bool
	}

	newClient := Client{
		name:   "Gustavo",
		age:    20,
		active: true,
	}

	fmt.Printf("My client is %s, your age is %d, and you active status is %t.\n", newClient.name, newClient.age, newClient.active)

	newClientTwo := Client{
		name:   "Millena",
		age:    19,
		active: false,
	}

	fmt.Printf("My client is %s, your age is %d, and you active status is %t\n.", newClientTwo.name, newClientTwo.age, newClientTwo.active)

	newClient.name = "Roberto"
	newClient.age = 30
	newClient.active = false

	fmt.Printf("My client is %s, your age is %d, and you active status is %t.\n", newClient.name, newClient.age, newClient.active)

	newClientTwo.name = "Gisele"
	newClientTwo.age = 49
	newClientTwo.active = true

	fmt.Printf("My client is %s, your age is %d, and you active status is %t\n.", newClientTwo.name, newClientTwo.age, newClientTwo.active)
}

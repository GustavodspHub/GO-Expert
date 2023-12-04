/*maps*/
package main

import "fmt"

func main() {
	salaries := map[string]int{"Gustavo": 10000, "Millena": 50000}
	fmt.Println(salaries["Gustavo"])
	delete(salaries, "Gustavo")
	salaries["Gustavo Teste"] = 100000
	fmt.Println(salaries)

	newSalary := make(map[string]int)
	newSalary["Richard Thompson"] = 100
	newSalaryTest := map[string]int{}
	newSalaryTest["Roberto"] = 200

	for name, salary := range salaries {
		fmt.Printf("O %s tem um salário de %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("Você recebera um salário de %d\n", salary)
	}

}

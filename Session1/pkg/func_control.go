package pkg

import (
	"errors"
	"fmt"
)

func RunFuncAndControl() {
	printValue := "Hello World"
	printFn(printValue)

	a := 0
	b := 2
	var add, sub, prod, div, err = calculation(a, b)
	if err != nil {
		fmt.Print(err.Error())
		return
	} else if a == 0 {
		fmt.Printf("Addition: %v\n", add)
		fmt.Println("Substraction: ", sub)
		fmt.Println("Product: ", prod)
	} else {
		fmt.Printf("Addition: %v\n", add)
		fmt.Println("Substraction: ", sub)
		fmt.Println("Product: ", prod)
		fmt.Println("Division: ", div)
	}

}

func printFn(value string) {
	fmt.Println(value + " :))))")
}

func calculation(a int, b int) (int, int, int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, 0, 0, err
	}
	add := a + b
	substract := a - b
	product := a * b
	div := a / b

	return add, substract, product, div, err
}

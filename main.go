package main

import (
	"fmt"
	"projetoTeste/calc"
)

func main() {
	var num1 int = 4
	num2 := 2

	result := calc.Dividir(num1, num2)

	fmt.Println(result)
}

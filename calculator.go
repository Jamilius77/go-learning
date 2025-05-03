package main

import "fmt"

func main() {
	var firstNumber int
	var secondNumber int
	var operator string
	fmt.Println("введите первое число")
	fmt.Scanln(&firstNumber)
	fmt.Println("введите второе число")
	fmt.Scanln(&secondNumber)
	fmt.Println("введите оператор +, -, /, *")
	fmt.Scanln(&operator)
	switch operator {
	case "*":
		multiply(firstNumber, secondNumber)
	case "+":
		add(firstNumber, secondNumber)
	case "/":
		split(firstNumber, secondNumber)
	case "-":
		sub(firstNumber, secondNumber)

	}

}
func add(a int, b int) {
	var result int = a + b
	fmt.Println("ответ", result)
}
func sub(a int, b int) {
	var result int = a - b
	fmt.Println("результат", result)
}
func multiply(a int, b int) {
	var result int = a * b
	fmt.Println("Результат", result)
}
func split(a int, b int) {
	var result int = a / b
	fmt.Println("Ответ", result)
}

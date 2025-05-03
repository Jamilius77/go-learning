package main

import "fmt"

// Определение високосного года
func main() {
	var number int
	fmt.Println("введите количество дней в году")
	fmt.Scanln(&number)
	dayYear(number)

}
func dayYear(a int) {
	if a != 366 {
		fmt.Println("Год не високосный")
	} else {
		fmt.Println("год високосный")
	}
}

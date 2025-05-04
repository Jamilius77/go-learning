package main

import "fmt"

func main() {
	var numbers = [5]int{1, 4, 5, 6, 8}
	sumArrayNumbers(numbers)
}

func sumArrayNumbers(array [5]int) {
	sum := 0
	for _, value := range array {
		sum += value
	}
	fmt.Println(sum)
}

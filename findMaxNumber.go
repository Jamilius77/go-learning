package main

import "fmt"

func main() {
	var numbers = [5]int{99, 24, 645, 22, 12}
	findMaxNumber(numbers)
}
func findMaxNumber(array [5]int) {
	var max int = 0
	for _, value := range array {
		if value > max {
			max = value
		}
	}
	fmt.Println("Самое большое число", max)
}

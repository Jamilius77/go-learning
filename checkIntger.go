package main

import "fmt"

func main() {
	var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, value := range numbers {
		checkIntger(value)

	}
}
func checkIntger(a int) {
	if a%2 == 0 {
		fmt.Println("четное", a)
	} else {
		fmt.Println("нечетное", a)
	}

}

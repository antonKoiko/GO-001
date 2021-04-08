package main

import (
	"fmt"
)

func main() {
	var num int
	var result string
	fmt.Println("Введите число от 1 до 4")
	fmt.Scan(&num)
	result = seasons(num)
	fmt.Println(result)
}

func seasons(num int) string {
	if num == 1 {
		return "Зима"
	} else if num == 2 {
		return "Весна"
	} else if num == 3 {
		return "Лето"
	} else if num == 4 {
		return "Осень"
	} else {
		return "Error"
	}
}

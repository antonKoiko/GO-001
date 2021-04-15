package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Print("Введите минуты: ")
	fmt.Scan(&i)
	if i >= 60 {
		fmt.Println("Error")
	} else if i >= 45 {
		fmt.Println("Четвёртая четверть часа")
	} else if i >= 30 {
		fmt.Println("Третья четверть часа")
	} else if i >= 15 {
		fmt.Println("Вторая четверть часа")
	} else if i >= 0 {
		fmt.Println("Первая четверть часа")
	}

}

//Задача №1 1. Есть слайс из 1000 элементов типа int,
//нужно его забить и отсортировать в обратном порядке без использования встроенных функций
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var users []int                  // Создаём пустой слайс
	rand.Seed(time.Now().UnixNano()) //Передают текущее время. что дает постоянно меняющуюся последовательность псевдослучайных чисел.
	for i := 0; i <= 1000; i++ {     // Эта функция наполнят слайс случайными элементами от 0 до 1000
		users = append(users, rand.Intn(1000))
	}
	BubbleSort(users)  // Сортировка сортирует слайс в обратном порядке
	fmt.Println(users) //Выводим на экран
}

func BubbleSort(users []int) {
	for i := 0; i < len(users); i++ {
		for j := i; j < len(users); j++ {
			if users[i] < users[j] {
				users[i], users[j] = users[j], users[i]
			}
		}
	}
}

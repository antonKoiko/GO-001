/* Задача №1
Вывести на экран ряд чисел Фибоначчи, состоящий из n элементов.*/

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите число:") //Получаем число от пользователя
	fmt.Scan(&n)
	if n < 0 { // Проверка
		fmt.Println("Error: Повторите ввод")
		return
	}
	fub(n)

}

func fub(n int) int { //Эта функция вычисляет Числа Фибоначчи
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		fmt.Println(a)
	}
	return a
}

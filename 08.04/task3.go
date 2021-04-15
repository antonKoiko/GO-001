// Задача №3 Вычислить факториал введенного числа.

// Факториалом числа называют произведение всех натуральных чисел до этого числа включительно.
// Например, факториал числа 4 равен 1*2*3*4 = 24. Записывается факториал так: 4! = 24.

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Введите число:")
	fmt.Scan(&n)
	if n < 0 { // Проверка неотрицательное ли число
		fmt.Println("Error: Повторите ввод")
		return
	}
	fmt.Println(facto(n)) // В параметр функции передаём искомый элемент

}

func facto(n int) int { //Эта функция вычисляет Факториал числа
	var a int = 1

	for ; n > 0; n-- {
		a *= n
	}
	return a

}
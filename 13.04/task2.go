// Задача 2. Напишите функцию для подсчета частоты упоминания слов в строке текста и
// возвращения карты со словами и числом, указывающем, сколько раз они употребляются.
// Функция должна конвертировать текст в нижний регистр и обрезать знаки препинания.
// Используйте пакет strings. Функции, которые пригодятся для выполнения данного задания: Fields, ToLower и Trim

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "!Напишите функцию для подсчета частоты упоминания слов в строке текста и возвращения карты со словами и числом, указывающем, сколько раз они употребляются. Функция должна конвертировать текст в нижний регистр и обрезать знаки препинания! Используйте пакет strings. Функции, которые пригодятся для выполнения данного задания: Fields, ToLower и Trim."
	fmt.Println("Частота упоминания слова в тексте: ")
	result := words(text)
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
}

func words(text string) map[string]int {
	text2 := strings.ToLower(text)
	slice_text := strings.Fields(text2)
	result := make(map[string]int)
	for _, v := range slice_text {
		value := strings.Trim(v, "!-,.")
		result[value]++
	}
	return result
}

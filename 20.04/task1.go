// Задание: из предыдущего задания -
// реализовать функции валидации проверки заполнения фигур на ненулевые значения,
// заполнить map рандомно неповторяющимися объектами (минимум 4), реализовать отдельную функцию,
// которая за один обход будет проверять валидность всех значения в map.
// Проверить на правильность программу с передачей пустых объектов

// В фигурах есть координаты, радиус и прочее,
// сделать у ваших объектов функции валидации на проверку,
// что координаты не нулевые потом заполнить тип данных map вашими объектами и написать функцию,
// которая будет принимать map и валидировать все ваши фигуры

package main

import (
	"fmt"
)

func main() {
	figures := []Figure{
		Rectangle{
			left:   500,
			top:    1000,
			right:  500,
			bottom: 1000,
		},

		Circle{
			radius: 5,
			centre: 1.5,
		},

		Rectangle{
			left:   20,
			top:    40,
			right:  20,
			bottom: 40,
		},

		Circle{
			radius: 50,
			centre: 15.4,
		},
	}

	for _, fig := range figures {
		s := fig.Format()
		fmt.Println(s)
	}
}

type Figure interface {
	Format() string
}

type Rectangle struct {
	left   int
	top    int
	right  int
	bottom int
}

func (r Rectangle) Format() string {
	return fmt.Sprintf(
		"paint rectangle, Rect {left = %v, top = %v, right = %v, bottom = %v}",
		r.left,
		r.top,
		r.right,
		r.bottom,
	)
}

type Circle struct {
	radius int
	centre float64
}

func (c Circle) Format() string {
	return fmt.Sprintf("paint circle, radius=%v and centre=(%v)", c.radius, c.centre)
}

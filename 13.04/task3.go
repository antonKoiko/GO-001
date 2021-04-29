//Задача 3. Есть графические примитивы(фигуры), с различными свойствами,
// но над каждой фигурой можно производить одинаковые действия.
// Примитивы должны уметь отдавать информацию о себе в определенном формате,
// которую некая функция будет выводить в какой-нибудь вывод, для простоты примера — в stdout.
// При этом некоторые примитивы могут быть вариациями других.

// Формат вывода информации, на примере прямоугольника и круга, должен быть вот таким:
// paint rectangle, Rect {left = 10, top = 20, right = 600, bottom = 400}
// paint circle, radius=150 and centre=(50,300)

// Кроме того, примитивы нужно уметь объединить в однородный список.

// Нужно просто сделать базовую фигуру, а потом у всех в остальные фигуры встраивать базовую
// + раз мы уже прошли в новом материале - то можно добавить функциональность у типов фигур методы,
// которые и будут добавлять необходимую функциональность

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

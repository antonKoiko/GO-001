package main

import "fmt"

var a int
var b int

func main() {
	a = 3
	b = 5
	if a <= 0 && b >= 3 {
		var z int = a + b
		fmt.Println("a + b =", z)
	} else {
		var w int = a - b
		fmt.Println("a - b =", w)
	}
}

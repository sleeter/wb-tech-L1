package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
// С помощью сложения и вычитания двух переменных
func main() {
	a := 1
	b := 2

	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
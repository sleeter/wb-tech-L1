package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/
// С помощью XOR
func main() {
	a := 1
	b := 2

	a ^= b
	b ^= a
	a ^= b

	fmt.Println(a, b)
}

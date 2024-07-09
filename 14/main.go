package main

import "fmt"

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/
func main() {
	// создаем переменные разных типов
	integer := 1
	boolean := true
	str := "string"
	ch := make(chan int)
	// передаем их в функцию
	typeOfVar(integer)
	typeOfVar(boolean)
	typeOfVar(str)
	typeOfVar(ch)
}

func typeOfVar(val interface{}) {
	// с помощью switch и .(type) можем проверить тип на его соответствие
	switch val.(type) {
	case int:
		fmt.Println("type int")
	case string:
		fmt.Println("type string")
	case bool:
		fmt.Println("type bool")
	case chan int:
		fmt.Println("type chan")
	default:
		fmt.Println("type is unknown")
	}
}

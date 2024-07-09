package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 2
	// меняем i-й и последний элмент
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	// делаем срез до последнего элемента
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}

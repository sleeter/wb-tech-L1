package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 2
	// делаем рез до i-го элемента
	res := arr[:i]
	// добавляем срез после i-го элмента
	res = append(res, arr[i+1:]...)
	fmt.Println(res)
}

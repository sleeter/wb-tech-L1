package main

import "fmt"

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
func main() {
	arr := []int{1, 5, 4, 7, 10, 2, 8, 3, 9, 0, 6}
	qSort(arr)
	fmt.Println(arr)
}

func qSort(arr []int) []int {
	// если элементов меньше, чем 2, значит массив отсортирован
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	// выбор опорного элемента
	pivotIndex := len(arr) / 2

	// перестановка опорного элемента в конец
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// если элмент меньше опорного, то он идет в левую часть
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// возврат опорного элемента на место
	arr[left], arr[right] = arr[right], arr[left]

	// рекурсивная сортировка левой и правой частей
	qSort(arr[:left])
	qSort(arr[left+1:])

	return arr
}

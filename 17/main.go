package main

import (
	"errors"
	"fmt"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/
func main() {
	// создаем массив
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// начение, которое будем искать
	val := 8
	ind, err := binarySearch(arr, val)
	// проверка ошибки и вывод резульатата
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ind)
	}

}
func binarySearch(arr []int, val int) (int, error) {
	// инициализируем левую и правую границу
	left := 0
	right := len(arr) - 1
	// условие выхода из цикла
	for left <= right {
		// значение, которое проверяем
		mid := (left + right) / 2
		// если попали - выход
		if arr[mid] == val {
			return mid, nil
		}
		// если значение, которое проверяем меньше, чем которое надо найти
		// то сдвигаем границы на правую часть
		if arr[mid] < val {
			left = mid + 1
		} else { // иначе сдвигаем границы на левую часть
			right = mid - 1
		}
	}
	// если ничего не нашлось, то возвращаем ошибку
	return 0, errors.New("this value is not in array")
}

package main

import "fmt"

/*
Дана последовательность температурных колебаний:
-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	// создаем массив
	arr := []float64{-25.0, -27.0, -21.0, 13.0, 19.0, 15.5, 24.5, 35.0, -5, 5, -15}
	// создаем мапу, в которой ключом будет являтся диапазон, а значением слайс температур
	m := make(map[int][]float64)
	// проходимся по массиву
	for _, val := range arr {
		// считаю, что в примере ошибка, т.к. при таком разбиении значения от -10 о 10 попадут в одну группу
		// вычисляем диапазон с помощью целочисленного деления значения на 10, заетм умножения на 10
		// для отрицательных значений ключ вычисляется похожим образом, только от значение вычитаем 10
		// тогда в диапазоне значения будут разбиты с шагом в 10 градусов
		var key int
		if val < 0 {
			key = int(val-10.0) / 10 * 10
		} else {
			key = int(val) / 10 * 10
		}
		// добавляем в мапу значение
		m[key] = append(m[key], val)
	}
	// выводим мапу в stdout
	fmt.Println(m)
}

package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….)
с использованием конкурентных вычислений.
*/

// Данное решение небезопасно, т.к. у нас может произойти состояние гонки
func main() {
	// создаем массив
	arr := []int{2, 4, 6, 8, 10}
	// объявляем переменную для хранения суммы, по умолчанию равна 0
	var sum int
	// объявляем переменную WainGroup
	var wg sync.WaitGroup
	// запускаем цикл по массиву
	for _, val := range arr {
		// увеличиваем счетчик перед запуском горутины
		wg.Add(1)
		// запускаем горутину, передавая туда число
		go func(n int) {
			// после завершения горутины счетчик уменьшится на 1
			defer wg.Done()
			// прибавляем к сумме квадрат числа
			sum += n * n
		}(val)
	}
	// дожидаемся выполнения всех горутин
	wg.Wait()
	// выводим значение суммы в stdout
	fmt.Println(sum)
}

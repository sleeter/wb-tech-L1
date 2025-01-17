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

// Данное решение безопасно, т.к. участок кода в котором происходит
// конкурентный доступ к памяти покрыт мьютексом
func main() {
	// создаем массив
	arr := []int{2, 4, 6, 8, 10}
	// объявляем переменную для хранения суммы, по умолчанию равна 0
	var sum int
	// объявляем переменную WainGroup
	var wg sync.WaitGroup
	// объявляем переменную Mutex
	var mutex sync.Mutex
	// запускаем цикл по массиву
	for _, val := range arr {
		// увеличиваем счетчик перед запуском горутины
		wg.Add(1)
		// запускаем горутину, передавая туда число
		go func(n int) {
			// после завершения горутины счетчик уменьшится на 1
			defer wg.Done()
			// блокируем Mutex
			mutex.Lock()
			// прибавляем к сумме квадрат числа
			sum += n * n
			// разблокируем Mutex
			mutex.Unlock()
		}(val)
	}
	// дожидаемся выполнения всех горутин
	wg.Wait()
	// выводим значение суммы в stdout
	fmt.Println(sum)
}

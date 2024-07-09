package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/
func main() {
	// создаем массив
	arr := []int{2, 4, 6, 8, 10}
	// создаем 1 буферезированный канал
	ch1 := make(chan int, len(arr))
	// создаем 2 буферезированный канал
	ch2 := make(chan int, len(arr))
	// объявляем WaitGroup
	var wg sync.WaitGroup
	// увеличиваем счетчик
	wg.Add(1)
	// запускаем горутину
	go func() {
		// проходимся по массиву
		for _, val := range arr {
			// записываем число в канал 1
			ch1 <- val
		}
		// уменьшаем счетчик
		wg.Done()
		// закрываем канал
		close(ch1)
	}()
	// увеличиваем счетчик
	wg.Add(1)
	// запускаем горутину
	go func() {
		// проходимся по буферу канала 1
		for x := range ch1 {
			// записываем квадрат числа в канал 2
			ch2 <- x * x
		}
		// уменьшаем счетчик
		wg.Done()
		// закрываем канал
		close(ch2)
	}()
	// проходимся по буферу канала 1
	for val := range ch2 {
		// выводим в stdout
		fmt.Println(val)
	}
	// дожидаемся завершения
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

// Counter - структура счетчик
type Counter struct {
	counter int32
}

// inc - функция инкремента, в данном случае реализованна с помощью атомик
func (c *Counter) inc() {
	atomic.AddInt32(&c.counter, 1)
}
func main() {
	// создаем структуру
	c := Counter{}
	// объявляем WaitGroup
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		// увеличиваем счетчик
		wg.Add(1)
		// в цикле создаем две горутины для инкрементирования счетчика
		go func() {
			for j := 0; j < 10; j++ {
				c.inc()
			}
			// уменьшаем счетчик
			wg.Done()
		}()
	}
	// дожидаемся выполнения всех горутин
	wg.Wait()
	// выводим значение в stdout
	fmt.Println(c.counter)
}

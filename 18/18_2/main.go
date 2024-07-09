package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
type Counter struct {
	counter int
	mutex   sync.Mutex
}

func (c *Counter) inc() {
	c.mutex.Lock()
	c.counter++
	c.mutex.Unlock()
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

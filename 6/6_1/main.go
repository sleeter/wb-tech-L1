package main

import (
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func main() {
	ch := make(chan int)
	defer close(ch)

	go func() {
		for {
			// получаем значение ok, если false, то канал закрыт
			v, ok := <-ch
			if !ok {
				return
			}
			fmt.Printf("value = %d\n", v)
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
}

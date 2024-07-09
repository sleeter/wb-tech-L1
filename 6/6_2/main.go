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
	stop := make(chan struct{}) // заводим дополнительный канал для остановки

	var val int
	go func() {
		for {
			select {
			case ch <- val:
				val++
			case <-stop: // если появилось значение, то закрываем канал
				close(ch)
				fmt.Println("chanel closed")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		stop <- struct{}{}
	}()

	for i := range ch {
		fmt.Printf("value = %d\n", i)
	}

}

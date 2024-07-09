package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
func main() {
	ch := make(chan int)
	// закрытие канала с помощью контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	var val int
	go func(ctx context.Context) {
		for {
			select {
			case ch <- val:
				val++
			case <-ctx.Done(): // если контекст отменили, закрываем канал
				close(ch)
				fmt.Println("chanel closed")
				return
			}
			time.Sleep(1 * time.Second)
		}
	}(ctx)

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	for i := range ch {
		fmt.Printf("value = %d\n", i)
	}
}

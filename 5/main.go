package main

import (
	"fmt"
	"os"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
func main() {
	var t int

	fmt.Println("Введите количество секунд работы программы")
	// считываем количество секунд
	_, err := fmt.Scan(&t)
	if err != nil {
		return
	}
	// отложенная запись в timer по истечению t секунд
	timer := time.After(time.Duration(t) * time.Second)
	// канал для значений
	values := make(chan int, 1)
	defer close(values)
	// канал для остановки
	stop := make(chan struct{}, 1)
	// горутина для получения значений и стоп сигнала
	go func(values chan int, stop chan struct{}) {
		for {
			select {
			case <-stop:
				fmt.Printf("программа остановлена")
				os.Exit(0)
			default:
				fmt.Printf("значение = %d\n", <-values)
			}
		}
	}(values, stop)
	var val int
	// основной цикл
	for {
		select { // либо записывай в канал val, либо по истечению времени завершаем программу
		case <-timer:
			fmt.Printf("программа завершается, прошло %d секунд", t)
			stop <- struct{}{}
			//return
		default:
			values <- val
			val++
			time.Sleep(time.Second)
		}
	}
}

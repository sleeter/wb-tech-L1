package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала
и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

// создается сигнальный канал для каждой из горутин
// дожидаемся выполнения горутин, потом останавливаем
// это позвоялет дочитать значения и корректно остановить выполнение

func worker(id int, val <-chan int, signal chan os.Signal) {
	defer close(signal)
	for {
		select {
		case v := <-val:
			fmt.Printf("горутина %d пишет %d\n", id, v)

		case <-signal:
			fmt.Printf("горутина с идентификатором %d остановлена\n", id)
			return
		}
	}
}

func main() {
	var count int
	fmt.Println("Введите количество горутин для воркера")
	// считываем количество горутин для воркера
	_, err := fmt.Scan(&count)
	if err != nil {
		return
	}

	// создаем буферезированный канал, где размер буфера равен количеству горутин
	values := make(chan int, count)
	defer close(values)
	// создаем канал для сигнала остановки
	stop := make(chan os.Signal, 1)
	defer close(stop)

	// объявляем какие сигналы надо слушать
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// в цикле создаем канал для каждой горутины и запускаем воркера
	for i := 0; i < count; i++ {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go worker(i, values, sigChan)
	}

	var val int
	// основной цикл
	for {
		select { // записываем либо в канал значений, либо завершаем выполнение цикла
		case <-stop:
			fmt.Printf("основной цикл завершен")
			return
		default:
			values <- val
			val++
			time.Sleep(250 * time.Millisecond)
		}
	}

}

package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/
func sleep(d time.Duration) {
	start := time.Now()
	// пока разница между стартом и текущим временем меньше d выполняем цикл
	for {
		curr := time.Now()
		if curr.Sub(start) >= d {
			return
		}
	}
}

func main() {
	d := time.Duration(3) * time.Second
	fmt.Println("time to sleep")
	sleep(d)
	fmt.Println("wake up")
}

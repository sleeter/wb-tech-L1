package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/
func sleep(d time.Duration) {
	t := time.NewTimer(d)
	// произойдет блокировка пока не придет значение
	if <-t.C; true {
		return
	}
}
func main() {
	d := time.Duration(3) * time.Second
	fmt.Println("time to sleep")
	sleep(d)
	fmt.Println("wake up")
}

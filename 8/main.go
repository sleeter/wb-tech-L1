package main

import (
	"errors"
	"fmt"
)

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
func main() {
	val := int64(1025)
	// вызываем фукнцию setBit
	newVal, err := setBit(val, 1, false)
	// проверяем есть ли ошибка
	if err != nil {
		return
	}
	// выводим в stdout
	fmt.Println(newVal)
}
func setBit(val int64, bit int, pos bool) (int64, error) {
	// проверяем удовлетворяет ли позиция бита для переменной int64
	if bit > 63 || bit < 0 {
		return 0, errors.New("number of bit should be greater than zero and less than 64")
	}
	// если true, то ставим i бит в 1 с помощью логического или
	if pos {
		val = val | 1<<bit
	} else { // если false, то ставим i бит в 0
		tmp := int64(1 << bit)
		// сначала сделаем так, чтобы этот бит был 1
		val = val | tmp
		// затем вычтем число, чтобы бит стал 0
		val -= tmp
	}
	return val, nil
}

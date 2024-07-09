package main

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}
*/
// Использование глобальной переменной
// Захардкоженное значение для функции createHugeString
// Срез происходит у строки (последовательности байтов), а не символов
// Можно сразу создавать строку необходимого размера, не использовать глобальную переменную
func createHugeString(size int) string {
	// алфавит, из которого будет составлена строка
	alphabet := []rune("abcdefg12345")
	// создаем слайс рун
	res := make([]rune, size)
	// в цикле присваиваем каждой руне значение
	for i := range res {
		res[i] = alphabet[rand.Intn(len(alphabet))]
	}
	// возвращаем строку
	return string(res)
}

func someFunc(size int) (string, error) {
	// валидация размера строки
	if size <= 0 {
		return "", errors.New("size should be greater than zero")
	}
	// вызываем метод, который сгенерирует строку определенного размера
	hugeStr := createHugeString(size)
	return hugeStr, nil
}

func main() {
	//вызываем функцию, которая вернет строку длиной 50 символов
	hugeStr, err := someFunc(50)
	if err != nil {
		fmt.Print(err)
		return
	}
	// выводим строку в stdout
	fmt.Print(hugeStr)
}

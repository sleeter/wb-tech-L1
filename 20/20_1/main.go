package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
func main() {
	// считываем строку с stdin
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// удаляем пробельные символы
	str = strings.TrimSpace(str)
	// разделяем по пробелу
	arr := strings.Split(str, " ")
	// создаем переменную для результата
	res := ""
	// проходим в цикле по массиву в обратном порядке и конкатенируем строки, добавляя пробелы
	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i]
		res += " "
	}
	// выводим в stdout
	fmt.Println(res)
}

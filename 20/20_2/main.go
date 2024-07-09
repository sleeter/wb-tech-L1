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
	// создаем стрингбилдер
	var builder strings.Builder
	// проходим в цикле по массиву в обратном порядке и пишем в стрингбилдер, добавляя пробелы
	for i := len(arr) - 1; i >= 0; i-- {
		builder.WriteString(arr[i])
		builder.WriteString(" ")
	}
	// выводим в stdout
	fmt.Println(builder.String())
}

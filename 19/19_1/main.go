package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
func main() {
	// считываем строку с stdin
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// приводим к массиву рун
	runes := []rune(str)
	// создаем переменную для реверснутой строки
	res := ""
	// проходим по массиву рун в обратном порядке и конкатенируем строки
	for i := len(runes) - 1; i >= 0; i-- {
		res += string(runes[i])
	}
	// выводим в stdout
	fmt.Println(res)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	// объявляем стрингбилдер
	var builder strings.Builder
	// проходим по массиву рун в обратном порядке и пишем в стрингбилдер
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteString(string(runes[i]))
	}
	// выводим в stdout
	fmt.Println(builder.String())
}

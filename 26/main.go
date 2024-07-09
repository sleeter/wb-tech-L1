package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
func main() {
	arr := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, val := range arr {
		res := uniqueSymbols(val)
		fmt.Println(val, "-", res)
	}
}
func uniqueSymbols(s string) bool {
	// создаем сет, в котором будут хранится символы
	set := make(map[rune]struct{})
	// преобразовываем строку к рунам
	runes := []rune(s)

	// проходимся по рунам, если в сете уже есть такой символ, то возврщаем false
	// иначе добавляем в сет
	for _, val := range runes {
		_, ok := set[val]
		if ok {
			return false
		}
		set[val] = struct{}{}
	}
	// если прошли весь цикл, возвращаем true
	return true
}

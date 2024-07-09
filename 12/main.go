package main

import "fmt"

/*
Имеется последовательность строк
(cat, cat, dog, cat, tree)
создать для нее собственное множество.
*/
func main() {
	// создаем массив
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	// создаем мапу, где ключ - строка, значение - пустая структура
	set := make(map[string]struct{})
	// проходимся по массиву
	for _, v := range arr {
		// для каждого ключа добавляем пустую структуру
		set[v] = struct{}{}
	}
	// у нас получается сет из ключей
	for k, _ := range set {
		// выводим в stdout
		fmt.Println(k)
	}
}

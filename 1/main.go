package main

import (
	"errors"
	"fmt"
)

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human(аналог наследования).
*/

// структура Human с полями name и age
type Human struct {
	name string
	age  int
}

// структура Action
type Action struct {
	// используем встраивание для *Human
	// используем *Human чтобы избежать копирования
	*Human
}

// метод walk для Human
func (h Human) walk() {
	fmt.Println("Human is walking")
}

// метод run для Human
func (h Human) run() {
	fmt.Println("Human is running")
}

// конструктор для Human
func newHuman(name string, age int) *Human {
	return &Human{
		name: name,
		age:  age,
	}
}

// конструктор для Action
func newAction(human *Human) *Action {
	return &Action{Human: human}
}

// метод для валидации полей Human
// валидация производится в отдельной функции, чтобы легче покрыть тестами
func validateHumanFields(name string, age int) error {
	if len(name) == 0 {
		return errors.New("Human name should not be empty")
	}
	if age <= 0 {
		return errors.New("Human age should be greater than zero")
	}
	return nil
}
func main() {
	name := "Timur"
	age := 21
	// производим валидацию полей
	err := validateHumanFields(name, age)
	if err != nil {
		return
	}
	// если валидация прошла успешно создаем Human
	human := newHuman(name, age)
	// затем создаем Action
	action := newAction(human)
	// вызываем метод run() у Action
	action.run()
}

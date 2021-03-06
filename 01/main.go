package main

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

import "fmt"

type Human struct {
	name string
	age int
}

// Создает объект Human
func NewHuman(name string, age int) Human {
	return Human{
		name: name,
		age: age,
	}
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, i'm %s\n", h.name)
}

func (h *Human) Kiss() {
	fmt.Println("(´ε｀)")
}

// Action содержит Human,
// т.е. содержит его методы (но не поля, т.к. они приватные)
type Action struct {
	Human
}

func main() {
	// объект Human
	h := NewHuman("Joe", 23)
	// объект Action, который содержит объект Human
	a := Action{h}
	a.SayHi() // Вывод: Hi, i'm Joe
	// или так:
	a.Human.Kiss() // Вывод: (´ε｀)
}

// наследование vs встраивание
//
// композиция - когда объект состоит их более мелких обектов
//
// При наследовании наследуемый (дочерний) тип B получает свойства и методы родителя A и его тип (B - подтип A). Т.е. is (является)
// При встраивании A содержит B и получает его возможности, но его тип остается A. Т.е. has (имеет/содержит)
//
// Значения и методы начинают искать с корня, потом во встроенных типах (shadowing)
// colliding (коллизия) - одинаковые имена/методы во встроенных типах на одинаковой вложенности
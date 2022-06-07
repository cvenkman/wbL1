package main

// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

// Рефлексия — способность программы исследовать собственную структуру, в особенности через типы.

type MyInt int

func main() {
	var i int
	printType(i) // int

	var j MyInt
	printType(j) // MyInt

	printType(false) // bool
	printType("deasfsdf") // string
	printType(make(chan int)) // chan int


	// второй вариант
	printType2(i) // int
	printType2(true) // bool
}

// метод принимает интерфейс и определяет тип v
func printType(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}

func printType2(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Println("error")
	}
}
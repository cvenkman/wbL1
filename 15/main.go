package main

import "fmt"

// К каким негативным последствиям может привести данный фрагмент
// кода, и как это исправить? Приведите корректный пример реализации.


// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func mainErr() {
// 	someFunc()
// }


// проблема: мы используем 100 элементов из выделенных 1024 (1 << 10 = 1024)

func main() {
	// 1 << 10 = 1024
	fmt.Println(1 << 10)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func mainErr() {
	someFunc()
}

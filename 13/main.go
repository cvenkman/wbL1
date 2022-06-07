package main

import "fmt"

func main() {
	swap1()
	swap2()
}

func swap1() {
	a, b := 11, 8
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)
}

// недостаток - переполнение при больишх a и b
func swap2() {
	a, b := 2, 6
	fmt.Println(a, b)

	a += b // a = 8
	b = a - b // b = 2
	a -= b // a = 6
	fmt.Println(a, b)
}

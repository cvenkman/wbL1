package main

// Удалить i-ый элемент из слайса.

import "fmt"

func main() {
	a := []rune{'a', 'b', 'c', 'c', 'd', 'e', 'f'}
	
	removeByIndex1(&a, 3)
	fmt.Println(string(a))

	a = removeByIndex2(a, 0)
	fmt.Println(string(a))
}

// только append
func removeByIndex1(a *[]rune, i int) {
	*a = append((*a)[:i], (*a)[i+1:]...)
}

// начинаю с нужной переменно передвигакем все переменные в массиве вперед
func removeByIndex2(a []rune, i int) []rune {
	for ; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a = a[:len(a)-1]
	return a
}

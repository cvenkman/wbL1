package main

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

import "fmt"

func reverseString(str string) string {
	// массив для перевернутого слова
	res := make([]rune, len(str))
	i := len(str) - 1

	for _, elem := range str {
		// начало строки кладем в конец массива
		res[i] = elem
		i--
	}
	return string(res)
}

func main() {
	fmt.Println(reverseString("главрыба")) // абырвалг
	fmt.Println(reverseString(""))
	fmt.Println(reverseString("sd")) // ds
}
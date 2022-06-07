package main

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	// делим строку на массив строк, разделитель - пробел
	arr := strings.Split(str, " ")

	i := len(arr) - 1
	j := 0

	for ; j < i; {
		// меняем слова местами
		arr[j], arr[i] = arr[i], arr[j]
		j++
		i--
	}
	// возвращаем массив превращенный в строку
	return strings.Join(arr, " ")
}

func main() {
	fmt.Println(reverseString("snow dog sun")) // sun dog snow
	fmt.Println(reverseString(""))
	fmt.Println(reverseString("sd")) // ds
}
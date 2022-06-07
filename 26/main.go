package main

import (
	"fmt"
)
//Разработать программу, которая проверяет,
//что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
/*
Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func main() {
	fmt.Println(containsDuplicate("cat")) // true
	fmt.Println(containsDuplicate("ccat")) // false
	fmt.Println(containsDuplicate("")) // true
}

func containsDuplicate(str string) bool {
	// ключ - символ в строке
	m := make(map[rune]bool)

	for _, char := range str {
		if _, ok := m[char]; ok {
			// если в мапе уже есть такой элемент, значит мы его положили 
			// на предыдущих итерациях - т.е. символ повторился
			return false
		}
		// если в мапе нет - добавляем и смотрим строку дальше
		m[char] = true
	}

	return true
}

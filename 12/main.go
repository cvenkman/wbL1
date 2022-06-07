package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

// Множество
// Почти то же самое, что и хеш-таблица, но без сохранения значения.
// Если нам нужно только быстрая проверка вхождения, то можно использовать встроенный map.
// Нужно лишь указать пустое значение, что бы указать, что важен только ключ.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	// добавление
	for _, elem := range arr {
		set[elem] = struct{}{}
	}
	// вывод
	for key := range set {
		fmt.Println(key)
	}
}
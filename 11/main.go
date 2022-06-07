package main

// Реализовать пересечение двух неупорядоченных множеств.
// пересечение двух массивов - переменные, которые находятся в обоих массивах

import (
	"fmt"
)

func main() {
	res := find(
		[]int{1, 2, 3, 4, -6, -4, -4, 7},
		[]int{1, 1, 7, 6, -2, 5, 5, 5, 3, -4},
	)
	
	for key := range res {
		fmt.Print(key, " ") // [1 3 -4 7]
	}
}

func find(a, b []int) map[int]struct{} {
	// создаем множество
	m := make(map[int]struct{})
	// и еще одно для храненения пересечения
	res := make(map[int]struct{})

	// проходимся по первому массиву и заносим все значения в m как ключ
	for _, elem := range a {
		m[elem] = struct{}{}
	}

	// проходимся по второму массиву
	for _, elem := range b {
		// смотрим есть ли элемент в мапе m
		_, exist := m[elem]
		if exist {
			// если есть - значит нашли пересечение, добавляем в res
			res[elem] = struct{}{}
		}
	}
	return res
}
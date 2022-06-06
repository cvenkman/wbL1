package main

// Реализовать бинарный поиск встроенными методами языка.

import "fmt"

func main() {
	// Массив должен быть отсортирован

	fmt.Println(binarySearch([]int{-4, -1, 0, 2, 3, 6, 7 , 9, 11}, 7)) // 6
	fmt.Println(binarySearch([]int{1, 2, 5}, 5)) // 2
	fmt.Println(binarySearch([]int{1, 2, 5}, 88)) // -1
}

// Алгоритм поиска элемента в отсортированном массиве,
// использующий дробление массива на половины.
// Возвращает индекс найденного элемента в массиве.
func  binarySearch(arr []int, toFind int) int {
	first := 0
	last := len(arr)-1

	for first <= last {
		// Делим массив пополам и находим середину.
		mid := (first + last) / 2
		//Далее сравниваем срединный элемент с заданным искомым элементом.

		// Если равно - число найдено
		if arr[mid] == toFind {
			return mid
		}
		// Если искомое больше — продолжаем поиск в правой
		// части массива. Если искомое меньше — в левой части массива.
		if toFind > arr[mid] {
			// передвигаем начало масива для поиска в середину
			first = mid + 1
		} else {
			// передвигаем конец масива для поиска в середину
			last = mid - 1
		}
	}
	// элемент не найден
	return -1
}

package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

func main() {
	fmt.Println(quicksort([]int{2, 5, 0, 6, 8, 1, -2, -8, 9}))
	fmt.Println(quicksort([]int{2, 1}))
	fmt.Println(quicksort([]int{2, 1, -4, 5, 0}))
}

// из массива выбирается некоторый опорный элемент a[i],
// 1) запускается процедура разделения массива, которая перемещает все ключи, меньшие,
//		либо равные a[i], влево от него, а все ключи, большие, либо равные a[i] - вправо,
// 2) теперь массив состоит из двух подмножеств, причем левое меньше, либо равно правого,
// 3)для обоих подмассивов: если в подмассиве более двух элементов,
// 		рекурсивно запускаем для него ту же процедуру.


func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	// опорный элемент
	pivot := (left + right) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	
	for it := range arr {
		if arr[it] < arr[right] {
			arr[left], arr[it] = arr[it], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	// сортируем левую часть
	quicksort(arr[:left])
	// сортируем правую часть
	quicksort(arr[left+1:])
	return arr
}


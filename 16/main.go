package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

import "fmt"

func main() {
	fmt.Println(([]int{2, 5, 0, 6, 8, 1, -2, -8, 9}))
	// fmt.Println(([]int{2, 5, 1}))
}

// из массива выбирается некоторый опорный элемент a[i],
// 1) запускается процедура разделения массива, которая перемещает все ключи, меньшие,
//		либо равные a[i], влево от него, а все ключи, большие, либо равные a[i] - вправо,
// 2) теперь массив состоит из двух подмножеств, причем левое меньше, либо равно правого,
// 3)для обоих подмассивов: если в подмассиве более двух элементов,
// 		рекурсивно запускаем для него ту же процедуру.
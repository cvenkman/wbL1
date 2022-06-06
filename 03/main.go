package main

// Дана последовательность чисел: 2,3,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	sum := 0
	
	go countSquares([]int{2, 3, 4, 6, 8, 10}, ch)
	for sq := range ch {
		// с предыдущего решения добавила только эту строку
		sum += sq
	}
	fmt.Println(sum) // 229


	//////// второе решение ////////
	arr := []int{2, 3, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	res := 0

	for _, number := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			// mutex
			res += number * number
		}(number)
	}
	wg.Wait()
	fmt.Println(res) // 229
}

func countSquares(arr []int, ch chan<- int) {
	wg := sync.WaitGroup{}

	for _, number := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			sq := number * number
			ch <- sq
		}(number)
	}
	wg.Wait()
	close(ch)
}

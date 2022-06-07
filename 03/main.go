package main

// Дана последовательность чисел: 2,3,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 3, 4, 6, 8, 10}

	//////// первое решение ////////
	ch := make(chan int)
	sum := 0

	// то же самое, что и в 02, но значения из канала прибавляются к sum
	go countSquares(arr, ch)
	for sq := range ch {
		sum += sq
	}
	fmt.Println(sum)

	//////// второе решение ////////
	chForDone := make(chan bool)
	go countSquares2(arr, chForDone)
	// канал, чтобы заблокировать main, пока
	// горутина countSquares2 не вычислит все значения
	<-chForDone
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

// решение без отправки значений в канал
func countSquares2(arr []int, chForDone chan<- bool) {
	wg := sync.WaitGroup{}
	res := 0
	for _, number := range arr {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			res += number * number
		}(number)
	}
	wg.Wait()
	chForDone <- true
	fmt.Println(res)
}

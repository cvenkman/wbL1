package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.


func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	arr := []int{1, 2, 3, 4, 5}
	
	go write(ch1, arr)
	go read(ch1, ch2)

	for elem := range ch2 {
		fmt.Println(elem)
	}
}

func write(ch chan<- int, arr []int) {
	for _, elem := range arr {
		ch <- elem
	}
	// заерываем канал, чтобы сказать read(), что данных больше не будет
	close(ch)
}

func read(chRead <-chan int, chWrite chan<- int) {
	// читаем из chRead и пишем в chWrite
	for elem := range chRead {
		chWrite <- elem * elem
	}
	// заерываем канал, чтобы сказать main(), что данных больше не будет
	close(chWrite)
}

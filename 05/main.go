package main

// Разработать программу, которая будет последовательно
// отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

import (
	"fmt"
	"time"
)

func main() {
	run(2)
}

// создает канал для чтения и записи и засыпает на n секунд
func run(n time.Duration) {
	ch := make(chan int)

	// постоянная запись
	go write(ch)
	// постоянное чтение того, что отправили в ch
	go read(ch)

	// main блокируется на n секунд, пока time.After
	// не отправит в <-chan time.Time текущее время
	<-time.After(n*time.Second)

	// Еще два способа подождать n секунд:

	// таймер
//	wait2(n)

	// просто цикл
//	wait3(n)

	fmt.Printf("Прошло %d секунд(ы)\n", n)
}

// пишет в канал только для записи
func write(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
}

// читает из канала только для чтения
func read(ch <-chan int) {
	for elem := range ch {
		fmt.Println(elem)
	}
}

// создвет таймер на n секунд и ждет
func wait2(n time.Duration) {
	// по сути то же самое, что и <-time.After(n*time.Second)

	// type Timer struct {
	// 	C <-chan Time
	// 	...
	// }

	// Текущее время отправится на C, когда таймер истечет
	timer := time.NewTimer(time.Second * n)
	<- timer.C
}

// стоит в цикле for пока не пройдет n секунд
func wait3(n time.Duration) {
	start := time.Now()
	for time.Since(start) < n*time.Second {
	}
}
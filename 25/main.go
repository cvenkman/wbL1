package main

import (
  "fmt"
  "time"
//   "log"
)

func main() {
	first := time.Now()
	sleep1(2 * time.Second)
	fmt.Println(time.Since(first)) // ~2.000255241s

	first = time.Now()
	sleep2(200 * time.Millisecond)
	fmt.Println(time.Since(first)) // ~200.000288ms
}

// функция After возвращает канал типа Time (<-chan Time)
func sleep1(needSleep time.Duration) {
	if needSleep <= 0 {
		return
	}

	// NewTimer creates a new Timer that will send
	// the current time on its channel after at least duration d.
	timer := time.NewTimer(needSleep)
	<-timer.C	
	timer.Stop()

	// или так: After также запускает таймер
	// <-time.After(needSleep)
}

// реализация без каналов
func sleep2(needSleep time.Duration) {
	// проверка на отрицательное время
	if needSleep <= 0 {
		return
	}
	// замеряем время перед циклом
	first := time.Now()
	// находимся в цикле пока время,
	// прошедшее с first, не достигнет needSleep
	for ; time.Since(first) < needSleep; {
	}
}
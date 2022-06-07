package main

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

import (
	"fmt"
	"sync"
)

type counter struct {
	// обычный мьютекс, т.к. параллельный нужно только писать
	m sync.Mutex
	c int
}

func (c *counter) Increment() {
	c.m.Lock()
	// зона к которой должна получить доступ только одна горутина
	c.c++
	c.m.Unlock()
}

func (c *counter) GetCount() int {
	return c.c
}

func main() {
	counter := counter{c: 0}

	i := 0
	for i = 0; i < 1000; i++ {
		counter.Increment()
	}

	fmt.Println(i) // 1000
	fmt.Println(counter.GetCount()) // 1000
}
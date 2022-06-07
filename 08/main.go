package main

import (
	"fmt"
	"log"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	var v int64
	v = 5
	changeByte(v, 2, 0)
	// поменять 3й бит на ноль
	// 5 - 101
	// 100 - 4
}

func changeByte(v int64, i, bit int) {
	// >64 - больше, чем бит в int64, <0 - меньше
	if i > 64 || i <= 0 {
		log.Fatal("error")
	}

	fmt.Println(v) // 5

	// << Сдвигает битовое представление числа, представленного первым операндом,
	// влево на определенное количество разрядов, которое задается вторым операндом.

	if bit == 0 {
		// & поразрядная конъюнкция (операция И или поразрядное умножение)
		v = v & 63 << i
	} else if bit == 1 {
		// | поразрядная дизъюнкция (операция ИЛИ или поразрядное сложение)
		v = v | 1 << i
	}

	fmt.Println(v) // 4
}

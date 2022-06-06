package main

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// SetString устанавливает a и b в значения, переданные сторокой
	// (строка интерпретируется в десятичной системе счисления)
	a.SetString("543513587914353153599", 10)
	b.SetString("4324235325252598702321323132", 10)
	
	fmt.Println(new(big.Int).Add(a, b)) // сложение
	fmt.Println(new(big.Int).Sub(a, b)) // вычитание
	fmt.Println(new(big.Int).Mul(a, b)) // умножение
	fmt.Println(new(big.Int).Div(a, b)) // деление
	fmt.Println(new(big.Int).Div(b, a))
}

//	тип big.Int:

//	type Int struct {
//		neg bool // знак
//		abs nat  // значение
//	}
//	type nat []Word
//	type Word uint
//	max uint - 18446744073709551615
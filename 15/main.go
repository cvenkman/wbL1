package main

// К каким негативным последствиям может привести данный фрагмент
// кода, и как это исправить? Приведите корректный пример реализации.


// var justString string
// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func mainErr() {
// 	someFunc()
// }


// проблема: мы используем 100 элементов из выделенных 1024 (1 << 10 = 1024)

func main() {
	// 1 << 10 = 1024
	someFunc()
}

func someFunc() string {
	v := createHugeString(1 << 10)
	justString := v[:100]
	return justString // возвращать вместо глобальной перемнной
}

func createHugeString(size int) string {
	str := make([]rune, size)
	for i := range str {
		str[i] = '~'
	}
	return string(str)
}
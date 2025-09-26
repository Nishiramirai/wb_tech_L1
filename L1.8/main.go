package main

import "fmt"

// Установка бита в числе

// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
//  Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

func main() {
	var n int64 = 5

	bitNumber := 8
	bitMode := 1

	switch bitMode {
	case 1:
		n |= 1 << bitNumber
		fmt.Println(n)
	case 0:
		// 0101	0100

		// 0101
		// 1011
		// 0100
	default:
		fmt.Println("Error")
	}

}

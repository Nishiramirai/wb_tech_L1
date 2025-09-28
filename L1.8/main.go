package main

import "fmt"

// Установка бита в числе

// Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
//  Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

func main() {
	var num int64
	var pos int64
	var bitValue uint

	fmt.Printf("Enter number, bit position and new bit value:")
	if _, err := fmt.Scanln(&num, &pos, &bitValue); err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// В условии задания как я понимаю биты считаются с 1, а не нуля, поэтому делаю -1
	res, err := setBit(num, pos-1, bitValue)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("Source number: %d - %b\n", num, num)
	fmt.Printf("Result: %d - %b\n", res, res)

}

func setBit(n int64, pos int64, value uint) (int64, error) {
	// Проверяем позицию на вхождение в диапазон
	if pos >= 64 || pos < 0 {
		return 0, fmt.Errorf("invalid bit position: %d. must be in range [0, 63]", pos)
	}

	switch value {
	case 1:
		n |= (1 << pos)
	case 0:
		n &^= (1 << pos)
	default:
		return 0, fmt.Errorf("invalid bit value")
	}

	return n, nil
}

package main

import "fmt"

// Удаление элемента слайса

// Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента (copy(slice[i:], slice[i+1:]))
//  и уменьшить длину слайса на 1.

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(s)
	s = remove(s, 5)
	fmt.Println(s)
}

func remove(slice []int, i int) []int {

	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])

	slice[len(slice)-1] = 0

	return slice[:len(slice)-1]
}

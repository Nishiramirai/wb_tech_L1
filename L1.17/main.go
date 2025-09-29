package main

import "fmt"

// Бинарный поиск

// Реализовать алгоритм бинарного поиска встроенными методами языка.
//  Функция должна принимать отсортированный слайс и искомый элемент,
//  возвращать индекс элемента или -1, если элемент не найден.

// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

func main() {
	lst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(lst, 5))
	fmt.Println("Hello world")
}

func binarySearch(list []int, item int) int {
	// Реализован ирератвно
	low := 0
	high := len(list) - 1

	for low <= high {
		// трюк для обхода переполнения. Но пу сути делает
		//тоже самое что и low + high / 2
		mid := low + (high-low)/2
		guess := list[mid]
		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

package main

import "fmt"

// Реализовать алгоритм быстрой сортировки массива
// встроенными средствами языка. Можно использовать рекурсию.

// Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел.
// Для выбора опорного элемента можно взять середину или первый элемент.

func main() {
	s := []int{7, 6, 4, 3, 5, 3, 1, 4, 5}
	res := quickSort(s)
	fmt.Println(s)
	fmt.Println(res)

}

func quickSort(s []int) []int {
	if len(s) == 0 {
		return []int{}
	}

	newSlice := make([]int, len(s))
	copy(newSlice, s)

	quickSortRecurcive(newSlice, 0, len(newSlice)-1)
	return newSlice
}

func quickSortRecurcive(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSortRecurcive(arr, low, p-1)
		quickSortRecurcive(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

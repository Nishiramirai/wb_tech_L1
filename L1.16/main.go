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

func quickSortInner(s []int, low, high int) []int {

	if low < high {

		var p int

		s, p = partition(s, low, high)

		s = quickSortInner(s, low, p-1)

		s = quickSortInner(s, p+1, high)

	}

	return s

}

func quickSort(s []int) []int {
	return quickSortInner(s, 0, len(s)-1)
}

func partition(arr []int, low, high int) ([]int, int) {

	pivot := arr[high]

	i := low

	for j := low; j < high; j++ {

		if arr[j] < pivot {

			arr[i], arr[j] = arr[j], arr[i]

			i++

		}

	}

	arr[i], arr[high] = arr[high], arr[i]

	return arr, i

}

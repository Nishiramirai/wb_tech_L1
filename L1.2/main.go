package main

import (
	"fmt"
	"sync"
)

// Конкурентное возведение в квадрат

// Написать программу, которая конкурентно рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

// Более современный подход, увидел в Go by Example
func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, val := range arr {
		// Добавили в 1.25, довольно удобно, сам делает wg.Add и wg.Done в горутине
		wg.Go(func() {
			printSquare(val)
		})
	}

	wg.Wait()
}

func printSquare(n int) {
	fmt.Println(n * n)
}

// Решил сделать еще одную версию более классическим способом
// func main() {
// 	arr := [...]int{2, 4, 6, 8, 10}
// 	var wg sync.WaitGroup

// 	// Сначала делал wg.Add по 1 на каждую итерацию, но ведь количество действий у нас заранее известно,
// 	// поэтому можно сделать так
// 	wg.Add(len(arr))

// 	// Проходим по массиву и запускаем каждую функцию подсчета квадрата в отдельной горутине
// 	for _, val := range arr {
// 		go printSquare(val, &wg)
// 	}

// 	// Ждем когда все горутины завершатся
// 	wg.Wait()
// }

// func printSquare(n int, wg *sync.WaitGroup) {
// 	// Закидываем wg.Done в defer, чтобы при выходе из функции он гарантированно выполнился
// 	defer wg.Done()

// 	fmt.Println(n * n)

// }

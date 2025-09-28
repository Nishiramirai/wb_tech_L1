package main

import "fmt"

// Конвейер чисел

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива,
//  во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout.
//  То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
//  Убедитесь, что чтение из второго канала корректно завершается.

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := gen(arr...)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()

	return out
}

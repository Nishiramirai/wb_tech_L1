package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"
)

// Работа нескольких воркеров

// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

func main() {
	const sendDelay = 250

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	// Получаем количество воркеров из argv и конвертируем в int
	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil || workerCount < 1 {
		fmt.Println("Invalid worker count, must be positive integer")
		return
	}

	ch := make(chan int)

	// Запускаем заданное число воркеров
	for i := 0; i < workerCount; i++ {
		go worker(i+1, ch)
	}

	// Отправляем сообщения в канал с определенной периодичностью
	for {
		ch <- rand.IntN(100)
		time.Sleep(time.Millisecond * sendDelay)
	}
}

func worker(id int, ch <-chan int) {
	fmt.Printf("Worker %d started\n", id)
	for v := range ch {
		fmt.Println(v)
	}
}

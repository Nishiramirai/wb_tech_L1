package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Завершение по Ctrl+C

// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

func main() {
	const sendDelay = 250 * time.Millisecond

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	workerCount, err := strconv.Atoi(os.Args[1])
	if err != nil || workerCount < 1 {
		fmt.Println("Invalid worker count, must be positive integer")
		return
	}

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(workerCount)

	// Запускаем заданное число воркеров
	for i := 0; i < workerCount; i++ {
		go worker(i+1, ch, wg)
	}

	// Наиболее идиоматичным способом является контекст, поэтому выбрал его.
	// Ждем SIGINT и SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Отправляем сообщения в канал с определенной периодичностью
	for {
		select {
		// Если получаем сигнал завершения, закрываем канал и ждем завершения всех воркеров
		case <-ctx.Done():
			fmt.Println("Shutting down gracefully ...")
			close(ch)
			wg.Wait()
			fmt.Println("All workers have finished. Program terminated")
			return
		default:
			ch <- rand.IntN(100)
			time.Sleep(sendDelay)
		}
	}
}

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)

	// Бесконечно читает канал и выводит его содержимое, пока канал не закроют
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Printf("Worker %d finished\n", id)
}

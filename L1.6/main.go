package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Остановка горутины

// Реализовать все возможные способы остановки выполнения горутины.
//  Классические подходы: выход по условию, через канал уведомления,
//  через контекст, прекращение работы runtime.Goexit() и др.

func main() {
	wg := sync.WaitGroup{}
	wg.Add(6)

	// 1)Просто что-то делает и после этого завершается
	go func() {
		defer wg.Done()
		num := 1
		for i := range 10 {
			num *= i
		}
		fmt.Println("Goroutine 1 done")
	}()

	// 2)Выход по условию
	go func() {
		defer wg.Done()
		i := 1
		for {
			i += 1
			if i%50 == 0 {
				fmt.Println("Goroutine 2 done")
				return
			}
		}
	}()

	// 3)Канал уведомления
	doneCh := make(chan struct{})
	time.AfterFunc(time.Second*3, func() {
		close(doneCh)
	})

	go func(doneCh <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-doneCh:
				// Устали, можно и отдохнуть
				fmt.Println("Goroutine 3 done")
				return
			default:
				// Симулируем работу в поте лица
				time.Sleep(1 * time.Second)
			}
		}
	}(doneCh)

	// 4)Контекст
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine 4 done")
				return
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 5)runtime.Goexit()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 1)

		fmt.Println("Goroutime 5 done")
		runtime.Goexit()
	}()

	// 6)Закрытие читаемого канала
	jobs := make(chan int, 5)

	go func() {
		defer wg.Done()

		// Будет читать канал, пока его не закроют
		for job := range jobs {
			fmt.Printf("Processing job %d\n", job)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Goroutine 6 done")
	}()

	// Отправляем задачи и закрываем канал, чтобы горутина консьюмер завершилась
	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Program finished")
}

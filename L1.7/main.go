package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Конкурентная запись в map

// Реализовать безопасную для конкуренции запись данных в структуру map.

type safeMap struct {
	data map[int]int
	mu   sync.Mutex
}

const workerCount = 10 // Очень не люблю magic numbers, так что вынес чтобы не дублировалось в коде

func main() {
	mp := &safeMap{data: make(map[int]int)}
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg.Add(workerCount)
	fmt.Println("Work started")
	for range workerCount {
		go worker(mp, ctx, wg)
	}

	wg.Wait()
	fmt.Println("Work done")
	for i := range workerCount {
		fmt.Printf("%d: %d\n", i, mp.data[i])
	}
}

func worker(mp *safeMap, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Генерируем случайный индекс и пишем в мапу по нему,
			// предварительно заблокировав мьютекс
			// Конкретно с тем что я тут делаю справились бы и атомики, но это просто как
			// пример работы с мапой из разных горутин
			i := rand.IntN(10)
			mp.mu.Lock()
			mp.data[i] += 1
			mp.mu.Unlock()
			time.Sleep(250 * time.Millisecond)
		}
	}
}

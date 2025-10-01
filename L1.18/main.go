package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Конкурентный счетчик

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
//  По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

type counter struct {
	cnt atomic.Int64 // Для счетчика лучше использовать гораздо более легковесный атомик, чем мьютекс
}

const workers int = 10
const incrementCount int = 1000

func main() {
	wg := &sync.WaitGroup{}
	c := counter{}

	wg.Add(workers)
	for i := range workers {
		go worker(i, &c, wg)
	}

	wg.Wait()
	fmt.Println("All workers finished work")
	fmt.Printf("Final counter value: %d\n", c.cnt.Load())
}

func worker(id int, c *counter, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("Worker %d done\n", id)

	for range incrementCount {
		c.cnt.Add(1)
	}
}

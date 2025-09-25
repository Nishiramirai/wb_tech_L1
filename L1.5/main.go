package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Таймаут на канал

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

const (
	timeoutValue = time.Second * 5
	sendDelay    = time.Second * 1
)

func main() {
	ch := make(chan int)

	go func() {
		for {
			ch <- rand.IntN(100)
			time.Sleep(sendDelay)
		}
	}()

	timeout := time.After(timeoutValue)

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			// По хорошему всегда стоит делать graceful shutdown (тут горутина остается незавершенной),
			// но мне кажется суть этого задания в том, что мы якобы ждем что-то от нас независящее
			// и тут это не особо уместно, так как в задании рекомендуется time.After.
			// Но можно было бы это сделать через context.WithTimeout
			fmt.Println("Timeout")
			return
		}
	}
}

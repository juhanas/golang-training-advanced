package helpers

import (
	"sync"
)

type ChanIntOrString interface {
	chan string | chan int
}

func CloseChan[C ChanIntOrString](chanToClose C, wg *sync.WaitGroup) {
	wg.Wait()
	close(chanToClose)
}

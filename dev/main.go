package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func main() {
	var count int64

	block := make(chan bool)
	tempo := make(chan float64)
	tick := time.NewTicker(time.Duration(60000/60.0) * time.Millisecond)
	go func() {
		// loop:
		for {
			select {
			case newTempo := <-tempo:
				tick.Reset(time.Duration(60000/newTempo) * time.Millisecond)
			case tim := <-tick.C:
				fmt.Println(tim)
				if count == 100 {
					tick.Stop()
					close(tempo)
					// break loop
					block <- true
				}
				log.Println(atomic.AddInt64(&count, 1))
			}
		}
	}()

	time.Sleep(10 * time.Second)
	tempo <- 120.5

	<-block
}

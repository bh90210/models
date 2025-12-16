package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan []byte, 1000)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- []byte{144, 60, 127}
		time.Sleep(2 * time.Second)
		ch <- []byte{1}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

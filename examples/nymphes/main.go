package main

import (
	"fmt"

	"github.com/bh90210/models/nymphes"
)

func main() {
	ns, err := nymphes.NewProject()
	if err != nil {
		panic(err)
	}
	defer ns.Close()

	in := ns.Incoming()
	go func() {
		for {
			val := <-in
			println("MIDI IN:", val)
		}
	}()

	for i := range 5 {
		ns.Note(nymphes.Channel, 50, 120, 1500)
		ns.Note(nymphes.Channel, 60, 120, 1500)
		ns.Note(nymphes.Channel, 70, 120, 1500)
		ns.Note(nymphes.Channel, 75, 120, 1500)

		fmt.Println(i)
	}
}

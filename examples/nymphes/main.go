package main

import (
	"fmt"

	"github.com/bh90210/models/nymphes"
)

func main() {
	nymphesSynth, err := nymphes.NewProject()
	if err != nil {
		panic(err)
	}
	defer nymphesSynth.Close()

	in := nymphesSynth.Incoming()
	go func() {
		for {
			val := <-in
			println("MIDI IN:", val)
		}
	}()

	for i := range 5 {
		nymphesSynth.Note(0, 50, 120, 1500)
		nymphesSynth.Note(0, 60, 120, 1500)
		nymphesSynth.Note(0, 70, 120, 1500)
		nymphesSynth.Note(0, 75, 120, 1500)

		fmt.Println(i)
	}
}

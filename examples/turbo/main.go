package main

import (
	"bytes"
	"fmt"

	m "github.com/bh90210/models/cycles"
	"github.com/bh90210/models/turbo"
)

func main() {
	turboAlesis, err := turbo.NewProject()
	if err != nil {
		panic(err)
	}
	defer turboAlesis.Close()

	p, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	in := p.Incoming()
	go func(in chan []byte) {
		for {
			val := <-in
			fmt.Println("MIDI IN:", val)
		}
	}(in)

	var noteLength int = 50

	in = turboAlesis.Incoming()
	for {
		d := <-in
		switch {
		case bytes.Contains(d, []byte{turbo.Drums["Snare"]}):
			println("Snare hit", d)
			if !bytes.Contains(d, []byte{153, 38, 0}) {
				p.Note(1, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["Kick"]}):
			println("Kick hit")
			if !bytes.Contains(d, []byte{153, 36, 0}) {
				p.Note(0, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["HiHatClosed"]}):
			println("HiHatClosed hit")
		case bytes.Contains(d, []byte{turbo.Drums["HiHatOpen"]}):
			println("HiHatOpen hit")
		case bytes.Contains(d, []byte{turbo.Drums["Tom1"]}):
			println("Tom1 hit")
			if !bytes.Contains(d, []byte{153, 48, 0}) {
				p.Note(2, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["Tom2"]}):
			println("Tom2 hit")
			if !bytes.Contains(d, []byte{153, 45, 0}) {
				p.Note(2, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["Tom3"]}):
			println("Tom3 hit")
			if !bytes.Contains(d, []byte{153, 43, 0}) {
				p.Note(3, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["Crash"]}):
			println("Crash hit")
			if !bytes.Contains(d, []byte{153, 49, 0}) {
				p.Note(4, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["Ride"]}):
			println("Ride hit")
			if !bytes.Contains(d, []byte{153, 51, 0}) {
				p.Note(5, 36, int8(d[2]), float64(noteLength))
			}

		case bytes.Contains(d, []byte{turbo.Drums["HiHatFoot"]}):
			println("HiHatFoot hit")
		default:
			println("Unknown MIDI IN:", d)
		}
	}
}

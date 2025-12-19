package main

import (
	"github.com/bh90210/models/midicom"
	"github.com/bh90210/models/nymphes"
	"github.com/bh90210/models/pattern"
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

	var poly pattern.Poly

	// Generate the notes (melody.)

	// Generate the duration + velocity (rhythm.) for each note.

	notes1 := []pattern.Note{
		{Note: 50, Duration: 500, Velocity: 100},
		{Note: 50 + midicom.Note(pattern.Major3rd), Duration: 500, Velocity: 100},
		{Note: 50 + midicom.Note(pattern.Perfect5th), Duration: 500, Velocity: 100},
		{Note: 50 + midicom.Note(pattern.Major7th), Duration: 500, Velocity: 100},
	}

	pat1 := pattern.Pattern{
		Midicom: ns,
		Notes:   notes1,
		Channel: nymphes.Channel,
		Meta: pattern.Meta{
			Synth: nymphes.Nymphes,
			Part:  "voice1-start",
		},
	}

	pat2 := pat1.Shift(pattern.Perfect5th)

	poly.AddPattern(0, pat1)
	poly.AddPattern(1, pat2)

	err = pattern.Play(poly)
}

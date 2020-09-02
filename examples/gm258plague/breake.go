package main

import (
	"time"

	elektron "github.com/bh90210/elektronmodels"
)

func Intro(p *elektron.Project) (*elektron.Pattern, error) {
	// var klimax []Note = []Note{C1, G1, F1, D1}
	// var chords []Chord = []Chord{Major6, MajorAdd9, MajorMajor7, MinorMinor9no5, MajorMajor7, MajorMinor9no5, MajorMajor76no5, MinorMinor7}

	// loopy := func() {
	// 	for {
	// 		c.Cycles.Note(NewNoteTrack(T6, time.Duration(100*time.Millisecond)), 50, 120,
	// 			map[Parameter]int64{
	// 				REVERB:       60,
	// 				REBERBTONE:   80,
	// 				REVERBZISE:   80,
	// 				DELAY:        0,
	// 				DECAY:        50,
	// 				SHAPE:        int64(5),
	// 				SWEEP:        10,
	// 				CHANCE:       100,
	// 				GATE:         0,
	// 				DELAYTIME:    10,
	// 				COLOR:        120,
	// 				LFODEST:      int64(5),
	// 				LFOWAVEFORM:  0,
	// 				LFOMULTIPIER: 14,
	// 				LFODEPTH:     120,
	// 			},
	// 		)
	// 	}
	// }

	p.Patterns[0].Trigs[0].Track = elektron.NewNoteTrack(elektron.T6, time.Duration(100*time.Millisecond))
	p.Patterns[0].Trigs[0].CC = elektron.NewCC(elektron.CCT6, map[elektron.Parameter]int64{
		elektron.REVERB: 100,
	})

	return p.Patterns[0], nil
}

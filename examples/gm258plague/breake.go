package main

import (
	"time"

	e "github.com/bh90210/elektronmodels"
)

func Intro() *e.Track {
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
	trig := e.NewTrig(0)
	trig.CC(
		map[e.Parameter]int64{
			e.REVERB: 100,
		})
	trig.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig2 := e.NewTrig(2)
	trig2.CC(
		map[e.Parameter]int64{
			e.REVERB: 100,
		})
	trig2.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig3 := e.NewTrig(4)
	trig3.CC(
		map[e.Parameter]int64{
			e.REVERB: 100,
		})
	trig3.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig4 := e.NewTrig(6)
	trig4.CC(
		map[e.Parameter]int64{
			e.REVERB: 100,
		})
	trig4.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig5 := e.NewTrig(8)
	trig5.CC(
		map[e.Parameter]int64{
			e.REVERB: 100,
		})
	trig5.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	endTrig := e.LastTrig(10)

	track1 := e.NewTrack(e.T1, trig, trig2, trig3, trig4, trig5, endTrig)

	return track1
}

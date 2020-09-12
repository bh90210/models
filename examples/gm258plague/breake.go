package main

import (
	"time"

	e "github.com/bh90210/elektronmodels"
)

func Intro() *e.Track {
	trig := e.NewTrig(0)
	// trig.CC(
	// 	map[e.Parameter]uint8{
	// 		// e.NOTE:   uint8(e.A4),
	// 		e.REBERBTONE:   80,
	// 		e.REVERBZISE:   80,
	// 		e.DELAY:        0,
	// 		e.DECAY:        50,
	// 		e.SHAPE:        uint8(5),
	// 		e.SWEEP:        10,
	// 		e.CHANCE:       0,
	// 		e.SWING:        0,
	// 		e.GATE:         0,
	// 		e.DELAYTIME:    10,
	// 		e.COLOR:        120,
	// 		e.LFODEST:      uint8(0),
	// 		e.LFOWAVEFORM:  0,
	// 		e.LFOMULTIPIER: 0,
	// 		e.LFODEPTH:     0,
	// 		e.VOLUMEDIST:   0,
	// 	})
	trig.Note(
		e.A3,
		100,
		time.Duration(100*time.Millisecond))

	trig2 := e.NewTrig(2)
	// trig2.CC(
	// 	map[e.Parameter]uint8{
	// 		e.REBERBTONE:   0,
	// 		e.REVERBZISE:   0,
	// 		e.DELAY:        0,
	// 		e.DECAY:        50,
	// 		e.SHAPE:        uint8(5),
	// 		e.SWEEP:        10,
	// 		e.CHANCE:       0,
	// 		e.SWING:        0,
	// 		e.GATE:         0,
	// 		e.DELAYTIME:    10,
	// 		e.COLOR:        120,
	// 		e.LFODEST:      uint8(0),
	// 		e.LFOWAVEFORM:  0,
	// 		e.LFOMULTIPIER: 0,
	// 		e.LFODEPTH:     0,
	// 		e.VOLUMEDIST:   0,
	// 	})
	trig2.Note(
		e.A4,
		100,
		time.Duration(100*time.Millisecond))

	endTrig := e.LastTrig(4)

	track1 := e.NewTrack(e.T1, trig, trig2, endTrig)

	return track1
}

package main

import (
	"time"

	e "github.com/bh90210/elektronmodels"
)

func Intro() *e.Track {
	var te []*e.Trig
	var it int
loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
				// e.NOTE:         uint8(it),
				// e.CPITCH:       uint8(it),
				// e.PAN:          uint8(it),
				e.CONTOUR:      uint8(it),
				e.REVERB:       uint8(it),
				e.REBERBTONE:   80,
				e.REVERBZISE:   80,
				e.DELAY:        uint8(it),
				e.DECAY:        uint8(it),
				e.SHAPE:        uint8(it),
				e.SWEEP:        uint8(it),
				e.CHANCE:       0,
				e.SWING:        0,
				e.GATE:         1,
				e.DELAYTIME:    uint8(it),
				e.COLOR:        uint8(it),
				e.LFODEST:      uint8(it),
				e.LFOWAVEFORM:  0,
				e.LFOMULTIPIER: 0,
				e.LFODEPTH:     uint8(it),
				e.VOLUMEDIST:   60,
			})
		trig.Note(
			e.A0+uint8(it),
			100,
			time.Duration(100*time.Millisecond))
		te = append(te, trig)
		if it == 99 {
			break loop
		}
		it++
	}
	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track1 := e.NewTrack(e.T1, te)

	return track1
}

package main

import (
	"log"
	"time"

	cycles "github.com/bh90210/elektronmodels"
)

func main() {
	example, _ := cycles.NewProject()
	defer example.Close()

	t1intro := Intro()

	example.Pattern(t1intro)
	example.Loop()

	if err := example.Play(); err != nil {
		log.Println(err)
	}
}

func Intro() *cycles.Track {
	trig := cycles.NewTrig(0)
	trig.CC(
		map[cycles.Parameter]uint8{
			// e.NOTE:   int64(e.A4),
			cycles.REBERBTONE:   80,
			cycles.REVERBZISE:   80,
			cycles.DELAY:        0,
			cycles.DECAY:        50,
			cycles.SHAPE:        5,
			cycles.SWEEP:        10,
			cycles.CHANCE:       0,
			cycles.SWING:        0,
			cycles.GATE:         0,
			cycles.DELAYTIME:    10,
			cycles.COLOR:        120,
			cycles.LFODEST:      0,
			cycles.LFOWAVEFORM:  0,
			cycles.LFOMULTIPIER: 0,
			cycles.LFODEPTH:     0,
		})
	trig.Note(
		cycles.A4,
		120,
		time.Duration(100*time.Millisecond))

	endTrig := cycles.LastTrig(10)

	var trigs []*cycles.Trig
	trigs = append(trigs, trig, trig, trig, trig, trig, endTrig)

	track1 := cycles.NewTrack(cycles.T1, trigs)

	return track1
}

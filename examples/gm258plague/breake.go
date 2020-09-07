package main

import (
	"time"

	e "github.com/bh90210/elektronmodels"
)

func Intro() *e.Track {
	trig := e.NewTrig(0)
	trig.CC(
		map[e.Parameter]int64{
			// e.NOTE:   int64(e.A4),
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig.Note(
		e.A4,
		120,
		time.Duration(100*time.Millisecond))

	trig2 := e.NewTrig(2)
	trig2.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig2.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig3 := e.NewTrig(4)
	trig3.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig3.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig4 := e.NewTrig(6)
	trig4.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig4.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	trig5 := e.NewTrig(8)
	trig5.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig5.Note(
		e.A0,
		100,
		time.Duration(100*time.Millisecond))

	endTrig := e.LastTrig(10)

	track1 := e.NewTrack(e.T4, trig, trig2, trig3, trig4, trig5, endTrig)

	return track1
}

func Intro2() *e.Track {
	trig := e.NewTrig(0)
	trig.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig.Note(
		e.A5,
		100,
		time.Duration(100*time.Millisecond))

	trig2 := e.NewTrig(2)
	trig2.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig2.Note(
		e.A4,
		100,
		time.Duration(100*time.Millisecond))

	trig3 := e.NewTrig(4)
	trig3.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig3.Note(
		e.A3,
		100,
		time.Duration(100*time.Millisecond))

	trig4 := e.NewTrig(6)
	trig4.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig4.Note(
		e.A4,
		100,
		time.Duration(100*time.Millisecond))

	trig5 := e.NewTrig(8)
	trig5.CC(
		map[e.Parameter]int64{
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        120,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig5.Note(
		e.A5,
		100,
		time.Duration(100*time.Millisecond))

	endTrig := e.LastTrig(50)

	track1 := e.NewTrack(e.T6, trig, trig2, trig3, trig4, trig5, endTrig)

	return track1
}

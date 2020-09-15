package main

import (
	"time"

	e "github.com/bh90210/elektronmodels"
)

var total int = 40

func t1Intro() *e.Track {
	var te []*e.Trig

	trig := e.NewTrig(0)
	trig.CC(
		map[e.Parameter]uint8{
			// e.PAN:          uint8(it),
			e.CYCLESPITCH:  64,
			e.CONTOUR:      0,
			e.REVERB:       120,
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        10,
			e.DECAY:        55,
			e.SHAPE:        e.MajorMajor7,
			e.SWEEP:        40,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         1,
			e.DELAYTIME:    10,
			e.COLOR:        0,
			e.LFODEST:      e.LFTUN,
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 1,
			e.LFODEPTH:     67,
			e.VOLUMEDIST:   60,
		})
	trig.Note(
		e.A1,
		100,
		time.Duration(80000*time.Millisecond))
	te = append(te, trig)

	var it int = 1
loop:
	for {
		trig2 := e.NewTrig(uint8(it))
		trig2.CC(
			map[e.Parameter]uint8{
				// e.PAN:          uint8(it),
				e.CONTOUR: uint8(it),
				e.SHAPE:   e.MajorMajor7,
				e.SWEEP:   uint8(it + 40),
				e.COLOR:   uint8(it),
			})
		te = append(te, trig2)

		if it == total {
			break loop
		}
		it++
	}

	endTrig := e.LastTrig(uint8(total) + 1)
	te = append(te, endTrig)

	track := e.NewTrack(e.T6, te)

	return track
}

func t2Intro() *e.Track {
	var te []*e.Trig
	var it int

	trig := e.NewTrig(uint8(0))
	trig.CC(
		map[e.Parameter]uint8{
			// e.PAN:          uint8(it),
			e.CONTOUR:      uint8(it),
			e.REVERB:       uint8(it),
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        uint8(it),
			e.DECAY:        uint8(it),
			e.SHAPE:        uint8(it),
			e.SWEEP:        70,
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

loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
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
		it++
		if it == total {
			break loop
		}
	}

	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track2 := e.NewTrack(e.T2, te)

	return track2
}

func t3Intro() *e.Track {
	var te []*e.Trig
	var it int

loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
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
		it++
		if it == total {
			break loop
		}
	}

	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track := e.NewTrack(e.T3, te)

	return track
}
func t4Intro() *e.Track {
	var te []*e.Trig
	var it int

loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
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
		it++
		if it == total {
			break loop
		}
	}

	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track := e.NewTrack(e.T4, te)

	return track
}
func t5Intro() *e.Track {
	var te []*e.Trig
	var it int

loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
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
		it++
		if it == total {
			break loop
		}
	}

	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track := e.NewTrack(e.T5, te)

	return track
}

func t6Intro() *e.Track {
	var te []*e.Trig
	var it int

loop:
	for {
		trig := e.NewTrig(uint8(it))
		trig.CC(
			map[e.Parameter]uint8{
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
		it++
		if it == total {
			break loop
		}
	}

	endTrig := e.LastTrig(uint8(it + 1))
	te = append(te, endTrig)
	track := e.NewTrack(e.T6, te)

	return track
}

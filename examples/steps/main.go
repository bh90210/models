package main

import (
	"log"
	"math/rand"
	"time"

	cycles "github.com/bh90210/elektronmodels"
)

type chord struct {
	note  uint8
	shape uint8
	bass  []uint8
}

var total int = 0

var (
	// 1-3-5 7-9-11-13
	// 5
	Bmaj7  = chord{note: cycles.B3, shape: cycles.MajorMajor76no5, bass: []uint8{cycles.B0, cycles.Fs1, cycles.As0}}
	D7     = chord{note: cycles.D4, shape: cycles.MajorMinor7, bass: []uint8{cycles.D1, cycles.A0, cycles.C1}}
	Gmaj7  = chord{note: cycles.G4, shape: cycles.MajorMajor7, bass: []uint8{cycles.G1, cycles.D1, cycles.A0, cycles.Fs1}}
	Bb7    = chord{note: cycles.Bf3, shape: cycles.MajorMinor7, bass: []uint8{cycles.Bf0, cycles.F1, cycles.Af1, cycles.C1}}
	Ebmaj7 = chord{note: cycles.Ef4, shape: cycles.MajorMajor9no5, bass: []uint8{cycles.Ef1, cycles.D1, cycles.F1}}
	Am7    = chord{note: cycles.A4, shape: cycles.MinorMinor9no5, bass: []uint8{cycles.A0, cycles.E1, cycles.B0}}
	Fs7    = chord{note: cycles.Fs4, shape: cycles.MajorMinor7, bass: []uint8{cycles.Fs1, cycles.Cs1, cycles.E1, cycles.Gs1}}
	Csm7   = chord{note: cycles.Cs4, shape: cycles.MinorMinor7Sus4, bass: []uint8{cycles.Cs1, cycles.B0, cycles.Ds1, cycles.As0}}
	Fm7    = chord{note: cycles.F4, shape: cycles.Minor6, bass: []uint8{cycles.F1, cycles.Af1, cycles.C1, cycles.Ef1, cycles.G1, cycles.D1}}

	array = map[int]chord{
		0:  Bmaj7,
		1:  D7,
		2:  Gmaj7,
		3:  Bb7,
		4:  Ebmaj7,
		5:  Ebmaj7,
		6:  Am7,
		7:  D7,
		8:  Gmaj7,
		9:  Bb7,
		10: Ebmaj7,
		11: Fs7,
		12: Bmaj7,
		13: Bmaj7,
		14: Fm7,
		15: Bb7,
		16: Ebmaj7,
		17: Ebmaj7,
		18: Am7,
		19: D7,
		20: Gmaj7,
		21: Gmaj7,
		22: Csm7,
		23: Fs7,
		24: Bmaj7,
		25: Bmaj7,
		26: Fm7,
		27: Bb7,
		28: Ebmaj7,
		29: Ebmaj7,
		30: Csm7,
		31: Fs7,
	}
)

func piano() *cycles.Track {
	var te []*cycles.Trig

	source := rand.NewSource(666)
	r1 := rand.New(source)

	for k, v := range array {
		trig := cycles.NewTrig(uint8(k))
		trig.CC(
			map[cycles.Parameter]uint8{
				cycles.CYCLESPITCH: 63,
				cycles.DECAY:       40,
				cycles.COLOR:       86,
				cycles.SHAPE:       v.shape,
				cycles.SWEEP:       1,
				cycles.CONTOUR:     uint8(r1.Intn(125)),
				cycles.DELAY:       10,
				cycles.REVERB:      25,
				cycles.LFODEST:     0,
				cycles.VOLUMEDIST:  63,
				cycles.GATE:        1,
				cycles.PUNCH:       0,
			})
		trig.Note(
			v.note,
			100,
			time.Duration(4000*time.Millisecond))
		te = append(te, trig)
	}

	trig := cycles.LastTrig(uint8(len(array)) + 1)
	te = append(te, trig)

	track := cycles.NewTrack(cycles.T6, te)

	return track
}

func bass() *cycles.Track {
	var te []*cycles.Trig

	// source := rand.NewSource(666)
	// r1 := rand.New(source)

	for k, v := range array {
		trig := cycles.NewTrig(uint8(k))
		trig.Note(
			v.bass[rand.Intn(2)],
			100,
			time.Duration(4000*time.Millisecond))
		te = append(te, trig)
	}

	trig := cycles.LastTrig(uint8(len(array)) + 1)
	te = append(te, trig)

	track := cycles.NewTrack(cycles.T4, te)

	return track
}

func main() {
	steps, err := cycles.NewProject()
	if err != nil {
		log.Fatal(err)
	}
	defer steps.Close()

	// steps.Pattern(piano())
	steps.Pattern(piano(), bass())

	steps.Loop()
	if err := steps.Play(); err != nil {
		log.Println(err)
	}

	time.Sleep(1 * time.Second)
}

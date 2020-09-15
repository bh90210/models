package main

import (
	"log"
	"math/rand"
	"time"

	cycles "github.com/bh90210/elektronmodels"
)

var total int = 0

func t1Intro() *cycles.Track {
	var te []*cycles.Trig

	type chord struct {
		note  uint8
		shape uint8
	}

	Bmaj7 := chord{note: cycles.B3, shape: cycles.MajorMajor76no5}
	D7 := chord{note: cycles.D4, shape: cycles.MajorMinor7b9no5}
	Gmaj7 := chord{note: cycles.G4, shape: cycles.MajorMajor7}
	Bb7 := chord{note: cycles.Bf3, shape: cycles.MajorMinor7}
	Ebmaj7 := chord{note: cycles.Ef4, shape: cycles.MajorMajor9no5}
	Am7 := chord{note: cycles.A4, shape: cycles.MinorMinor9no5}
	Fs7 := chord{note: cycles.Fs4, shape: cycles.MajorMinor7}
	Csm7 := chord{note: cycles.Cs4, shape: cycles.MinorMinor7Sus4}
	Fm7 := chord{note: cycles.F4, shape: cycles.Minor6}

	array := map[int]chord{
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

func main() {
	gm258plague, err := cycles.NewProject()
	if err != nil {
		log.Fatal(err)
	}
	defer gm258plague.Close()

	// gm258plague.Pattern(t1Intro(), t2Intro(), t3Intro(), t4Intro(), t5Intro(), t6Intro())
	gm258plague.Pattern(t1Intro())
	gm258plague.Pattern(t1Intro())

	gm258plague.Loop()
	if err := gm258plague.Play(); err != nil {
		log.Println(err)
	}

	time.Sleep(1 * time.Second)
}

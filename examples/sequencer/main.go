package main

import (
	m "github.com/bh90210/models"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

func main() {
	// new project
	project, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer project.Close()

	// intro pattern
	p0 := project.Pattern(INTRO).
		Scale(m.PTN, 30, 1.0, 0).
		Tempo(110.0)

	// track 1
	preT1 := make(map[m.Parameter]int8)
	preT1[m.DECAY] = 125
	preT1[m.COLOR] = 0
	preT1[m.SHAPE] = 126

	preT1[m.PUNCH] = 0
	preT1[m.GATE] = 1

	preT1[m.SWEEP] = 1
	preT1[m.CONTOUR] = 1

	preT1[m.DELAY] = 10
	preT1[m.DELAYTIME] = 17
	preT1[m.DELAYFEEDBACK] = 20
	preT1[m.REVERB] = 125
	preT1[m.REVERBSIZE] = 100
	preT1[m.REVERBTONE] = 75

	preT1[m.LFODEST] = 0

	preT12 := make(map[m.Parameter]int8)
	for k, v := range preT1 {
		switch k {
		case m.GATE:
			preT12[m.GATE] = 0
		case m.DECAY:
			preT12[m.DECAY] = 120
		case m.DELAY:
			preT12[m.DELAY] = 90
		case m.DELAYFEEDBACK:
			preT12[m.DELAYFEEDBACK] = 30
		case m.DELAYTIME:
			preT12[m.DELAYTIME] = 90
		default:
			preT12[k] = v
		}
	}
	preT12[m.MACHINE] = int8(m.SNARE)

	// intro pattern
	p0t1 := p0.Track(m.T1).Preset(preT1)

	// kick := []int{0, 2, 3}
	// for _, v := range kick {
	// 	p0t1.Trig(v).Note(m.A3, 1000, 125)
	// }

	p0t1.Trig(0).Note(m.C4, 5000, 125)
	p0t1.Trig(4).Note(m.D4, 5000, 125).Lock(preT12)
	p0t1.Trig(5).Lock(preT12)
	p0t1.Trig(6).Lock(preT12)
	p0t1.Trig(7).Lock(preT12)
	p0t1.Trig(8).Lock(preT12)
	p0t1.Trig(9).Lock(preT12)
	p0t1.Trig(10).Lock(preT12)
	p0t1.Trig(11).Lock(preT12)
	p0t1.Trig(12).Lock(preT12)
	// preT12 = preT1
	// p0t1.Trig(14).Note(m.G2, 10, 125).Lock(preT12)
	// p0t1.Trig(3).Note(m.G4, 1000, 125)
	// p0t1.Trig(5).Note(m.G3, 1000, 125)

	// track 2
	// p0t2 := p0.Track(m.T2).Preset(preT12)
	// p0t2.Trig(5).Note(m.D4, 100, 125)
	// p0t2.Trig(5).Note(m.D4, 100, 125).Lock(preT12)
	// snare := []int{2, 6, 10, 14}
	// for _, v := range snare {
	// 	p0t2.Trig(v).Note(m.A4, 50, 60)
	// }

	// // track 3
	// p0t3 := p0.Track(m.T3)
	// hh := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	// for _, v := range hh {
	// 	p0t3.Trig(v).Note(m.A4, 50, 60).Nudge(float64(rand.Intn(12)))
	// }

	// // track 4
	// p0t4 := p0.Track(m.T4)
	// perc := []int{1, 3, 5, 7, 9, 11, 13, 15}
	// for _, v := range perc {
	// 	p0t4.Trig(v).Note(m.A4, 50, 60).Nudge(float64(rand.Intn(20)))
	// }

	// // track 5
	// p0t5 := p0.Track(m.T5)
	// tone := []int{1, 5, 9, 13}
	// for _, v := range tone {
	// 	p0t5.Trig(v).Note(m.A4, 50, 60).Nudge(float64(rand.Intn(15)))
	// }

	// // track 6
	// p0t6 := p0.Track(m.T6)
	// chord := []int{3, 15}
	// for _, v := range chord {
	// 	p0t6.Trig(v).Note(m.A4, 50, 60)
	// }

	//verse

	// track 1

	// track 2

	// track 3

	// track 4

	// track 5

	// track 6

	// chorus

	// track 1

	// track 2

	// track 3

	// track 4

	// track 5

	// track 6

	// outro

	// track 1

	// track 2

	// track 3

	// track 4

	// track 5

	// track 6

	// play
	project.Play(INTRO)
	// project.Chain(INTRO, VERSE, CHORUS, OUTRO).Play()
}

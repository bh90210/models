package main

import (
	"math/rand"

	m "github.com/athenez/models"
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
		Scale(m.TRK, 4, 1.0, 0). // problem: scale does not work
		Tempo(120.0)

	// track 1
	p0t1 := p0.Track(m.T1)
	p0t1.Scale(15, 2.0) // problem: len\gth does not work, it follows pattern's setting
	kick := []int{0, 4, 8, 12}
	for _, v := range kick {
		p0t1.Trig(v).Note(m.A4, 50, 60)
	}

	// // track 2
	// p0t2 := p0.Track(m.T2)
	// snare := []int{2, 6, 10, 14}
	// for _, v := range snare {
	// 	p0t2.Trig(v).Note(m.A4, 50, 60)
	// }

	// track 3
	p0t3 := p0.Track(m.T3)
	p0t3.Scale(15, 2.0)
	hh := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	for _, v := range hh {
		p0t3.Trig(v).Note(m.A4, 50, 60).Nudge(float64(rand.Intn(12)))
	}

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
	// chord := []int{3, 7, 15}
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

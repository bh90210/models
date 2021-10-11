package main

import (
	m "github.com/bh90210/models"
)

const (
	INTRO int = iota
	VERSE
)

func main() {
	// presets
	// kick := make(map[m.Parameter]uint8)
	// kick[m.]

	project, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer project.Close()

	intro := project.Pattern(INTRO)

	t1 := intro.Track(m.T1)
	t1.Trig(0)
	t1.Trig(8)

	t2 := intro.Track(m.T2)
	t2.Trig(4)
	t2.Trig(12)

	t3 := intro.Track(m.T3)
	t3.Trig(0)
	t3.Trig(4)
	t3.Trig(8)
	t3.Trig(12)

	t4 := intro.Track(m.T4)
	t4.Trig(5)

	verse := project.Pattern(VERSE)

	t1 = verse.Track(m.T1)
	t1.Trig(0)
	t1.Trig(6)
	t1.Trig(8)
	t1.Trig(14)

	t2 = verse.Track(m.T2)
	t2.Trig(4)
	t2.Trig(12)

	t3 = verse.Track(m.T3)
	t3.Trig(0)
	t3.Trig(2)
	t3.Trig(4)
	t3.Trig(6)
	t3.Trig(8)
	t3.Trig(10)
	t3.Trig(12)
	t3.Trig(14)

	t4 = verse.Track(m.T4)
	t4.Trig(5)
	t4.Trig(11)

	project.Chain(INTRO, INTRO, INTRO, VERSE).Play()
}

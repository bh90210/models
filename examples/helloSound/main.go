package main

import (
	em "github.com/bh90210/elektronmodels"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

func main() {
	// start a new project
	p, err := em.NewProject(em.CYCLES)
	if err != nil {
		panic(err)
	}

	// test
	// p.CopyPattern(CHORUS, VERSE)
	p.Pattern(INTRO).Scale(em.PTN, 12, 1.0, 0).Tempo(150.0)
	// p.Pattern(INTRO).ScaleMod(em.TRK)

	// p.Pattern(INTRO).Track(em.T1)

	// p.Tempo(150.0)

	p.Pattern(INTRO).Track(em.T1)

	p.Pattern(INTRO).Track(em.T1).Trig(0)
	// p.Pattern(INTRO).Track(em.T1).Tempo(125.0)
	// p.Pattern(INTRO).Track(em.T1).Trig(1)

	// p.Pattern(INTRO).Track(em.T2).Trig(2)
	// p.Pattern(INTRO).Track(em.T2).Trig(3)

	// p.Pattern(INTRO).Track(em.T3).Trig(4)
	// p.Pattern(INTRO).Track(em.T3).Trig(5)

	// p.Pattern(INTRO).Track(em.T4).Trig(6)
	// p.Pattern(INTRO).Track(em.T4).Trig(7)

	// p.Pattern(INTRO).Track(em.T5).Trig(8)
	// p.Pattern(INTRO).Track(em.T5).Trig(9)

	// p.Pattern(INTRO).Track(em.T6).Trig(10)
	// p.Pattern(INTRO).Track(em.T6).Trig(11)

	// p.Pattern(INTRO).CopyTrack(em.T3, em.T4)

	preset := make(map[em.Parameter]int8)
	// p.Pattern(INTRO).Track(em.T1).Preset(preset)
	// p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset)

	// p.Chain(INTRO, INTRO, VERSE, CHORUS).Play()

	p.Play(INTRO)
	// p.Next(INTRO)

	// p.Stop()

	p.Free.Preset(preset).Note(em.A0, 0.5, 120).Trig(em.T1)

	p.Close()
}

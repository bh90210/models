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
	preset := make(em.Preset)
	preset[0] = 8
	preset[em.COLOR] = 100
	preset.SetParameter(em.COLOR, 100)

	ll := new(em.Lock)
	ll.Preset = preset

	// locks := make([]*em.Lock, 0)
	// locks = append(locks, ll)

	// start a new project
	p := em.NewProject(em.CYCLES)
	p.PatternInit(INTRO)
	p.Pattern[INTRO].T1.Preset = preset
	t1 := p.Pattern[INTRO].T1
	t1.Preset = preset

	// p.Patterns[1].Tracks[em.T6].CopyTrack(p0.Tracks[em.T1])
	// p.Patterns[0].Tracks[em.T1].Trigs[0].SetLock(ll)
	p.Play()

	// can be used without a number too - if used without a number and there is no next currently playing pattern keeps on looping
	// if used and not found, an empty default pattern should be returned - silence
	// p.Next()
	// p.Next(2)
	// p.Next(END)
	// Second number indicates jump to specific pattern number rather the next in line.
	// p.Next(END, 5)

	// p.Pause()
	// p.Stop()

	p.Close()
}

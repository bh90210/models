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
	preset[em.COLOR] = 100

	lock := new(em.Lock)
	lock.Preset = preset

	// locks := make([]*em.Lock, 0)
	// locks = append(locks, ll)

	// start a new project
	p := em.NewProject(em.CYCLES)
	p.InitPattern(INTRO)
	p.Pattern[INTRO].T1.Preset = preset

	// copy pattern
	p.InitPattern(VERSE)
	p.Pattern[INTRO].CopyPattern(p.Pattern[VERSE])

	// track
	p.Pattern[INTRO].T1.SetScale(em.PTN, 16, 4, 0)
	p.Pattern[INTRO].T1.SetPreset(preset)
	p.Pattern[INTRO].T1.CopyTrack(p.Pattern[INTRO].T2)
	p.Pattern[INTRO].T1.InitTrig(0)
	p.Pattern[INTRO].T1.InitTrig(2)
	p.Pattern[INTRO].T1.InitTrig(4)

	// scale
	p.Pattern[INTRO].T1.Scale.SetMod(em.PTN)
	p.Pattern[INTRO].T1.Scale.SetLen(15)
	p.Pattern[INTRO].T1.Scale.SetScl(4)
	p.Pattern[INTRO].T1.Scale.SetChg(0)

	// preset

	// trig
	p.Pattern[INTRO].T1.Trig[0].SetNote(em.A4, 4, 125)
	p.Pattern[INTRO].T1.Trig[0].SetLock(lock)
	p.Pattern[INTRO].T1.Trig[2].SetNote(em.E4, 4, 125)
	p.Pattern[INTRO].T1.Trig[2].SetLock(lock)
	p.Pattern[INTRO].T1.Trig[4].SetNote(em.C4, 4, 125)
	p.Pattern[INTRO].T1.Trig[4].SetLock(lock)

	// note
	p.Pattern[INTRO].T1.Trig[0].Note.SetKey(em.A5)
	p.Pattern[INTRO].T1.Trig[0].Note.SetLength(4)
	p.Pattern[INTRO].T1.Trig[0].Note.SetVelocity(126)

	// lock
	p.Pattern[INTRO].T1.Trig[0].Lock.SetPreset(preset)
	p.Pattern[INTRO].T1.Trig[0].Lock.SetMachine(em.KICK)

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

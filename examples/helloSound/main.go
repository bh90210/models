package main

import (
	"log"
	"time"

	em "github.com/bh90210/elektronmodels"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

func main() {
	// preset
	preset := make(em.Preset)
	preset[em.COLOR] = 100

	// lock
	lock := new(em.Lock)
	lock.Preset = preset

	// locks := make([]*em.Lock, 0)
	// locks = append(locks, ll)

	// start a new project
	p, err := em.NewProject(em.CYCLES)
	if err != nil {
		log.Fatal(err)
	}
	// pattern
	p.InitPattern(INTRO)
	p0 := p.Pattern[INTRO]

	// track
	p0.T1.SetScale(em.PTN, 16, 1.0, 0)
	p0.T1.Preset = preset
	p0.T1.InitTrig(0)
	p0.T1.InitTrig(2)
	p0.T1.InitTrig(4)
	// copy track
	p0.T2.CopyTrack(p0.T1)

	// scale
	p0.T1.Scale.SetMod(em.PTN)
	p0.T1.Scale.SetLen(15)
	p0.T1.Scale.SetScl(1.0)
	// inf = 127
	p0.T1.Scale.SetChg(0)

	// preset

	// trig
	p0.T1.Trig[0].SetNote(em.A4, 0.4, 127)
	p0.T1.Trig[0].SetLock(lock)

	// copy trig
	p0.T1.InitTrig(6)
	p0.T1.Trig[6].CopyTrig(p0.T1.Trig[4])

	// note
	p0.T1.Trig[0].Note.SetKey(em.A5)
	// inf = 0
	p0.T1.Trig[0].Note.SetLength(0.4)
	p0.T1.Trig[0].Note.SetVelocity(126)
	// copy note
	p0.T1.Trig[2].Note.CopyNote(p0.T1.Trig[0].Note)
	p0.T1.Trig[4].Note.CopyNote(p0.T1.Trig[0].Note)

	// lock
	p0.T1.Trig[0].Lock.Preset = preset
	p0.T1.Trig[0].Lock.SetMachine(em.KICK)

	// copy pattern
	p.InitPattern(VERSE)
	p.Pattern[VERSE].CopyPattern(p.Pattern[INTRO])

	p.Play()

	// can be used without a number too - if used without a number and there is no next currently playing pattern keeps on looping
	// if used and not found, an empty default pattern should be returned - silence
	p.Next()
	p.Next(2)
	p.Next(CHORUS)
	// Second number indicates jump to specific pattern number rather the next in line.
	p.Next(OUTRO, 5)

	p.Pause()
	time.Sleep(2 * time.Second)
	p.Play()

	p.Stop()
	p.Play()

	p.Close()
}

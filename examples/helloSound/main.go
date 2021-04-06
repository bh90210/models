package main

import (
	em "github.com/bh90210/elektronmodels"
)

func main() {
	// lock := e.NewLock()
	preset := em.NewPreset()
	scale := em.NewScale(true, 16, 1, 16)

	// start a new project
	p := em.NewProject(em.CYCLES)
	// create first pattern
	p1 := em.NewPattern(scale)
	// create a new tack
	kick := em.NewTrack()
	// set preset for track
	kick.SetPreset(preset)
	// create a new trig
	trig1 := em.NewTrig()
	// optionally set a lock on it
	// trig1.SetLock(lock)
	// add trig to track
	kick.AddTrigs(trig1)
	// assign track for pattern
	p1.T1(kick)
	// add pattern to project
	p.AddPattern(p1)
	// play the project
	// p.Play()
	// time.Sleep(2 * time.Second)
	// can be used without a number too - if used without a number and there is no next currently playing pattern keeps on looping
	// if used and not found, an empty default pattern should be returned - silence
	// p.Next(2)
	// time.Sleep(2 * time.Second)
	p.Stop()

	p.Close()
}

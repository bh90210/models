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

	p.Play()
	// test
	// p.CopyPattern(CHORUS, VERSE)
	// p.Pattern(INTRO)

	// p.Pattern(INTRO).Track(em.T1)

	// p.Pattern(INTRO).Track(em.T1).SetScale(em.PTN, 16, 1.0, 8)

	// preset := make(map[em.Parameter]int8)
	// p.Pattern(INTRO).Track(em.T1).SetPreset(preset)
	// p.Pattern(INTRO).Track(em.T1).Trig(0).Lock(preset)

	// todo: delete

	// // pattern
	// p.InitPattern(INTRO)
	// p0 := p.Pattern[INTRO]
	// p0.SetTempo(90)

	// // track
	// p0.T1.SetScale(em.PTN, 8, 1.0, 0)
	// // p0.T1.Preset = preset
	// var t prelock
	// p0.T1.SetPreset(t.NewPreset())
	// p0.T1.InitTrig(0)
	// p0.T1.InitTrig(2)
	// p0.T1.InitTrig(4)
	// // copy track
	// p0.T2.CopyTrack(p0.T1)

	// // scale
	// p0.T1.Scale.SetMod(em.PTN)
	// p0.T1.Scale.SetLen(8)
	// p0.T1.Scale.SetScl(1.0)
	// // inf = 0
	// p0.T1.Scale.SetChg(0)

	// // preset
	// p0.T1.Preset = preset
	// p0.T1.Preset.SetParameter(em.COLOR, 120)
	// p0.T1.Preset.DelParameter(em.COLOR)

	// // trig
	// p0.T1.Trig[0].SetNote(em.A4, 0.4, 127)
	// p0.T1.Trig[0].Lock = lock

	// // copy trig
	// p0.T1.InitTrig(6)
	// p0.T1.Trig[6].CopyTrig(p0.T1.Trig[4])

	// // note
	// p0.T1.Trig[0].Note.SetKey(em.A5)
	// // inf = 0
	// p0.T1.Trig[0].Note.SetLength(0.4)
	// p0.T1.Trig[0].Note.SetVelocity(126)
	// // copy note
	// p0.T1.Trig[2].Note.CopyNote(p0.T1.Trig[0].Note)
	// p0.T1.Trig[4].Note.CopyNote(p0.T1.Trig[0].Note)

	// // lock
	// p0.T1.Trig[0].Lock.Preset = preset
	// p0.T1.Trig[0].Lock.Preset[em.COLOR] = 12
	// p0.T1.Trig[0].Lock.SetMachine(em.KICK)

	// // copy pattern
	// p.InitPattern(VERSE)
	// p.Pattern[VERSE].CopyPattern(p.Pattern[INTRO])
	// p.Pattern[VERSE].SetTempo(150)

	// // sequencer
	// s, err := p.Sequencer()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = s.Play()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// s.Tempo(140.9)

	// s.Volume(127)

	// p.Next()
	// p.Next(2)
	// p.Next(CHORUS)
	// p.Next(OUTRO, 5)

	// p.Pause()
	// time.Sleep(2 * time.Second)
	// p.Play()

	// p.Stop()
	// p.Play()

	p.Close()
}

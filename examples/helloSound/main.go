package main

import (
	"log"

	em "github.com/bh90210/elektronmodels"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

type prelock int

func (p *prelock) NewPreset() map[em.Parameter]int8 {
	return nil
}

func (p *prelock) NewLock() (map[em.Parameter]int8, em.Machine) {
	return nil, 0
}

func main() {
	// preset
	preset := make(em.Preset)
	preset[em.TRACKLEVEL] = 127
	preset[em.MUTE] = 0
	preset[em.PAN] = 65

	preset[em.SWEEP] = 60
	preset[em.CONTOUR] = 60

	preset[em.DELAY] = 10
	preset[em.DELAYTIME] = 25
	preset[em.DELAYFEEDBACK] = 25

	preset[em.REVERB] = 25
	preset[em.REVERBSIZE] = 25
	preset[em.REVERBTONE] = 25

	preset[em.VOLUMEDIST] = 40

	preset[em.CYCLESPITCH] = 40
	preset[em.DECAY] = 40
	preset[em.COLOR] = 40
	preset[em.SHAPE] = 40

	preset[em.PUNCH] = 0
	preset[em.GATE] = 0

	// lock
	lock := new(em.lock)
	lock.Preset = preset
	lock.Machine = em.METAL

	// locks := make([]*em.Lock, 0)
	// locks = append(locks, ll)

	// start a new project
	p := em.NewProject(em.CYCLES)

	// test
	p.Sequencer().CopyPattern(CHORUS, VERSE)
	p.Sequencer().Pattern(INTRO)

	p.Sequencer().Pattern(INTRO).Track(em.T1)

	p.Pattern(INTRO).Track(em.T1).SetScale(em.PTN, 16, 1.0, 8)

	// pattern
	p.InitPattern(INTRO)
	p0 := p.Pattern[INTRO]
	p0.SetTempo(90)

	// track
	p0.T1.SetScale(em.PTN, 8, 1.0, 0)
	// p0.T1.Preset = preset
	var t prelock
	p0.T1.SetPreset(t.NewPreset())
	p0.T1.InitTrig(0)
	p0.T1.InitTrig(2)
	p0.T1.InitTrig(4)
	// copy track
	p0.T2.CopyTrack(p0.T1)

	// scale
	p0.T1.Scale.SetMod(em.PTN)
	p0.T1.Scale.SetLen(8)
	p0.T1.Scale.SetScl(1.0)
	// inf = 0
	p0.T1.Scale.SetChg(0)

	// preset
	p0.T1.Preset = preset
	p0.T1.Preset.SetParameter(em.COLOR, 120)
	p0.T1.Preset.DelParameter(em.COLOR)

	// trig
	p0.T1.Trig[0].SetNote(em.A4, 0.4, 127)
	p0.T1.Trig[0].Lock = lock

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
	p0.T1.Trig[0].Lock.Preset[em.COLOR] = 12
	p0.T1.Trig[0].Lock.SetMachine(em.KICK)

	// copy pattern
	p.InitPattern(VERSE)
	p.Pattern[VERSE].CopyPattern(p.Pattern[INTRO])
	p.Pattern[VERSE].SetTempo(150)

	// sequencer
	s, err := p.Sequencer()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Play()
	if err != nil {
		log.Fatal(err)
	}

	s.Tempo(140.9)

	s.Volume(127)

	// p.Next()
	// p.Next(2)
	// p.Next(CHORUS)
	// p.Next(OUTRO, 5)

	// p.Pause()
	// time.Sleep(2 * time.Second)
	// p.Play()

	// p.Stop()
	// p.Play()

	s.Close()
}

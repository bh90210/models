package main

import (
	"time"

	m "github.com/bh90210/models"
)

func main() {
	p, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	var noteLength int = 250

	defaultPresetT1 := m.PT1()
	p.Free.Preset(m.T1, defaultPresetT1)
	p.Free.Note(m.T1, m.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Note(m.T2, m.C4, 120, float64(noteLength), m.PT2())
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Preset(m.T3, m.PT3())
	p.Free.CC(m.T3, m.DELAY, 0)
	p.Free.Note(m.T3, m.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	preset4 := m.PT4()
	p.Free.Preset(m.T4, preset4)
	preset4 = make(map[m.Parameter]int8)
	preset4[m.DELAY] = 0
	p.Free.Preset(m.T4, preset4)
	p.Free.Note(m.T4, m.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Note(m.T5, m.C4, 120, float64(noteLength), m.PT5())
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	chord := m.PT6()
	chord[m.SHAPE] = int8(m.MajorMinor9no5)
	p.Free.Note(m.T6, m.C4, 120, float64(noteLength), chord)
	time.Sleep(time.Duration(noteLength) * time.Millisecond)
}

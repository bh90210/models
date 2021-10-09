package main

import (
	"time"

	em "github.com/bh90210/models"
)

func main() {
	// start a new project
	p, err := em.NewProject(em.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	// reproduce helloSound example using Free API
	var noteLength int = 500

	defaultPresetT1 := em.PT1()
	p.Free.Preset(em.T1, defaultPresetT1)
	p.Free.Note(em.T1, em.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Note(em.T2, em.C4, 120, float64(noteLength), em.PT2())
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Preset(em.T3, em.PT3())
	p.Free.CC(em.T3, em.DELAY, 0)
	p.Free.Note(em.T3, em.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	preset4 := em.PT4()
	p.Free.Preset(em.T4, preset4)
	preset4 = make(map[em.Parameter]int8)
	preset4[em.DELAY] = 0
	p.Free.Preset(em.T4, preset4)
	p.Free.Note(em.T4, em.C4, 120, float64(noteLength))
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Note(em.T5, em.C4, 120, float64(noteLength), em.PT5())
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Free.Note(em.T6, em.C4, 120, float64(noteLength), em.PT6())
	time.Sleep(time.Duration(noteLength) * time.Millisecond)

	p.Close()
}

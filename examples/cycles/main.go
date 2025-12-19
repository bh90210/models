package main

import (
	"fmt"
	"time"

	m "github.com/bh90210/models/cycles"
)

func main() {
	p, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	in := p.Incoming()
	go func() {
		for {
			val := <-in
			fmt.Println("MIDI IN:", val)
		}
	}()

	var noteLength int = 50

	// defaultPresetT1 := m.PT1()
	// p.Preset(m.T1, defaultPresetT1)
	p.CC(0, 16, 100)
	p.Note(0, 36, 120, float64(noteLength))
	time.Sleep(time.Duration(100 * time.Millisecond))
	p.CC(0, 16, 10)
	p.Note(0, 36, 120, float64(noteLength))
	// time.Sleep(time.Duration(1 * time.Second))
	// p.Note(1, m.C4, 120, float64(noteLength))
	// time.Sleep(time.Duration(1 * time.Second))
	// p.Note(1, m.C4, 120, float64(noteLength))
	// time.Sleep(time.Duration(1 * time.Second))

	// p.Note(m.T2, m.C4, 120, float64(noteLength), m.PT2())
	// time.Sleep(time.Duration(noteLength) * time.Millisecond)

	// p.Preset(m.T3, m.PT3())
	// p.CC(m.T3, m.DELAY, 0)
	// p.Note(m.T3, m.C4, 120, float64(noteLength))
	// time.Sleep(time.Duration(noteLength) * time.Millisecond)

	// preset4 := m.PT4()
	// p.Preset(m.T4, preset4)
	// preset4 = make(map[m.Parameter]int8)
	// preset4[m.DELAY] = 0
	// p.Preset(m.T4, preset4)
	// p.Note(m.T4, m.C4, 120, float64(noteLength))
	// time.Sleep(time.Duration(noteLength) * time.Millisecond)

	// p.Note(m.T5, m.C4, 120, float64(noteLength), m.PT5())
	// time.Sleep(time.Duration(noteLength) * time.Millisecond)

	// chord := m.PT6()
	// chord[m.SHAPE] = int8(m.MajorMinor9no5)
	// p.Note(m.T6, m.C4, 120, float64(noteLength), chord)
	// time.Sleep(time.Duration(noteLength) * time.Millisecond)
}

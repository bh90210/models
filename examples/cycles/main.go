package main

import (
	"fmt"

	"github.com/bh90210/models/cycles"
	m "github.com/bh90210/models/cycles"
	"github.com/bh90210/models/midicom"
	"github.com/bh90210/models/pattern"
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

	allVoices := make(map[int][]pattern.Pattern)

	note := pattern.Note{Note: midicom.Note(cycles.C4), Duration: 200, Velocity: 100}

	// BD
	pat := pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{note},
		Channel: midicom.Channel(cycles.T1),
		Meta: pattern.Meta{
			Synth: string(cycles.CYCLES),
			Part:  "BD",
		},
	}

	pat.Notes[0].CC = cycles.PT1()

	allVoices[0] = append(allVoices[0], pat)

	// SN
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{note},
		Channel: midicom.Channel(cycles.T2),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "SN",
		},
	}

	pat.Notes[0].CC = cycles.PT2()

	allVoices[0] = append(allVoices[0], pat)

	// Metal
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{note},
		Channel: midicom.Channel(cycles.T3),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "Metal",
		},
	}

	pat.Notes[0].CC = cycles.PT3()

	allVoices[0] = append(allVoices[0], pat)

	// Perc
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{note},
		Channel: midicom.Channel(cycles.T4),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "Perc",
		},
	}

	pat.Notes[0].CC = cycles.PT4()

	allVoices[0] = append(allVoices[0], pat)

	// Tone
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{note},
		Channel: midicom.Channel(cycles.T5),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "Tone",
		},
	}

	pat.Notes[0].CC = cycles.PT5()

	allVoices[0] = append(allVoices[0], pat)

	// Chord
	for o := range 37 {
		chordNote := note
		chordNote.Note = midicom.Note(int(note.Note-5) + o)
		pat = pattern.Pattern{
			Midicom: p,
			Notes:   []pattern.Note{chordNote},
			Channel: midicom.Channel(cycles.T6),
			Meta: pattern.Meta{
				Synth: pat.Synth,
				Part:  "Intro Chord",
			},
		}

		pat.Notes[0].Duration = pat.Notes[0].Duration + float64(o*10)
		pat.Notes[0].CC = cycles.PT6()
		pat.Notes[0].CC[cycles.SWEEP] = int8(o * 2)
		pat.Notes[0].CC[cycles.CONTOUR] = int8(o * 2)
		pat.Notes[0].CC[cycles.REVERB] = int8(o * 2)
		pat.Notes[0].CC[cycles.COLOR] = int8(o * 2)
		pat.Notes[0].CC[cycles.DECAY] = int8(o * 2)
		pat.Notes[0].CC[cycles.DELAY] = int8(o)
		pat.Notes[0].CC[cycles.REVERBSIZE] = int8(o * 2)
		pat.Notes[0].CC[cycles.REVERBTONE] = int8(o * 2)
		pat.Notes[0].CC[cycles.SHAPE] = int8(cycles.Unisonx2) + int8(o)

		allVoices[0] = append(allVoices[0], pat)
	}

	// 	Print
	pr := pattern.NewPrint(allVoices)
	pr.Print(pattern.PatternPosition)

	err = pattern.Play(allVoices)
	if err != nil {
		panic(err)
	}
}

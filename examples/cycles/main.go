package main

import (
	"flag"
	"fmt"

	"github.com/bh90210/models/cycles"
	m "github.com/bh90210/models/cycles"
	"github.com/bh90210/models/midicom"
	"github.com/bh90210/models/pattern"
)

func main() {
	startFrom := flag.Int("position", 0, "start the audio playback form the designated pattern position.")
	flag.Parse()

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

	compo(allVoices, p)

	// Trim patterns to start from designated position.
	trimmedVoices := make(map[int][]pattern.Pattern)
	for k, v := range allVoices {
		trimmedVoices[k] = v[*startFrom:]
	}

	// Print the patterns.
	pr := pattern.NewPrint(trimmedVoices)
	pr.Print(pattern.PatternPosition)

	err = pattern.Play(trimmedVoices)
	if err != nil {
		panic(err)
	}
}

func baseNote() pattern.Note {
	return pattern.Note{
		Note:     midicom.Note(cycles.C4),
		Duration: 200,
		Velocity: 100,
	}
}

func compo(allVoices map[int][]pattern.Pattern, p midicom.MidiCom) {
	// BD
	pat := pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{baseNote()},
		Channel: midicom.Channel(cycles.T1),
		Meta: pattern.Meta{
			Synth: string(cycles.CYCLES),
			Part:  "BD",
		},
	}

	pat.Notes[0].Duration = 50
	pat.Notes[0].CC = cycles.PT1()
	pat.Notes[0].CC[cycles.VOLUMEDIST] = int8(80)

	allVoices[0] = append(allVoices[0], pat)

	// SN
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{baseNote()},
		Channel: midicom.Channel(cycles.T2),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "SN",
		},
	}

	pat.Notes[0].CC = cycles.PT2()
	pat.Notes[0].CC[cycles.DELAY] = int8(120)

	allVoices[0] = append(allVoices[0], pat)

	// Metal
	pat = pattern.Pattern{
		Midicom: p,
		Notes:   []pattern.Note{baseNote()},
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
		Notes:   []pattern.Note{baseNote()},
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
		Notes:   []pattern.Note{baseNote()},
		Channel: midicom.Channel(cycles.T5),
		Meta: pattern.Meta{
			Synth: pat.Synth,
			Part:  "Tone",
		},
	}

	pat.Notes[0].Note = midicom.Note(cycles.C2)
	pat.Notes[0].Duration = 3000

	pat.Notes[0].CC = cycles.PT5()
	pat.Notes[0].CC[cycles.REVERB] = int8(100)
	pat.Notes[0].CC[cycles.COLOR] = int8(50)
	pat.Notes[0].CC[cycles.SHAPE] = int8(40)
	pat.Notes[0].CC[cycles.DECAY] = int8(70)

	allVoices[0] = append(allVoices[0], pat)

	// Chord
	for o := range 37 {
		chordNote := baseNote()
		chordNote.Note = midicom.Note(int(chordNote.Note-10) + o)
		pat = pattern.Pattern{
			Midicom: p,
			Notes:   []pattern.Note{chordNote},
			Channel: midicom.Channel(cycles.T6),
			Meta: pattern.Meta{
				Synth: pat.Synth,
				Part:  "Intro Chord",
			},
		}

		if o == 0 {
			pat.Notes[0].Velocity = 0
		}

		pat.Notes[0].Duration = pat.Notes[0].Duration + float64(o*10)
		pat.Notes[0].CC = cycles.PT6()
		pat.Notes[0].CC[cycles.SWEEP] = int8(o * 2)
		pat.Notes[0].CC[cycles.CONTOUR] = int8(o * 2)
		pat.Notes[0].CC[cycles.REVERB] = int8(o * 2)
		pat.Notes[0].CC[cycles.COLOR] = int8(o * 2)
		pat.Notes[0].CC[cycles.DECAY] = int8(o*2) + 10
		pat.Notes[0].CC[cycles.DELAY] = int8(o)
		pat.Notes[0].CC[cycles.REVERBSIZE] = int8(o * 2)
		pat.Notes[0].CC[cycles.REVERBTONE] = int8(o * 2)
		pat.Notes[0].CC[cycles.SHAPE] = int8(cycles.Unisonx2) + int8(o)

		allVoices[0] = append(allVoices[0], pat)
	}
}

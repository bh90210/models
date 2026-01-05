// Package pattern provides structures and functions to create and manipulate musical patterns.
// It is based on the midicom.MidiCom interface defined in the midicom package.
package pattern

import (
	"fmt"
	"sync"

	"github.com/bh90210/models/midicom"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Degree represents a musical interval in terms of semitones.
type Degree int

const (
	// Minor2nd Degree represents a minor second interval (1 semitone).
	Minor2nd Degree = iota + 1
	// Major2nd Degree represents a major second interval (2 semitones).
	Major2nd
	// Minor3rd Degree represents a minor third interval (3 semitones).
	Minor3rd
	// Major3rd Degree represents a major third interval (4 semitones).
	Major3rd
	// Perfect4th Degree represents a perfect fourth interval (5 semitones).
	Perfect4th
	// Tritone Degree represents a tritone interval (6 semitones).
	Tritone
	// Perfect5th Degree represents a perfect fifth interval (7 semitones).
	Perfect5th
	// Minor6th Degree represents a minor sixth interval (8 semitones).
	Minor6th
	// Major6th Degree represents a major sixth interval (9 semitones).
	Major6th
	// Minor7th Degree represents a minor seventh interval (10 semitones).
	Minor7th
	// Major7th Degree represents a major seventh interval (11 semitones).
	Major7th
	// Octave Degree represents an octave interval (12 semitones).
	Octave
)

type Note struct {
	Note     midicom.Note
	Duration float64 // In milliseconds.
	Velocity int8
	CC       map[midicom.Parameter]int8 // CC changes to apply when playing this note.
	PC       *int8                      // Program Change to apply when playing this note.
}

type Meta struct {
	Synth string
	Part  string
}

// Note to self: the way we achive polypoly is by stacking multple patterns.
// For example, on Nymphes we would play in parallel 3 patterns, each assigned to
// channel 0 (since Nymphes only has one channel). Each pattern would have its own
// set of notes, durations, and velocities, togeteher forming a chord.
// In the case of Model Cycles the is no polymphony so we need to stack the patterns
// against the different channels of the synth (6 channels.)
type Pattern struct {
	Midicom midicom.MidiCom
	Channel midicom.Channel
	Notes   []Note
	Meta
}

// Shift shifts all notes in the pattern by the given degree
// and returns a new Pattern with the shifted notes.
// It does not modify the original pattern.
func (p *Pattern) Shift(shift Degree) Pattern {
	var shiftedNotes []Note
	for _, note := range p.Notes {
		shiftedNotes = append(shiftedNotes, Note{
			Note:     midicom.Note(int8(note.Note) + int8(shift)),
			Duration: note.Duration,
			Velocity: note.Velocity,
		})
	}

	return Pattern{
		Midicom: p.Midicom,
		Notes:   shiftedNotes,
		Channel: p.Channel,
		Meta: Meta{
			Synth: p.Meta.Synth,
			Part:  fmt.Sprintf("%s_shifted_%d", p.Meta.Part, shift),
		},
	}
}

func (p *Pattern) Print() {
	t := table.NewWriter()

	t.SetStyle(table.StyleRounded)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Synth", WidthMax: 30},
		{Name: "Part", WidthMax: 30},
		{Name: "Notes", WidthMax: 30},
		{Name: "Durations", WidthMax: 30},
		{Name: "Velocities", WidthMax: 30},
	})

	t.AppendHeader(table.Row{"Synth", "Part", "Channel", "Notes", "Durations", "Velocities"})

	var notes []midicom.Note
	var durations []float64
	var velocities []int8

	for _, n := range p.Notes {
		notes = append(notes, n.Note)
		durations = append(durations, n.Duration)
		velocities = append(velocities, n.Velocity)
	}

	t.AppendRow(table.Row{
		p.Meta.Synth,
		p.Meta.Part,
		p.Channel,
		notes,
		durations,
		velocities,
	})

	fmt.Println(t.Render())
}

type Print struct {
	allVoices map[int][]Pattern
	voices    int
	// t is the table writer.
	t table.Writer
}

func NewPrint(allVoices map[int][]Pattern) *Print {
	p := &Print{
		allVoices: allVoices,
	}

	// Find how many voices we have.
	p.voices = len(p.allVoices)

	p.t = table.NewWriter()

	p.t.SetStyle(table.StyleRounded)
	p.t.SetColumnConfigs([]table.ColumnConfig{
		{Name: "Synth", WidthMax: 30},
		{Name: "Part", WidthMax: 30},
		{Name: "Notes", WidthMax: 30},
		{Name: "Durations", WidthMax: 30},
		{Name: "Velocities", WidthMax: 30},
	})

	p.t.AppendHeader(table.Row{"Synth", "Part", "Voice", "Position", "Channel", "Notes", "Durations", "Velocities"})

	return p
}

type Sorter int

const (
	Voice Sorter = iota
	PatternPosition
)

func (p *Print) Print(sorter Sorter) {
	defer p.t.ResetRows()
	defer p.t.ResetFooters()

	switch sorter {
	case Voice:
		p.voiceSorter()
	case PatternPosition:
		p.patternPositionSorter()
	}

	fmt.Println(p.t.Render())
}

func (p *Print) voiceSorter() {
	var rows []table.Row

	for voice, patterns := range p.allVoices {
		for patternPosition, pat := range patterns {
			var notes []midicom.Note
			var durations []float64
			var velocities []int8

			for _, n := range pat.Notes {
				notes = append(notes, n.Note)
				durations = append(durations, n.Duration)
				velocities = append(velocities, n.Velocity)
			}

			switch voice % 5 {
			case 0:
				rows = append(rows, table.Row{
					text.FgHiRed.Sprint(pat.Meta.Synth),
					text.FgHiRed.Sprint(pat.Meta.Part),
					text.FgHiRed.Sprint(voice),
					text.FgHiRed.Sprint(patternPosition),
					text.FgHiRed.Sprint(pat.Channel),
					text.FgHiRed.Sprint(notes),
					text.FgHiRed.Sprint(durations),
					text.FgHiRed.Sprint(velocities),
				})
			case 1:
				rows = append(rows, table.Row{
					text.FgHiGreen.Sprint(pat.Meta.Synth),
					text.FgHiGreen.Sprint(pat.Meta.Part),
					text.FgHiGreen.Sprint(voice),
					text.FgHiGreen.Sprint(patternPosition),
					text.FgHiGreen.Sprint(pat.Channel),
					text.FgHiGreen.Sprint(notes),
					text.FgHiGreen.Sprint(durations),
					text.FgHiGreen.Sprint(velocities),
				})
			case 2:
				rows = append(rows, table.Row{
					text.FgHiBlue.Sprint(pat.Meta.Synth),
					text.FgHiBlue.Sprint(pat.Meta.Part),
					text.FgHiBlue.Sprint(voice),
					text.FgHiBlue.Sprint(patternPosition),
					text.FgHiBlue.Sprint(pat.Channel),
					text.FgHiBlue.Sprint(notes),
					text.FgHiBlue.Sprint(durations),
					text.FgHiBlue.Sprint(velocities),
				})
			case 3:
				rows = append(rows, table.Row{
					text.FgHiCyan.Sprint(pat.Meta.Synth),
					text.FgHiCyan.Sprint(pat.Meta.Part),
					text.FgHiCyan.Sprint(voice),
					text.FgHiCyan.Sprint(patternPosition),
					text.FgHiCyan.Sprint(pat.Channel),
					text.FgHiCyan.Sprint(notes),
					text.FgHiCyan.Sprint(durations),
					text.FgHiCyan.Sprint(velocities),
				})
			case 4:
				rows = append(rows, table.Row{
					text.FgHiMagenta.Sprint(pat.Meta.Synth),
					text.FgHiMagenta.Sprint(pat.Meta.Part),
					text.FgHiMagenta.Sprint(voice),
					text.FgHiMagenta.Sprint(patternPosition),
					text.FgHiMagenta.Sprint(pat.Channel),
					text.FgHiMagenta.Sprint(notes),
					text.FgHiMagenta.Sprint(durations),
					text.FgHiMagenta.Sprint(velocities),
				})
			}
		}
	}

	for _, v := range rows {
		p.t.AppendRow(v)
	}

	p.t.AppendSeparator()

	p.t.AppendFooter(table.Row{"voice", "sorter"})
}

func (p *Print) patternPositionSorter() {
	patternsSort := make(map[int][]Pattern)

	for _, patterns := range p.allVoices {
		for patternPosition, pat := range patterns {
			patternsSort[patternPosition] = append(patternsSort[patternPosition], pat)
		}
	}

	var rows []table.Row
	patternsLength := len(patternsSort)
	for patternPosition := range patternsLength {
		patterns := patternsSort[patternPosition]
		for voice, pat := range patterns {
			var notes []midicom.Note
			var durations []float64
			var velocities []int8

			for _, n := range pat.Notes {
				notes = append(notes, n.Note)
				durations = append(durations, n.Duration)
				velocities = append(velocities, n.Velocity)
			}

			switch patternPosition % 5 {
			case 0:
				rows = append(rows, table.Row{
					text.FgHiRed.Sprint(pat.Meta.Synth),
					text.FgHiRed.Sprint(pat.Meta.Part),
					text.FgHiRed.Sprint(voice),
					text.FgHiRed.Sprint(patternPosition),
					text.FgHiRed.Sprint(pat.Channel),
					text.FgHiRed.Sprint(notes),
					text.FgHiRed.Sprint(durations),
					text.FgHiRed.Sprint(velocities),
				})
			case 1:
				rows = append(rows, table.Row{
					text.FgHiGreen.Sprint(pat.Meta.Synth),
					text.FgHiGreen.Sprint(pat.Meta.Part),
					text.FgHiGreen.Sprint(voice),
					text.FgHiGreen.Sprint(patternPosition),
					text.FgHiGreen.Sprint(pat.Channel),
					text.FgHiGreen.Sprint(notes),
					text.FgHiGreen.Sprint(durations),
					text.FgHiGreen.Sprint(velocities),
				})
			case 2:
				rows = append(rows, table.Row{
					text.FgHiBlue.Sprint(pat.Meta.Synth),
					text.FgHiBlue.Sprint(pat.Meta.Part),
					text.FgHiBlue.Sprint(voice),
					text.FgHiBlue.Sprint(patternPosition),
					text.FgHiBlue.Sprint(pat.Channel),
					text.FgHiBlue.Sprint(notes),
					text.FgHiBlue.Sprint(durations),
					text.FgHiBlue.Sprint(velocities),
				})
			case 3:
				rows = append(rows, table.Row{
					text.FgHiCyan.Sprint(pat.Meta.Synth),
					text.FgHiCyan.Sprint(pat.Meta.Part),
					text.FgHiCyan.Sprint(voice),
					text.FgHiCyan.Sprint(patternPosition),
					text.FgHiCyan.Sprint(pat.Channel),
					text.FgHiCyan.Sprint(notes),
					text.FgHiCyan.Sprint(durations),
					text.FgHiCyan.Sprint(velocities),
				})
			case 4:
				rows = append(rows, table.Row{
					text.FgHiMagenta.Sprint(pat.Meta.Synth),
					text.FgHiMagenta.Sprint(pat.Meta.Part),
					text.FgHiMagenta.Sprint(voice),
					text.FgHiMagenta.Sprint(patternPosition),
					text.FgHiMagenta.Sprint(pat.Channel),
					text.FgHiMagenta.Sprint(notes),
					text.FgHiMagenta.Sprint(durations),
					text.FgHiMagenta.Sprint(velocities),
				})
			}
		}
	}

	for _, v := range rows {
		p.t.AppendRow(v)
	}

	p.t.AppendSeparator()

	p.t.AppendFooter(table.Row{"pattern", "sorter"})
}

// Play
//
// allVoices argument holds the polyphonic patterns to be played.
// The key of the map is the voice. Note that this is
// independent of the channel, as multiple voices can
// share the same channel, for example Nymphes.
func Play(allVoices map[int][]Pattern) error {
	// Find how many voices we have.
	length := len(allVoices)

	// Play all voices in parallel.
	var wg sync.WaitGroup
	wg.Add(length)
	for voice := range length {
		// Read each voice's patterns concurrently.
		go func(patterns []Pattern) {
			defer wg.Done()

			// Read each voice's patterns serially.
			for _, p := range patterns {
				if p.Midicom == nil {
					fmt.Println("No MidiCom assigned to pattern", p.Meta)
					return
				}

				// TODO: prolly remove as showing the meta while playing in concurent
				// fashion will rather be chaotic, but try it first.
				// p.Print(voice, i)

				for _, n := range p.Notes {
					// First apply any PC or CC changes.
					if n.PC != nil {
						err := p.Midicom.PC(p.Channel, *n.PC)
						if err != nil {
							fmt.Println("Error sending PC:", err)
							return
						}
					}

					if n.CC != nil {
						for cc, val := range n.CC {
							err := p.Midicom.CC(p.Channel, cc, val)
							if err != nil {
								fmt.Println("Error sending CC:", err)
								return
							}
						}
					}

					// Now play the note.
					err := p.Midicom.Note(p.Channel, n.Note, n.Velocity, n.Duration)
					if err != nil {
						fmt.Println("Error playing note:", err)
						return
					}
				}
			}
		}(allVoices[voice])
	}

	// Wait for all voices to finish.
	wg.Wait()

	return nil
}

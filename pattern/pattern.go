package pattern

import (
	"fmt"
	"sync"

	"github.com/bh90210/models"
)

type Degree int

const (
	Minor2nd Degree = iota + 1
	Major2nd
	Minor3rd
	Major3rd
	Perfect4th
	Tritone
	Perfect5th
	Minor6th
	Major6th
	Minor7th
	Major7th
	Octave
)

// Note to self: the way we achive polypoly is by stacking multple patterns.
// For example, on Nymphes we would play in parallel 3 patterns, each assigned to
// channel 0 (since Nymphes only has one channel). Each pattern would have its own
// set of notes, durations, and velocities, togeteher forming a chord.
// In the case of Model Cycles the is no polymphony so we need to stack the patterns
// against the different channels of the synth (6 channels.)
type Pattern struct {
	Midicom models.MidiCom
	Notes   []Note
	// Notes      []models.Notes
	// Durations  []float64 // In milliseconds.
	// Velocities []uint8
	Channel models.Channel

	Meta
}

type Note struct {
	Note     models.Note
	Duration float64
	Velocity int8
}

func (p *Pattern) Shift(shift Degree) Pattern {
	var shiftedNotes []Note
	for _, note := range p.Notes {
		shiftedNotes = append(shiftedNotes, Note{
			Note:     models.Note(int8(note.Note) + int8(shift)),
			Duration: note.Duration,
			Velocity: note.Velocity,
		})
	}

	// Add shift info to meta.
	p.Meta.Part = fmt.Sprintf("%s_shifted_%d", p.Meta.Part, shift)

	return Pattern{
		Midicom: p.Midicom,
		Notes:   shiftedNotes,
		Channel: p.Channel,
		Meta:    p.Meta,
	}
}

type Poly struct {
	// patterns hols the polyphonic patterns to be played.
	// The key of the map is the voice. Note that this is
	// independent of the channel, as multiple voices can
	// share the same channel, for example Nymphes.
	patterns map[int][]Pattern
}

func (p *Poly) AddPattern(voice int, pattern Pattern) {
	p.patterns[voice] = append(p.patterns[voice], pattern)
}

func (p *Poly) GetPatterns() map[int][]Pattern {
	return p.patterns
}

type Meta struct {
	Synth string
	Part  string
}

func Play(patterns Poly) error {
	allVoices := patterns.GetPatterns()
	length := len(allVoices)

	wg := sync.WaitGroup{}
	for voice := 0; voice < length; voice++ {
		wg.Add(1)
		go func(voice []Pattern) {
			for _, pat := range voice {
				for _, n := range pat.Notes {
					err := pat.Midicom.Note(pat.Channel, n.Note, n.Velocity, n.Duration)
					if err != nil {
						fmt.Println("Error playing note:", err)
						return
					}
				}
			}

			wg.Done()

		}(allVoices[voice])
	}

	wg.Wait()

	return nil
}

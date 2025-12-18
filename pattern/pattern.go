package pattern

import (
	"fmt"
	"sync"

	"github.com/bh90210/models"
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

// Note to self: the way we achive polypoly is by stacking multple patterns.
// For example, on Nymphes we would play in parallel 3 patterns, each assigned to
// channel 0 (since Nymphes only has one channel). Each pattern would have its own
// set of notes, durations, and velocities, togeteher forming a chord.
// In the case of Model Cycles the is no polymphony so we need to stack the patterns
// against the different channels of the synth (6 channels.)
type Pattern struct {
	Midicom models.MidiCom
	Notes   []Note
	Channel models.Channel

	Meta
}

type Note struct {
	Note     models.Note
	Duration float64 // In milliseconds.
	Velocity int8
	CC       map[int8]int8 // CC changes to apply when playing this note.
	PC       *int8         // Program Change to apply when playing this note.
}

type Meta struct {
	Synth string
	Part  string
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

func Play(patterns Poly) error {
	// Get all voices and their patterns.
	allVoices := patterns.GetPatterns()
	// Find how many voices we have.
	length := len(allVoices)

	wg := sync.WaitGroup{}
	// Play all voices in parallel.
	for voice := range length {
		wg.Add(1)

		// Read each voice's patterns concurrently.
		go func(voice []Pattern) {
			for _, pat := range voice {
				if pat.Midicom == nil {
					fmt.Println("No MidiCom assigned to pattern", pat.Meta)
					wg.Done()
					return
				}

				for _, n := range pat.Notes {
					// First apply any PC or CC changes.
					if n.PC != nil {
						err := pat.Midicom.PC(pat.Channel, *n.PC)
						if err != nil {
							fmt.Println("Error sending PC:", err)
							return
						}
					}

					if n.CC != nil {
						for cc, val := range n.CC {
							err := pat.Midicom.CC(pat.Channel, models.Parameter(cc), val)
							if err != nil {
								fmt.Println("Error sending CC:", err)
								return
							}
						}
					}

					// Now play the note.
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

	// Wait for all voices to finish.
	wg.Wait()

	return nil
}

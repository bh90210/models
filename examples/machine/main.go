package main

import (
	"math/rand"
	"time"

	"github.com/bh90210/models/machine"
	"github.com/bh90210/models/machine/bd"
	"github.com/bh90210/models/machine/synthesis"
)

func main() {
	// Connect to the synth.
	m, err := machine.New()
	if err != nil {
		panic(err)
	}

	defer m.Close()

	// Start the song.
	s := &synthesis.Synthesis{
		BDSampleRateReduction: 97,
	}

	// Initiate BD track.
	t1 := new(machine.Track)

	// Settings for track 1.
	synthesis.Position1(m)

	dursEq7 := synthesis.EqualDuration(7, time.Millisecond*time.Duration(rand.Intn(500)))
	for i, d := range dursEq7 {
		t1.Add(machine.Event{
			Duration: d,
			Action: func(duration time.Duration) {
				s.SampleReduction(m)

				m.CC(bd.CHANNEL, bd.DELAY, 0)
				m.CC(bd.CHANNEL, bd.REVERB, 0)

				if i > 2 && i < 5 {
					m.CC(bd.CHANNEL, bd.REVERB, int8(126))
					m.CC(bd.CHANNEL, bd.DELAY, int8(126))
				}

				synthesis.Position2(m)

				m.Note(bd.CHANNEL, bd.NOTE, 120, duration)
			},
		})
	}

	for i, d := range dursEq7 {
		t1.Add(machine.Event{
			Duration: d / time.Duration(i+1),
			Action: func(duration time.Duration) {
				s.SampleReduction(m)

				m.CC(bd.CHANNEL, bd.DELAY, 0)
				m.CC(bd.CHANNEL, bd.REVERB, 0)

				if i > 2 && i < 5 {
					m.CC(bd.CHANNEL, bd.REVERB, int8(126))
					m.CC(bd.CHANNEL, bd.DELAY, int8(126))
				}

				synthesis.Position2(m)

				m.Note(bd.CHANNEL, bd.NOTE, 120, duration)
			},
		})
	}

	t1.Run()
}

// 1. Just kick
// 2. All channels + Changing machines
// 4. With nord
// 5. With semi modulars

package main

import (
	"slices"
	"time"

	"github.com/bh90210/models/machine"
	"github.com/bh90210/models/machine/bd"
	"github.com/bh90210/models/machine/sd"
	"github.com/bh90210/models/machine/synthesis"
)

var cycle = []int8{20, 40, 20, 40, 50, 20, 40, 55, 45}

func main() {
	m, err := machine.New()
	if err != nil {
		panic(err)
	}

	defer m.Close()

	// Start the song.
	s := &synthesis.Synthesis{
		Tracks: []*machine.Track{
			bd.TRACK: new(machine.Track),
			sd.TRACK: new(machine.Track),
		},
		Events: []machine.Events{
			bd.TRACK: make(machine.Events, 0),
			sd.TRACK: make(machine.Events, 0),
		},
	}

	s.Position1()

	dursEq256 := synthesis.EqualDuration(185, time.Millisecond*time.Duration(300))

	for i, d := range dursEq256 {
		// Set the rhythm.
		d = synthesis.Rhythm1(i, d)

		switch i {
		case 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190:
			s.Tracks[bd.TRACK].Param3 -= 8
			s.Tracks[bd.TRACK].Param4 -= 4
			s.Tracks[bd.TRACK].LFOAmount -= 1
		}

		switch i {
		case 30, 60, 90, 100, 110, 130, 140, 160, 170, 190:
			s.Tracks[bd.TRACK].LFOSpeed += 4
			s.Tracks[bd.TRACK].LFOShape += 4
		}

		s.Events[bd.CHANNEL].Add(machine.Event{
			Duration: d,
			Track:    *s.Tracks[bd.CHANNEL],
			Action: func(duration time.Duration, t machine.Track) {
				// s.Position1()

				// Assign baseline. We want no additional computation.
				// t.Param1 = uint8((cycle[i%len(cycle)] + int8(i))) % 126
				// t.Param1 = uint8((cycle[i%len(cycle)]))

				// t.LFOShape = (t.LFOShape + uint8(i)) % 126
				// Baseline.
				// t.FilterBaseFrq = (t.FilterBaseFrq + uint8(i)) % 80
				// t.FilterWidth = t.FilterWidth + uint8(i)%80
				// t.FilterQ = (t.FilterQ + uint8(i)) % 80
				// t.Pan = (t.Pan + uint8(i)) % 126
				// t.Level = 120
				// if i > 20 && duration < time.Millisecond*51 {
				// 	t.SampleRateReduction = 120
				// 	t.Level = 0
				// 	t.Reverb = 35
				// 	t.Delay = 30
				// 	t.Param3 = 50
				// 	t.Param4 = 120
				// }

				synthesis.BDCC(m, t)
				// t.Level -= 90
				synthesis.SDCC(m, t)

				// Send the note.
				m.Note(0, bd.NOTE, 126, duration)
				m.Note(0, sd.NOTE, 126, duration)

				time.Sleep(duration)
			},
		})
	}

	var part machine.Events
	for _, v := range s.Events[bd.CHANNEL] {
		part = append(part, v)
	}

	slices.Reverse(part)

	for _, v := range part {
		v.Duration /= 15
		s.Events[bd.CHANNEL].Add(v)
	}

	// s.Events[bd.CHANNEL] = s.Events[bd.CHANNEL][185:]
	s.Run()
}

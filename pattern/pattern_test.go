package pattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	tests := map[string]struct {
		notes          []Note
		shift          Degree
		expectedShifts []Note
		meta           Meta
		expectedMeta   Meta
	}{
		"positive": {
			notes:          []Note{{Note: 60}, {Note: 62}, {Note: 64}},
			shift:          Major2nd,
			expectedShifts: []Note{{Note: 62}, {Note: 64}, {Note: 66}},
			meta: Meta{
				Part: "test_part",
			},
			expectedMeta: Meta{
				Part: "test_part_shifted_2",
			},
		},
		"negative": {
			notes:          []Note{{Note: 60}, {Note: 62}, {Note: 64}},
			shift:          -Major2nd,
			expectedShifts: []Note{{Note: 58}, {Note: 60}, {Note: 62}},
			meta: Meta{
				Part: "test_part",
			},
			expectedMeta: Meta{
				Part: "test_part_shifted_-2",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := Pattern{
				Notes: tc.notes,
				Meta:  tc.meta,
			}
			shiftedPattern := p.Shift(tc.shift)
			assert.Equal(t, tc.expectedShifts, shiftedPattern.Notes)
			assert.Equal(t, tc.expectedMeta, shiftedPattern.Meta)
		})
	}
}

func TestPatternPrint(t *testing.T) {
	tests := map[string]struct {
		patterns []Pattern
	}{
		"simple": {
			patterns: []Pattern{
				{
					Notes: []Note{
						{Note: 60, Duration: 500, Velocity: 100},
						{Note: 62, Duration: 500, Velocity: 100},
					},
					Channel: 0,
					Meta: Meta{
						Synth: "TestSynth",
						Part:  "TestPart",
					},
				},
			},
		},
		"myltiple patterns": {
			patterns: []Pattern{
				{
					Notes: []Note{
						{Note: 65, Duration: 250, Velocity: 90},
						{Note: 67, Duration: 250, Velocity: 90},
						{Note: 69, Duration: 500, Velocity: 100},
						{Note: 67, Duration: 500, Velocity: 100},
						{Note: 65, Duration: 1000, Velocity: 110},
						{Note: 64, Duration: 500, Velocity: 100},
						{Note: 62, Duration: 500, Velocity: 90},
						{Note: 60, Duration: 2000, Velocity: 80},
						{Note: 62, Duration: 500, Velocity: 90},
						{Note: 64, Duration: 500, Velocity: 100},
						{Note: 65, Duration: 1000, Velocity: 110},
						{Note: 67, Duration: 500, Velocity: 100},
						{Note: 69, Duration: 500, Velocity: 90},
						{Note: 71, Duration: 2000, Velocity: 80},
						{Note: 69, Duration: 500, Velocity: 90},
						{Note: 67, Duration: 500, Velocity: 100},
					},
					Channel: 0,
					Meta: Meta{
						Synth: "Nymphes",
						Part:  "Intro",
					},
				},
				{
					Notes: []Note{
						{Note: 50, Duration: 1250, Velocity: 90},
						{Note: 60, Duration: 1250, Velocity: 90},
						{Note: 70, Duration: 1500, Velocity: 100},
						{Note: 80, Duration: 1500, Velocity: 100},
					},
					Channel: 0,
					Meta: Meta{
						Synth: "Nymphes",
						Part:  "Intro 2",
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			for _, p := range tc.patterns {
				p.Print()
			}
		})
	}
}

func TestPrint(t *testing.T) {
	tests := map[string]struct {
		allVoices map[int][]Pattern
	}{
		"multiple voices": {
			allVoices: map[int][]Pattern{
				0: {
					{
						Notes: []Note{
							{Note: 60, Duration: 500, Velocity: 100},
							{Note: 62, Duration: 500, Velocity: 100},
						},
						Channel: 0,
						Meta: Meta{
							Synth: "cycles",
							Part:  "intro BD",
						},
					},
					{
						Notes: []Note{
							{Note: 64, Duration: 500, Velocity: 100},
							{Note: 65, Duration: 500, Velocity: 100},
						},
						Channel: 0,
						Meta: Meta{
							Synth: "cycles",
							Part:  "verse BD",
						},
					},
				},
				1: {
					{
						Notes: []Note{
							{Note: 62, Duration: 250, Velocity: 90},
							{Note: 64, Duration: 250, Velocity: 90},
						},
						Channel: 1,
						Meta: Meta{
							Synth: "cycles",
							Part:  "intro SD",
						},
					},
					{
						Notes: []Note{
							{Note: 65, Duration: 500, Velocity: 100},
							{Note: 64, Duration: 500, Velocity: 100},
						},
						Channel: 1,
						Meta: Meta{
							Synth: "cycles",
							Part:  "verse SD",
						},
					},
				},
				2: {
					{
						Notes: []Note{
							{Note: 65, Duration: 250, Velocity: 90},
							{Note: 67, Duration: 250, Velocity: 90},
						},
						Channel: 0,
						Meta: Meta{
							Synth: "nymphes",
							Part:  "intro",
						},
					},
					{
						Notes: []Note{
							{Note: 69, Duration: 500, Velocity: 100},
							{Note: 67, Duration: 500, Velocity: 100},
						},
						Channel: 0,
						Meta: Meta{
							Synth: "nymphes",
							Part:  "verse",
						},
					},
					{
						Notes: []Note{
							{Note: 65, Duration: 1000, Velocity: 110},
							{Note: 64, Duration: 500, Velocity: 100},
						},
						Channel: 0,
						Meta: Meta{
							Synth: "nymphes",
							Part:  "outro",
						},
					},
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			p := NewPrint(tc.allVoices)
			p.Print(Voice)
			p.Print(PatternPosition)
		})
	}
}

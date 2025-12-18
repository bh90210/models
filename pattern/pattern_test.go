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
			got := shiftedPattern.Notes
			assert.Equal(t, tc.expectedShifts, got)
			assert.Equal(t, tc.expectedMeta, shiftedPattern.Meta)
		})
	}
}

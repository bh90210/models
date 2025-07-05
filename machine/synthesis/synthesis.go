package synthesis

import (
	"math/rand"
	"time"

	"github.com/bh90210/models/machine"
	"github.com/bh90210/models/machine/bd"
)

type Synthesis struct {
	BDSampleRateReduction int
}

func (s *Synthesis) SampleReduction(m *machine.Machine) {
	if s.BDSampleRateReduction == 0 {
		s.BDSampleRateReduction = 127
	}

	m.CC(bd.CHANNEL, bd.SAMPLERATEREDUCTION, int8(s.BDSampleRateReduction))

	s.BDSampleRateReduction -= 5
}

func EqualDuration(x int, dur time.Duration) []time.Duration {
	var durations []time.Duration
	for range x {
		durations = append(durations, dur)
	}

	return durations
}

func Position1(m *machine.Machine) {
	m.CC(bd.CHANNEL, bd.LEVEL, 100)
	m.CC(bd.CHANNEL, bd.PARAM1, 0)
	m.CC(bd.CHANNEL, bd.PARAM2, 57)
	m.CC(bd.CHANNEL, bd.PARAM3, 0)
	m.CC(bd.CHANNEL, bd.PARAM4, 126)
	m.CC(bd.CHANNEL, bd.PARAM5, 64)
	m.CC(bd.CHANNEL, bd.PARAM6, 85)
	m.CC(bd.CHANNEL, bd.PARAM7, 0)
	m.CC(bd.CHANNEL, bd.PARAM8, 99)
	m.CC(bd.CHANNEL, bd.AMDEPTH, 100)
	m.CC(bd.CHANNEL, bd.AMRATE, 100)
	m.CC(bd.CHANNEL, bd.EQFREQ, 98)
	m.CC(bd.CHANNEL, bd.EQGAIN, 37)
	m.CC(bd.CHANNEL, bd.FILTERBASEFRQ, 100)
	m.CC(bd.CHANNEL, bd.FILTERWIDTH, 100)
	m.CC(bd.CHANNEL, bd.FILTERQ, 100)
	m.CC(bd.CHANNEL, bd.SAMPLERATEREDUCTION, 97)
	m.CC(bd.CHANNEL, bd.DISTORTION, 0)
	// m.CC(bd.CHANNEL, bd.VOLUME, 100)
	m.CC(bd.CHANNEL, bd.PAN, 60)
	m.CC(bd.CHANNEL, bd.DELAY, 0)
	m.CC(bd.CHANNEL, bd.REVERB, 0)
	m.CC(bd.CHANNEL, bd.LFOSPEED, 69)
	m.CC(bd.CHANNEL, bd.LFOAMOUNT, 100)
	m.CC(bd.CHANNEL, bd.LFOSHAPE, 100)
}

func Position2(m *machine.Machine) {
	n := rand.Intn(127)
	// m.CC(bd.CHANNEL, bd.LEVEL, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM1, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM2, int8(n))
	// m.CC(bd.CHANNEL, bd.PARAM3, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM4, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM5, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM6, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM7, int8(n))
	m.CC(bd.CHANNEL, bd.PARAM8, int8(n))
	// m.CC(bd.CHANNEL, bd.AMDEPTH, int8(n))
	// m.CC(bd.CHANNEL, bd.AMRATE, int8(n))
	// m.CC(bd.CHANNEL, bd.EQFREQ, int8(n))
	// m.CC(bd.CHANNEL, bd.EQGAIN, int8(n))
	// m.CC(bd.CHANNEL, bd.FILTERBASEFRQ, int8(n))
	// m.CC(bd.CHANNEL, bd.FILTERWIDTH, int8(n))
	// m.CC(bd.CHANNEL, bd.FILTERQ, int8(n))
	// m.CC(bd.CHANNEL, bd.SAMPLERATEREDUCTION, int8(n))
	// m.CC(bd.CHANNEL, bd.DISTORTION, int8(n))
	// // m.CC(bd.CHANNEL, bd.VOLUME, 100)
	m.CC(bd.CHANNEL, bd.PAN, int8(n))
	// m.CC(bd.CHANNEL, bd.DELAY, int8(n))
	// m.CC(bd.CHANNEL, bd.REVERB, int8(n))
	// m.CC(bd.CHANNEL, bd.LFOSPEED, int8(n))
	// m.CC(bd.CHANNEL, bd.LFOAMOUNT, int8(n))
	// m.CC(bd.CHANNEL, bd.LFOSHAPE, int8(n))
}

package synthesis

import (
	"time"

	"github.com/bh90210/models/machine"
	"github.com/bh90210/models/machine/bd"
	"github.com/bh90210/models/machine/sd"
)

type Synthesis struct {
	Tracks []*machine.Track
	Events []machine.Events
}

func (s *Synthesis) Run() {
	for _, events := range s.Events {
		events.Run()
	}
}

func EqualDuration(x int, dur time.Duration) []time.Duration {
	var durations []time.Duration
	for range x {
		durations = append(durations, dur)
	}

	return durations
}

func (s *Synthesis) Position1() {
	for _, t := range s.Tracks {
		t.Level = 127
		t.Volume = 127
		t.Param1 = 95
		t.Param2 = 70
		t.Param3 = 126
		t.Param4 = 126
		t.Param5 = 83
		t.Param6 = 64
		t.Param7 = 1
		t.Param8 = 98
		t.LFOSpeed = 8
		t.LFOAmount = 45
		t.LFOShape = 0
		t.Pan = 60
		t.Delay = 10
		t.Reverb = 15

		t.AMDepth = 0
		t.AMRate = 0
		t.EQFreq = 0
		t.EQGain = 0
		t.FilterBaseFrq = 0
		t.FilterWidth = 127
		t.FilterQ = 0
		t.SampleRateReduction = 0
		t.Distortion = 0
	}
}

func BDCC(m *machine.Machine, t machine.Track) {
	m.CC(bd.CHANNEL, bd.LEVEL, t.Level)
	m.CC(bd.CHANNEL, bd.PARAM1, t.Param1)
	m.CC(bd.CHANNEL, bd.PARAM2, t.Param2)
	m.CC(bd.CHANNEL, bd.PARAM3, t.Param3)
	m.CC(bd.CHANNEL, bd.PARAM4, t.Param4)
	m.CC(bd.CHANNEL, bd.PARAM5, t.Param5)
	m.CC(bd.CHANNEL, bd.PARAM6, t.Param6)
	m.CC(bd.CHANNEL, bd.PARAM7, t.Param7)
	m.CC(bd.CHANNEL, bd.PARAM8, t.Param8)
	m.CC(bd.CHANNEL, bd.AMDEPTH, t.AMDepth)
	m.CC(bd.CHANNEL, bd.AMRATE, t.AMRate)
	m.CC(bd.CHANNEL, bd.EQFREQ, t.EQFreq)
	m.CC(bd.CHANNEL, bd.EQGAIN, t.EQGain)
	m.CC(bd.CHANNEL, bd.FILTERBASEFRQ, t.FilterBaseFrq)
	m.CC(bd.CHANNEL, bd.FILTERWIDTH, t.FilterWidth)
	m.CC(bd.CHANNEL, bd.FILTERQ, t.FilterQ)
	m.CC(bd.CHANNEL, bd.SAMPLERATEREDUCTION, t.SampleRateReduction)
	m.CC(bd.CHANNEL, bd.DISTORTION, t.Distortion)
	m.CC(bd.CHANNEL, bd.VOLUME, t.Volume)
	m.CC(bd.CHANNEL, bd.PAN, t.Pan)
	m.CC(bd.CHANNEL, bd.DELAY, t.Delay)
	m.CC(bd.CHANNEL, bd.REVERB, t.Reverb)
	m.CC(bd.CHANNEL, bd.LFOSPEED, t.LFOSpeed)
	m.CC(bd.CHANNEL, bd.LFOAMOUNT, t.LFOAmount)
	m.CC(bd.CHANNEL, bd.LFOSHAPE, t.LFOShape)
}

func SDCC(m *machine.Machine, t machine.Track) {
	m.CC(sd.CHANNEL, sd.LEVEL, t.Level)
	m.CC(sd.CHANNEL, sd.PARAM1, t.Param1)
	m.CC(sd.CHANNEL, sd.PARAM2, t.Param2)
	m.CC(sd.CHANNEL, sd.PARAM3, t.Param3)
	m.CC(sd.CHANNEL, sd.PARAM4, t.Param4)
	m.CC(sd.CHANNEL, sd.PARAM5, t.Param5)
	m.CC(sd.CHANNEL, sd.PARAM6, t.Param6)
	m.CC(sd.CHANNEL, sd.PARAM7, t.Param7)
	m.CC(sd.CHANNEL, sd.PARAM8, t.Param8)
	m.CC(sd.CHANNEL, sd.AMDEPTH, t.AMDepth)
	m.CC(sd.CHANNEL, sd.AMRATE, t.AMRate)
	m.CC(sd.CHANNEL, sd.EQFREQ, t.EQFreq)
	m.CC(sd.CHANNEL, sd.EQGAIN, t.EQGain)
	m.CC(sd.CHANNEL, sd.FILTERBASEFRQ, t.FilterBaseFrq)
	m.CC(sd.CHANNEL, sd.FILTERWIDTH, t.FilterWidth)
	m.CC(sd.CHANNEL, sd.FILTERQ, t.FilterQ)
	m.CC(sd.CHANNEL, sd.SAMPLERATEREDUCTION, t.SampleRateReduction)
	m.CC(sd.CHANNEL, sd.DISTORTION, t.Distortion)
	m.CC(sd.CHANNEL, sd.VOLUME, t.Volume)
	m.CC(sd.CHANNEL, sd.PAN, t.Pan)
	m.CC(sd.CHANNEL, sd.DELAY, t.Delay)
	m.CC(sd.CHANNEL, sd.REVERB, t.Reverb)
	m.CC(sd.CHANNEL, sd.LFOSPEED, t.LFOSpeed)
	m.CC(sd.CHANNEL, sd.LFOAMOUNT, t.LFOAmount)
	m.CC(sd.CHANNEL, sd.LFOSHAPE, t.LFOShape)
}

func Rhythm1(i int, d time.Duration) time.Duration {
	switch {
	case i%23 == 0:
		d = d * 2
	case i%17 == 0:
		d = d / 2
	}

	return d
}

func Rhythm2(i int, d time.Duration) time.Duration {
	switch {
	case i%7 == 0:
		d = d * 7
	case i%6 == 0:
		d = d / 6
	case i%5 == 0:
		d = d * 5
	case i%4 == 0:
		d = d / 4
	case i%3 == 0:
		d = d * 3
	case i%2 == 0:
		d = d / 2
	}

	return d
}

func Rhythm3(i int, d time.Duration) time.Duration {
	switch {
	case i%9 == 0:
		d = d * 9
	case i%8 == 0:
		d = d / 8
	case i%7 == 0:
		d = d * 7
	case i%6 == 0:
		d = d / 6
	case i%5 == 0:
		d = d * 5
	case i%4 == 0:
		d = d / 4
	case i%3 == 0:
		d = d * 3
	case i%2 == 0:
		d = d / 2
	}

	return d
}

func Rhythm4(i int, d time.Duration) time.Duration {
	switch {
	case i%11 == 0:
		d = d * 11
	case i%9 == 0:
		d = d * 9
	case i%8 == 0:
		d = d / 8
	case i%7 == 0:
		d = d * 7
	case i%6 == 0:
		d = d / 6
	case i%5 == 0:
		d = d * 5
	case i%4 == 0:
		d = d / 4
	case i%3 == 0:
		d = d * 3
	case i%2 == 0:
		d = d / 2
	}

	return d
}

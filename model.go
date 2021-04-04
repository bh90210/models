package elektronmodels

import (
	"time"
)

type track uint8

const (
	T1 track = iota
	T2
	T3
	T4
	T5
	T6
)

type cctrack uint8

const (
	cct1 cctrack = 0xB0
	cct2 cctrack = 0xB1
	cct3 cctrack = 0xB2
	cct4 cctrack = 0xB3
	cct5 cctrack = 0xB4
	cct6 cctrack = 0xB5
)

type cc struct {
	pamVal map[Parameter]uint8
}

const (
	A0 uint8 = iota + 21
	As0
	B0
	C1
	Cs1
	D1
	Ds1
	E1
	F1
	Fs1
	G1
	Gs1
	A1
	As1
	B1
	C2
	Cs2
	D2
	Ds2
	E2
	F2
	Fs2
	G2
	Gs2
	A2
	As2
	B2
	C3
	Cs3
	D3
	Ds3
	E3
	F3
	Fs3
	G3
	Gs3
	A3
	As3
	B3
	C4
	Cs4
	D4
	Ds4
	E4
	F4
	Fs4
	G4
	Gs4
	A4
	As4
	B4
	C5
	Cs5
	D5
	Ds5
	E5
	F5
	Fs5
	G5
	Gs5
	A5
	As5
	B5
	C6
	Cs6
	D6
	Ds6
	E6
	F6
	Fs6
	G6
	Gs6
	A6
	As6
	B6
	C7
	Cs7
	D7
	Ds7
	E7
	F7
	Fs7
	G7
	Gs7
	A7
	As7
	B7
	C8

	Bf0 uint8 = As0
	Df1 uint8 = Cs1
	Ef1 uint8 = Ds1
	Gf1 uint8 = Fs1
	Af1 uint8 = Gs1
	Bf1 uint8 = As1
	Df2 uint8 = Cs2
	Ef2 uint8 = Ds2
	Gf2 uint8 = Fs2
	Af2 uint8 = Gs2
	Bf2 uint8 = As2
	Df3 uint8 = Cs3
	Ef3 uint8 = Ds3
	Gf3 uint8 = Fs3
	Af3 uint8 = Gs3
	Bf3 uint8 = As3
	Df4 uint8 = Cs4
	Ef4 uint8 = Ds4
	Gf4 uint8 = Fs4
	Af4 uint8 = Gs4
	Bf4 uint8 = As4
	Df5 uint8 = Cs5
	Ef5 uint8 = Ds5
	Gf5 uint8 = Fs5
	Af5 uint8 = Gs5
	Bf5 uint8 = As5
	Df6 uint8 = Cs6
	Ef6 uint8 = Ds6
	Gf6 uint8 = Fs6
	Af6 uint8 = Gs6
	Bf6 uint8 = As6
	Df7 uint8 = Cs7
	Ef7 uint8 = Ds7
	Gf7 uint8 = Fs7
	Af7 uint8 = Gs7
	Bf7 uint8 = As7
)

const (
	Unisonx2 uint8 = iota
	Unisonx3
	Unisonx4
	Minor
	Major
	Sus2
	Sus4
	MinorMinor7
	MajorMinor7
	MinorMajor7
	MajorMajor7
	MinorMinor7Sus4
	Dim7
	MinorAdd9
	MajorAdd9
	Minor6
	Major6
	Minorb5
	Majorb5
	MinorMinor7b5
	MajorMinor7b5
	MajorAug5
	MinorMinor7Aug5
	MajorMinor7Aug5
	Minorb6
	MinorMinor9no5
	MajorMinor9no5
	MajorAdd9b5
	MajorMajor7b5
	MajorMinor7b9no5
	Sus4Aug5b9
	Sus4AddAug5
	MajorAddb5
	Major6Add4no5
	MajorMajor76no5
	MajorMajor9no5
	Fourths
	Fifths
)

type note struct {
	dur *time.Duration
	key uint8
}

type Parameter uint8

const (
	NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10

	// model:cycles
	CYCLESPITCH Parameter = 65
	DECAY       Parameter = 80
	COLOR       Parameter = 16
	SHAPE       Parameter = 17

	// model:samples
	PITCH        Parameter = 16
	SAMPLESTART  Parameter = 19
	SAMPLELENGTH Parameter = 20
	CUTOFF       Parameter = 74
	RESONANCE    Parameter = 71

	// model:cycles
	PUNCH Parameter = 66
	GATE  Parameter = 67

	// model:samples
	LOOP    Parameter = 17
	REVERSE Parameter = 18

	SWEEP   Parameter = 18
	CONTOUR Parameter = 19
	DELAY   Parameter = 12
	REVERB  Parameter = 13

	VOLUMEDIST Parameter = 7
	SWING      Parameter = 15
	CHANCE     Parameter = 14

	// LFO section
	LFOSPEED Parameter = iota + 102
	LFOMULTIPIER
	LFOFADE
	LFODEST
	LFOWAVEFORM
	LFOSTARTPHASE
	LFORESET
	LFODEPTH
	// LFOMULTIPIER  Parameter = 103
	// LFOFADE       Parameter = 104
	// LFODEST       Parameter = 105
	// LFOWAVEFORM   Parameter = 106
	// LFOSTARTPHASE Parameter = 107
	// LFORESET      Parameter = 108
	// LFODEPTH      Parameter = 109

	// FX section
	DELAYTIME Parameter = iota + 85
	DELAYFEEDBACK
	REVERBZISE
	REBERBTONE
	// DELAYFEEDBACK Parameter = 86
	// REVERBZISE    Parameter = 87
	// REBERBTONE    Parameter = 88
)

const (
	LNONE  uint8 = 0
	LPITCH uint8 = 9

	LCOLOR uint8 = iota + 9
	LSHAPE
	LSWEEP
	LCONTOUR
	LPAW
	LGATE
	LFTUN
	LDECAY
	LDIST
	LDELAY
	LREVERB
	LPAN
)

// Project .
type Project struct {
	Name     string
	Patterns []*Pattern
}

// Pattern .
type Pattern struct {
	*Scale
	// Cycle manuel '9.11 Scale Menu'.
	// If true Scale Mode is set to PATTERN
	// if false to TRACK.
	ScaleMode bool
	Tracks    [6]*Track
}

// Track .
type Track struct {
	*Scale
	*Preset
	Trigs []*Trig
}

// Trig .
type Trig struct {
	*Preset
	*Lock
}

// Preset .
type Preset struct {
	*Parameters
}

// Param .
type Parameters struct {
	Parameter []*Parameter
}

// Lock .
type Lock struct {
	*Preset
}

// Scale .
type Scale struct {
	Len int
	Scl int
}

// Scale .

// type Trig struct {
// 	track
// 	note
// 	cctrack
// 	level uint8
// 	dur   *time.Duration
// 	*cc
// 	beat     uint8
// 	lastBeat bool
// 	hasCC    bool
// 	hasNote  bool
// }

// func NewTrig(beat uint8) *Trig {
// 	return &Trig{
// 		beat: beat,
// 	}
// }

// func LastTrig(beat uint8) *Trig {
// 	trig := Trig{
// 		beat:     beat,
// 		lastBeat: true,
// 	}
// 	return &trig
// }

// func (t *Trig) Note(key uint8, level uint8, dur time.Duration) {
// 	tn := &note{}
// 	tn.key = key
// 	tn.dur = &dur

// 	t.note = *tn
// 	t.level = level
// 	t.hasNote = true
// }

// func (t *Trig) CC(values map[Parameter]uint8) {
// 	t.cc = &cc{
// 		pamVal: values,
// 	}
// 	t.hasCC = true
// }

// type Track struct {
// 	number *track
// 	trig   []*Trig
// }

// func NewTrack(trackNumber track, variadicTracks []*Trig) *Track {
// 	for _, v := range variadicTracks {
// 		switch trackNumber {
// 		case T1:
// 			v.track = T1
// 		case T2:
// 			v.track = T2
// 		case T3:
// 			v.track = T3
// 		case T4:
// 			v.track = T4
// 		case T5:
// 			v.track = T5
// 		case T6:
// 			v.track = T6
// 		}
// 	}
// 	track := Track{
// 		number: &trackNumber,
// 		trig:   variadicTracks,
// 	}

// 	return &track
// }

// type pattern struct {
// 	tracks []*Track
// }

// type Project struct {
// 	pattern []*pattern

// 	drv midi.Driver
// 	mu  *sync.Mutex

// 	inPorts  map[int]midi.In
// 	outPorts map[int]midi.Out

// 	wr *writer.Writer

// 	loop bool
// }

// func NewProject() (*Project, error) {
// 	drv, mu, err := newmidi()
// 	if err != nil {
// 		return nil, err
// 	}

// 	project := &Project{
// 		drv: drv,
// 		mu:  mu,
// 	}

// 	ii := make(map[int]midi.In)
// 	oo := make(map[int]midi.Out)

// 	mu.Lock()
// 	ins, _ := drv.Ins()
// 	for i, in := range ins {
// 		ii[i] = in
// 	}

// 	outs, _ := drv.Outs()
// 	for i, out := range outs {
// 		oo[i] = out
// 	}

// 	project.inPorts = ii
// 	project.outPorts = oo

// 	// TODO: FIX THIS MESS
// 	project.outPorts[1].Open()

// 	wr := writer.New(project.outPorts[1])

// 	project.wr = wr
// 	mu.Unlock()

// 	return project, nil
// }

// func (p *Project) Loop() {
// 	p.loop = true
// }

// func (p *Project) Play() error {
// 	// calculate how many patterns are present
// 	totalPatters := len(p.pattern)
// 	fmt.Println("Total Patterns: ", totalPatters)

// 	// map[pattern'sBeatCounting]triggersForAll6tracks(ifPresent)
// 	type tri map[int][]*Trig

// 	// map[patternsLength]collectionOfArrangedTriggers
// 	type timeline map[int]*tri

// 	// init a key/value map acting as timeline for the song
// 	timlin := make(timeline)

// 	type patternsEnd map[int]int
// 	ends := make(patternsEnd)

// 	// range over patterns slice []pattern
// 	for i, pattern := range p.pattern {
// 		fmt.Println("Pattern: ", i)

// 		patternTimeline := make(tri)
// 		patternEndingHelper := make(patternsEnd)

// 		totalTracks := len(pattern.tracks)
// 		fmt.Println("Total Tracks: ", totalTracks)

// 		// range over individual pattern's tracks
// 		for tid, track := range pattern.tracks {
// 			// count total trigs in track
// 			trigsLength := len(track.trig)
// 			fmt.Println("Track: ", track.number, " Total Trigs: ", trigsLength)

// 			for ii := 0; ii < trigsLength; ii++ {
// 				patternTimeline[int(track.trig[ii].beat)] = append(patternTimeline[int(track.trig[ii].beat)], track.trig[ii])

// 				if track.trig[ii].lastBeat == true {
// 					patternEndingHelper[tid] = int(track.trig[ii].beat)
// 				}
// 			}
// 		}

// 		// compare last beats to find where pattern ends
// 		var longestBeat []int
// 		for _, v := range patternEndingHelper {
// 			longestBeat = append(longestBeat, int(v))
// 		}
// 		sorted := sort.IntSlice(longestBeat)

// 		// assign pattern end to relevant map
// 		last := len(sorted)
// 		ends[i] = sorted[last-1]
// 		timlin[i] = &patternTimeline
// 	}

// 	// range over timeline (map)
// 	// each iteration return song's individual patterns in order
// 	for k, pattern := range timlin {
// 		beat := *pattern

// 		// loop over patterns' trigs (beats) until ends[pattern]longestBeat
// 		lastBeat := ends[k]
// 		it := 0
// 		tick := time.NewTicker(4000 * time.Millisecond)

// 	loop:
// 		for {
// 			<-tick.C
// 			totalTrigsForBeat := beat[it]
// 			if len(totalTrigsForBeat) != 0 {
// 				// total trigs for current beat range
// 				for _, trig := range totalTrigsForBeat {
// 					if trig.hasCC == true {
// 						p.cc(&trig.track, trig.cc.pamVal)
// 					}

// 					if trig.hasNote == true {
// 						p.note(&trig.track, &trig.note, trig.note.key, trig.level)
// 					}
// 				}

// 				fmt.Println("lastbeat", lastBeat, "current beat", it)
// 			}

// 			if it == lastBeat-1 {
// 				break loop
// 			}
// 			it++
// 		}
// 	}

// 	return nil
// }

// func (c *Project) Pause() {
// }

// func (p *Project) Close() {
// 	p.drv.Close()
// }

// func (t *Project) Pattern(variadicTracks ...*Track) {
// 	pattern := &pattern{
// 		tracks: variadicTracks,
// 	}

// 	t.pattern = append(t.pattern, pattern)
// }

// func (c *Project) note(track *track, note *note, key uint8, velocity uint8) {
// 	timer := time.NewTimer(*note.dur)
// 	// note on
// 	c.mu.Lock()
// 	c.wr.SetChannel(uint8(*track))
// 	writer.NoteOn(c.wr, key, velocity)
// 	c.mu.Unlock()

// 	go func() {
// 		<-timer.C
// 		// note off
// 		c.mu.Lock()
// 		writer.NoteOff(c.wr, key)
// 		c.mu.Unlock()
// 	}()
// }

// func (c *Project) cc(track *track, ccvalues map[Parameter]uint8) {
// 	for k, v := range ccvalues {
// 		c.mu.Lock()
// 		c.wr.SetChannel(uint8(*track))
// 		// writer.ProgramChange()
// 		writer.ControlChange(c.wr, uint8(k), v)
// 		c.mu.Unlock()
// 	}
// }

// // func (c *Cycles) pc(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
// // }

// func newmidi() (*driver.Driver, *sync.Mutex, error) {
// 	var err error
// 	drv, err := driver.New()
// 	if err != nil {
// 		panic("can't initialize driver")
// 	}

// 	mutex := &sync.Mutex{}

// 	return drv, mutex, nil
// }

//
// Timing
//

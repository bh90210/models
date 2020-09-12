package elektronmodels

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

//
// Mapping
//

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

// func (c *cctrack) int64() int64 {
// 	return int64(*c)
// }

type cc struct {
	pamVal map[Parameter]uint8
}

type notes uint8

const (
	A0 notes = iota + 21
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

	Bf0 notes = As0
	Df1 notes = Cs1
	Ef1 notes = Ds1
	Gf1 notes = Fs1
	Af1 notes = Gs1
	Bf1 notes = As1
	Df2 notes = Cs2
	Ef2 notes = Ds2
	Gf2 notes = Fs2
	Af2 notes = Gs2
	Bf2 notes = As2
	Df3 notes = Cs3
	Ef3 notes = Ds3
	Gf3 notes = Fs3
	Af3 notes = Gs3
	Bf3 notes = As3
	Df4 notes = Cs4
	Ef4 notes = Ds4
	Gf4 notes = Fs4
	Af4 notes = Gs4
	Bf4 notes = As4
	Df5 notes = Cs5
	Ef5 notes = Ds5
	Gf5 notes = Fs5
	Af5 notes = Gs5
	Bf5 notes = As5
	Df6 notes = Cs6
	Ef6 notes = Ds6
	Gf6 notes = Fs6
	Af6 notes = Gs6
	Bf6 notes = As6
	Df7 notes = Cs7
	Ef7 notes = Ds7
	Gf7 notes = Fs7
	Af7 notes = Gs7
	Bf7 notes = As7
)

// func (n *notes) int64() int64 {
// 	return int64(*n)
// }

type chord uint8

const (
	Unisonx2 chord = iota
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
	on  noteOn
	off noteOff
	dur *time.Duration
	key notes
}

type noteOn uint8

const (
	t1on noteOn = 0x90
	t2on noteOn = 0x91
	t3on noteOn = 0x92
	t4on noteOn = 0x93
	t5on noteOn = 0x94
	t6on noteOn = 0x95
)

// func (n *noteOn) int64() int64 {
// 	return int64(*n)
// }

type noteOff uint8

const (
	t1off noteOff = 0x80
	t2off noteOff = 0x81
	t3off noteOff = 0x82
	t4off noteOff = 0x83
	t5off noteOff = 0x84
	t6off noteOff = 0x85
)

// func (n *noteOff) int64() int64 {
// 	return int64(*n)
// }

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
	LFOSPEED      Parameter = 102
	LFOMULTIPIER  Parameter = 103
	LFOFADE       Parameter = 104
	LFODEST       Parameter = 105
	LFOWAVEFORM   Parameter = 106
	LFOSTARTPHASE Parameter = 107
	LFORESET      Parameter = 108
	LFODEPTH      Parameter = 109

	// FX section
	DELAYTIME     Parameter = 85
	DELAYFEEDBACK Parameter = 86
	REVERBZISE    Parameter = 87
	REBERBTONE    Parameter = 88
)

// func (p *Parameter) int64() int64 {
// 	return int64(*p)
// }

type lfoDest uint8

const (
	LNONE  lfoDest = 0
	LPITCH lfoDest = 9

	LCOLOR lfoDest = iota + 9
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

//
// Project
//

type Trig struct {
	track
	note
	cctrack
	level uint8
	dur   *time.Duration
	*cc
	beat     uint8
	lastBeat bool
	hasCC    bool
	hasNote  bool
}

func NewTrig(beat uint8) *Trig {
	return &Trig{
		beat: beat,
	}
}

func LastTrig(beat uint8) *Trig {
	trig := Trig{
		beat:     beat,
		lastBeat: true,
	}
	return &trig
}

func (t *Trig) Note(key notes, level uint8, dur time.Duration) {
	tn := &note{}
	tn.key = key
	tn.dur = &dur

	t.note = *tn
	t.level = level
	t.hasNote = true
}

func (t *Trig) CC(values map[Parameter]uint8) {
	t.cc = &cc{
		pamVal: values,
	}
	t.hasCC = true
}

type Track struct {
	number *track
	trig   []*Trig
}

func NewTrack(trackNumber track, variadicTracks ...*Trig) *Track {
	for _, v := range variadicTracks {
		switch trackNumber {
		case T1:
			v.track = T1
			v.note.on = t1on
			v.note.off = t1off
			v.cctrack = cct1
		case T2:
			v.track = T2
			v.note.on = t2on
			v.note.off = t2off
			v.cctrack = cct2
		case T3:
			v.track = T3
			v.note.on = t3on
			v.note.off = t3off
			v.cctrack = cct3
		case T4:
			v.track = T4
			v.note.on = t4on
			v.note.off = t4off
			v.cctrack = cct4
		case T5:
			v.track = T5
			v.note.on = t5on
			v.note.off = t5off
			v.cctrack = cct5
		case T6:
			v.track = T6
			v.note.on = t6on
			v.note.off = t6off
			v.cctrack = cct6
		}
	}
	track := Track{
		number: &trackNumber,
		trig:   variadicTracks,
	}

	return &track
}

type pattern struct {
	tracks []*Track
}

type Project struct {
	pattern []*pattern

	drv midi.Driver
	mu  *sync.Mutex

	inPorts  map[int]midi.In
	outPorts map[int]midi.Out

	wr *writer.Writer

	loop bool
}

func NewProject() (*Project, error) {
	drv, mu, err := newmidi()
	if err != nil {
		return nil, err
	}

	project := &Project{
		drv: &drv,
		mu:  mu,
	}

	ii := make(map[int]midi.In)
	oo := make(map[int]midi.Out)

	mu.Lock()
	ins, _ := drv.Ins()
	for i, in := range ins {
		ii[i] = in
	}

	outs, _ := drv.Outs()
	for i, out := range outs {
		oo[i] = out
	}

	project.inPorts = ii
	project.outPorts = oo

	project.outPorts[2].Open()
	// outs[2].Open()
	wr := writer.New(project.outPorts[2])
	fmt.Println(outs[2])
	project.wr = wr
	mu.Unlock()

	return project, nil
}

func (p *Project) Loop() {
	p.loop = true
}

func (p *Project) Play() error {
	// calculate how many patterns are present
	totalPatters := len(p.pattern)
	fmt.Println("Total Patterns: ", totalPatters)

	// map[pattern'sBeatCounting]triggersForAll6tracks(ifPresent)
	type tri map[int][]*Trig

	// map[patternsLength]collectionOfArrangedTriggers
	type timeline map[int]*tri

	// init a key/value map acting as timeline for the song
	timlin := make(timeline)

	type patternsEnd map[int]int
	ends := make(patternsEnd)

	// range over patterns slice []pattern
	for i, pattern := range p.pattern {
		fmt.Println("Pattern: ", i)

		patternTimeline := make(tri)
		patternEndingHelper := make(patternsEnd)

		totalTracks := len(pattern.tracks)
		fmt.Println("Total Tracks: ", totalTracks)

		// range over individual pattern's tracks
		for tid, track := range pattern.tracks {
			// count total trigs in track
			trigsLength := len(track.trig)
			fmt.Println("Track: ", track.number, " Total Trigs: ", trigsLength)

			for ii := 0; ii < trigsLength; ii++ {
				patternTimeline[int(track.trig[ii].beat)] = append(patternTimeline[int(track.trig[ii].beat)], track.trig[ii])

				if track.trig[ii].lastBeat == true {
					patternEndingHelper[tid] = int(track.trig[ii].beat)
				}
			}
		}

		// compare last beats to find where pattern ends
		var longestBeat []int
		for _, v := range patternEndingHelper {
			longestBeat = append(longestBeat, int(v))
		}
		sorted := sort.IntSlice(longestBeat)

		// assign pattern end to relevant map
		last := len(sorted)
		ends[i] = sorted[last-1]
		timlin[i] = &patternTimeline
	}

	// range over timeline (map)
	// each iteration return song's individual patterns in order
	for k, pattern := range timlin {
		// var pat represents a single pattern (NewPattern())
		// beat := *pattern
		beat := *pattern

		// loop over patterns' trigs (beats) until ends[pattern]longestBeat
		lastBeat := ends[k]
		it := 0

	loop:
		for {
			totalTrigsForBeat := beat[it]
			if len(totalTrigsForBeat) != 0 {
				// total trigs for current beat range
				for _, trig := range totalTrigsForBeat {
					if trig.hasCC == true {
						p.cc(&trig.track, &trig.cctrack, trig.cc.pamVal)
					}

					if trig.hasNote == true {
						p.note(&trig.track, &trig.note, trig.note.key, trig.level)
					}
				}

				fmt.Println("lastbeat", lastBeat, "current beat", it)
			}

			if it == lastBeat-1 {
				break loop
			}

			fmt.Println("Sleeping", it)
			time.Sleep(700 * time.Millisecond)
			it++
		}
	}

	return nil
}

func (c *Project) Pause() {
}

func (p *Project) Close() {
	p.drv.Close()
}

func (t *Project) Pattern(variadicTracks ...*Track) {
	pattern := &pattern{
		tracks: variadicTracks,
	}

	t.pattern = append(t.pattern, pattern)
}

func (c *Project) note(track *track, note *note, key notes, velocity uint8) {
	timer := time.NewTimer(*note.dur)
	// note on
	c.mu.Lock()
	c.wr.SetChannel(uint8(*track))
	writer.NoteOn(c.wr, 60, 100)
	time.Sleep(time.Millisecond * 800)
	writer.NoteOff(c.wr, 60)
	fmt.Println(uint8(key), uint8(*track), velocity, "<<<<<<<<<<<<<<<<<<")
	writer.NoteOn(c.wr, uint8(key), velocity)
	// c.pm.WriteShort(track.on.int64(), note.int64(), intensity)
	c.mu.Unlock()

	go func() {
		<-timer.C

		// note off
		c.mu.Lock()
		writer.NoteOff(c.wr, uint8(key))
		// c.pm.WriteShort(track.off.int64(), note.int64(), intensity)
		c.mu.Unlock()
	}()
}

func (c *Project) cc(track *track, cctrack *cctrack, ccvalues map[Parameter]uint8) {
	wr := writer.New(c.outPorts[2])

	for k, v := range ccvalues {
		c.mu.Lock()
		wr.SetChannel(uint8(*track))
		writer.ControlChange(c.wr, uint8(k), v)
		// writer.CcOn(wr, uint8(k))
		// c.pm.WriteShort(track.int64(), k.int64(), v)
		c.mu.Unlock()
	}
}

// func (c *Cycles) pc(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
// }

func newmidi() (driver.Driver, *sync.Mutex, error) {
	var err error
	drv, err := driver.New()
	if err != nil {
		panic("can't initialize driver")
	}

	mutex := &sync.Mutex{}

	return *drv, mutex, nil
}

//
// Timing
//

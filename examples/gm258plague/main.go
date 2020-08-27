package main

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/rakyll/portmidi"
)

type Composition interface {
	HyperMeasure()
	Measure()
	Beat()
}

type GM258plague struct {
}

func (c *GM258plague) HyperMeasure() {

}

func (c *GM258plague) Measure() {

}

func (c *GM258plague) Beat() {

}

type HyperMeasure struct {
}

type Measure struct {
}

type Beat struct {
}

type LFODest int64

const (
	LNONE    LFODest = 0
	LPITCH   LFODest = 9
	LCOLOR   LFODest = 11
	LSHAPE   LFODest = 12
	LSWEEP   LFODest = 13
	LCONTOUR LFODest = 14
	LPAW     LFODest = 15
	LGATE    LFODest = 16
	LFTUN    LFODest = 17
	LDECAY   LFODest = 18
	LDIST    LFODest = 19
	LDELAY   LFODest = 20
	LREVERB  LFODest = 21
	LPAN     LFODest = 22
)

type Chord int64

const (
	Unisonx2 Chord = iota
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
	MagorMinor7Aug5
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

type Note int64

const (
	A0 Note = iota + 21
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
)

func (n *Note) Int64() int64 {
	return int64(*n)
}

type Track int64

const (
	T1 Track = iota
	T2
	T3
	T4
	T5
	T6
)

type NoteTrack struct {
	ON  NoteOn
	OFF NoteOff
	Dur *time.Duration
}

func NewNoteTrack(track Track, dur time.Duration) *NoteTrack {
	noteTrack := &NoteTrack{}
	noteTrack.Dur = &dur
	switch track {
	case T1:
		noteTrack.ON = T1ON
		noteTrack.OFF = T1OFF
	case T2:
		noteTrack.ON = T2ON
		noteTrack.OFF = T2OFF
	case T3:
		noteTrack.ON = T3ON
		noteTrack.OFF = T3OFF
	case T4:
		noteTrack.ON = T4ON
		noteTrack.OFF = T4OFF
	case T5:
		noteTrack.ON = T5ON
		noteTrack.OFF = T5OFF
	case T6:
		noteTrack.ON = T6ON
		noteTrack.OFF = T6OFF
	}
	return noteTrack
}

type NoteOn int64

const (
	T1ON NoteOn = 0x90
	T2ON NoteOn = 0x91
	T3ON NoteOn = 0x92
	T4ON NoteOn = 0x93
	T5ON NoteOn = 0x94
	T6ON NoteOn = 0x95
)

func (n *NoteOn) Int64() int64 {
	return int64(*n)
}

type NoteOff int64

const (
	T1OFF NoteOff = 0x80
	T2OFF NoteOff = 0x81
	T3OFF NoteOff = 0x82
	T4OFF NoteOff = 0x83
	T5OFF NoteOff = 0x84
	T6OFF NoteOff = 0x85
)

func (n *NoteOff) Int64() int64 {
	return int64(*n)
}

type CCtrack int64

const (
	CCT1 CCtrack = 0xB0
	CCT2 CCtrack = 0xB1
	CCT3 CCtrack = 0xB2
	CCT4 CCtrack = 0xB3
	CCT5 CCtrack = 0xB4
	CCT6 CCtrack = 0xB5
)

func (c *CCtrack) Int64() int64 {
	return int64(*c)
}

type Parameter int64

const (
	NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10

	PITCH Parameter = 65
	DECAY Parameter = 80
	COLOR Parameter = 16
	SHAPE Parameter = 17

	SWEEP   Parameter = 18
	CONTOUR Parameter = 19
	DELAY   Parameter = 12
	REVERB  Parameter = 13

	VOLUMEDIST Parameter = 7
	SWING      Parameter = 15
	CHANCE     Parameter = 14

	PUNCH Parameter = 66
	GATE  Parameter = 67

	LFOSPEED      Parameter = 102
	LFOMULTIPIER  Parameter = 103
	LFOFADE       Parameter = 104
	LFODEST       Parameter = 105
	LFOWAVEFORM   Parameter = 106
	LFOSTARTPHASE Parameter = 107
	LFORESET      Parameter = 108
	LFODEPTH      Parameter = 109

	DELAYTIME     Parameter = 85
	DELAYFEEDBACK Parameter = 86
	REVERBZISE    Parameter = 87
	REBERBTONE    Parameter = 88
)

func (p *Parameter) Int64() int64 {
	return int64(*p)
}

type Cycles struct {
	PM *portmidi.Stream
	// Chord

	Mu *sync.Mutex

	// Ranto int64
	// Dur   time.Duration
	// Tick *time.Ticker
}

func NewCycles() (*Cycles, error) {
	out, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		return nil, err
	}
	mutex := &sync.Mutex{}
	// tick := time.NewTicker(time.Duration(500 * time.Millisecond))
	return &Cycles{
		PM: out,
		Mu: mutex,
		// Dur:  500,
		// Tick: tick,
	}, nil
}

// Note .
func (c *Cycles) Note(track *NoteTrack, note Note, intensity int64, values map[Parameter]int64) {
	switch track.ON {
	case T1ON:
		c.CC(CCT1, values)
	case T2ON:
		c.CC(CCT2, values)
	case T3ON:
		c.CC(CCT3, values)
	case T4ON:
		c.CC(CCT4, values)
	case T5ON:
		c.CC(CCT5, values)
	case T6ON:
		c.CC(CCT6, values)
	}

	timer := time.NewTimer(*track.Dur)

	// note on
	c.Mu.Lock()
	c.PM.WriteShort(track.ON.Int64(), note.Int64(), intensity)
	c.Mu.Unlock()

	go func() {
		<-timer.C

		// note off
		c.Mu.Lock()
		c.PM.WriteShort(track.OFF.Int64(), note.Int64(), intensity)
		c.Mu.Unlock()
	}()
}

type CC struct {
	Track     CCtrack
	Parameter Parameter
	Value     int64
}

func NewCC(track CCtrack, parameter Parameter, value int64) *CC {
	return &CC{
		Track:     track,
		Parameter: parameter,
		Value:     value,
	}
}

// CC .
func (c *Cycles) CC(track CCtrack, value map[Parameter]int64) {
	for k, v := range value {
		c.Mu.Lock()
		c.PM.WriteShort(track.Int64(), k.Int64(), v)
		c.Mu.Unlock()
	}
}

func (c *Cycles) PC(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
}

func main() {
	cycles, err := NewCycles()
	if err != nil {
		log.Fatal(err)
	}

	// intro(cycles, 0)
	// intro(cycles, 1)
	// intro(cycles, 2)
	// intro(cycles, 3)
	// intro(cycles, 4)

	breake(cycles)

	cycles.PM.Close()
	os.Exit(0)

	select {}
}

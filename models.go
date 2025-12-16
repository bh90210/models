package models

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

type model string

// Model
const (
	CYCLES  model = "Model:Cycles"
	SAMPLES model = "Model:Samples"
)

// Channel represents a physical midi channel.
type Channel int8

// Voices/Tracks
const (
	T1 Channel = iota
	T2
	T3
	T4
	T5
	T6
)

// Notes are all notes reproducible by the machines.
type Notes int8

// Keys/letter notes
const (
	A0 Notes = iota + 21
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
	Cs8
	D8
	Ds8
	E8
	F8
	Fs8
	G8
	Gs8
	A8
	As8
	B8

	Bf0 Notes = As0
	Df1 Notes = Cs1
	Ef1 Notes = Ds1
	Gf1 Notes = Fs1
	Af1 Notes = Gs1
	Bf1 Notes = As1
	Df2 Notes = Cs2
	Ef2 Notes = Ds2
	Gf2 Notes = Fs2
	Af2 Notes = Gs2
	Bf2 Notes = As2
	Df3 Notes = Cs3
	Ef3 Notes = Ds3
	Gf3 Notes = Fs3
	Af3 Notes = Gs3
	Bf3 Notes = As3
	Df4 Notes = Cs4
	Ef4 Notes = Ds4
	Gf4 Notes = Fs4
	Af4 Notes = Gs4
	Bf4 Notes = As4
	Df5 Notes = Cs5
	Ef5 Notes = Ds5
	Gf5 Notes = Fs5
	Af5 Notes = Gs5
	Bf5 Notes = As5
	Df6 Notes = Cs6
	Ef6 Notes = Ds6
	Gf6 Notes = Fs6
	Af6 Notes = Gs6
	Bf6 Notes = As6
	Df7 Notes = Cs7
	Ef7 Notes = Ds7
	Gf7 Notes = Fs7
	Af7 Notes = Gs7
	Bf7 Notes = As7
	Df8 Notes = Cs8
	Ef8 Notes = Ds8
	Gf8 Notes = Fs8
	Af8 Notes = Gs8
	Bf8 Notes = As8
)

// Chords are all chords supported by the machines mapped custom type.
type Chords int8

// Chords
const (
	Unisonx2 Chords = iota
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

// Parameter is all track parameters of the physical machine.
// Sample has certain different key/values than Cycles.
type Parameter int8

const (
	NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10
	SWEEP      Parameter = 18
	CONTOUR    Parameter = 19
	DELAY      Parameter = 12
	REVERB     Parameter = 13
	VOLUMEDIST Parameter = 7
	// SWING      Parameter = 15
	// CHANCE     Parameter = 14

	// model:cycles
	MACHINE     Parameter = 64
	CYCLESPITCH Parameter = 65
	DECAY       Parameter = 80
	COLOR       Parameter = 16
	SHAPE       Parameter = 17
	PUNCH       Parameter = 66
	GATE        Parameter = 67

	// model:samples
	PITCH        Parameter = 16
	SAMPLESTART  Parameter = 19
	SAMPLELENGTH Parameter = 20
	CUTOFF       Parameter = 74
	RESONANCE    Parameter = 71
	LOOP         Parameter = 17
	REVERSE      Parameter = 18
)

// Reverb & Delay settings
const (
	DELAYTIME Parameter = iota + 85
	DELAYFEEDBACK
	REVERBSIZE
	REVERBTONE
)

// LFO settings
const (
	LFOSPEED Parameter = iota + 102
	LFOMULTIPIER
	LFOFADE
	LFODEST
	LFOWAVEFORM
	LFOSTARTPHASE
	LFORESET
	LFODEPTH
)

type machine int8

// Machines
const (
	KICK machine = iota
	SNARE
	METAL
	PERC
	TONE
	CHORD
)

var ErrNotImplemented = fmt.Errorf("not implemented")

type MidiCom interface {
	Preset(channel Channel, preset Preset) error
	Note(channel Channel, note Notes, velocity int8, duration float64) error
	CC(channel Channel, parameter Parameter, value int8) error
	PC(channel Channel, pc int8) error
	Incoming() chan []byte
	Close()
}

var _ MidiCom = (*Project)(nil)

// Project long description of the data structure, methods, behaviors and useage.
type Project struct {
	model

	mu *sync.Mutex
	// midi fields
	drv midi.Driver
	in  midi.In
	out midi.Out
	wr  *writer.Writer
}

// NewProject initiates and returns a *Project struct.
func NewProject(m model) (*Project, error) {
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	p := &Project{
		model: m,
		mu:    new(sync.Mutex),
		drv:   drv,
	}

	p.mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), string(m)) {
			p.in = in
		}
	}

	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), string(m)) {
			p.out = out
		}
	}

	// check if nothing found
	if p.in == nil && p.out == nil {
		return nil, fmt.Errorf("device %s not found", m)
	}

	err = p.in.Open()
	if err != nil {
		return nil, err
	}

	err = p.out.Open()
	if err != nil {
		return nil, err
	}

	wr := writer.New(p.out)
	p.wr = wr
	p.mu.Unlock()

	return p, nil
}

// Preset immediately sets (CC) provided parameters.
func (p *Project) Preset(track Channel, preset Preset) error {
	for parameter, value := range preset {
		err := p.CC(track, parameter, value)
		if err != nil {
			fmt.Println("Error sending CC:", err)
			return err
		}
	}

	return nil
}

// Note fires immediately a midi note on signal followed by a note off specified duration in milliseconds (ms).
// Optionally user can pass a preset too for convenience.
func (p *Project) Note(track Channel, note Notes, velocity int8, duration float64) error {
	p.wr.SetChannel(uint8(track))
	err := writer.NoteOn(p.wr, uint8(note), uint8(velocity))
	if err != nil {
		fmt.Println("Error sending NoteOn:", err)
		return err
	}

	time.Sleep(time.Millisecond * time.Duration(duration))

	p.wr.SetChannel(uint8(track))
	err = writer.NoteOff(p.wr, uint8(note))
	if err != nil {
		fmt.Println("Error sending NoteOff:", err)
		return err
	}

	return nil
}

// CC control change.
func (p *Project) CC(track Channel, parameter Parameter, value int8) error {
	p.wr.SetChannel(uint8(track))
	return writer.ControlChange(p.wr, uint8(parameter), uint8(value))
}

// PC Project control change.
func (p *Project) PC(t Channel, pc int8) error {
	p.wr.SetChannel(uint8(t))
	return writer.ProgramChange(p.wr, uint8(pc))
}

// Read reads incoming midi data.
func (p *Project) Incoming() chan []byte {
	ch := make(chan []byte)
	p.in.SetListener(func(p []byte, deltaMicroseconds int64) {
		select {
		case ch <- p:
		default:
			fmt.Println("no models receiver")
		}
	})

	return ch
}

// Close midi connection. Use it with defer after creating a new project.
func (p *Project) Close() {
	p.in.Close()
	p.out.Close()
	p.drv.Close()
}

// Preset represents a machine's preset.
type Preset map[Parameter]int8

// PT1 is the cycles preset for track 1.
func PT1() Preset {
	p := make(map[Parameter]int8)
	p[MACHINE] = int8(KICK)
	p[TRACKLEVEL] = int8(120)
	p[MUTE] = int8(0)
	p[PAN] = int8(63)
	p[SWEEP] = int8(16)
	p[CONTOUR] = int8(24)
	p[DELAY] = int8(0)
	p[REVERB] = int8(0)
	p[VOLUMEDIST] = int8(60)
	p[CYCLESPITCH] = int8(64)
	p[DECAY] = int8(29)
	p[COLOR] = int8(10)
	p[SHAPE] = int8(16)
	p[PUNCH] = int8(0)
	p[GATE] = int8(0)
	return p
}

// PT2 is the cycles preset for track 2.
func PT2() Preset {
	p := PT1()
	p[MACHINE] = int8(SNARE)
	p[SWEEP] = int8(8)
	p[CONTOUR] = int8(0)
	p[DECAY] = int8(40)
	p[COLOR] = int8(0)
	p[SHAPE] = int8(127)
	return p
}

// PT3 is the cycles preset for track 3.
func PT3() Preset {
	p := PT1()
	p[MACHINE] = int8(METAL)
	p[SWEEP] = int8(48)
	p[CONTOUR] = int8(0)
	p[DECAY] = int8(20)
	p[COLOR] = int8(16)
	p[SHAPE] = int8(46)
	return p
}

// PT4 is the cycles preset for track 4.
func PT4() Preset {
	p := PT1()
	p[MACHINE] = int8(PERC)
	p[SWEEP] = int8(100)
	p[CONTOUR] = int8(64)
	p[DECAY] = int8(26)
	p[COLOR] = int8(15)
	p[SHAPE] = int8(38)
	return p
}

// PT5 is the cycles preset for track 5.
func PT5() Preset {
	p := PT1()
	p[MACHINE] = int8(TONE)
	p[SWEEP] = int8(38)
	p[CONTOUR] = int8(52)
	p[DECAY] = int8(42)
	p[COLOR] = int8(22)
	p[SHAPE] = int8(40)
	return p
}

// PT6 is the cycles preset for track 6.
func PT6() Preset {
	p := PT1()
	p[MACHINE] = int8(CHORD)
	p[SWEEP] = int8(43)
	p[CONTOUR] = int8(24)
	p[DECAY] = int8(64)
	p[COLOR] = int8(20)
	p[SHAPE] = int8(4)
	return p
}

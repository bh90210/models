package elektronmodels

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

type model uint8

const (
	CYCLES model = iota
	SAMPLES
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

type Chords uint8

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
	MACHINE     Parameter = 64
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
)

const (
	// LFO section
	LFOSPEED Parameter = iota + 102
	LFOMULTIPIER
	LFOFADE
	LFODEST
	LFOWAVEFORM
	LFOSTARTPHASE
	LFORESET
	LFODEPTH
)

const (
	// FX section
	DELAYTIME Parameter = iota + 85
	DELAYFEEDBACK
	REVERBZISE
	REBERBTONE
)

const (
	LNONE  Parameter = 0
	LPITCH Parameter = 9

	LCOLOR Parameter = iota + 9
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

// Machine section
const (
	KICK Parameter = iota
	SNARE
	METAL
	PERC
	TONE
	CHORD
)

type ScaleMode bool

const (
	PTN ScaleMode = true
	TRK ScaleMode = false
)

// Project .
type Project struct {
	patterns []*Pattern

	drv midi.Driver
	mu  *sync.Mutex

	in  midi.In
	out midi.Out

	wr *writer.Writer
}

// Pattern .
type Pattern struct {
	scale  *Scale
	tracks [6]*Track
}

// Track .
type Track struct {
	scale  *Scale
	preset *Preset
	trigs  []*Trig
}

// Preset .
type Preset struct {
	parameters map[Parameter]uint8
}

// Scale .
type Scale struct {
	// Cycle manuel '9.11 Scale Menu'.
	// If true Scale Mode is set to PATTERN
	// if false to TRACK.
	mod ScaleMode
	// Length sets the step length of thew pattern/track.
	len int
	// Scale controls the speed of the playback in multiples of the current tempo.
	scl int
	chg int
}

// Trig .
type Trig struct {
	preset *Preset
	lock   *Lock
}

// Lock .
type Lock struct {
	preset *Preset
}

func NewProject(m model) *Project {
	drv, err := driver.New()
	if err != nil {
		panic(err)
	}

	mu := &sync.Mutex{}
	project := &Project{
		drv: drv,
		mu:  mu,
	}

	// find elektron and assign it to in/out
	mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), "Model:Cycles") || strings.Contains(in.String(), "Model:Samples") {
			project.in = in
		}
	}
	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), "Model:Cycles") || strings.Contains(out.String(), "Model:Samples") {
			project.out = out
		}
	}
	project.in.Open()
	project.out.Open()
	wr := writer.New(project.out)
	project.wr = wr
	mu.Unlock()

	return project
}

func (p *Project) AddPattern(pattern ...*Pattern) {
	p.patterns = append(p.patterns, pattern...)
}

func (p *Project) Play() {
	// writer.NoteOn(p.wr, 64, 120)
	for _, pat := range p.patterns {
		// range through every track of currect pattern
		for j := 0; j < 6; j++ {
			// if current track (of current pattern) is not empty
			if (&Track{}) != pat.tracks[j] {

			}
		}
		// if len(pat.tracks[0]) > 0 {

		// }
		// for t := pat.tracks {

		// }
	}
}

func (p *Project) Stop() {
	for {
		p.cc(T1, CYCLESPITCH, 50)
		p.noteon(T1, E5, 126)
		time.Sleep(750 * time.Millisecond)
		p.noteoff(T1, E5)

		p.cc(T1, CYCLESPITCH, 70)
		p.noteon(T1, C4, 127)
		// p.cc(T1, MACHINE, 1)
		time.Sleep(500 * time.Millisecond)
		p.noteoff(T1, C4)

		p.cc(T1, MACHINE, uint8(rand.Intn(5)))
		p.cc(T1, CYCLESPITCH, 80)
		p.noteon(T1, F4, 127)
		// p.cc(T1, MACHINE, 2)
		time.Sleep(500 * time.Millisecond)
		p.noteoff(T1, F4)

		p.cc(T1, MACHINE, 0)
		// p.cc(T1, CYCLESPITCH, 90)
		p.cc(T1, CYCLESPITCH, uint8(rand.Intn(126)))
		p.cc(T1, DECAY, uint8(rand.Intn(126)))
		p.cc(T1, COLOR, uint8(rand.Intn(126)))
		p.cc(T1, SHAPE, uint8(rand.Intn(126)))
		p.cc(T1, SWEEP, uint8(rand.Intn(126)))
		p.cc(T1, CONTOUR, uint8(rand.Intn(126)))
		p.noteon(T1, A4, 127)
		time.Sleep(250 * time.Millisecond)
		p.noteoff(T1, A4)

		// p.cc(T1, MACHINE, uint8(PERC))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, C4, 126)
		// time.Sleep(500 * time.Millisecond)
		// p.noteoff(T1, C4)

		// p.cc(T1, MACHINE, uint8(METAL))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, C5, 127)
		// time.Sleep(1000 * time.Millisecond)
		// p.noteoff(T1, C5)

		// p.cc(T1, MACHINE, uint8(TONE))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, F4, 126)
		// time.Sleep(1000 * time.Millisecond)
		// p.noteoff(T1, F4)

		// p.cc(T1, MACHINE, uint8(SNARE))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, G6, 126)
		// time.Sleep(500 * time.Millisecond)
		// p.noteoff(T1, G4)
	}
}

func (p *Project) Next(patternNumber ...int) {

}

func (p *Project) SetVolume() {

}

// Close midi connection.
func (p *Project) Close() {
	p.out.Close()
}

func (p *Project) noteon(t track, n notes, vel uint8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOn(p.wr, uint8(n), vel)
	p.mu.Unlock()
}

func (p *Project) noteoff(t track, n notes) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOff(p.wr, uint8(n))
	p.mu.Unlock()
}

func (p *Project) cc(t track, par Parameter, val uint8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ControlChange(p.wr, uint8(par), val)
	p.mu.Unlock()
}

func (p *Project) pc(t track, pc uint8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ProgramChange(p.wr, pc)
	p.mu.Unlock()
}

func NewPattern(scale *Scale) *Pattern {
	var tracks [6]*Track
	// for i := range tracks {
	// 	tracks[i] = new(Track)
	// }
	pattern := &Pattern{scale, tracks}

	return pattern
}

func NewPatternFrom(pattern *Pattern) *Pattern {
	// var copy *Pattern
	// *copy = *pattern
	copy := pattern
	return copy
}

func (p *Pattern) ScaleSetup(s *Scale) {
	p.scale = s
}

func (p *Pattern) T1(t *Track) {
	p.tracks[0] = t
}

func (p *Pattern) T2(t *Track) {
	p.tracks[1] = t
}

func (p *Pattern) T3(t *Track) {
	p.tracks[2] = t
}

func (p *Pattern) T4(t *Track) {
	p.tracks[3] = t
}

func (p *Pattern) T5(t *Track) {
	p.tracks[4] = t
}

func (p *Pattern) T6(t *Track) {
	p.tracks[5] = t
}

func NewTrack() *Track {
	p := NewPreset()
	newt := &Track{preset: p}
	return newt
}

func (t *Track) SetScale(s *Scale) {
	t.scale = s
}

func (t *Track) SetPreset(p *Preset) {
	t.preset = p
}

func (t *Track) AddTrigs(trigs ...*Trig) {
	t.trigs = append(t.trigs, trigs...)
}

func NewPreset(preset ...map[Parameter]uint8) *Preset {
	if preset != nil {
		return &Preset{parameters: preset[0]}
	}
	p := make(map[Parameter]uint8)
	defaultPreset(p)
	return &Preset{parameters: p}
}

func (p *Preset) SetParameter(param Parameter, value uint8) {
	p.parameters[param] = value
}

// maybe?
func (p *Project) CopyPreset(pat *Pattern) *Pattern {
	newp := pat
	return newp
}

func NewScale(mod ScaleMode, len, scl, chg int) *Scale {
	scale := &Scale{mod, len, scl, chg}
	return scale
}

func (s *Scale) SetMod(m ScaleMode) {
	s.mod = m
}

func (s *Scale) SetLen(l int) {
	s.len = l
}

func (s *Scale) SetScl(scl int) {
	s.scl = scl
}

func (s *Scale) SetChg(c int) {
	s.chg = c
}

func NewTrig() *Trig {
	return &Trig{}
}

func (t *Trig) SetPreset(p *Preset) {
	t.preset = p
}

func (t *Trig) SetLock(l *Lock) {
	t.lock = l
}

func NewLock() *Lock {
	return &Lock{}
}

func (l *Lock) SetPreset(p *Preset) {
	l.preset = p
}

func defaultPreset(p map[Parameter]uint8) {
	// p[]
}

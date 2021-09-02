package elektronmodels

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

//
// constants
//

type model string

const (
	CYCLES  model = "Model:Cycles"
	SAMPLES model = "Model:Samples"
)

type voice int8

const (
	T1 voice = iota
	T2
	T3
	T4
	T5
	T6
)

type notes int8

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
	Df8 notes = Cs8
	Ef8 notes = Ds8
	Gf8 notes = Fs8
	Af8 notes = Gs8
	Bf8 notes = As8
)

type chords int8

const (
	Unisonx2 chords = iota
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
	SWING      Parameter = 15
	CHANCE     Parameter = 14

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

const (
	DELAYTIME Parameter = iota + 85
	DELAYFEEDBACK
	REVERBSIZE
	REVERBTONE
)

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

// ??
// const (
// 	LNONE  Parameter = 0
// 	LPITCH Parameter = 9

// 	LCOLOR Parameter = iota + 9
// 	LSHAPE
// 	LSWEEP
// 	LCONTOUR
// 	LPAW
// 	LGATE
// 	LFTUN
// 	LDECAY
// 	LDIST
// 	LDELAY
// 	LREVERB
// 	LPAN
// )

type machine int8

const (
	KICK machine = iota + 1
	SNARE
	METAL
	PERC
	TONE
	CHORD
)

type scaleMode bool

const (
	PTN scaleMode = true
	TRK scaleMode = false
)

//
// data structures
//

// Project .
type Project struct {
	model
	*sequencer
	// Free allows to bypass the sequencer and send triggers in real time.
	Free *free
}

// Sequencer .
type sequencer struct {
	// midi fields
	drv midi.Driver
	in  midi.In
	out midi.Out
	wr  *writer.Writer

	pattern map[int]*pattern

	// playtime fields
	tempo  chan float64
	chains []int
	// fillMode chan bool
	// chances chan float64
	// swing chan float64
}

type free struct {
	preset
	track voice
	note  *note
	midi  *sequencer
}

type pattern struct {
	track map[voice]*track
	scale *scale
}

type track struct {
	preset
	scale *scale
	trig  map[int]*trig
}

type scale struct {
	mod scaleMode
	len int
	scl float64
	chg int8

	tempo float64
}

type preset map[Parameter]int8

type trig struct {
	note *note
	lock preset

	// nudge float64
	// condition float64
}

type note struct {
	key      notes
	length   float64
	velocity int8
}

//
// Project
//

var mu *sync.Mutex

// NewProject initiates and returns a *Project struct.
// TODO: better documentation
func NewProject(m model) (*Project, error) {
	mu = new(sync.Mutex)

	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	sequencer := &sequencer{
		drv: drv,
	}

	// find elektron and assign it to in/out
	var helperIn, helperOut bool
	mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), string(m)) {
			sequencer.in = in
			helperIn = true
		}
	}
	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), string(m)) {
			sequencer.out = out
			helperOut = true
		}
	}
	// check if nothing found
	if !helperIn && !helperOut {
		return nil, fmt.Errorf("device %s not found", m)
	}

	err = sequencer.in.Open()
	if err != nil {
		return nil, err
	}

	err = sequencer.out.Open()
	if err != nil {
		return nil, err
	}

	wr := writer.New(sequencer.out)
	sequencer.wr = wr
	mu.Unlock()

	sequencer.pattern = make(map[int]*pattern)

	return &Project{model: m, sequencer: sequencer, Free: &free{midi: sequencer}}, nil
}

//
// sequencer
//

func (s *sequencer) Play(id ...int) {
	var (
		pattern    *pattern
		individual bool
	)

	// if user did not specify a pattern neither Chain method used, print an error
	if len(id) == 0 && len(s.chains) == 0 {
		fmt.Println("error: no pattern selected")
		return
	}

	if len(s.chains) != 0 {
		pattern = s.pattern[s.chains[0]]
	} else {
		pattern = s.pattern[id[0]]
	}

	// check playPattern exists
	if pattern == nil {
		fmt.Println("error: pattern does not exist")
		return
	}

	// check scaleMode
	for _, track := range pattern.track {
		if track.scale.mod != PTN {
			individual = true
			break
		}
	}

	// current pattern play

	var count int64

	block := make(chan bool)
	s.tempo = make(chan float64)

	tick := time.NewTicker(time.Duration(60000/60.0) * time.Millisecond)
	go func() {
	loop:
		for {
			select {
			case newTempo := <-s.tempo:
				tick.Reset(time.Duration(60000/newTempo) * time.Millisecond)
			case <-tick.C:
				if count == 2 {
					tick.Stop()
					close(s.tempo)
					// break loop
					block <- true
					break loop
				}
				log.Println(atomic.AddInt64(&count, 1))
				s.cc(T1, CYCLESPITCH, 50)
				s.noteon(T1, E5, 126)
				time.Sleep(750 * time.Millisecond)
				s.noteoff(T1, E5)
			}
		}
	}()

	<-block

	// return s

	// for {
	// 	p.cc(T1, CYCLESPITCH, 50)
	// 	p.noteon(T1, E5, 126)
	// 	time.Sleep(750 * time.Millisecond)
	// 	p.noteoff(T1, E5)

	// 	p.cc(T1, CYCLESPITCH, 70)
	// 	p.noteon(T1, C4, 127)
	// 	// p.cc(T1, MACHINE, 1)
	// 	time.Sleep(500 * time.Millisecond)
	// 	p.noteoff(T1, C4)

	// 	p.cc(T1, MACHINE, int(rand.Intn(5)))
	// 	p.cc(T1, CYCLESPITCH, 80)
	// 	p.noteon(T1, F4, 127)
	// 	// p.cc(T1, MACHINE, 2)
	// 	time.Sleep(500 * time.Millisecond)
	// 	p.noteoff(T1, F4)

	// 	p.cc(T1, MACHINE, 0)
	// 	// p.cc(T1, CYCLESPITCH, 90)
	// 	p.cc(T1, CYCLESPITCH, int(rand.Intn(126)))
	// 	p.cc(T1, DECAY, int(rand.Intn(126)))
	// 	p.cc(T1, COLOR, int(rand.Intn(126)))
	// 	p.cc(T1, SHAPE, int(rand.Intn(126)))
	// 	p.cc(T1, SWEEP, int(rand.Intn(126)))
	// 	p.cc(T1, CONTOUR, int(rand.Intn(126)))
	// 	p.noteon(T1, A4, 127)
	// 	time.Sleep(250 * time.Millisecond)
	// 	p.noteoff(T1, A4)
	// }
}

// Next 	// can be used without a number too - if used without a number and there is no next currently playing pattern keeps on looping
// if used and not found, an empty default pattern should be returned - silence
// Second number indicates jump to specific pattern number rather the next in line.
func (s *sequencer) Next(pattern int) *sequencer {
	return s
}

// Pause .
func (s *sequencer) Pause() {

}

// Stop .
func (s *sequencer) Stop() {

}

// func (s *sequencer) Tempo(bpm float64) {
//
// }

// SetVolume .
func (s *sequencer) Volume(value int8) {

}

// Close midi connection.
func (s *sequencer) Close() {
	s.in.Close()
	s.out.Close()
	s.drv.Close()
}

func (s *sequencer) noteon(t voice, n notes, vel int8) {
	mu.Lock()
	s.wr.SetChannel(uint8(t))
	writer.NoteOn(s.wr, uint8(n), uint8(vel))
	mu.Unlock()
}

func (s *sequencer) noteoff(t voice, n notes) {
	mu.Lock()
	s.wr.SetChannel(uint8(t))
	writer.NoteOff(s.wr, uint8(n))
	mu.Unlock()
}

func (s *sequencer) cc(t voice, par Parameter, val int8) {
	mu.Lock()
	s.wr.SetChannel(uint8(t))
	writer.ControlChange(s.wr, uint8(par), uint8(val))
	mu.Unlock()
}

func (s *sequencer) pc(t voice, pc int8) {
	mu.Lock()
	s.wr.SetChannel(uint8(t))
	writer.ProgramChange(s.wr, uint8(pc))
	mu.Unlock()
}

func (s *sequencer) unlockPreset() {

}

func (s *sequencer) unlockMachine() {

}

func (s *sequencer) Chain(patterns ...int) *sequencer {
	for _, pattern := range patterns {
		s.chains = append(s.chains, pattern)
	}

	return s
}

func (s *sequencer) Tempo(bpm float64) *sequencer {
	return s
}

// // CopyPattern copies the input source pattern to caller destination.
// func (s *sequencer) CopyPattern(src, dst int) *sequencer {
// 	if s.pattern[src] == nil {
// 		fmt.Println("warning: can not copy empty source pattern")
// 		return nil
// 	}

// 	dest := s.Pattern(dst)

// 	dest.scale.mod = s.pattern[src].scale.mod
// 	dest.scale.len = s.pattern[src].scale.len
// 	dest.scale.scl = s.pattern[src].scale.scl
// 	dest.scale.chg = s.pattern[src].scale.chg
// 	dest.scale.tempo = s.pattern[src].scale.tempo

// 	return s
// }

// Pattern returns the specified pattern out of project's pattern collection.
// Allows to access pattern's methods.
func (s *sequencer) Pattern(id int) *pattern {
	if s.pattern[id] == nil {
		s.pattern[id] = &pattern{
			track: make(map[voice]*track),
			scale: &scale{PTN, 15, 1.0, 0, 120.0},
		}
	}

	return s.pattern[id]
}

//
// free
//

func (f *free) Trig(id voice) {
	f.track = id
}

func (f *free) Preset(preset preset) *free {
	f.preset = preset
	return f
}

func (f *free) Note(key notes, length float64, velocity int8) *free {
	f.note.key = key
	f.note.length = length
	f.note.velocity = velocity
	return f
}

//
// pattern
//

// // CopyTrack copies a track from input source to destination.
// func (p *pattern) CopyTrack(src, dst voice) *pattern {
// 	if p.track[src] == nil {
// 		fmt.Println("warning: can not copy empty source track")
// 		return nil
// 	}

// 	p.track[dst] = p.track[src]
// 	return p
// }

func (p *pattern) Scale(mod scaleMode, length int, scale float64, chg int8) *pattern {
	p.scale.mod = mod
	p.scale.len = length
	p.scale.scl = scale
	p.scale.chg = chg
	return p
}

func (p *pattern) ScaleMod(mod scaleMode) *pattern {
	p.scale.mod = mod
	return p
}

func (p *pattern) Tempo(tempo float64) *pattern {
	p.scale.tempo = tempo
	return p
}

func (p *pattern) Track(id voice) *track {
	if p.track[id] == nil {
		p.track[id] = &track{
			scale:  &scale{mod: PTN, len: 15, scl: 1.0, chg: 0},
			preset: defaultPreset(id),
			trig:   make(map[int]*trig),
		}
	}

	return p.track[id]
}

//
// Track
//

// SetScale sets a new scale for the track.
// If not set a default one is used.
func (t *track) Scale(length int, scl float64, chg int8) *track {
	t.scale.len = length
	t.scale.scl = scl
	t.scale.chg = chg
	return t
}

func (t *track) Preset(p preset) *track {
	t.preset = p
	return t
}

// SetParameter assigned a parameter to the preset.
// First argument is a Parameter type and second value an int8.
func (t *track) Parameter(parameter Parameter, value int8) *track {
	mu.Lock()
	t.preset[parameter] = value
	mu.Unlock()
	return t
}

// SetParameter assigned a parameter to the preset.
// First argument is a Parameter type and second value an int8.
func (t *track) DelParameter(parameter Parameter) *track {
	mu.Lock()
	delete(t.preset, parameter)
	mu.Unlock()
	return t
}

func (t *track) Tempo(tempo float64) *track {
	t.scale.tempo = tempo
	return t
}

func (t *track) Trig(id int) *trig {
	mu.Lock()
	t.trig[id] = &trig{&note{C4, 4, 126}, make(map[Parameter]int8)}
	mu.Unlock()

	return t.trig[id]
}

//
// scale
//

// SetMod Mode can be set to either PTN (pattern) or TRK (track). In PTN mode all tracks share the same
// SCALE and LENGTH settings. In TRK mode, all tracks can have individual SCALE and LENGTH settings.
func (s *scale) Mod(mod scaleMode) *scale {
	s.mod = mod
	return s
}

// SetLen sets the step length (amount of steps) of the pattern/track.
func (s *scale) Len(length int) *scale {
	s.len = length
	return s
}

// SetScl controls the speed the playback in multiples of the current tempo. It offers seven possible
// settings, 1/8X, 1/4X, 1/2X, 3/4X, 1X, 3/2X and 2X. A setting of 1/8X plays back the pattern at one-eighth of
// the set tempo. 3/4X plays the pattern back at three-quarters of the tempo; 3/2X plays back the pattern
// twice as fast as the 3/4X setting. 2X makes the pattern play at twice the BPM.
func (s *scale) Scl(scl float64) *scale {
	s.scl = scl
	return s
}

// SetChg controls for how long the active pattern plays before it loops or a cued (the next selected) pattern begins to play. If CHG is set to 64, the pattern behaves like a pattern consisting of 64 steps
// regarding cueing and chaining. If CHG is set to OFF, the default change length is INF (infinite) in TRACK
// mode and the same value as LEN in PATTERN mode.
func (s *scale) Chg(chg int8) *scale {
	s.chg = chg
	return s
}

// InitTrig initiates a trigger note and places it to designated position of trigs map (map[int]*trig).
// All triggers need to be initiated first so the appropriate memeroy allocation takes place.
// If you do not init your trigs you will get panic: runtime error.

//
// preset
//

//
// trig
//

func (t *trig) Lock(preset preset) *trig {
	mu.Lock()
	t.lock = preset
	mu.Unlock()
	return t
}

// SetMachine .
func (t *trig) LockMachine(m machine) *trig {
	// l.lockMachine = lockMachine
	// l.Machine = m
	return t
}

// SetNote .
func (t *trig) Note(key notes, length float64, velocity int8) {
	t.note.key = key
	t.note.length = length
	t.note.velocity = velocity
}

// // CopyTrig .
// func (t *trig) CopyTrig(src *trig) {
// 	*t = *src
// }

// func (t *trig) ClearTrig(src *trig) {
// 	*t = *src
// }

//
// note
//

// SetKey .
func (n *note) Key(key notes) {
	n.key = key
}

// SetLength Trig Length sets the duration of the notes. When a note has finished playing a NOTE OFF command
// is sent. The INF setting equals infinite note length. This parameter only applies if GATE is set to ON or
// when sending trig length data over MIDI. (0.125–128, INF)
func (n *note) Length(length float64) {
	n.length = length
}

// SetVelocity .
func (n *note) Velocity(velocity int8) {
	n.velocity = velocity
}

// // CopyNote .
// func (n *note) CopyNote(src *note) {
// 	*n = *src
// }

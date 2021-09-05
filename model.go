package elektronmodels

import (
	"fmt"
	"strings"
	"sync"
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

type machine int8

const (
	KICK machine = iota
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

// Project long description of the data structure, methods, behaviors and useage.
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
	pause  chan bool
	chains []int
	// fillMode chan bool
	// chances chan float64
	// swing chan float64
}

type free struct {
	midi *sequencer
}

type pattern struct {
	track map[voice]*track
	scale *scale
	tempo float64
}

type track struct {
	preset
	scale *scale
	trig  map[int]*trig
}

type scale struct {
	length int
	scale  float64
	change int8
}

type preset map[Parameter]int8

type trig struct {
	note *note
	lock preset

	scale *scale
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
		drv:   drv,
		tempo: make(chan float64),
	}

	// find elektron and assign it to in/out
	var helperIn, helperOut bool
	// mu.Lock()
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
	// mu.Unlock()

	sequencer.pattern = make(map[int]*pattern)
	sequencer.tempo = make(chan float64)

	return &Project{
		model:     m,
		sequencer: sequencer,
		Free: &free{
			midi: sequencer,
		}}, nil
}

//
// sequencer
//

// Play
// Is blocking
func (s *sequencer) Play(ids ...int) {
	// if user did not specify a pattern neither Chain method used, print an error
	if len(ids) == 0 && len(s.chains) == 0 {
		fmt.Println("error: no pattern selected")
		return
	}

	var pattern *pattern
	var id int = ids[0]

	if len(s.chains) != 0 {
		pattern = s.pattern[s.chains[0]]
	} else {
		pattern = s.pattern[id]
	}

	// check pattern exists
	if pattern == nil {
		fmt.Println("error: pattern does not exist")
		return
	}

	// check tracks for scale settings

	// check tempo

	// var count int

	for i := 0; i <= 5; i++ {
		voice := voice(i)
		if track, ok := s.pattern[id].track[voice]; ok {
			tick := time.NewTicker(
				time.Duration(60000/(pattern.tempo*pattern.scale.scale)) * time.Millisecond)

			go func() {
				var count int

				go func() {
					// check if track has preset
					// if not set default preset for track
					if len(track.preset) == 0 {
						track.preset = defaultPreset(voice)
					}

					// apply preset
					for parameter, value := range track.preset {
						s.cc(voice, parameter, value)
					}
				}()

			loop:
				for {
					select {
					case newTempo := <-s.tempo:
						tick.Reset(time.Duration(60000/(newTempo*pattern.scale.scale)) * time.Millisecond)
					case <-tick.C:
						if count > s.pattern[id].scale.length {
							count = 0
						}

						if trig, ok := track.trig[count]; ok {
							s.noteon(voice,
								trig.note.key,
								trig.note.velocity)
							go func() {
								time.Sleep(time.Millisecond * time.Duration(trig.note.length))
								s.noteoff(voice, trig.note.key)
							}()
							// fmt.Println("----new----")
							// fmt.Println("track: ", voice)
							// fmt.Println("trig: ", count)
							// fmt.Println()
						}

						count++

					case <-s.pause:
						break loop
					}
				}
			}()
		}
	}

	block := make(chan bool)
	<-block
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

// SetVolume .
func (s *sequencer) Volume(value int8) {

}

func (s *sequencer) Chain(patterns ...int) *sequencer {
	for _, pattern := range patterns {
		s.chains = append(s.chains, pattern)
	}

	return s
}

// Pattern returns the specified pattern out of project's pattern collection.
// Allows to access pattern's methods.
func (s *sequencer) Pattern(id int) *pattern {
	if _, ok := s.pattern[id]; !ok {
		s.pattern[id] = &pattern{
			track: make(map[voice]*track),
			scale: &scale{15, 1.0, 15},
		}
	}

	return s.pattern[id]
}

// Close midi connection. Use it with defer after creating a new project.
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

//
// free
//

func (f *free) Preset(track voice, preset preset) {
	for parameter, value := range preset {
		f.midi.cc(track, parameter, value)
	}
}

// Note
// duration in milliseconds (ms)
func (f *free) Note(track voice, note notes, velocity int8, duration float64) {
	f.midi.noteon(track, note, velocity)
	go func() {
		time.Sleep(time.Millisecond * time.Duration(duration))
		f.midi.noteoff(track, note)
	}()
}

func (f *free) CC(track voice, parameter Parameter, value int8) {
	f.midi.cc(track, parameter, value)
}

func (f *free) PC(t voice, pc int8) {
	f.midi.pc(t, pc)
}

//
// pattern
//

// Scale sets the scale for the pattern.
// If scaleMode is set to track TRK the provided scale settings are used as default to the rest of the tracks.
// This mimics synth's own functionality.
func (p *pattern) Scale(length int, scale float64, chg int8) *pattern {
	p.scale.length = length
	p.scale.scale = scale
	p.scale.change = chg
	return p
}

func (p *pattern) Tempo(tempo float64) *pattern {
	p.tempo = tempo
	return p
}

func (p *pattern) Track(id voice) *track {
	if _, ok := p.track[id]; !ok {
		p.track[id] = &track{
			// scale:  &scale{length: 15, scale: 1.0},
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
func (t *track) Scale(length int, factor float64) *track {
	t.scale = &scale{length: length, scale: factor}
	return t
}

func (t *track) Preset(p preset) *track {
	t.preset = p
	return t
}

// SetParameter assigned a parameter to the preset.
// First argument is a Parameter type and second value an int8.
func (t *track) Parameter(parameter Parameter, value int8) *track {
	// mu.Lock()
	t.preset[parameter] = value
	// mu.Unlock()
	return t
}

func (t *track) Trig(id int) *trig {
	if _, ok := t.trig[id]; !ok {
		t.trig[id] = &trig{note: &note{C4, 4, 126}}
	}

	return t.trig[id]
}

//
// scale
//

// Length sets the step length (amount of steps) of the pattern/track.
func (s *scale) Length(length int) *scale {
	s.length = length
	return s
}

// Scale controls the speed the playback in multiples of the current tempo. It offers seven possible
// settings, 1/8X, 1/4X, 1/2X, 3/4X, 1X, 3/2X and 2X. A setting of 1/8X plays back the pattern at one-eighth of
// the set tempo. 3/4X plays the pattern back at three-quarters of the tempo; 3/2X plays back the pattern
// twice as fast as the 3/4X setting. 2X makes the pattern play at twice the BPM.
func (s *scale) Scale(scl float64) *scale {
	s.scale = scl
	return s
}

// Change controls for how long the active pattern plays before it loops or a cued (the next selected) pattern begins to play. If CHG is set to 64, the pattern behaves like a pattern consisting of 64 steps
// regarding cueing and chaining. If CHG is set to OFF, the default change length is INF (infinite) in TRACK
// mode and the same value as LEN in PATTERN mode.
func (s *scale) Change(chg int8) *scale {
	s.change = chg
	return s
}

//
// trig
//

func (t *trig) Lock(preset preset) *trig {
	t.lock = preset
	return t
}

// Note .
func (t *trig) Note(key notes, length float64, velocity int8) {
	t.note.key = key
	t.note.length = length
	t.note.velocity = velocity
}

func (t *trig) Scale(factor float64) *trig {
	t.scale = &scale{scale: factor}
	return t
}

//
// note
//

// Key .
func (n *note) Key(key notes) {
	n.key = key
}

// Length Trig Length sets the duration of the notes. When a note has finished playing a NOTE OFF command
// is sent. The INF setting equals infinite note length. This parameter only applies if GATE is set to ON or
// when sending trig length data over MIDI. (0.125–128, INF)
func (n *note) Length(length float64) {
	n.length = length
}

// Velocity .
func (n *note) Velocity(velocity int8) {
	n.velocity = velocity
}

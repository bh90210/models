package elektronmodels

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
)

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

type parameter int8

const (
	NOTE       parameter = 3
	TRACKLEVEL parameter = 17
	MUTE       parameter = 94
	PAN        parameter = 10
	SWEEP      parameter = 18
	CONTOUR    parameter = 19
	DELAY      parameter = 12
	REVERB     parameter = 13
	VOLUMEDIST parameter = 7
	SWING      parameter = 15
	CHANCE     parameter = 14

	// model:cycles
	MACHINE     parameter = 64
	CYCLESPITCH parameter = 65
	DECAY       parameter = 80
	COLOR       parameter = 16
	SHAPE       parameter = 17
	PUNCH       parameter = 66
	GATE        parameter = 67

	// model:samples
	PITCH        parameter = 16
	SAMPLESTART  parameter = 19
	SAMPLELENGTH parameter = 20
	CUTOFF       parameter = 74
	RESONANCE    parameter = 71
	LOOP         parameter = 17
	REVERSE      parameter = 18
)

const (
	DELAYTIME parameter = iota + 85
	DELAYFEEDBACK
	REVERBSIZE
	REVERBTONE
)

const (
	LFOSPEED parameter = iota + 102
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
// 	LNONE  parameter = 0
// 	LPITCH parameter = 9

// 	LCOLOR parameter = iota + 9
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

// Project .
type Project struct {
	// Pattern map[int]*pattern
	model
	sequencer
}

// Sequencer .
type sequencer struct {
	fillMode bool

	// midi fields
	drv midi.Driver
	in  midi.In
	out midi.Out
	wr  *writer.Writer

	pattern map[int]*pattern

	// playtime fields
	tempo          chan float64
	patternLength  int
	patternRunning int
}

type pattern struct {
	track map[voice]*track
	// T1     *track
	// T2     *track
	// T3     *track
	// T4     *track
	// T5     *track
	// T6     *track
	// tempo float64
}

type track struct {
	*scale
	preset
	trig map[int]*trig
}

type scale struct {
	mod scaleMode
	len int
	scl float64
	chg int8

	tempo float64
}

type preset map[parameter]int8

// type preset struct {
// 	parameter map[Parameter]int8
// }

type trig struct {
	*note
	*lock
}

type note struct {
	key      notes
	length   float64
	velocity int8
}

type lock struct {
	// conditional *Condition
	// parameter, preset, machine preset
	// machine           Machine
	preset
}

var mu *sync.Mutex

// NewProject initiates and returns a *Project struct.
// TODO: better documentation
func NewProject(m model) *Project {

	// func (p *Project) initMidi() error {
	// 	drv, err := driver.New()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	sequencer := &sequencer{
	// 		drv: drv,
	// 		// mu:  p.mu,
	// 	}

	// 	// find elektron and assign it to in/out
	// 	var helperIn, helperOut bool
	// 	mu.Lock()
	// 	ins, _ := drv.Ins()
	// 	for _, in := range ins {
	// 		if strings.Contains(in.String(), string(p.model)) {
	// 			sequencer.in = in
	// 			helperIn = true
	// 		}
	// 	}
	// 	outs, _ := drv.Outs()
	// 	for _, out := range outs {
	// 		if strings.Contains(out.String(), string(p.model)) {
	// 			sequencer.out = out
	// 			helperOut = true
	// 		}
	// 	}
	// 	// check if nothing found
	// 	if !helperIn && !helperOut {
	// 		return nil, fmt.Errorf("device %s not found", p.model)
	// 	}

	// 	err = sequencer.in.Open()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	err = sequencer.out.Open()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	wr := writer.New(sequencer.out)
	// 	sequencer.wr = wr
	// 	mu.Unlock()

	// 	return sequencer, nil
	// }
	mu = new(sync.Mutex)
	return &Project{model: m, pattern: make(map[int]*pattern)}
}

// InitPattern initiates a new pattern for the selected position.
// The equivalent of storing a pattern on ie. T1 trig 1.
// func (p *Project) InitPattern(position int) {
// 	mu.Lock()
// 	p.Pattern[position] = &pattern{
// 		T1: &track{
// 			Scale:  &scale{PTN, 15, 4.0, 0},
// 			Preset: defaultT1(),
// 			Trig:   make(map[int]*trig),
// 		},
// 		T2: &track{Scale: &scale{PTN, 15, 4.0, 0}, Preset: defaultT2(), Trig: make(map[int]*trig)},
// 		T3: &track{Scale: &scale{PTN, 15, 4.0, 0}, Preset: defaultT3(), Trig: make(map[int]*trig)},
// 		T4: &track{Scale: &scale{PTN, 15, 4.0, 0}, Preset: defaultT4(), Trig: make(map[int]*trig)},
// 		T5: &track{Scale: &scale{PTN, 15, 4.0, 0}, Preset: defaultT5(), Trig: make(map[int]*trig)},
// 		T6: &track{Scale: &scale{PTN, 15, 4.0, 0}, Preset: defaultT6(), Trig: make(map[int]*trig)},
// 	}
// 	mu.Unlock()
// }

func (p *Project) Sequencer() *sequencer {
	return &p.sequencer
}

// CopyPattern copies the input source pattern to caller destination.
func (s *sequencer) CopyPattern(src, dst int) *sequencer {
	s.pattern[dst] = s.pattern[src]
	return s
}

// Pattern returns the specified pattern out of project's pattern collection.
// Allows to access pattern's methods.
func (s *sequencer) Pattern(pos int) *pattern {
	return s.pattern[pos]
}

// SetTempo .
func (p *pattern) SetTempo(tempo float64) *pattern {
	p.tempo = tempo
	return p
}

// CopyTrack copies a track from input source to caller destination.
func (p *pattern) CopyTrack(src, dst voice) *pattern {
	p.track[dst] = p.track[src]
	return p
}

func (p *pattern) Track(id voice) *track {
	return p.track[id]
}

// SetScale sets a new scale for the track.
// If not set a default one is used.
func (t *track) SetScale(mod scaleMode, length int, scl float64, chg int8) *track {
	t.mod = mod
	t.len = length
	t.scl = scl
	t.chg = chg
	return t
}

// InitTrig initiates a trigger note and places it to designated position of trigs map (map[int]*trig).
// All triggers need to be initiated first so the appropriate memeroy allocation takes place.
// If you do not init your trigs you will get panic: runtime error.
func (t *track) InitTrig(position int) {
	mu.Lock()
	t.Trig[position] = &trig{&note{C4, 4, 126}, &lock{}}
	mu.Unlock()
}

func (t *track) SetPreset(preset preset) {
	mu.Lock()
	// t.Trig[position] = &trig{&note{C4, 4, 126}, &lock{}}
	mu.Unlock()
}

// SetMod Mode can be set to either PTN (pattern) or TRK (track). In PTN mode all tracks share the same
// SCALE and LENGTH settings. In TRK mode, all tracks can have individual SCALE and LENGTH settings.
func (s *scale) SetMod(mod scaleMode) {
	s.mod = mod
}

// SetLen sets the step length (amount of steps) of the pattern/track.
func (s *scale) SetLen(length int) {
	s.len = length
}

// SetScl controls the speed the playback in multiples of the current tempo. It offers seven possible
// settings, 1/8X, 1/4X, 1/2X, 3/4X, 1X, 3/2X and 2X. A setting of 1/8X plays back the pattern at one-eighth of
// the set tempo. 3/4X plays the pattern back at three-quarters of the tempo; 3/2X plays back the pattern
// twice as fast as the 3/4X setting. 2X makes the pattern play at twice the BPM.
func (s *scale) SetScl(scl float64) {
	s.scl = scl
}

// SetChg controls for how long the active pattern plays before it loops or a cued (the next selected) pattern begins to play. If CHG is set to 64, the pattern behaves like a pattern consisting of 64 steps
// regarding cueing and chaining. If CHG is set to OFF, the default change length is INF (infinite) in TRACK
// mode and the same value as LEN in PATTERN mode.
func (s *scale) SetChg(chg int8) {
	s.chg = chg
}

// SetParameter assigned a parameter to the preset.
// First argument is a Parameter type and second value an int8.
func (p *track) SetParameter(parameter parameter, value int8) {
	mu.Lock()
	p.preset[parameter] = value
	mu.Unlock()
}

// SetParameter assigned a parameter to the preset.
// First argument is a Parameter type and second value an int8.
func (p *track) DelParameter(parameter parameter) {
	mu.Lock()
	delete(p.preset, parameter)
	mu.Unlock()
}

// SetNote .
func (t *trig) SetNote(key notes, length float64, velocity int8) {
	t.key = key
	t.length = length
	t.velocity = velocity
}

// CopyTrig .
func (t *trig) CopyTrig(src *trig) {
	*t = *src
}

func (t *trig) ClearTrig(src *trig) {
	*t = *src
}

// SetKey .
func (n *note) SetKey(key notes) {
	n.key = key
}

// SetLength Trig Length sets the duration of the notes. When a note has finished playing a NOTE OFF command
// is sent. The INF setting equals infinite note length. This parameter only applies if GATE is set to ON or
// when sending trig length data over MIDI. (0.125–128, INF)
func (n *note) SetLength(length float64) {
	n.length = length
}

// SetVelocity .
func (n *note) SetVelocity(velocity int8) {
	n.velocity = velocity
}

// CopyNote .
func (n *note) CopyNote(src *note) {
	*n = *src
}

// SetMachine .
func (l *lock) SetMachine(m machine) {
	// l.lockMachine = lockMachine
	// l.Machine = m
}

func (l *lock) setMachine(m machine) {
	// var tt lockMachine
	// tt = m
	// l = m
	// l.Machine = m
}

func (p *sequencer) Play(position ...int) error {
	// check for errors in current pattern

	// check for warnings of the existing and incoming patterns

	// analyze scale

	// current pattern play

	var count int64

	block := make(chan bool)
	p.tempo = make(chan float64)

	tick := time.NewTicker(time.Duration(60000/60.0) * time.Millisecond)
	go func() {
	loop:
		for {
			select {
			case newTempo := <-p.tempo:
				tick.Reset(time.Duration(60000/newTempo) * time.Millisecond)
			case <-tick.C:
				if count == 20 {
					tick.Stop()
					close(p.tempo)
					// break loop
					block <- true
					break loop
				}
				log.Println(atomic.AddInt64(&count, 1))
			}
		}
	}()

	time.Sleep(10 * time.Second)
	p.tempo <- 120.5

	<-block

	return nil

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
func (p *sequencer) Next(patternNumber ...int) {

}

// Pause .
func (p *sequencer) Pause() {

}

// Stop .
func (p *sequencer) Stop() {

}

func (p *sequencer) Tempo(bpm float64) {

}

// SetVolume .
func (p *sequencer) Volume(value int8) {

}

// Close midi connection.
func (p *sequencer) Close() {
	p.out.Close()
}

func (p *sequencer) noteon(t voice, n notes, vel int8) {
	mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOn(p.wr, uint8(n), uint8(vel))
	mu.Unlock()
}

func (p *sequencer) noteoff(t voice, n notes) {
	mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOff(p.wr, uint8(n))
	mu.Unlock()
}

func (p *sequencer) cc(t voice, par parameter, val int8) {
	mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ControlChange(p.wr, uint8(par), uint8(val))
	mu.Unlock()
}

func (p *sequencer) pc(t voice, pc int8) {
	mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ProgramChange(p.wr, uint8(pc))
	mu.Unlock()
}

func (p *sequencer) unlockPreset() {

}

func (p *sequencer) unlockMachine() {

}

// default presets for voices 1-6
func defaultT1() preset {
	p := make(map[parameter]int8)
	p[COLOR] = 10
	return p
}

func defaultT2() preset {
	d := make(map[parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT3() preset {
	d := make(map[parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT4() preset {
	d := make(map[parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT5() preset {
	d := make(map[parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT6() preset {
	p := make(preset)
	p[COLOR] = 10
	return p
}

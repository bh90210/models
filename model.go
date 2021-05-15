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

type Chords int8

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

// type note struct {
// 	dur *time.Duration
// 	key int
// }

type Parameter int8

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

type machine int8

// Machine section
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
	Pattern map[int]*pattern

	// midi fields
	drv midi.Driver
	mu  *sync.Mutex
	in  midi.In
	out midi.Out
	wr  *writer.Writer

	// playtime fields
	patternLength  int
	patternRunning int
	clock          chan int64
}

type pattern struct {
	T1 *track
	T2 *track
	T3 *track
	T4 *track
	T5 *track
	T6 *track
}

type track struct {
	Scale  *scale
	Preset Preset
	Trig   map[int]*trig
}

type scale struct {
	// Cycle manual '9.11 Scale Menu'.
	mod scaleMode
	len int
	scl float64
	chg int8
}

// Preset .
type Preset map[Parameter]int8

type trig struct {
	Note *note
	Lock *Lock
}

type note struct {
	key      notes
	length   float64
	velocity int8
}

// Lock .
type Lock struct {
	// conditional *Condition
	Preset  Preset
	Machine machine
}

// NewProject initiates and returns a *Project struct.
// TODO: better documentation
func NewProject(m model) (*Project, error) {
	mu := &sync.Mutex{}
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	project := &Project{
		drv:     drv,
		mu:      mu,
		Pattern: make(map[int]*pattern),
	}

	// find elektron and assign it to in/out
	var helperIn, helperOut bool
	mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), string(m)) {
			project.in = in
			helperIn = true
		}
	}
	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), string(m)) {
			project.out = out
			helperOut = true
		}
	}
	// check if nothing found
	if !helperIn && !helperOut {
		return nil, fmt.Errorf("device %s not found", m)
	}

	err = project.in.Open()
	if err != nil {
		return nil, err
	}

	err = project.out.Open()
	if err != nil {
		return nil, err
	}

	wr := writer.New(project.out)
	project.wr = wr
	mu.Unlock()

	return project, nil
}

// InitPattern initiates a new pattern for the selected position.
// The equivalent of storing a pattern on ie. T1 trig 1.
func (p *Project) InitPattern(position int) error {
	p.Pattern[position] = &pattern{
		T1: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT1(), Trig: make(map[int]*trig)},
		T2: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT2(), Trig: make(map[int]*trig)},
		T3: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT3(), Trig: make(map[int]*trig)},
		T4: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT4(), Trig: make(map[int]*trig)},
		T5: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT5(), Trig: make(map[int]*trig)},
		T6: &track{Scale: &scale{PTN, 15, 4, 0}, Preset: defaultT6(), Trig: make(map[int]*trig)},
	}

	return nil
}

func (p *Project) Play() error {
	// check for errors in current pattern

	// check for warnings of the existing and incoming patterns

	// analyze scale

	// current pattern play

	var count int64

	block := make(chan bool)
	tempo := make(chan float64)
	tick := time.NewTicker(time.Duration(60000/60.0) * time.Millisecond)
	go func() {
	loop:
		for {
			select {
			case newTempo := <-tempo:
				tick.Reset(time.Duration(60000/newTempo) * time.Millisecond)
			case <-tick.C:
				if count == 20 {
					tick.Stop()
					close(tempo)
					// break loop
					block <- true
					break loop
				}
				log.Println(atomic.AddInt64(&count, 1))
			}
		}
	}()

	time.Sleep(10 * time.Second)
	tempo <- 120.5

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

// Pause .
func (p *Project) Pause() {

}

// Stop .
func (p *Project) Stop() {

}

// Next .
func (p *Project) Next(patternNumber ...int) {

}

// SetVolume .
func (p *Project) SetVolume() {

}

// Close midi connection.
func (p *Project) Close() {
	p.out.Close()
}

func (p *Project) noteon(t voice, n notes, vel int8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOn(p.wr, uint8(n), uint8(vel))
	p.mu.Unlock()
}

func (p *Project) noteoff(t voice, n notes) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOff(p.wr, uint8(n))
	p.mu.Unlock()
}

func (p *Project) cc(t voice, par Parameter, val int8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ControlChange(p.wr, uint8(par), uint8(val))
	p.mu.Unlock()
}

func (p *Project) pc(t voice, pc int8) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ProgramChange(p.wr, uint8(pc))
	p.mu.Unlock()
}

func (p *Project) unlockPreset() {

}

func (p *Project) unlockMachine() {

}

// CopyPattern copies the input source pattern to caller destination.
func (p *pattern) CopyPattern(src *pattern) {
	*p = *src
}

// SetScale sets a new scale for the track.
// If not set a default one is used.
func (t *track) SetScale(mod scaleMode, length int, scl float64, chg int8) {
	t.Scale.mod = mod
	t.Scale.len = length
	t.Scale.scl = scl
	t.Scale.chg = chg
}

// CopyTrack copies a track from input source to caller destination.
func (t *track) CopyTrack(src *track) {
	*t = *src
}

// InitTrig initiates a trigger note and places it to designated position of trigs map (map[int]*trig).
// All triggers need to be initiated first so the appropriate memeroy allocation takes place.
// If you do not init your trigs you will get panic: runtime error.
func (t *track) InitTrig(position int) {
	t.Trig[position] = &trig{&note{C4, 4, 126}, &Lock{}}
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

// preset

// ??
// func (p Preset) CopyPreset(pat *pattern) *pattern {
// 	return pat
// }

// SetNote .
func (t *trig) SetNote(key notes, length float64, velocity int8) {
	t.Note.key = key
	t.Note.length = length
	t.Note.velocity = velocity
}

// SetLock .
func (t *trig) SetLock(l *Lock) {
	t.Lock = l
}

// CopyTrig .
func (t *trig) CopyTrig(src *trig) {
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
func (l *Lock) SetMachine(m machine) {
	l.Machine = m
}

// default presets for voices 1-6
func defaultT1() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT2() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT3() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT4() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT5() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func defaultT6() Preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

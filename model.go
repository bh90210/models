package elektronmodels

import (
	"log"
	"math/rand"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

// type Model interface {
// 	Play(Project)
// }

// type synth func()

// func (s synth) Play() {
// 	s()
// }

type model int

const (
	CYCLES model = iota
	SAMPLES
)

type track int

const (
	T1 track = iota
	T2
	T3
	T4
	T5
	T6
)

type notes int

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

type Chords int

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
	key int
}

type Parameter int

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

type Machine int

// Machine section
const (
	KICK Machine = iota + 1
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

// type Length int

// const (
// 	14 Length = iota
// )

// Project .
type Project struct {
	patterns []*Pattern

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

// Pattern .
type Pattern struct {
	tracks [6]*Track
}

// Track .
type Track struct {
	id     track
	scale  *Scale
	preset Preset
	trigs  []*Trig
}

// Scale .
type Scale struct {
	// Cycle manual '9.11 Scale Menu'.
	// If true Scale Mode is set to PATTERN
	// if false to TRACK.
	// 	MOD Mode can be set to either PATTERN or TRACK. In PATTERN mode all tracks share the same
	// SCALE and LENGTH settings. In TRACK mode, all tracks can have individual SCALE and LENGTH settings. Press [T1–6] to select the track to set the scale for.
	mod ScaleMode
	// LEN Length sets the step length (amount of steps) of the pattern/track.
	len int
	// 	SCL Scale controls the speed the playback in multiples of the current tempo. It offers seven possible
	// settings, 1/8X, 1/4X, 1/2X, 3/4X, 1X, 3/2X and 2X. A setting of 1/8X plays back the pattern at one-eighth of
	// the set tempo. 3/4X plays the pattern back at three-quarters of the tempo; 3/2X plays back the pattern
	// twice as fast as the 3/4X setting. 2X makes the pattern play at twice the BPM.
	scl int
	// 	CHG Change controls for how long the active pattern plays before it loops or a cued (the next selected) pattern begins to play. If CHG is set to 64, the pattern behaves like a pattern consisting of 64 steps
	// regarding cueing and chaining. If CHG is set to OFF, the default change length is INF (infinite) in TRACK
	// mode and the same value as LEN in PATTERN mode.
	chg int
}

// Preset .
type Preset map[Parameter]int

// Trig .
type Trig struct {
	note *Note
	lock *Lock
}

// Note .
type Note struct {
	key notes
	// 0.125–128, INF
	length   int
	velocity int
}

// Lock .
type Lock struct {
	// conditional *Condition
	preset  *Preset
	machine *Machine
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

func (p *Project) Play() error {
	// check for errors in current pattern

	// check for warnings of the existing and incoming patterns

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
}

func clock() {

}

func (p *Project) playTrack(t *track) {
	// set track preset
	p.setPreset(p.patterns[0].tracks[*t].preset)

	// play trigs
	for i, trig := range p.patterns[0].tracks[*t].trigs {
		<-p.clock

		// check for machine lock for next trig
		if *p.patterns[0].tracks[*t].trigs[i+1].lock.machine != 0 {
			// m := (*trig.lock.preset)[MACHINE]
			defer p.unlockMachine()
		}

		// check for preset lock
		if len(*trig.lock.preset) != 0 {
			for k, v := range *trig.lock.preset {
				p.cc(*t, k, v)
			}
			defer p.unlockPreset()
		}

		// play note
		p.noteon(*t, trig.note.key, trig.note.velocity)
		time.Sleep(time.Duration(trig.note.length))
		p.noteoff(*t, trig.note.key)
	}
}

func (p *Project) setPreset(Preset) {

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

		p.cc(T1, MACHINE, int(rand.Intn(5)))
		p.cc(T1, CYCLESPITCH, 80)
		p.noteon(T1, F4, 127)
		// p.cc(T1, MACHINE, 2)
		time.Sleep(500 * time.Millisecond)
		p.noteoff(T1, F4)

		p.cc(T1, MACHINE, 0)
		// p.cc(T1, CYCLESPITCH, 90)
		p.cc(T1, CYCLESPITCH, int(rand.Intn(126)))
		p.cc(T1, DECAY, int(rand.Intn(126)))
		p.cc(T1, COLOR, int(rand.Intn(126)))
		p.cc(T1, SHAPE, int(rand.Intn(126)))
		p.cc(T1, SWEEP, int(rand.Intn(126)))
		p.cc(T1, CONTOUR, int(rand.Intn(126)))
		p.noteon(T1, A4, 127)
		time.Sleep(250 * time.Millisecond)
		p.noteoff(T1, A4)

		// p.cc(T1, MACHINE, int(PERC))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, C4, 126)
		// time.Sleep(500 * time.Millisecond)
		// p.noteoff(T1, C4)

		// p.cc(T1, MACHINE, int(METAL))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, C5, 127)
		// time.Sleep(1000 * time.Millisecond)
		// p.noteoff(T1, C5)

		// p.cc(T1, MACHINE, int(TONE))
		// p.cc(T1, CYCLESPITCH, 0)
		// p.noteon(T1, F4, 126)
		// time.Sleep(1000 * time.Millisecond)
		// p.noteoff(T1, F4)

		// p.cc(T1, MACHINE, int(SNARE))
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

func (p *Project) noteon(t track, n notes, vel int) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOn(p.wr, uint8(n), uint8(vel))
	p.mu.Unlock()
}

func (p *Project) noteoff(t track, n notes) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.NoteOff(p.wr, uint8(n))
	p.mu.Unlock()
}

func (p *Project) cc(t track, par Parameter, val int) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ControlChange(p.wr, uint8(par), uint8(val))
	p.mu.Unlock()
}

func (p *Project) pc(t track, pc int) {
	p.mu.Lock()
	p.wr.SetChannel(uint8(t))
	writer.ProgramChange(p.wr, uint8(pc))
	p.mu.Unlock()
}

func (p *Project) unlockPreset() {

}

func (p *Project) unlockMachine() {

}

func NewPattern() (newPattern *Pattern) {
	return
}

func NewPatternFrom(pattern *Pattern) (newPattern *Pattern) {
	return pattern
}

// func (p *Pattern) ScaleSetup(s *Scale) {
// 	p.scale = s
// }

func NewTrack(id track) (newTrack *Track) {
	newTrack.id = id
	return
}

// SetScale sets a new scale for the track.
// If not set a default one is used.
func (t *Track) SetScale(s *Scale) {
	t.scale = s
}

// SetPreset sets a new scale for the track.
// If not set a default one is used.
func (t *Track) SetPreset(p Preset) {
	t.preset = p
}

func (t *Track) SetTrackID(newId track) {
	t.id = newId
}

func (t *Track) CopyTrack(newId track) (newTrack *Track) {
	newTrack = t
	newTrack.id = newId
	return
}

func (t *Track) AddTrigs(trigs ...*Trig) {
	t.trigs = append(t.trigs, trigs...)
}

func NewPreset(inPreset ...map[Parameter]int) (newPreset Preset) {
	if inPreset != nil {
		newPreset = inPreset[0]
	} else {
		newPreset = make(map[Parameter]int)
	}
	return
}

func (p Preset) SetParameter(param Parameter, value int) {
	p[param] = value
}

// maybe?
func (p *Project) CopyPreset(pat *Pattern) *Pattern {
	newp := pat
	return newp
}

func NewScale(mod ScaleMode, len, scl, chg int) *Scale {
	return &Scale{mod, len, scl, chg}
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

// func (t *Trig) SetPreset(p *Preset) {
// 	t.preset = p
// }

func (t *Trig) SetLock(l *Lock) {
	t.lock = l
}

func NewLock() *Lock {
	return &Lock{}
}

func (l *Lock) SetPreset(p *Preset) {
	l.preset = p
}

func (l *Lock) SetMachine(m *Machine) {
	l.machine = m
}

func defaultPreset(p map[Parameter]int) {
	// p[]
}

package elektronmodels

import (
	"fmt"
	"sync"
	"time"

	"github.com/rakyll/portmidi"
)

type Trig struct {
	*note
	level int64
	dur   *time.Duration
	*cc
	beat     int64
	lastBeat bool
}

func NewTrig(beat int64) *Trig {
	return &Trig{
		beat: beat,
	}
}

func LastTrig(beat int64) *Trig {
	trig := Trig{
		beat:     beat,
		lastBeat: true,
	}
	return &trig
}

func (t *Trig) Note(key notes, level int64, dur time.Duration) {
	tn := &note{}
	// check track
	tn.on = t1on
	tn.off = t1off
	tn.key = key
	tn.dur = &dur

	t.note = tn
	t.level = level
}

func (t *Trig) CC(values map[Parameter]int64) {
	t.cc = &cc{
		pamVal: values,
	}
}

type Track struct {
	number track
	trig   []*Trig
}

func NewTrack(trackNumber track, variadicTracks ...*Trig) *Track {
	track := Track{
		number: trackNumber,
		trig:   variadicTracks,
	}

	return &track
}

type pattern struct {
	track []*Track
}

type Project struct {
	pattern []*pattern
	pm      *portmidi.Stream
	mu      *sync.Mutex
}

func NewProject() (*Project, error) {
	pm, mu, err := newportmidi()
	if err != nil {
		return nil, err
	}

	project := &Project{
		pm: pm,
		mu: mu,
	}
	return project, nil
}

func (p *Project) Play() error {
	fmt.Println("Analyzing")

	totalPatters := len(p.pattern)
	fmt.Println("Total patterns", totalPatters)

	for i := 0; i != totalPatters; i++ {
		tracks := len(p.pattern[i].track)
		fmt.Println("Pattern ", i, " Tracks used: ", tracks)

		for _, track := range p.pattern[i].track {
			fmt.Println("Total Trigs ", len(track.trig), " track: ", track.number, " trig: ", track.trig)
		}
	}

	return nil
}

func (c *Project) Pause() {
}

func (c *Project) Close() {
	c.pm.Close()
}

func (t *Project) NewPattern(variadicTracks ...*Track) {
	pattern := &pattern{}

	for _, track := range variadicTracks {
		pattern.track = append(pattern.track, track)
	}

	t.pattern = append(t.pattern, pattern)
}

func (c *Project) note(track *note, note notes, intensity int64) {
	timer := time.NewTimer(*track.dur)

	// note on
	c.mu.Lock()
	c.pm.WriteShort(track.on.int64(), note.int64(), intensity)
	c.mu.Unlock()

	go func() {
		<-timer.C

		// note off
		c.mu.Lock()
		c.pm.WriteShort(track.off.int64(), note.int64(), intensity)
		c.mu.Unlock()
	}()
}

func (c *Project) cc(track cctrack, ccvalues map[Parameter]int64) {
	for k, v := range ccvalues {
		c.mu.Lock()
		c.pm.WriteShort(track.int64(), k.int64(), v)
		c.mu.Unlock()
	}
}

// func (c *Cycles) pc(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
// }

func newportmidi() (*portmidi.Stream, *sync.Mutex, error) {
	out, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		return nil, nil, err
	}
	mutex := &sync.Mutex{}

	return out, mutex, nil
}

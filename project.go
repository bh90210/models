package elektronmodels

import (
	"sync"
	"time"

	"github.com/rakyll/portmidi"
)

type Project struct {
	model   *model
	Trig    []*time.Timer
	Pattern []*Pattern
}

type Pattern struct {
	Track []*Tracks
}

type Tracks struct {
	Trig []*Trigger
}

type Trigger struct {
	*Note
	Level int64
	Dur   *time.Duration
	*CC
}

type Model int

const (
	CYCLES Model = iota
	SAMLPES
)

func NewProject(model Model) (*Project, error) {
	project := &Project{}
	switch model {
	case CYCLES:
		cycles, err := newmodel()
		if err != nil {
			return nil, err
		}
		project.model = cycles
	case SAMLPES:

	}

	return project, nil
}

func (p *Project) Play(map[int64]*Pattern) error {
	return nil
}

func (c *Project) Pause() {
}

func (c *Project) Close() {
	c.model.pm.Close()
}

type model struct {
	pm *portmidi.Stream
	mu *sync.Mutex
}

func newmodel() (*model, error) {
	out, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		return nil, err
	}
	mutex := &sync.Mutex{}
	return &model{
		pm: out,
		mu: mutex,
	}, nil
}

func (p *Project) NewNote(note Notes, value int64, dur time.Duration, ccvalues ...map[Parameter]int64) *Note {
	tnote := &Note{}
	tnote.Dur = &dur
	tnote.CC.PamVal = ccvalues[0]
	return tnote
}

// note .
func (c *Project) note(track *Note, note Notes, intensity int64, ccvalues map[Parameter]int64) {
	switch track.ON {
	case T1ON:
		c.cc(CCT1, ccvalues)
	case T2ON:
		c.cc(CCT2, ccvalues)
	case T3ON:
		c.cc(CCT3, ccvalues)
	case T4ON:
		c.cc(CCT4, ccvalues)
	case T5ON:
		c.cc(CCT5, ccvalues)
	case T6ON:
		c.cc(CCT6, ccvalues)
	}

	timer := time.NewTimer(*track.Dur)

	// note on
	c.model.mu.Lock()
	c.model.pm.WriteShort(track.ON.Int64(), note.Int64(), intensity)
	c.model.mu.Unlock()

	go func() {
		<-timer.C

		// note off
		c.model.mu.Lock()
		c.model.pm.WriteShort(track.OFF.Int64(), note.Int64(), intensity)
		c.model.mu.Unlock()
	}()
}

func (p *Project) NewCC(values map[Parameter]int64) *CC {
	return &CC{
		PamVal: values,
	}
}

// cc .
func (c *Project) cc(track CCtrack, ccvalues map[Parameter]int64) {
	for k, v := range ccvalues {
		c.model.mu.Lock()
		c.model.pm.WriteShort(track.Int64(), k.Int64(), v)
		c.model.mu.Unlock()
	}
}

// func (c *Cycles) pc(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
// }

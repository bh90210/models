package machine

import (
	"strings"
	"sync"
	"time"

	"github.com/bh90210/models"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

type Parameter int

type Machine struct {
	drv *driver.Driver
	mu  sync.Mutex
	in  midi.In
	out midi.Out
	wr  *writer.Writer
}

func New() (*Machine, error) {
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	p := &Machine{
		drv: drv,
	}

	p.mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		p.in = in
	}
	outs, _ := drv.Outs()
	for _, out := range outs {
		// fmt.Println("Found MIDI output:", out.String())
		if strings.Contains(out.String(), "Elektron") {
			p.out = out
		}
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

func (s *Machine) PC(t models.Channel, pc int8) {
	s.wr.SetChannel(uint8(t))
	writer.ProgramChange(s.wr, uint8(pc))
}

func (s *Machine) CC(t Parameter, par Parameter, val uint8) {
	s.wr.SetChannel(uint8(t))
	writer.ControlChange(s.wr, uint8(par), val)
}

// Close midi connection. Use it with defer after creating a new project.
func (s *Machine) Close() {
	s.in.Close()
	s.out.Close()
	s.drv.Close()
}

func (s *Machine) Note(track Parameter, note Parameter, velocity int8, duration time.Duration) {
	s.NoteOn(track, note, velocity)
	go func() {
		time.Sleep(duration)
		s.NoteOff(track, note)
	}()
}

func (s *Machine) NoteOn(t Parameter, n Parameter, vel int8) {
	s.wr.SetChannel(uint8(t))
	writer.NoteOn(s.wr, uint8(n), uint8(vel))
}

func (s *Machine) NoteOff(t Parameter, n Parameter) {
	s.wr.SetChannel(uint8(t))
	writer.NoteOff(s.wr, uint8(n))
}

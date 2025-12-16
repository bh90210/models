package turbo

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/bh90210/models"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const Turbo string = "Alesis Turbo"

var Drums = map[uint8]string{
	38: "Snare",
	48: "Tom1",
	45: "Tom2",
	43: "Tom3",
	4:  "HiHat Closed",
	46: "HH",
	49: "Crash",
	51: "Ride",
	36: "Kick",
	21: "HH Foot",
}

var _ models.MidiCom = (*Project)(nil)

type Project struct {
	mu *sync.Mutex
	// midi fields
	drv      midi.Driver
	in       midi.In
	wr       *writer.Writer
	listener chan []byte
}

func NewProject() (*Project, error) {
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	p := &Project{
		mu:  new(sync.Mutex),
		drv: drv,
	}

	p.mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), Turbo) {
			p.in = in
		}
	}

	if p.in == nil {
		return nil, fmt.Errorf("device %s not found", Turbo)
	}

	err = p.in.Open()
	if err != nil {
		return nil, err
	}

	p.listener = make(chan []byte)

	p.in.SetListener(func(d []byte, deltaMicroseconds int64) {
		// We are getting a non stop signal from 248. We don't know why this is happening. Needs some further investigation.
		// Also we do nothing for the 44 messages as these are part of the foor high hat. This too needs further investigation as to what exactly is for.
		if bytes.Equal(d, []byte{248}) || d[1] == 44 {
			return
		}

		select {
		case p.listener <- d:
		default:
			fmt.Println("no turbo receiver", d)
		}
	})
	p.mu.Unlock()

	return p, nil
}

func (p *Project) Preset(track models.Channel, preset models.Preset) error {
	return models.ErrNotImplemented
}

func (p *Project) Note(track models.Channel, note models.Notes, velocity int8, duration float64) error {
	return models.ErrNotImplemented
}

func (p *Project) CC(track models.Channel, parameter models.Parameter, value int8) error {
	return models.ErrNotImplemented
}

func (p *Project) PC(t models.Channel, pc int8) error {
	return models.ErrNotImplemented
}

func (p *Project) Incoming() chan []byte {
	return p.listener
}

func (p *Project) Close() {
	p.in.Close()
	p.drv.Close()
}

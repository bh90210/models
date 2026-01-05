// Package turbo implements the midicom.MidiCom interface for the
// Alesis Turbo drum module.
package turbo

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/bh90210/models/midicom"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const Turbo string = "Alesis Turbo"

var Drums = map[string]uint8{
	"Snare":       38,
	"Tom1":        48,
	"Tom2":        45,
	"Tom3":        43,
	"HiHatClosed": 4,
	"HiHatOpen":   46,
	"Crash":       49,
	"Ride":        51,
	"Kick":        36,
	"HiHatFoot":   21,
}

var _ midicom.MidiCom = (*Project)(nil)

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

		fmt.Println(d)

		p.listener <- d
	})
	p.mu.Unlock()

	return p, nil
}

func (p *Project) Note(track midicom.Channel, note midicom.Note, velocity int8, duration float64) error {
	return midicom.ErrNotImplemented
}

func (p *Project) CC(track midicom.Channel, parameter midicom.Parameter, value int8) error {
	return midicom.ErrNotImplemented
}

func (p *Project) PC(t midicom.Channel, pc int8) error {
	return midicom.ErrNotImplemented
}

func (p *Project) Incoming() chan []byte {
	return p.listener
}

func (p *Project) Close() {
	p.in.Close()
	p.drv.Close()
}

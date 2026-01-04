package nymphes

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bh90210/models/midicom"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const (
	Nymphes string          = "Nymphes"
	Channel midicom.Channel = 0
)

var _ midicom.MidiCom = (*Project)(nil)

type Project struct {
	mu *sync.Mutex
	// midi fields
	drv      midi.Driver
	in       midi.In
	out      midi.Out
	wr       *writer.Writer
	listener chan []byte
}

func NewProject() (*Project, error) {
	// Initialize MIDI driver.
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	p := &Project{
		mu:  new(sync.Mutex),
		drv: drv,
	}

	p.mu.Lock()
	// Find the correct MIDI input and output ports.
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), Nymphes) {
			p.in = in
		}
	}

	outs, _ := drv.Outs()
	for _, out := range outs {
		if strings.Contains(out.String(), Nymphes) {
			p.out = out
		}
	}

	// Check if nothing found.
	if p.in == nil && p.out == nil {
		return nil, fmt.Errorf("device %s not found", Nymphes)
	}

	// Open MIDI input port.
	err = p.in.Open()
	if err != nil {
		return nil, err
	}

	// Set up MIDI listener.
	p.listener = make(chan []byte)
	p.in.SetListener(func(d []byte, deltaMicroseconds int64) {
		select {
		case p.listener <- d:
		default:
			fmt.Println("no nymphes receiver", d)
		}
	})

	// Open MIDI output port.
	err = p.out.Open()
	if err != nil {
		return nil, err
	}

	// Initialize MIDI writer.
	wr := writer.New(p.out)
	p.wr = wr
	p.wr.SetChannel(uint8(Channel))
	p.mu.Unlock()

	return p, nil
}

func (p *Project) Note(_ midicom.Channel, note midicom.Note, velocity int8, duration float64) error {
	// If velocity is 0, we just wait for the duration (rest)
	// but do not send any note on/off messages.
	if velocity == 0 {
		time.Sleep(time.Millisecond * time.Duration(duration))
		return nil
	}

	err := writer.NoteOn(p.wr, uint8(note), uint8(velocity))
	if err != nil {
		fmt.Println("Error sending NoteOn:", err)
		return err
	}

	time.Sleep(time.Millisecond * time.Duration(duration))

	err = writer.NoteOff(p.wr, uint8(note))
	if err != nil {
		fmt.Println("Error sending NoteOff:", err)
		return err
	}

	return nil
}

func (p *Project) CC(_ midicom.Channel, parameter midicom.Parameter, value int8) error {
	return writer.ControlChange(p.wr, uint8(parameter), uint8(value))
}

func (p *Project) PC(_ midicom.Channel, pc int8) error {
	return writer.ProgramChange(p.wr, uint8(pc))
}

func (p *Project) Incoming() chan []byte {
	return p.listener
}

func (p *Project) Close() {
	p.in.Close()
	p.out.Close()
	p.drv.Close()
}

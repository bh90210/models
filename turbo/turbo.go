package turbo

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/bh90210/models"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

const (
	turbo string = "Alesis Turbo"
)

var _ models.MidiCom = (*Project)(nil)

// Project long description of the data structure, methods, behaviors and useage.
type Project struct {
	mu *sync.Mutex
	// midi fields
	drv midi.Driver
	in  midi.In
	out midi.Out
	wr  *writer.Writer
}

// NewProject initiates and returns a *Project struct.
func NewProject() (*Project, error) {
	drv, err := driver.New()
	if err != nil {
		return nil, err
	}

	p := &Project{
		mu:  new(sync.Mutex),
		drv: drv,
	}

	// Find turbo and assign it to in/out.
	var helperIn bool

	p.mu.Lock()
	ins, _ := drv.Ins()
	for _, in := range ins {
		if strings.Contains(in.String(), turbo) {
			p.in = in
			helperIn = true
		}
	}

	// check if nothing found
	if !helperIn {
		return nil, fmt.Errorf("device %s not found", turbo)
	}

	err = p.in.Open()
	if err != nil {
		return nil, err
	}

	drums := map[uint8]string{
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

	p.in.SetListener(func(p []byte, deltaMicroseconds int64) {
		if bytes.Equal(p, []byte{248}) || p[1] == 44 {
			return
		}

		fmt.Println("MIDI IN:", drums[p[1]], p[2])
		// fmt.Println("MIDI IN:", p)
	})

	// err = p.out.Open()
	// if err != nil {
	// 	return nil, err
	// }

	// wr := writer.New(p.out)
	// p.wr = wr

	p.mu.Unlock()

	return p, nil
}

// Preset immediately sets (CC) provided parameters.
func (p *Project) Preset(track models.Channel, preset models.Preset) error {
	for parameter, value := range preset {
		err := p.CC(track, parameter, value)
		if err != nil {
			fmt.Println("Error sending CC:", err)
			return err
		}
	}

	return nil
}

// Note fires immediately a midi note on signal followed by a note off specified duration in milliseconds (ms).
// Optionally user can pass a preset too for convenience.
func (p *Project) Note(track models.Channel, note models.Notes, velocity int8, duration float64) error {
	p.wr.SetChannel(uint8(track))
	err := writer.NoteOn(p.wr, uint8(note), uint8(velocity))
	if err != nil {
		fmt.Println("Error sending NoteOn:", err)
		return err
	}

	time.Sleep(time.Millisecond * time.Duration(duration))
	p.wr.SetChannel(uint8(track))
	err = writer.NoteOff(p.wr, uint8(note))
	if err != nil {
		fmt.Println("Error sending NoteOff:", err)
		return err
	}

	return nil

}

// CC control change.
func (p *Project) CC(track models.Channel, parameter models.Parameter, value int8) error {
	p.wr.SetChannel(uint8(track))
	return writer.ControlChange(p.wr, uint8(parameter), uint8(value))
}

// PC Project control change.
func (p *Project) PC(t models.Channel, pc int8) error {
	p.wr.SetChannel(uint8(t))
	return writer.ProgramChange(p.wr, uint8(pc))
}

// Incoming reads incoming midi data.
func (p *Project) Incoming() chan []byte {
	ch := make(chan []byte)
	p.in.SetListener(func(p []byte, deltaMicroseconds int64) {
		select {
		case ch <- p:
		default:
			fmt.Println("no turbo receiver")
		}
	})

	return ch
}

// Close midi connection. Use it with defer after creating a new project.
func (p *Project) Close() {
	p.in.Close()
	p.out.Close()
	p.drv.Close()
}

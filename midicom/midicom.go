// Package midicom defines the interface for MIDI communication with
// various MIDI-capable devices.
package midicom

import "fmt"

// Parameter is all track parameters of the physical machine.
type Parameter int8

// Note are all notes reproducible by the machines.
type Note int8

// Channel represents a physical midi channel.
type Channel int8

// ErrNotImplemented is returned when a method is not implemented.
var ErrNotImplemented = fmt.Errorf("not implemented")

type MidiCom interface {
	Note(channel Channel, note Note, velocity int8, duration float64) error
	CC(channel Channel, parameter Parameter, value int8) error
	PC(channel Channel, pc int8) error
	Incoming() chan []byte
	Close()
}

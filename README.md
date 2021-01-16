<img src="https://i.imgur.com/omIKbjp.jpg" alt="drawing" width="350"/> <img src="https://i.imgur.com/pJbgSUh.png" alt="drawing" width="350"/>

# elektron:models

A small library that allows you to programmatically interact with [elektron](https://www.elektron.se/)'s **model:cycles** & **model:samples** via midi written in Go.

_WARNING: still in active development, things might not work, things might change._

## Prerequisites

### Go

Install Go https://golang.org/doc/install

### RtMidi

For Ubuntu 20.04+ run `apt install librtmidi4 librtmidi-dev`
For older versions take a look [here](https://launchpad.net/ubuntu/+source/rtmidi).

Instructions for other operating systems coming soon.

## Usage

_complete examples can be found under [examples](https://github.com/bh90210/elektronmodels/tree/master/examples/) folder_

If you haven't already, download either cycles/samples manual from elektron's website.
The relevant part for this library is the `APPENDIX A: MIDI SPECIFICATIONS`.

<img src="https://i.imgur.com/Yrs6YS3.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/cmil9NG.png" alt="drawing" width="350"/>

### Quick use

Cycles example: 
```go
package main

import (
	"log"

	cycles "github.com/bh90210/elektronmodels"
)

func main() {
	example, _ := cycles.NewProject()
	defer example.Close()

	t1intro := Intro()

	example.Pattern(t1intro)
	example.Loop()

	if err := example.Play(); err != nil {
		log.Println(err)
	}
}

func Intro() *cycles.Track {
	trig := cycles.NewTrig(0)
	trig.CC(
		map[cycles.Parameter]uint8{
			// e.NOTE:   int64(e.A4),
			cycles.REBERBTONE:   80,
			cycles.REVERBZISE:   80,
			cycles.DELAY:        0,
			cycles.DECAY:        50,
			cycles.SHAPE:        5,
			cycles.SWEEP:        10,
			cycles.CHANCE:       0,
			cycles.SWING:        0,
			cycles.GATE:         0,
			cycles.DELAYTIME:    10,
			cycles.COLOR:        120,
			cycles.LFODEST:      0,
			cycles.LFOWAVEFORM:  0,
			cycles.LFOMULTIPIER: 0,
			cycles.LFODEPTH:     0,
		})
	trig.Note(
		cycles.A4,
		120,
		time.Duration(100*time.Millisecond))

	endTrig := cycles.LastTrig(2)

	var trigs []*cycles.Trig
	trigs = append(trigs, trig, endTrig)

	track1 := cycles.NewTrack(cycles.T1, trigs)

	return track1
}

```

### Pattern system explained

### Timing system explained

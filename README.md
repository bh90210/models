<img src="https://www.elektron.se/wp-content/uploads/2020/03/ModelCycles_Above_2400.png " alt="drawing" width="350"/> <img src="https://images-na.ssl-images-amazon.com/images/I/91mGXDflYmL._AC_SL1500_.jpg" alt="drawing" width="350"/>

# elektron:models

A small library that allows you to interact with elektron's model:cycles & model:samples via midi written in Go.

## Prerequisites

### Go

### Portmidi

## Usage

_complete examples can be found under [examples](https://github.com/bh90210/elektronmodels/tree/master/examples/) folder_

```go
package main

import (
	"log"

	cycles "github.com/bh90210/elektronmodels"
)

func main() {
	gm258plague, err := cycles.NewProject()
	if err != nil {
		log.Fatal(err)
	}
	defer gm258plague.Close()

	t1intro := Intro()

	gm258plague.NewPattern(t1intro)

	gm258plague.Loop()
	if err := gm258plague.Play(); err != nil {
		log.Println(err)
	}
}

func Intro() *cycles.Track {
	trig := e.NewTrig(0)
	trig.CC(
		map[e.Parameter]int64{
			// e.NOTE:   int64(e.A4),
			e.REBERBTONE:   80,
			e.REVERBZISE:   80,
			e.DELAY:        0,
			e.DECAY:        50,
			e.SHAPE:        int64(5),
			e.SWEEP:        10,
			e.CHANCE:       0,
			e.SWING:        0,
			e.GATE:         0,
			e.DELAYTIME:    10,
			e.COLOR:        120,
			e.LFODEST:      int64(0),
			e.LFOWAVEFORM:  0,
			e.LFOMULTIPIER: 0,
			e.LFODEPTH:     0,
		})
	trig.Note(
		e.A4,
		120,
		time.Duration(100*time.Millisecond))

	endTrig := e.LastTrig(2)

	track1 := e.NewTrack(e.T1, trig, endTrig)

	return track1
}

```

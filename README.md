<img src="https://user-images.githubusercontent.com/22690219/130872109-150ac61f-ad69-4bfb-8f10-3337abcb6551.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/pJbgSUh.png" alt="drawing" width="350"/>

[![Go Reference](https://pkg.go.dev/badge/github.com/bh90210/models.svg)](https://pkg.go.dev/github.com/bh90210/models)

# elektron:models

Go package to programmatically control [Elektron's](https://www.elektron.se/) **model:cycles** & **model:samples** via midi.

## Prerequisites

### Go

Install Go https://golang.org/doc/install.

### RtMidi

#### Ubuntu 20.04+

```console
apt install librtmidi4 librtmidi-dev
```
For older versions take a look [here](https://launchpad.net/ubuntu/+source/rtmidi).

#### MacOS

```console
brew install rtmidi
```
For more information see the [formulae page](https://formulae.brew.sh/formula/rtmidi).

#### Windows

`Help needed.`

## Quick Use

_complete examples can be found in the [examples](https://github.com/bh90210/elektronmodels/tree/master/examples/) folder._

_The relevant cycles/samples manuals' part for this library is the `APPENDIX A: MIDI SPECIFICATIONS`._

<img src="https://i.imgur.com/Yrs6YS3.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/cmil9NG.png" alt="drawing" width="350"/>

### Free

Code to get a single kick drum hit at C4 key, with velocity set at `120` and length at 200 milliseconds:
```go
package main

import (
	"time"

	m "github.com/bh90210/models"
)

func main() {
	p, _ := m.NewProject(em.CYCLES)
	defer p.Close()

	p.Free.Note(m.T1, m.C4, 120, 200, m.PT1())
	time.Sleep(200 * time.Millisecond)
}

```
There are four Free methods to use, `Preset` to set preset on the fly, `Note` to fire a note on/off for given duration, `CC` to send a single control change message && `PC` for program changes. 

### Sequencer

_For an in-depth tutorial on how to use the sequencer [read this]() Hackernoon post._

The sequencer is implemented to mirror the original machine's functionality as close as reasonable/possible. You can expect the `Play` method to play given loop indefinitely, `Change` to change you to a new pattern, `Chain` to allow for multiple patterns to be chained together in serries etc.

Here is a single 16-bar example utilizing all six tracks:
```go
package main

import (
	m "github.com/bh90210/models"
)

const (
	INTRO int = iota
	VERSE
)

func main() {
	project, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer project.Close()

	intro := project.Pattern(INTRO)

	t1 := intro.Track(m.T1)
	t1.Trig(0)
	t1.Trig(8)

	t2 := intro.Track(m.T2)
	t2.Trig(4)
	t2.Trig(12)

	t3 := intro.Track(m.T3)
	t3.Trig(0)
	t3.Trig(4)
	t3.Trig(8)
	t3.Trig(12)

	t4 := intro.Track(m.T4)
	t4.Trig(5)

	verse := project.Pattern(VERSE)

	t1 = verse.Track(m.T1)
	t1.Trig(0)
	t1.Trig(6)
	t1.Trig(8)
	t1.Trig(14)

	t2 = verse.Track(m.T2)
	t2.Trig(4)
	t2.Trig(12)

	t3 = verse.Track(m.T3)
	t3.Trig(0)
	t3.Trig(2)
	t3.Trig(4)
	t3.Trig(6)
	t3.Trig(8)
	t3.Trig(10)
	t3.Trig(12)
	t3.Trig(14)

	t4 = verse.Track(m.T4)
	t4.Trig(5)
	t4.Trig(11)

	project.Chain(INTRO, VERSE).Play()
```

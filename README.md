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

_A complete example can be found in the [example](https://github.com/bh90210/elektronmodels/tree/master/example/) folder._

_The relevant cycles/samples manuals' part for this library is the `APPENDIX A: MIDI SPECIFICATIONS`._

<img src="https://i.imgur.com/Yrs6YS3.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/cmil9NG.png" alt="drawing" width="350"/>


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

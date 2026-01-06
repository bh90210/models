[![Go Reference](https://pkg.go.dev/badge/github.com/bh90210/models.svg)](https://pkg.go.dev/github.com/bh90210/models)

# Models

Go package to programmatically control via midi:
* Elektron model:cycles 
* Elektron model:samples
* Nord Lead x2
* Dreadbox Nymphes

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

_Complete examples can be found in the [examples](https://github.com/bh90210/elektronmodels/tree/master/examples/) folder._

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

	p.Note(m.T1, m.C4, 120, 200, m.PT1())
	time.Sleep(200 * time.Millisecond)
}

```

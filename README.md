<img src="https://user-images.githubusercontent.com/22690219/130872109-150ac61f-ad69-4bfb-8f10-3337abcb6551.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/pJbgSUh.png" alt="drawing" width="350"/>

[![DeepSource](https://deepsource.io/gh/bh90210/elektronmodels.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/bh90210/elektronmodels/?ref=repository-badge)

# elektron:models

A Go package to programmatically interact with [elektron](https://www.elektron.se/)'s **model:cycles** & **model:samples** via midi.

## Prerequisites

### Go

Install Go https://golang.org/doc/install

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

## Usage

_complete examples can be found under [examples](https://github.com/bh90210/elektronmodels/tree/master/examples/) folder_

If you haven't already, download cycles/samples manuals from elektron's website.
The relevant part for this library is the `APPENDIX A: MIDI SPECIFICATIONS`.

<img src="https://i.imgur.com/Yrs6YS3.png" alt="drawing" width="350"/> <img src="https://i.imgur.com/cmil9NG.png" alt="drawing" width="350"/>

### Free

### Sequencer
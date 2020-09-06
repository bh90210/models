package elektronmodels

import (
	"time"
)

type track int64

const (
	T1 track = iota
	T2
	T3
	T4
	T5
	T6
)

type cctrack int64

const (
	cct1 cctrack = 0xB0
	cct2 cctrack = 0xB1
	cct3 cctrack = 0xB2
	cct4 cctrack = 0xB3
	cct5 cctrack = 0xB4
	cct6 cctrack = 0xB5
)

func (c *cctrack) int64() int64 {
	return int64(*c)
}

type cc struct {
	pamVal map[Parameter]int64
}

type notes int64

const (
	A0 notes = iota + 21
	As0
	B0
	C1
	Cs1
	D1
	Ds1
	E1
	F1
	Fs1
	G1
	Gs1
	A1
	As1
	B1
	C2
	Cs2
	D2
	Ds2
	E2
	F2
	Fs2
	G2
	Gs2
	A2
	As2
	B2
	C3
	Cs3
	D3
	Ds3
	E3
	F3
	Fs3
	G3
	Gs3
	A3
	As3
	B3
	C4
	Cs4
	D4
	Ds4
	E4
	F4
	Fs4
	G4
	Gs4
	A4
	As4
	B4
	C5
	Cs5
	D5
	Ds5
	E5
	F5
	Fs5
	G5
	Gs5
	A5
	As5
	B5
	C6
	Cs6
	D6
	Ds6
	E6
	F6
	Fs6
	G6
	Gs6
	A6
	As6
	B6
	C7
	Cs7
	D7
	Ds7
	E7
	F7
	Fs7
	G7
	Gs7
	A7
	As7
	B7
	C8

	Bf0 notes = As0
	Df1 notes = Cs1
	Ef1 notes = Ds1
	Gf1 notes = Fs1
	Af1 notes = Gs1
	Bf1 notes = As1
	Df2 notes = Cs2
	Ef2 notes = Ds2
	Gf2 notes = Fs2
	Af2 notes = Gs2
	Bf2 notes = As2
	Df3 notes = Cs3
	Ef3 notes = Ds3
	Gf3 notes = Fs3
	Af3 notes = Gs3
	Bf3 notes = As3
	Df4 notes = Cs4
	Ef4 notes = Ds4
	Gf4 notes = Fs4
	Af4 notes = Gs4
	Bf4 notes = As4
	Df5 notes = Cs5
	Ef5 notes = Ds5
	Gf5 notes = Fs5
	Af5 notes = Gs5
	Bf5 notes = As5
	Df6 notes = Cs6
	Ef6 notes = Ds6
	Gf6 notes = Fs6
	Af6 notes = Gs6
	Bf6 notes = As6
	Df7 notes = Cs7
	Ef7 notes = Ds7
	Gf7 notes = Fs7
	Af7 notes = Gs7
	Bf7 notes = As7
)

func (n *notes) int64() int64 {
	return int64(*n)
}

type chord int64

const (
	Unisonx2 chord = iota
	Unisonx3
	Unisonx4
	Minor
	Major
	Sus2
	Sus4
	MinorMinor7
	MajorMinor7
	MinorMajor7
	MajorMajor7
	MinorMinor7Sus4
	Dim7
	MinorAdd9
	MajorAdd9
	Minor6
	Major6
	Minorb5
	Majorb5
	MinorMinor7b5
	MajorMinor7b5
	MajorAug5
	MinorMinor7Aug5
	MagorMinor7Aug5
	Minorb6
	MinorMinor9no5
	MajorMinor9no5
	MajorAdd9b5
	MajorMajor7b5
	MajorMinor7b9no5
	Sus4Aug5b9
	Sus4AddAug5
	MajorAddb5
	Major6Add4no5
	MajorMajor76no5
	MajorMajor9no5
	Fourths
	Fifths
)

type note struct {
	on  noteOn
	off noteOff
	dur *time.Duration
	key notes
}

type noteOn int64

const (
	t1on noteOn = 0x90
	t2on noteOn = 0x91
	t3on noteOn = 0x92
	t4on noteOn = 0x93
	t5on noteOn = 0x94
	t6on noteOn = 0x95
)

func (n *noteOn) int64() int64 {
	return int64(*n)
}

type noteOff int64

const (
	t1off noteOff = 0x80
	t2off noteOff = 0x81
	t3off noteOff = 0x82
	t4off noteOff = 0x83
	t5off noteOff = 0x84
	t6off noteOff = 0x85
)

func (n *noteOff) int64() int64 {
	return int64(*n)
}

type Parameter int64

const (
	NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10

	// model:cycles
	CYCLESPITCH Parameter = 65
	DECAY       Parameter = 80
	COLOR       Parameter = 16
	SHAPE       Parameter = 17

	// model:samples
	PITCH        Parameter = 16
	SAMPLESTART  Parameter = 19
	SAMPLELENGTH Parameter = 20
	CUTOFF       Parameter = 74
	RESONANCE    Parameter = 71

	// model:cycles
	PUNCH Parameter = 66
	GATE  Parameter = 67

	// model:samples
	LOOP    Parameter = 17
	REVERSE Parameter = 18

	SWEEP   Parameter = 18
	CONTOUR Parameter = 19
	DELAY   Parameter = 12
	REVERB  Parameter = 13

	VOLUMEDIST Parameter = 7
	SWING      Parameter = 15
	CHANCE     Parameter = 14

	// LFO section
	LFOSPEED      Parameter = 102
	LFOMULTIPIER  Parameter = 103
	LFOFADE       Parameter = 104
	LFODEST       Parameter = 105
	LFOWAVEFORM   Parameter = 106
	LFOSTARTPHASE Parameter = 107
	LFORESET      Parameter = 108
	LFODEPTH      Parameter = 109

	// FX section
	DELAYTIME     Parameter = 85
	DELAYFEEDBACK Parameter = 86
	REVERBZISE    Parameter = 87
	REBERBTONE    Parameter = 88
)

func (p *Parameter) int64() int64 {
	return int64(*p)
}

type lfoDest int64

const (
	LNONE    lfoDest = 0
	LPITCH   lfoDest = 9
	LCOLOR   lfoDest = 11
	LSHAPE   lfoDest = 12
	LSWEEP   lfoDest = 13
	LCONTOUR lfoDest = 14
	LPAW     lfoDest = 15
	LGATE    lfoDest = 16
	LFTUN    lfoDest = 17
	LDECAY   lfoDest = 18
	LDIST    lfoDest = 19
	LDELAY   lfoDest = 20
	LREVERB  lfoDest = 21
	LPAN     lfoDest = 22
)

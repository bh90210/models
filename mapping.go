package elektronmodels

import (
	"time"
)

type Track int64

const (
	T1 Track = iota
	T2
	T3
	T4
	T5
	T6
)

type Notes int64

const (
	A0 Notes = iota + 21
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

	Bf0 Notes = As0
	Df1 Notes = Cs1
	Ef1 Notes = Ds1
	Gf1 Notes = Fs1
	Af1 Notes = Gs1
	Bf1 Notes = As1
	Df2 Notes = Cs2
	Ef2 Notes = Ds2
	Gf2 Notes = Fs2
	Af2 Notes = Gs2
	Bf2 Notes = As2
	Df3 Notes = Cs3
	Ef3 Notes = Ds3
	Gf3 Notes = Fs3
	Af3 Notes = Gs3
	Bf3 Notes = As3
	Df4 Notes = Cs4
	Ef4 Notes = Ds4
	Gf4 Notes = Fs4
	Af4 Notes = Gs4
	Bf4 Notes = As4
	Df5 Notes = Cs5
	Ef5 Notes = Ds5
	Gf5 Notes = Fs5
	Af5 Notes = Gs5
	Bf5 Notes = As5
	Df6 Notes = Cs6
	Ef6 Notes = Ds6
	Gf6 Notes = Fs6
	Af6 Notes = Gs6
	Bf6 Notes = As6
	Df7 Notes = Cs7
	Ef7 Notes = Ds7
	Gf7 Notes = Fs7
	Af7 Notes = Gs7
	Bf7 Notes = As7
)

func (n *Notes) Int64() int64 {
	return int64(*n)
}

type Chord int64

const (
	Unisonx2 Chord = iota
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

type Note struct {
	ON  NoteOn
	OFF NoteOff
	Dur *time.Duration
	*CC
}

type NoteOn int64

const (
	T1ON NoteOn = 0x90
	T2ON NoteOn = 0x91
	T3ON NoteOn = 0x92
	T4ON NoteOn = 0x93
	T5ON NoteOn = 0x94
	T6ON NoteOn = 0x95
)

func (n *NoteOn) Int64() int64 {
	return int64(*n)
}

type NoteOff int64

const (
	T1OFF NoteOff = 0x80
	T2OFF NoteOff = 0x81
	T3OFF NoteOff = 0x82
	T4OFF NoteOff = 0x83
	T5OFF NoteOff = 0x84
	T6OFF NoteOff = 0x85
)

func (n *NoteOff) Int64() int64 {
	return int64(*n)
}

type CCtrack int64

const (
	CCT1 CCtrack = 0xB0
	CCT2 CCtrack = 0xB1
	CCT3 CCtrack = 0xB2
	CCT4 CCtrack = 0xB3
	CCT5 CCtrack = 0xB4
	CCT6 CCtrack = 0xB5
)

func (c *CCtrack) Int64() int64 {
	return int64(*c)
}

type CC struct {
	PamVal map[Parameter]int64
}

type Parameter int64

const (
	NOTE       Parameter = 3
	TRACKLEVEL Parameter = 17
	MUTE       Parameter = 94
	PAN        Parameter = 10

	PITCH Parameter = 65
	DECAY Parameter = 80
	COLOR Parameter = 16
	SHAPE Parameter = 17

	SWEEP   Parameter = 18
	CONTOUR Parameter = 19
	DELAY   Parameter = 12
	REVERB  Parameter = 13

	VOLUMEDIST Parameter = 7
	SWING      Parameter = 15
	CHANCE     Parameter = 14

	PUNCH Parameter = 66
	GATE  Parameter = 67

	LFOSPEED      Parameter = 102
	LFOMULTIPIER  Parameter = 103
	LFOFADE       Parameter = 104
	LFODEST       Parameter = 105
	LFOWAVEFORM   Parameter = 106
	LFOSTARTPHASE Parameter = 107
	LFORESET      Parameter = 108
	LFODEPTH      Parameter = 109

	DELAYTIME     Parameter = 85
	DELAYFEEDBACK Parameter = 86
	REVERBZISE    Parameter = 87
	REBERBTONE    Parameter = 88
)

func (p *Parameter) Int64() int64 {
	return int64(*p)
}

type LFODest int64

const (
	LNONE    LFODest = 0
	LPITCH   LFODest = 9
	LCOLOR   LFODest = 11
	LSHAPE   LFODest = 12
	LSWEEP   LFODest = 13
	LCONTOUR LFODest = 14
	LPAW     LFODest = 15
	LGATE    LFODest = 16
	LFTUN    LFODest = 17
	LDECAY   LFODest = 18
	LDIST    LFODest = 19
	LDELAY   LFODest = 20
	LREVERB  LFODest = 21
	LPAN     LFODest = 22
)

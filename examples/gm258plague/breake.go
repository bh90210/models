package main

import (
	"log"
	"time"
)

func Breake(c *GM258plague) (*[]byte, error) {
	var klimax []Note = []Note{C1, G1, F1, D1}
	var chords []Chord = []Chord{Major6, MajorAdd9, MajorMajor7, MinorMinor9no5, MajorMajor7, MajorMinor9no5, MajorMajor76no5, MinorMinor7}
	var i int
	var o int

	var1 := make(chan int64, 1)
	var1 <- 0

	loopy := func() {
		var intensity int64
		for {
			select {
			case newVal := <-var1:
				intensity = newVal
			default:
			}

			c.Cycles.Note(NewNoteTrack(T6, time.Duration(100*time.Millisecond)), klimax[i], 120,
				map[Parameter]int64{
					REVERB:       60,
					REBERBTONE:   80,
					REVERBZISE:   80,
					DELAY:        0,
					DECAY:        50,
					SHAPE:        int64(chords[o]),
					SWEEP:        10,
					CHANCE:       100,
					GATE:         0,
					DELAYTIME:    10,
					COLOR:        120,
					LFODEST:      int64(LCONTOUR),
					LFOWAVEFORM:  0,
					LFOMULTIPIER: 14,
					LFODEPTH:     120,
				},
			)
			i++
			if i == 4 {
				i = 0
			}
			o++
			if o == 8 {
				o = 0
			}
		}
	return nil, nil
}

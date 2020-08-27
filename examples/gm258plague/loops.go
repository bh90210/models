package main

// import (
// 	"log"
// 	"math/rand"
// 	"time"
// )

// func intro(cycles *Cycles, repeat int) {
// 	var dur time.Duration
// 	switch repeat {
// 	// intro of intro
// 	case 0:
// 		dur = time.Duration(15000 * time.Millisecond)
// 		t := time.NewTicker(dur)

// 		cycles.Note(NewNoteTrack(T6, time.Duration(120000*time.Millisecond)), C3, 120,
// 			NewCC(CCT6, REVERB, 80),
// 			NewCC(CCT6, DELAY, 10),
// 			NewCC(CCT6, DECAY, 20),
// 			NewCC(CCT6, CHANCE, 100),
// 			NewCC(CCT6, GATE, 1),
// 			NewCC(CCT6, PUNCH, 1),
// 			NewCC(CCT6, VOLUMEDIST, 60),

// 			NewCC(CCT6, SHAPE, int64(Major6Add4no5)),
// 			NewCC(CCT6, COLOR, 125),
// 			NewCC(CCT6, SWEEP, 0),
// 			NewCC(CCT6, CONTOUR, 125),
// 		)

// 		go func() {
// 			var (
// 				up   int64 = 0
// 				down int64 = 126
// 			)
// 			slide := time.NewTicker(500 * time.Millisecond)
// 			stop := time.NewTimer(120000 * time.Millisecond)

// 		loop:
// 			for {
// 				<-slide.C
// 				down = down - 1
// 				cycles.CC(
// 					NewCC(CCT6, COLOR, rand.Int63n(120)),
// 					NewCC(CCT6, CONTOUR, down),
// 				)
// 				log.Println(down)

// 				<-slide.C
// 				up = up + 1
// 				cycles.CC(
// 					// NewCC(CCT6, COLOR, rand.Int63n(120)),
// 					NewCC(CCT6, SWEEP, up),
// 				)
// 				log.Println(up)

// 				select {
// 				case <-stop.C:
// 					slide.Stop()
// 					break loop
// 				default:
// 				}
// 			}
// 		}()

// 		// 15-30
// 		<-t.C
// 		// 30-45
// 		<-t.C
// 		// 45-60
// 		<-t.C

// 		<-t.C

// 		cycles.CC(
// 			NewCC(CCT6, NOTE, int64(D3)),
// 			NewCC(CCT6, SHAPE, int64(MinorAdd9)),
// 		)

// 		<-t.C

// 		cycles.CC(
// 			NewCC(CCT6, NOTE, int64(E3)),
// 			NewCC(CCT6, SHAPE, int64(MinorMinor9no5)),
// 		)

// 		<-t.C

// 		return
// 	case 1:
// 		dur = time.Duration(650 * time.Millisecond)
// 	case 2:
// 		dur = time.Duration(700 * time.Millisecond)
// 	case 3:
// 		dur = time.Duration(850 * time.Millisecond)
// 	case 4:
// 		dur = time.Duration(1000 * time.Millisecond)
// 	}

// 	t := time.NewTicker(dur)

// 	<-t.C
// 	cycles.Note(NewNoteTrack(T6, time.Duration(600*time.Millisecond)), C3, 120,
// 		NewCC(CCT6, REVERB, 80),
// 		NewCC(CCT6, DELAY, 10),
// 		NewCC(CCT6, DECAY, 50),
// 		NewCC(CCT6, SHAPE, int64(Major6Add4no5)),
// 		NewCC(CCT6, SWEEP, 80),
// 		NewCC(CCT6, CHANCE, 100),
// 		NewCC(CCT6, GATE, 1),
// 	)

// 	<-t.C
// 	cycles.Note(NewNoteTrack(T6, time.Duration(600*time.Millisecond)), D3, 120,
// 		NewCC(CCT6, REVERB, 80),
// 		NewCC(CCT6, DELAY, 10),
// 		NewCC(CCT6, DECAY, 50),
// 		NewCC(CCT6, SHAPE, int64(MinorAdd9)),
// 		NewCC(CCT6, SWEEP, 80),
// 		NewCC(CCT6, CHANCE, 100),
// 		NewCC(CCT6, GATE, 1),
// 	)

// 	<-t.C
// 	cycles.Note(NewNoteTrack(T6, time.Duration(580*time.Millisecond)), E3, 120,
// 		NewCC(CCT6, REVERB, 80),
// 		NewCC(CCT6, DELAY, 10),
// 		NewCC(CCT6, DECAY, 50),
// 		NewCC(CCT6, SHAPE, int64(MinorMinor9no5)),
// 		NewCC(CCT6, SWEEP, 80),
// 		NewCC(CCT6, CHANCE, 100),
// 		NewCC(CCT6, GATE, 1),
// 	)

// 	<-t.C
// 	cycles.Note(NewNoteTrack(T6, time.Duration(600*time.Millisecond)), D3, 120,
// 		NewCC(CCT6, REVERB, 80),
// 		NewCC(CCT6, DELAY, 10),
// 		NewCC(CCT6, DECAY, 50),
// 		NewCC(CCT6, SHAPE, int64(MinorAdd9)),
// 		NewCC(CCT6, SWEEP, 80),
// 		NewCC(CCT6, CHANCE, 100),
// 		NewCC(CCT6, GATE, 1),
// 	)
// 	t.Stop()
// }

package main

import (
	m "github.com/athenez/models"
)

const (
	INTRO int = iota
	VERSE
	CHORUS
	OUTRO
)

func main() {
	// start a new project
	p, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	p.Pattern(INTRO).
		Scale(m.TRK, 15, 4.0, 0).
		Tempo(250.0)

	p.Pattern(INTRO).Track(m.T2).Scale(15, 3.0)

	// // kick
	kick := []int{0, 4, 8, 12}
	for _, v := range kick {
		p.Pattern(INTRO).Track(m.T2).Trig(v).Note(m.A4, 200, 120)
	}

	// preset1 := m.PT3()
	// preset1[m.GATE] = 1
	// preset1[m.PUNCH] = 1
	// preset1[m.DECAY] = 120
	// // p.Pattern(INTRO).Track(m.T1).Trig(0).Lock(preset1)
	// // p.Pattern(INTRO).Track(m.T1).Trig(4).Lock(preset1)

	// p.Pattern(INTRO).Track(m.T2).Trig(8).Lock(preset1)
	// p.Pattern(INTRO).Track(m.T2).Trig(0).Nudge(250)
	// // p.Pattern(INTRO).Track(m.T1).Trig(12).Lock(preset2)

	// // snare
	// snare := []int{0, 4, 8, 12}
	// for _, v := range snare {
	// 	p.Pattern(INTRO).Track(m.T1).Trig(v)
	// }

	// // hi-hat
	// for i := 0; i <= 15; i++ {
	// 	p.Pattern(INTRO).Track(m.T3).Trig(i).Scale(float64(i))
	// 	// p.Pattern(INTRO).Track(m.T3).Trig(i)
	// }

	// // tom
	// tom := []int{2, 3, 6, 7, 10, 11, 14, 15}
	// for _, v := range tom {
	// 	p.Pattern(INTRO).Track(m.T4).Trig(v)
	// }

	// // tone
	// tone := []int{5, 9, 10}
	// for _, v := range tone {
	// 	p.Pattern(INTRO).Track(m.T5).Trig(v)
	// }

	// // synth
	// synth := []int{3, 7, 9}
	// for _, v := range synth {
	// 	p.Pattern(INTRO).Track(m.T6).Trig(v)
	// }

	// preset := make(map[m.Parameter]int8)
	// preset := m.PT1()
	// preset[m.COLOR] = 100
	// p.Pattern(INTRO).Track(m.T1).Preset(preset)
	// p.Pattern(INTRO).Track(m.T1).Trig(0).Lock(preset)
	// p.Pattern(VERSE).
	// 	Scale(m.TRK, 15, 4.0, 4).
	// 	Tmpo(250.0)

	p.Pattern(VERSE).Scale(m.TRK, 15, 1.0, 7).Tempo(250.0).Track(m.T2).Scale(15, 3.0)

	// kick = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, v := range kick {
		p.Pattern(VERSE).Track(m.T2).Trig(v).Note(m.A4, 200, 120)
	}

	p.Chain(INTRO, VERSE).Play()
}

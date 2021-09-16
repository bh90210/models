package main

import (
	"time"

	em "github.com/athenez/models"
)

func main() {
	// start a new project
	p, err := em.NewProject(em.CYCLES)
	if err != nil {
		panic(err)
	}
	defer p.Close()
	// free
	var i int = 500
	p.Free.CC(em.T1, em.MACHINE, 1)
	p.Free.CC(em.T1, em.GATE, 0)
	p.Free.CC(em.T1, em.COLOR, 100)
	p.Free.CC(em.T1, em.CONTOUR, 0)
	p.Free.CC(em.T1, em.SWEEP, 100)
	p.Free.CC(em.T1, em.REVERB, 0)
	p.Free.Note(em.T1, em.C4, 120, 200)
	p.Free.CC(em.T1, em.MACHINE, 2)
	time.Sleep(time.Duration(i) * time.Millisecond)

	p.Free.CC(em.T1, em.GATE, 1)
	p.Free.CC(em.T1, em.COLOR, 120)
	p.Free.CC(em.T1, em.CONTOUR, 0)
	p.Free.CC(em.T1, em.SWEEP, 120)
	p.Free.CC(em.T1, em.REVERB, 120)
	p.Free.CC(em.T1, em.REVERBSIZE, 120)
	p.Free.Note(em.T1, em.A4, 120, 500)
	p.Free.CC(em.T1, em.MACHINE, 6)
	time.Sleep(time.Duration(i*4) * time.Millisecond)

	p.Free.CC(em.T1, em.GATE, 0)
	p.Free.CC(em.T1, em.COLOR, 100)
	p.Free.CC(em.T1, em.CONTOUR, 0)
	p.Free.CC(em.T1, em.SWEEP, 100)
	p.Free.CC(em.T1, em.REVERB, 0)
	p.Free.CC(em.T1, em.REVERBSIZE, 10)
	p.Free.Note(em.T1, em.D6, 120, 1500)
	go func() {
		var color, contour, sweep em.Parameter
		color = 100
		sweep = 100
		contour = 0
		for {
			color -= 1
			sweep -= 1
			contour += 1
			time.Sleep(10 * time.Millisecond)
			p.Free.CC(em.T1, em.COLOR, int8(color))
			p.Free.CC(em.T1, em.CONTOUR, int8(contour))
			p.Free.CC(em.T1, em.SWEEP, int8(sweep))
			p.Free.CC(em.T1, em.SHAPE, int8(sweep))
		}
	}()
	time.Sleep(2 * time.Second)
}

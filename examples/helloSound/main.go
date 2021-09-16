package main

import (
	"time"

	m "github.com/athenez/models"
)

func main() {
	project, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer project.Close()

	pattern := project.Pattern(0)

	pattern.Track(m.T1).Trig(0)
	pattern.Track(m.T2).Trig(2)
	pattern.Track(m.T3).Trig(4)
	pattern.Track(m.T4).Trig(6)
	pattern.Track(m.T5).Trig(8)
	pattern.Track(m.T6).Trig(10)

	go project.Play(0)

	time.Sleep(time.Second * 6)
}

package main

import (
	"time"

	m "github.com/bh90210/models"
)

func main() {
	project, err := m.NewProject(m.CYCLES)
	if err != nil {
		panic(err)
	}
	defer project.Close()

	pattern := project.Pattern(0)

	pattern.Track(m.T1).Trig(0)
	pattern.Track(m.T2).Trig(1)
	pattern.Track(m.T3).Trig(2)
	pattern.Track(m.T4).Trig(3)
	pattern.Track(m.T5).Trig(4)
	pattern.Track(m.T6).Trig(5)

	go project.Play(0)

	time.Sleep(time.Second * 3)
}

package main

import (
	"fmt"
	"time"

	"github.com/bh90210/models/machine"
	"github.com/bh90210/models/machine/bd"
)

func main() {
	m, err := machine.New()
	if err != nil {
		panic(err)
	}

	defer m.Close()

	t1 := new(machine.Track)

	dursEq7 := machine.EqualDuration(7, time.Millisecond*500)
	for _, d := range dursEq7 {
		t1.Add(machine.Event{
			Duration: d,
			Action: func(duration time.Duration) {
				m.CC(bd.CHANNEL, bd.LEVEL, 100)
				m.Note(bd.CHANNEL, bd.NOTE, 120, duration)
			},
		})
	}

	for i, d := range dursEq7 {
		fmt.Println("dur: ", d/time.Duration(i+1))
		t1.Add(machine.Event{
			Duration: d / time.Duration(i+1),
			Action: func(duration time.Duration) {
				m.CC(bd.CHANNEL, bd.LEVEL, 0)
				m.Note(bd.CHANNEL, bd.NOTE, 0, duration)
			},
		})
	}

	t1.Run()
}

// 1. Just kick
// 2. All channels + Changing machines
// 4. With nord
// 5. With semi modulars

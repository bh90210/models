package main

import (
	"fmt"
	"time"

	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

func main() {
	err := midiConnection()
	if err != nil {
		panic(err)
	}

	select {}
}

func midiConnection() error {
	drv, err := driver.New()
	if err != nil {
		return err
	}

	outs, _ := drv.Outs()
	for i, out := range outs {
		// if strings.Contains(out.String(), string(m)) {
		// }
		fmt.Println(out)

		if i == 1 {
			fmt.Println(out, i)
			err = out.Open()
			if err != nil {
				return err
			}

			wr := writer.New(out)

			wr.SetChannel(0)

			for j := 0; j < 5; j++ {
				writer.NoteOn(wr, 50, 120)
				time.Sleep(1500 * time.Millisecond)
				writer.NoteOff(wr, 50)

				writer.NoteOn(wr, 60, 120)
				time.Sleep(1500 * time.Millisecond)
				writer.NoteOff(wr, 60)

				writer.NoteOn(wr, 70, 120)
				time.Sleep(1500 * time.Millisecond)
				writer.NoteOff(wr, 70)

				writer.NoteOn(wr, 75, 120)
				time.Sleep(1500 * time.Millisecond)
				writer.NoteOff(wr, 75)

				fmt.Println(j)
			}
			// writer.ControlChange(wr, 20, 20)
		}
	}

	err = p.out.Open()
	if err != nil {
		return nil, err
	}

	return nil
}

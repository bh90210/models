package main

import (
	"bytes"
	"fmt"
	"strings"

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

	ins, _ := drv.Ins()
	for _, in := range ins {
		// if strings.Contains(in.String(), string(m)) {
		// }
		fmt.Println(in)

		if strings.Contains(in.String(), "Turbo:Alesis") {
			in.Open()

			drums := map[uint8]string{
				38: "Snare",
				48: "Tom1",
				45: "Tom2",
				43: "Tom3",
				4:  "HiHat Closed",
				46: "HH",
				49: "Crash",
				51: "Ride",
				36: "Kick",
				21: "HH Foot",
			}

			in.SetListener(func(p []byte, deltaMicroseconds int64) {
				if bytes.Equal(p, []byte{248}) || p[1] == 44 {
					return
				}

				fmt.Println("MIDI IN:", drums[p[1]], p[2])
				// fmt.Println("MIDI IN:", p)
			})
			// time.Sleep(1 * time.Minute)
		}
	}

	// outs, _ := drv.Outs()
	// for i, out := range outs {
	// 	// if strings.Contains(out.String(), string(m)) {
	// 	// }
	// 	fmt.Println(out)

	// 	if i == 1 {
	// 		fmt.Println(out, i)
	// 		err = out.Open()
	// 		if err != nil {
	// 			return err
	// 		}

	// 		wr := writer.New(out)

	// 		wr.SetChannel(0)

	// 		for j := 0; j < 5; j++ {
	// 			writer.NoteOn(wr, 50, 120)
	// 			time.Sleep(1500 * time.Millisecond)
	// 			writer.NoteOff(wr, 50)

	// 			writer.NoteOn(wr, 60, 120)
	// 			time.Sleep(1500 * time.Millisecond)
	// 			writer.NoteOff(wr, 60)

	// 			writer.NoteOn(wr, 70, 120)
	// 			time.Sleep(1500 * time.Millisecond)
	// 			writer.NoteOff(wr, 70)

	// 			writer.NoteOn(wr, 75, 120)
	// 			time.Sleep(1500 * time.Millisecond)
	// 			writer.NoteOff(wr, 75)

	// 			fmt.Println(j)
	// 		}
	// 		// writer.ControlChange(wr, 20, 20)
	// 	}
	// }

	// err = p.out.Open()
	// if err != nil {
	// 	return nil, err
	// }

	return nil
}

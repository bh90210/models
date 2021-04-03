// package variation1 demonstrates that rtmidi is indeed working. This is a very streamlined example, for full examples see variation2,3,4...
package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/midi/reader"
	"gitlab.com/gomidi/midi/writer"
	driver "gitlab.com/gomidi/rtmididrv"
)

var (
	portsMx sync.Mutex
	drv     midi.Driver

	inPorts  = map[int]midi.In{}
	outPorts = map[int]midi.Out{}
)

func init() {
	var err error
	drv, err = driver.New()
	if err != nil {
		panic("can't initialize driver")
	}
}

func main() {
	// make sure to close all open ports at the end
	defer drv.Close()

	var block = make(chan int, 10)

	go checkPorts()

	// interrupt with ctrl+c
	<-block
}

func greet(out midi.Out) {
	out.Open()

	started := time.Now()
	var ended time.Time

	// INTRO
	for loop := 0; loop <= 480; loop = loop + 2 {
		wr := writer.New(out)

		go func() {
			// kick
			writer.NoteOn(wr, 60, 100)
			time.Sleep(time.Millisecond * 10)
			writer.NoteOff(wr, 60)

			// perc
			wr.SetChannel(3)
			writer.NoteOn(wr, 70, 100)
			time.Sleep(time.Millisecond * 5)
			writer.NoteOff(wr, 70)

			time.Sleep(time.Millisecond * time.Duration(loop))

			// snare
			wr.SetChannel(1)
			writer.NoteOn(wr, 70, 100)
			time.Sleep(time.Millisecond * 10)
			writer.NoteOff(wr, 70)

			// tone
			wr.SetChannel(4)
			writer.NoteOn(wr, 70, 100)
			time.Sleep(time.Millisecond * 5)
			writer.NoteOff(wr, 70)

			time.Sleep(time.Millisecond * time.Duration(loop-1))

			// metal
			wr.SetChannel(2)
			writer.NoteOn(wr, 70, 100)
			time.Sleep(time.Millisecond * 10)
			writer.NoteOff(wr, 70)
		}()

		log.Println(loop)
		if loop == 480 {
			ended = time.Now()
			diff := started.Sub(ended)
			fmt.Println(diff)
			break
		}

		time.Sleep(time.Millisecond * 480)
	}
}

func listen(in midi.In) {
	in.Open()
	rd := reader.New(reader.NoLogger(),
		reader.Each(func(_ *reader.Position, msg midi.Message) {
			fmt.Printf("got message %s from in port %s\n", msg.String(), in.String())
		}),
	)
	rd.ListenTo(in)
}

func checkPorts() {
	var elektron int
	portsMx.Lock()
	ins, _ := drv.Ins()

	for _, in := range ins {
		// if strings.Contains(in.String(), "Client") {
		// 	continue
		// }
		// if inPorts[in.Number()] != nil {
		// 	if inPorts[in.Number()].String() != in.String() {
		// 		inPorts[in.Number()].StopListening()
		// 		inPorts[in.Number()].Close()
		// 		fmt.Printf("closing in port: [%v] %s\n", in.Number(), inPorts[in.Number()].String())
		// 		inPorts[in.Number()] = in
		// 		fmt.Printf("new in port: [%v] %s\n", in.Number(), in.String())
		// 		// if strings.Contains(in.String(), "Elektron") {
		// 		// 	fmt.Printf("Found Elektron")
		// 		// 	go listen(in)
		// 		// 	elektron = in.Number()
		// 		// }
		// 	} else {
		// 		continue
		// 	}
		// } else {
		// inPorts[in.Number()] = in
		fmt.Printf("new in port: [%v] %s\n", in.Number(), in.String())
		if strings.Contains(in.String(), "Model:Cycles") {
			fmt.Println("Found Elektron")
			go listen(in)
			elektron = in.Number()
		}
		// go listen(in)
		// }
	}

	outs, _ := drv.Outs()

	go greet(outs[elektron])
	portsMx.Unlock()
}

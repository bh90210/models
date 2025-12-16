package main

import "github.com/bh90210/models/turbo"

func main() {
	turboAlesis, err := turbo.NewProject()
	if err != nil {
		panic(err)
	}
	defer turboAlesis.Close()

	in := turboAlesis.Incoming()
	for {
		val := <-in
		println("MIDI IN:", val)
	}
}

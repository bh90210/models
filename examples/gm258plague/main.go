package main

import (
	"log"

	cycles "github.com/bh90210/elektronmodels"
)

func main() {
	gm258plague, err := cycles.NewProject()
	if err != nil {
		log.Fatal(err)
	}
	defer gm258plague.Close()

	t1intro := Intro()
	// t5intro := Intro2()

	gm258plague.Pattern(t1intro)
	// gm258plague.NewPattern(t6intro, t5intro)

	gm258plague.Loop()
	if err := gm258plague.Play(); err != nil {
		log.Println(err)
	}
}

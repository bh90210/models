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

	gm258plague.NewPattern(t1intro)
	gm258plague.NewPattern(t1intro)

	if err := gm258plague.Play(); err != nil {
		log.Println(err)
	}
}

package main

import (
	"log"
	"os"

	elektron "github.com/bh90210/elektronmodels"
)

func main() {
	gm258plague, err := elektron.NewProject(elektron.CYCLES)
	if err != nil {
		log.Fatal(err)
	}
	defer gm258plague.Close()

	intro, err := Intro(gm258plague)
	if err != nil {
		log.Fatal(err)
	}

	patterns := map[int64]*elektron.Pattern{
		1: intro,
	}

	if err := gm258plague.Play(patterns); err != nil {
		log.Println(err)
	}

	gm258plague.Close()
	os.Exit(0)
}

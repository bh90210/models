package main

import (
	"log"
	"time"

	cycles "github.com/bh90210/elektronmodels"
)

func main() {
	gm258plague, err := cycles.NewProject()
	if err != nil {
		log.Fatal(err)
	}
	defer gm258plague.Close()

	// gm258plague.Pattern(t1Intro(), t2Intro(), t3Intro(), t4Intro(), t5Intro(), t6Intro())
	gm258plague.Pattern(t1Intro())

	gm258plague.Loop()
	if err := gm258plague.Play(); err != nil {
		log.Println(err)
	}

	time.Sleep(1 * time.Second)
}

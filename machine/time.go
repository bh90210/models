package machine

import (
	"fmt"
	"time"
)

type Track struct {
	Level               uint8
	Param1              uint8
	Param2              uint8
	Param3              uint8
	Param4              uint8
	Param5              uint8
	Param6              uint8
	Param7              uint8
	Param8              uint8
	AMDepth             uint8
	AMRate              uint8
	EQFreq              uint8
	EQGain              uint8
	FilterBaseFrq       uint8
	FilterWidth         uint8
	FilterQ             uint8
	SampleRateReduction uint8
	Distortion          uint8
	Volume              uint8
	Pan                 uint8
	Delay               uint8
	Reverb              uint8
	LFOSpeed            uint8
	LFOAmount           uint8
	LFOShape            uint8
}

type Events []Event

type Event struct {
	Duration time.Duration
	Action   func(duration time.Duration, track Track)
	Track    Track
}

func (e *Events) Add(event Event) {
	*e = append(*e, event)
}

func (e *Events) Run() {
	for i, event := range *e {
		fmt.Printf("Running event %d with duration %s\n", i+1, event.Duration)
		event.Action(event.Duration, event.Track)
	}
}

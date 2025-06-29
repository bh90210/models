package machine

import "time"

type Track struct {
	Events []Event
}

type Event struct {
	Duration time.Duration
	Action   func(duration time.Duration)
}

func (t *Track) Add(event Event) {
	t.Events = append(t.Events, event)
}

func (t *Track) Run() {
	for _, event := range t.Events {
		go event.Action(event.Duration)
		time.Sleep(event.Duration)
	}
}

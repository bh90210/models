package elektronmodels

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/rakyll/portmidi"
)

type Trig struct {
	track
	note
	cctrack
	level int64
	dur   *time.Duration
	*cc
	beat     int64
	lastBeat bool
	hasCC    bool
	hasNote  bool
}

func NewTrig(beat int64) *Trig {
	return &Trig{
		beat: beat,
	}
}

func LastTrig(beat int64) *Trig {
	trig := Trig{
		beat:     beat,
		lastBeat: true,
	}
	return &trig
}

func (t *Trig) Note(key notes, level int64, dur time.Duration) {
	tn := &note{}
	tn.key = key
	tn.dur = &dur

	t.note = *tn
	t.level = level
	t.hasNote = true
}

func (t *Trig) CC(values map[Parameter]int64) {
	t.cc = &cc{
		pamVal: values,
	}
	t.hasCC = true
}

type Track struct {
	number *track
	trig   []*Trig
}

func NewTrack(trackNumber track, variadicTracks ...*Trig) *Track {
	for _, v := range variadicTracks {
		switch trackNumber {
		case T1:
			v.track = T1
			v.note.on = t1on
			v.note.off = t1off
			v.cctrack = cct1
		case T2:
			v.track = T2
			v.note.on = t2on
			v.note.off = t2off
			v.cctrack = cct2
		case T3:
			v.track = T3
			v.note.on = t3on
			v.note.off = t3off
			v.cctrack = cct3
		case T4:
			v.track = T4
			v.note.on = t4on
			v.note.off = t4off
			v.cctrack = cct4
		case T5:
			v.track = T5
			v.note.on = t5on
			v.note.off = t5off
			v.cctrack = cct5
		case T6:
			v.track = T6
			v.note.on = t6on
			v.note.off = t6off
			v.cctrack = cct6
		}
	}
	track := Track{
		number: &trackNumber,
		trig:   variadicTracks,
	}

	return &track
}

type pattern struct {
	tracks []*Track
}

type Project struct {
	pattern []*pattern
	pm      *portmidi.Stream
	mu      *sync.Mutex
	loop    bool
}

func NewProject() (*Project, error) {
	pm, mu, err := newportmidi()
	if err != nil {
		return nil, err
	}

	project := &Project{
		pm: pm,
		mu: mu,
	}
	return project, nil
}

func (p *Project) Loop() {
	p.loop = true
}

func (p *Project) Play() error {
	// calculate how many patterns are present
	totalPatters := len(p.pattern)
	fmt.Println("Total Patterns: ", totalPatters)

	// map[pattern'sBeatCounting]triggersForAll6tracks(ifPresent)
	type tri map[int][]*Trig

	// map[patternsLength]collectionOfArrangedTriggers
	type timeline map[int]*tri

	// init a key/value map acting as timeline for the song
	timlin := make(timeline)

	type patternsEnd map[int]int
	ends := make(patternsEnd)

	// range over patterns slice []pattern
	for i, pattern := range p.pattern {
		fmt.Println("Pattern: ", i)

		patternTimeline := make(tri)
		patternEndingHelper := make(patternsEnd)

		totalTracks := len(pattern.tracks)
		fmt.Println("Total Tracks: ", totalTracks)

		// range over individual pattern's tracks
		for tid, track := range pattern.tracks {
			// count total trigs in track
			trigsLength := len(track.trig)
			fmt.Println("Track: ", track.number, " Total Trigs: ", trigsLength)

			for ii := 0; ii < trigsLength; ii++ {
				patternTimeline[int(track.trig[ii].beat)] = append(patternTimeline[int(track.trig[ii].beat)], track.trig[ii])

				if track.trig[ii].lastBeat == true {
					patternEndingHelper[tid] = int(track.trig[ii].beat)
				}
			}
		}

		// compare last beats to find where pattern ends
		var longestBeat []int
		for _, v := range patternEndingHelper {
			longestBeat = append(longestBeat, int(v))
		}
		sorted := sort.IntSlice(longestBeat)

		// assign pattern end to relevant map
		last := len(sorted)
		ends[i] = sorted[last-1]
		timlin[i] = &patternTimeline
	}

	// range over timeline (map)
	// each iteration return song's individual patterns in order
	for k, pattern := range timlin {
		// var pat represents a single pattern (NewPattern())
		// beat := *pattern
		beat := *pattern

		// loop over patterns' trigs (beats) until ends[pattern]longestBeat
		lastBeat := ends[k]
		it := 0

	loop:
		for {
			totalTrigsForBeat := beat[it]
			if len(totalTrigsForBeat) != 0 {
				// total trigs for current beat range
				for _, trig := range totalTrigsForBeat {
					if trig.hasCC == true {
						p.cc(trig.cctrack, trig.cc.pamVal)
					}

					if trig.hasNote == true {
						p.note(&trig.note, trig.note.key, trig.level)
					}
				}

				fmt.Println("lastbeat", lastBeat, "current beat", it)
			}

			if it == lastBeat-1 {
				break loop
			}

			fmt.Println("Sleeping", it)
			time.Sleep(700 * time.Millisecond)
			it++
		}
	}

	return nil
}

func (c *Project) Pause() {
}

func (c *Project) Close() {
	c.pm.Close()
}

func (t *Project) NewPattern(variadicTracks ...*Track) {
	pattern := &pattern{
		tracks: variadicTracks,
	}

	t.pattern = append(t.pattern, pattern)
}

func (c *Project) note(track *note, note notes, intensity int64) {
	timer := time.NewTimer(*track.dur)

	// note on
	c.mu.Lock()
	c.pm.WriteShort(track.on.int64(), note.int64(), intensity)
	c.mu.Unlock()

	go func() {
		<-timer.C

		// note off
		c.mu.Lock()
		c.pm.WriteShort(track.off.int64(), note.int64(), intensity)
		c.mu.Unlock()
	}()
}

func (c *Project) cc(track cctrack, ccvalues map[Parameter]int64) {
	for k, v := range ccvalues {
		c.mu.Lock()
		c.pm.WriteShort(track.int64(), k.int64(), v)
		c.mu.Unlock()
	}
}

// func (c *Cycles) pc(out *portmidi.Stream, track CCtrack, parameter Parameter, value int64) {
// }

func newportmidi() (*portmidi.Stream, *sync.Mutex, error) {
	out, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
	if err != nil {
		return nil, nil, err
	}
	mutex := &sync.Mutex{}

	return out, mutex, nil
}

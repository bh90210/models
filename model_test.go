package elektronmodels

import (
	"testing"
)

var p *Project = NewProject(CYCLES)

// TestNewProject evaluates package's only constructor function.
func TestNewProject(t *testing.T) {
	if p.model != "Model:Cycles" {
		t.Errorf("want Model:Cycles got %s", p.model)
	}

	// Without p.InitPattern(0)
	pat1 := p.Pattern[0]
	if pat1 != nil {
		t.Errorf("got %v want nil", pat1)
	}

	pat2 := new(pattern)
	if pat2.tempo != 0 {
		t.Errorf("got %v want 0", pat2)
	}
}

// TestInitPattern evaluates pattern initialization.
func TestInitPattern(t *testing.T) {
	t.Parallel()

	p.InitPattern(0)

	var wantScale = &scale{
		PTN, 15, 4.0, 0,
	}
	// check Scale defaults are set
	if *p.Pattern[0].T1.Scale != *wantScale {
		t.Errorf("got %v want %v", *p.Pattern[0].T1.Scale, *wantScale)
	}

	// are presets filled alright?
	wantPreset := make(map[Parameter]int8)
	wantPreset[COLOR] = 10
	if p.Pattern[0].T1.Preset[COLOR] != wantPreset[COLOR] {
		t.Errorf("got %v want %v", p.Pattern[0].T1.Preset, wantPreset)
	}

	// nil trig
	wantMap := make(map[int]*trig)
	if len(p.Pattern[0].T1.Trig) != len(wantMap) {
		t.Errorf("got %v want %v", p.Pattern[0].T1.Trig, wantMap)
	}
}

// TestInitTrig evaluates trigger initialization.
func TestInitTrig(t *testing.T) {
	// t.Parallel()

	p.InitPattern(1)
	p.Pattern[1].T1.InitTrig(0)

	want := &trig{&note{C4, 4, 126}, &Lock{}}
	if *p.Pattern[1].T1.Trig[0].Note != *want.Note {
		t.Errorf("got %v want %v", *p.Pattern[1].T1.Trig[0], *want)
	}
}

package main

import (
	"github.com/bh90210/models/midicom"
	"github.com/bh90210/models/nymphes"
	"github.com/bh90210/models/pattern"
)

func main() {
	ns, err := nymphes.NewProject()
	if err != nil {
		panic(err)
	}
	defer ns.Close()

	in := ns.Incoming()
	go func() {
		for {
			val := <-in
			println("MIDI IN:", val)
		}
	}()

	notes1 := []pattern.Note{
		{Note: 50, Duration: 500, Velocity: 100, CC: baseNoteCC()},
		{Note: 50 + midicom.Note(pattern.Major3rd), Duration: 500, Velocity: 100},
		{Note: 50 + midicom.Note(pattern.Perfect5th), Duration: 500, Velocity: 100},
		{Note: 50 + midicom.Note(pattern.Major7th), Duration: 500, Velocity: 100},
	}

	pat1 := pattern.Pattern{
		Midicom: ns,
		Notes:   notes1,
		Channel: nymphes.Channel,
		Meta: pattern.Meta{
			Synth: nymphes.Nymphes,
			Part:  "voice1-start",
		},
	}

	pat2 := pat1.Shift(pattern.Perfect5th)

	allVoices := make(map[int][]pattern.Pattern)

	allVoices[0] = []pattern.Pattern{pat1}
	allVoices[1] = []pattern.Pattern{pat2}

	p := pattern.NewPrint(allVoices)
	p.Print(pattern.Voice)

	err = pattern.Play(allVoices)
	if err != nil {
		panic(err)
	}
}

func baseNoteCC() map[midicom.Parameter]int8 {
	cc := make(map[midicom.Parameter]int8)

	cc[midicom.Parameter(nymphes.OSCLevel)] = 100
	cc[midicom.Parameter(nymphes.SubLebel)] = 80
	cc[midicom.Parameter(nymphes.NoiseLevel)] = 60
	cc[midicom.Parameter(nymphes.PulseWidth)] = 70
	cc[midicom.Parameter(nymphes.LDOPitchDepth)] = 50
	cc[midicom.Parameter(nymphes.EGPitchDepth)] = 40
	cc[midicom.Parameter(nymphes.Detune)] = 20
	cc[midicom.Parameter(nymphes.ChordSelector)] = 10
	cc[midicom.Parameter(nymphes.PlayMode)] = 2 // Uni-B

	cc[midicom.Parameter(nymphes.LFO1Rate)] = 80
	cc[midicom.Parameter(nymphes.LFO1Wave)] = 2
	cc[midicom.Parameter(nymphes.LFO1Delay)] = 30
	cc[midicom.Parameter(nymphes.LFO1Fade)] = 40
	cc[midicom.Parameter(nymphes.LFO1Type)] = 1    // Low
	cc[midicom.Parameter(nymphes.LFO1KeySync)] = 1 // On
	cc[midicom.Parameter(nymphes.LFO2Rate)] = 70
	cc[midicom.Parameter(nymphes.LFO2Wave)] = 3
	cc[midicom.Parameter(nymphes.LFO2Delay)] = 20
	cc[midicom.Parameter(nymphes.LFO2Fade)] = 30
	cc[midicom.Parameter(nymphes.LFO2Type)] = 0    // BPM
	cc[midicom.Parameter(nymphes.LFO2KeySync)] = 0 // Off

	cc[midicom.Parameter(nymphes.ModSourceSelector)] = 1 // ModWheel

	cc[midicom.Parameter(nymphes.ModSourceOSCWaveDepth)] = 20
	cc[midicom.Parameter(nymphes.ModSourceOSCLevelDepth)] = 30
	cc[midicom.Parameter(nymphes.ModSourceSubLevelDepth)] = 40
	cc[midicom.Parameter(nymphes.ModSourceNoiseLevelDepth)] = 50
	cc[midicom.Parameter(nymphes.ModSourceLFOPitchDepthDepth)] = 60
	cc[midicom.Parameter(nymphes.ModSourcePulseWidthDepth)] = 70
	cc[midicom.Parameter(nymphes.ModSourceGlideDepth)] = 80

	cc[midicom.Parameter(nymphes.ModSourceDetuneDepth)] = 90
	cc[midicom.Parameter(nymphes.ModSourceChordSelectorDepth)] = 100
	cc[midicom.Parameter(nymphes.ModSourceEGPitchDepthDepth)] = 110
	cc[midicom.Parameter(nymphes.ModSourceLPFCutoffDepth)] = 127
	cc[midicom.Parameter(nymphes.ModSourceResonanceDepth)] = 120
	cc[midicom.Parameter(nymphes.ModSourceLPFEGDepthDepth)] = 110
	cc[midicom.Parameter(nymphes.ModSourceHPFCutoffDepth)] = 100
	cc[midicom.Parameter(nymphes.ModSourceLPFTrackingDepth)] = 90
	cc[midicom.Parameter(nymphes.ModSourceLPFCutoffLFODepthDepth)] = 80
	cc[midicom.Parameter(nymphes.ModSourceFilterEGAttackDepth)] = 70
	cc[midicom.Parameter(nymphes.ModSourceFilterEGDecayDepth)] = 60
	cc[midicom.Parameter(nymphes.ModSourceFilterEGSustainDepth)] = 50
	cc[midicom.Parameter(nymphes.ModSourceFilterEGReleaseDepth)] = 40

	cc[midicom.Parameter(nymphes.ModSourceFilterAMPAttackDepth)] = 30
	cc[midicom.Parameter(nymphes.ModSourceFilterAMPDecayDepth)] = 20
	cc[midicom.Parameter(nymphes.ModSourceFilterAMPSustainDepth)] = 10
	cc[midicom.Parameter(nymphes.ModSourceFilterAMPReleaseDepth)] = 0
	cc[midicom.Parameter(nymphes.ModSourceLFO1RateDepth)] = 15
	cc[midicom.Parameter(nymphes.ModLFO1WaveDepth)] = 25
	cc[midicom.Parameter(nymphes.ModSourceLFO1DelayDepth)] = 35
	cc[midicom.Parameter(nymphes.ModSourceLFO1FadeDepth)] = 45
	cc[midicom.Parameter(nymphes.ModSourceLFO2RateDepth)] = 55
	cc[midicom.Parameter(nymphes.ModSourceLFO2WaveDepth)] = 65
	cc[midicom.Parameter(nymphes.ModSourceLFO2DelayDepth)] = 75
	cc[midicom.Parameter(nymphes.ModSourceLFO2FadeDepth)] = 85

	cc[midicom.Parameter(nymphes.ModSourceReverbSizeDepth)] = 40
	cc[midicom.Parameter(nymphes.ModSourceReverbDecayDepth)] = 50
	cc[midicom.Parameter(nymphes.ModSourceReverbFilterDepth)] = 60
	cc[midicom.Parameter(nymphes.ModSourceReverbMixDepth)] = 70

	cc[midicom.Parameter(nymphes.SustainPedal)] = 127
	cc[midicom.Parameter(nymphes.Legato)] = 64

	cc[midicom.Parameter(nymphes.OSCWave)] = 2
	cc[midicom.Parameter(nymphes.Resonance)] = 80
	cc[midicom.Parameter(nymphes.AMPEGRelease)] = 70
	cc[midicom.Parameter(nymphes.AMPEGAttack)] = 60
	cc[midicom.Parameter(nymphes.LPFCutoff)] = 90
	cc[midicom.Parameter(nymphes.ReverbSize)] = 50
	cc[midicom.Parameter(nymphes.ReverbDecay)] = 40
	cc[midicom.Parameter(nymphes.ReverbFilter)] = 30
	cc[midicom.Parameter(nymphes.ReverbMix)] = 20
	cc[midicom.Parameter(nymphes.FilterEGAttack)] = 10
	cc[midicom.Parameter(nymphes.FilterEGDecay)] = 20
	cc[midicom.Parameter(nymphes.HPFCutoff)] = 80
	cc[midicom.Parameter(nymphes.FilterEGSustain)] = 80
	cc[midicom.Parameter(nymphes.FilterEGRelease)] = 60
	cc[midicom.Parameter(nymphes.AMPEGDecay)] = 90
	cc[midicom.Parameter(nymphes.AMPEGSustain)] = 90

	return cc
}

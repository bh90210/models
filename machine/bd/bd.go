package bd

import "github.com/bh90210/models/machine"

const (
	TRACK   = 0
	CHANNEL = 0  // Default channel for BD
	NOTE    = 36 // BD
)

const (
	LEVEL               machine.Parameter = 8
	MUTE                machine.Parameter = 12
	PARAM1              machine.Parameter = 16
	PARAM2              machine.Parameter = 17
	PARAM3              machine.Parameter = 18
	PARAM4              machine.Parameter = 19
	PARAM5              machine.Parameter = 20
	PARAM6              machine.Parameter = 21
	PARAM7              machine.Parameter = 22
	PARAM8              machine.Parameter = 23
	AMDEPTH             machine.Parameter = 24
	AMRATE              machine.Parameter = 25
	EQFREQ              machine.Parameter = 26
	EQGAIN              machine.Parameter = 27
	FILTERBASEFRQ       machine.Parameter = 28
	FILTERWIDTH         machine.Parameter = 29
	FILTERQ             machine.Parameter = 30
	SAMPLERATEREDUCTION machine.Parameter = 31
	DISTORTION          machine.Parameter = 32
	VOLUME              machine.Parameter = 33
	PAN                 machine.Parameter = 34
	DELAY               machine.Parameter = 35
	REVERB              machine.Parameter = 36
	LFOSPEED            machine.Parameter = 37
	LFOAMOUNT           machine.Parameter = 38
	LFOSHAPE            machine.Parameter = 39
)

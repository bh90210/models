package sd

import "github.com/bh90210/models/machine"

const (
	TRACK   = 1
	CHANNEL = 0
	NOTE    = 38 // SD
)

const (
	LEVEL               machine.Parameter = 9
	MUTE                machine.Parameter = 13
	PARAM1              machine.Parameter = 40
	PARAM2              machine.Parameter = 41
	PARAM3              machine.Parameter = 42
	PARAM4              machine.Parameter = 43
	PARAM5              machine.Parameter = 44
	PARAM6              machine.Parameter = 45
	PARAM7              machine.Parameter = 46
	PARAM8              machine.Parameter = 47
	AMDEPTH             machine.Parameter = 48
	AMRATE              machine.Parameter = 49
	EQFREQ              machine.Parameter = 50
	EQGAIN              machine.Parameter = 51
	FILTERBASEFRQ       machine.Parameter = 52
	FILTERWIDTH         machine.Parameter = 53
	FILTERQ             machine.Parameter = 54
	SAMPLERATEREDUCTION machine.Parameter = 55
	DISTORTION          machine.Parameter = 56
	VOLUME              machine.Parameter = 57
	PAN                 machine.Parameter = 58
	DELAY               machine.Parameter = 59
	REVERB              machine.Parameter = 60
	LFOSPEED            machine.Parameter = 61
	LFOAMOUNT           machine.Parameter = 62
	LFOSHAPE            machine.Parameter = 63
)

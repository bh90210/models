package elektronmodels

//
// default presets
//

func defaultPreset(id voice) preset {
	switch id {
	case T1:
		return PT1()
	case T2:
		return PT2()
	case T3:
		return PT3()
	case T4:
		return PT4()
	case T5:
		return PT5()
	case T6:
		return PT6()

	}

	return nil
}

func PT1() preset {
	p := make(map[Parameter]int8)
	p[MACHINE] = int8(KICK)
	return p
}

func PT2() preset {
	p := PT1()
	p[MACHINE] = int8(SNARE)
	return p
}

func PT3() preset {
	p := PT1()
	p[MACHINE] = int8(METAL)
	return p
}

func PT4() preset {
	p := PT1()
	p[MACHINE] = int8(PERC)
	return p
}

func PT5() preset {
	p := PT1()
	p[MACHINE] = int8(TONE)
	return p
}

func PT6() preset {
	p := PT1()
	p[MACHINE] = int8(CHORD)
	return p
}

package elektronmodels

//
// default presets
//

func defaultPreset(id voice) preset {
	switch id {
	case T1:
		return pT1()
	case T2:
		return pT2()
	case T3:
		return pT3()
	case T4:
		return pT4()
	case T5:
		return pT5()
	case T6:
		return pT6()

	}

	return nil
}

func pT1() preset {
	p := make(map[Parameter]int8)
	p[COLOR] = 10
	return p
}

func pT2() preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func pT3() preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func pT4() preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func pT5() preset {
	d := make(map[Parameter]int8)
	d[COLOR] = 10
	return d
}

func pT6() preset {
	p := make(preset)
	p[COLOR] = 10
	return p
}

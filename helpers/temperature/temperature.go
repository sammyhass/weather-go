package temperature

type Unit int

const (
	F Unit = 0
	C Unit = 1
	K Unit = 2
)

type Temp struct {
	Val float64
	Unit Unit
}

func (t Temp) To(unit Unit) Temp {
	var resTemp Temp
	resTemp.Unit = unit
	switch t.Unit {
	case K:
		if unit == F { resTemp.Val = kToF(t.Val) } else { resTemp.Val = kToC(t.Val) }
	case C:
		if unit == K { resTemp.Val = cToK(t.Val) } else { resTemp.Val = cToF(t.Val) }
	case F:
		if unit == C { resTemp.Val = fToC(t.Val) } else { resTemp.Val = fToK(t.Val) }
	}
	return resTemp
}

func cToK(c float64) float64 {
	return c + 273.15
}

func cToF(c float64) float64 {
	return c * 9/5 + 32
}

func fToC(f float64) float64 {
	return (f - 32) * 5/9
}

func fToK(f float64) float64 {
	return fToC(f) + 273.15
}

func kToC(k float64) float64 {
	return k - 273.15
}

func kToF(k float64) float64 {
	return cToF(kToC(k))
}
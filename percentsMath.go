package MathLib

func reducingFloatByPercent(a float64, pUser float64) float64 {
	p := pUser / 100
	percentOfA := p * a
	reduceF := a - percentOfA
	return reduceF
}

func reducingIntByPercent(a int, pUser int) int {
	p := pUser / 100
	percentOfA := p * a
	reduceI := a - percentOfA
	return reduceI
}

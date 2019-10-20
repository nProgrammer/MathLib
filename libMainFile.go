package MathLib

func AddFloat(a float64, b float64) float64 {
	return a + b
}

func AddInt(a int, b int) int {
	return a + b
}

func MultiplyFloat(a float64, b float64) float64 {
	return a * b
}

func MultiplyInt(a int, b int) int {
	return a * b
}

func DivideFloat(a float64, b float64) float64 {
	return a / b
}

func DivideInt(a int, b int) int {
	return a / b
}

func IsIntDivisible(a int, b int) bool {
	answerID := false
	m := a % b
	if m == 0 {
		answerID = true
	}
	return answerID
}

func IsIntEven(a int) bool {
	answerIE := false
	num := a % 2
	if num == 0 {
		answerIE = true
	}
	return answerIE
}

func ReducingFloatByPercent(a float64, pUser float64) float64 {
	p := pUser / 100
	percentOfA := p * a
	reduceF := a - percentOfA
	return reduceF
}

func ReducingIntByPercent(a int, pUser int) int {
	p := pUser / 100
	percentOfA := p * a
	reduceI := a - percentOfA
	return reduceI
}

func SqrtFloat(a float64) float64 {
	return a * a
}

func SqrtInt(a int) int {
	return a * a
}

func ToThePowerInt(a int, n int) int {
	score := a
	for i := 0; i < n; i++ {
		score = score * a
	}
	return score
}

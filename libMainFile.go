package MathLib

type ConstantMathValues struct {
	PI float64
}

func AddFloat(a float64, b float64) float64 {
	return a + b
}

func PI() float64 {
	constantMathValues := ConstantMathValues{}
	constantMathValues.PI = 3.14
	return constantMathValues.PI
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
	scorePI := a
	if n >= 1 {
		for i := 1; i < n; i++ {
			scorePI = scorePI * a
		}
	} else if n == 0 {
		scorePI = 1
	} else {
		scorePI = 0
	}
	return scorePI
}

func ToThePowerFloat(a float64, n float64) float64 {
	scorePF := a
	if n >= 1 {
		for i := 1; i < n; i++ {
			scorePF = scorePF * a
		}
	} else if n == 0 {
		scorePF = 1
	} else {
		scorePF = 0
	}
	return scorePF
}

package MathLib

type constantMathValues struct {
	pi           float64
	goldDivision float64
}

func GoldDivision() float64 {
	constantMathValues := constantMathValues{}
	constantMathValues.goldDivision = 1.618033988749894848204586834366
	return constantMathValues.goldDivision
}

func PI() float64 {
	constantMathValues := constantMathValues{}
	constantMathValues.pi = 3.14159265359
	return constantMathValues.pi
}

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
		for i := 1.0; i < n; i++ {
			scorePF = scorePF * a
		}
	} else if n == 0 {
		scorePF = 1
	} else {
		scorePF = 0
	}
	return scorePF
}

func CircleField(r float64) float64 {
	return PI() * SqrtFloat(r)
}

func RectangleFieldFloat(a float64, b float64) float64 {
	return a * b
}

func RectangleFieldInt(a int, b int) int {
	return a * b
}

func SquareFieldFloat(a float64) float64 {
	return SqrtFloat(a)
}

func SquareFieldInt(a int) int {
	return SqrtInt(a)
}

func TriangleFieldFloat(a float64, h float64) float64 {
	return 0.5 * (a * h)
}

func HexagonFieldFloat(a float64) float64 {
	return 6 * SqrtFloat(a)
}

func HexagonFieldInt(a int) int {
	return 6 * SqrtInt(a)
}

func CorrectCuboidFieldFloat(a float64, h float64) float64 {
	return 2*SqrtFloat(a) + 4*MultiplyFloat(a, h)
}

func CorrectCuboidFieldInt(a int, h int) int {
	return 2*SqrtInt(a) + 4*MultiplyInt(a, h)
}

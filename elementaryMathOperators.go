package MathLib

func addFloat(a float64, b float64) float64 {
	return a + b
}

func addInt(a int, b int) int {
	return a + b
}

func multiplyFloat(a float64, b float64) float64 {
	return a * b
}

func MultiplyInt(a int, b int) int {
	return a * b
}

func divideFloat(a float64, b float64) float64 {
	return a / b
}

func divideInt(a int, b int) int {
	return a / b
}

func isIntDivisible(a int, b int) bool {
	answerID := false
	m := a % b
	if m == 0 {
		answerID = true
	}
	return answerID
}

func isIntEven(a int) bool {
	answerIE := false
	num := a % 2
	if num == 0 {
		answerIE = true
	}
	return answerIE
}

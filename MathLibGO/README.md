# ***Math GO lang library - MathLib 1.1.1***
*Math lib is simple to use GOlang library*
*To import this library you should write:*
   ```go
 import (
 	    "github.com/nProgrammer/MathLib/MathLibGO"
 )
```

*Now you can use our 26 functions:*
```go
	a := MathLibGO.AddInt(10, 11) // this function is used for adding ints
	b := MathLibGO.AddFloat(10.2, 11.1) // this function is used for adding floats
	c := MathLibGO.MultiplyFloat(10.12, 11.1) // this function is used for multiplication floats
	d := MathLibGO.MultiplyInt(10, 11) // this function is used for multiplication ints
	e := MathLibGO.DivideFloat(10.12, 11.213) // this function is used for dividing floats
	f := MathLibGO.DivideInt(10, 11) // this function is used for dividing ints
	g := MathLibGO.IsIntDivisible(10, 11) // this function is used to check if int is divisible by another int
	h := MathLibGO.IsIntEven(10) // this function is used to check if int is even
	i := MathLibGO.ReducingFloatByPercent(100, 23) // this function is used to reducing float by some percent
	j := MathLibGO.ReducingIntByPercent(100, 23) // this function is used to reducing int by some percent
	k := MathLibGO.SqrtFloat(11.1) // this function is used to make sqrt of float
	l := MathLibGO.SqrtInt(11) // this function is used to make sqrt of int
	m := MathLibGO.ToThePowerInt(11, 2) // this function is used to calculate the power of any exponent of any int
	n := MathLibGO.ToThePowerFloat(11.0, 2.0) // this function is used to calculate the power of any exponent of any float
	p := MathLibGO.PI() // this function return number of PI (3.14159265359)
	r := MathLibGO.CircleField(12.0) // this function return value of circle field
	s := MathLibGO.RectangleFieldFloat(10.1, 1.0) // this function return value of rectangle field
	t := MathLibGO.RectangleFieldInt(10, 1) // this function return value of rectangle field
	u := MathLibGO.SquareFieldFloat(12.1) // this function return value of square field
	w := MathLibGO.SquareFieldInt(12) // this function return value of square field
	y := MathLibGO.TriangleFieldFloat(1.2, 12.2) // this function return value of triangle field
	z := MathLibGO.HexagonFieldFloat(12.1) // this function return value of hexagon field
	aa := MathLibGO.HexagonFieldInt(12) // this function return value of hexagon field
	ab := MathLibGO.CorrectCuboidFieldFloat(11.1, 0.1) // this function return value of correct cuboid field
	ac := MathLibGO.CorrectCuboidFieldInt(11, 1) // this function return value of correct cuboid field
        ad := MathLibGO.GoldDivision() // this function return the value of Gold Division
```

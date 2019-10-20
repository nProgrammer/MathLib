# ***Math GO lang library - MathLib 1.0***
*Math lib is simple to use GOlang library*
*To import this library you should write:*
   ```go
 import (
 	    "github.com/nProgrammer/MathLib/MathLib"
 )
```

*Now you can use our 15 functions:*
```go
	a := MathLib.AddInt(10, 11) // this function is used for adding ints
	b := MathLib.AddFloat(10.2, 11.1) // this function is used for adding floats
	c := MathLib.MultiplyFloat(10.12, 11.1) // this function is used for multiplication floats
	d := MathLib.MultiplyInt(10, 11) // this function is used for multiplication ints
	e := MathLib.DivideFloat(10.12, 11.213) // this function is used for dividing floats
	f := MathLib.DivideInt(10, 11) // this function is used for dividing ints
	g := MathLib.IsIntDivisible(10, 11) // this function is used to check if int is divisible by another int
	h := MathLib.IsIntEven(10, 11) // this function is used to check if int is even
	i := MathLib.ReducingFloatByPercent(100, 23) // this function is used to reducing float by some percent
	j := MathLib.ReducingIntByPercent(100, 23) // this function is used to reducing int by some percent
        k := MathLib.SqrtFloat(11.1) // this function is used to make sqrt of float
        l := MathLib.SqrtInt(11) // this function is used to make sqrt of int
        m := MathLib.ToThePowerInt(11, 2) // this function is used to calculate the power of any exponent of any int
        n := MathLib.ToThePowerFloat(11.0, 2.0) // this function is used to calculate the power of any exponent of any float
        p := MathLib.PI() // this function return number of PI (3.14159265359)
        r := MathLib.CircleField(12.0) // this function return value of circle field
        s := MathLib.RectangleFieldFloat(10.1, 1.0) // this function return value of rectangle field
        t := MathLib.RectangleFieldInt(10, 1) // this function return value of rectangle field
        u := MathLib.SquareFieldFloat(12.1, 123.1) // this function return value of square field
        w := MathLib.SquareFieldInt(12, 123) // this function return value of square field
        y := MathLib.TriangleFieldInt(1, 23) // this function return value of triangle field
        z := MathLib.TriangleFieldFloat(1.2, 12.2) // this function return value of triangle field
        aa := MathLib.HexagonFieldFloat(12.1) // this function return value of hexagon field
        ab := MathLib.HexagonFieldInt(12) // this function return value of hexagon field
        ac := MathLib.CorrectCuboidFieldFloat(11.1, 0.1) // this function return value of correct cuboid field
        ad := MathLib.CorrectCuboidFieldInt(11, 1) // this function return value of correct cuboid field
```

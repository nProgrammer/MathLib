# ***Math GO lang library - MathLib***
*Math lib is simple to use GOlang library*
*To import this library you should write:*
   ```go
 import (
 	"github.com/nProgrammer/MathLib"
    	_ "github.com/nProgrammer/MathLib"
    )
```

*Now you can use our 10 functions:*
```go
	a := MathLib.addInt(10, 11) // this function is used for adding ints
	b := MathLib.addFloat(10.2, 11.1) // this function is used for adding floats
	c := MathLib.multiplyFloat(10.12, 11.1) // this function is used for multiplication floats
	d := MathLib.multiplyInt(10, 11) // this function is used for multiplication ints
	e := MathLib.divideFloat(10.12, 11.213) // this function is used for dividing floats
	f := MathLib.divideInt(10, 11) // this function is used for dividing ints
	g := MathLib.isIntDivisible(10, 11) // this function is used to check if int is divisible by another int
	h := MathLib.isIntEven(10, 11) // this function is used to check if int is even
	i := MathLib.reducingFloatByPercent(100, 23) // this function is used to reducing float by some percent
	j := MathLib.reducingIntByPercent(100, 23) // this function is used to reducing int by some percent
```

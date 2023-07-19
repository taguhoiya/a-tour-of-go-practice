//1. Define an interface named `Abser` that requires a single method `Abs() float64`. This method should return the absolute value of a number.
//2. Create a new type `MyFloat` based on the `float64` type.
//3. Create a struct type `Vertex` that has two float64 fields `X` and `Y`.
//4. Implement the `Abs()` method for the `MyFloat` type. The method should return the absolute value of the `MyFloat`.
//5. Implement the `Abs()` method for the `Vertex` type. The method should return the Euclidean distance of the vertex from the origin (0,0).
//6. In the `main()` function, create a variable `a` of type `Abser`.
//7. Initialize `f` as a `MyFloat` with the value of the negative square root of 2.
//8. Initialize `v` as a `Vertex` with the X and Y values of 3 and 4, respectively.
//9. Assign `f` to `a`.
//10. Assign the address of `v` to `a`.
//11. Uncomment the line `// a = v`.
//12. Print the result of calling the `Abs()` method on `a`.

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
}

// interfaceで共通の振る舞いを定義する

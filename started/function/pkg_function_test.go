package function

import (
	"fmt"
	"math"
	"testing"
)

func TestVertex_Abs(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func TestVertex_Abs2(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

func TestVertex_Abs3(t *testing.T) {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func TestFactorial(t *testing.T) {
	fmt.Println(Factorial(5))
}

func TestVertex_ScalePointer(t *testing.T) {
	v := Vertex{3, 4}
	v.ScalePointer(10)
	fmt.Println(v.Abs()) // 50
}

func TestVertex_ScaleValue(t *testing.T) {
	v := Vertex{3, 4}
	v.ScaleValue(10)
	fmt.Println(v.Abs()) // 5
}

func TestScalePointer1(t *testing.T) {
	v := Vertex{3, 4}
	ScalePointer1(&v, 10)
	fmt.Println(Abs(v)) // 50
}

func TestScaleValue1(t *testing.T) {
	v := Vertex{3, 4}
	ScaleValue1(v, 10)
	fmt.Println(Abs(v)) // 5
}

// 方法与指针重定向
func TestScaleFunc(t *testing.T) {
	v := Vertex{3, 4}
	// 即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。
	// 也就是说，由于 ScalePointer 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(2) 解释为 (&v).Scale(2)。
	v.ScalePointer(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.ScalePointer(3)
	ScaleFunc(p, 7)

	fmt.Println(v, p)  // {60 80} &{84 63}
	fmt.Println(v, *p) // {60 80} {84 63}
}

//

func TestAbsFunc(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())    // 5
	fmt.Println(AbsFunc(v)) // 5
	//fmt.Println(AbsFunc(&v)) // 编译错误！

	p := &Vertex{4, 3}
	// 以值为接收者的方法被调用时，接收者既能为值又能为指针
	// 这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
	fmt.Println(p.Abs())     // 5
	fmt.Println(AbsFunc(*p)) // 5
}

func TestMethodsWithPointerReceivers(t *testing.T) {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	// Before scaling: &{X:3 Y:4}, Abs: 5
	v.ScalePointer(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	// After scaling: &{X:15 Y:20}, Abs: 25
}

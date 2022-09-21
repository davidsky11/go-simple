package function

import "math"

// 方法就是一类带特殊的 接收者 参数的函数。
//
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

type Vertex struct {
	X, Y float64
}

// Abs 方法
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 记住：方法只是个带接收者参数的函数。

// Abs 方法即函数，调用方式与上一个 Abs() 不同
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 为非结构体类型声明方法
// 只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
// 就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

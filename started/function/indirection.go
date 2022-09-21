package function

import "math"

// 方法与指针重定向

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 接受一个值作为参数的函数必须接受一个指定类型的值

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

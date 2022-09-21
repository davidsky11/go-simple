package errors

// Go 程序使用 error 值来表示错误状态。
// 与 fmt.Stringer 类似，error 类型是一个内建接口：
//	type error interface {
//  	  Error() string
//	}
//
// （与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）
//
// 通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。
// 		error 为 nil 时表示成功；非 nil 的 error 表示失败。

import (
	"fmt"
	"math"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func ErrorRun() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func ErrorSqrt(x float64) (float64, error) {
	return 0, nil
}

// ************************************************************************************

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
	}
	return ""
}

func ErrorNegativeSqrtRun(x float64) (float64, ErrNegativeSqrt) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	var ens ErrNegativeSqrt
	return math.Sqrt(x), ens
}

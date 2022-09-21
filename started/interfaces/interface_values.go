package interfaces

import (
	"fmt"
)

type Inter interface {
	M()
}

type Tct struct {
	S string
}

// **************************************************************************************
// 即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
// 在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。
// **************************************************************************************

func (t *Tct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type Fat float64

func (f Fat) M() {
	fmt.Println(f)
}

func Describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

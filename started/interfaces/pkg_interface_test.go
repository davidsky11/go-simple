package interfaces

import (
	"fmt"
	"github.com/katyusha/go-simple/started/function"
	"math"
	"testing"
)

func TestInterface(t *testing.T) {
	var a Abser
	f := function.MyFloat(-math.Sqrt2)
	v := function.Vertex{X: 3, Y: 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	fmt.Println(a.Abs())
}

func TestT_M(t *testing.T) {
	var i I = T{"hello"}
	i.M()
}

func TestDescribe(t *testing.T) {
	var inter Inter

	inter = &Tct{"Hello"}
	Describe(inter) // (&{Hello}, *interfaces.Tct)
	inter.M()       // Hello

	inter = Fat(math.Pi)
	Describe(inter) // (3.141592653589793, interfaces.Fat)
	inter.M()       // 3.141592653589793
}

func TestDescribeWithNil(t *testing.T) {
	var inter Inter

	var tct *Tct
	inter = tct
	Describe(inter) // (<nil>, *interfaces.Tct)
	inter.M()       // <nil>
	// 如果 Tct#M() 不特殊处理的话，panic: using nil *T pointer

	inter = &Tct{"hello go"}
	Describe(inter) // (&{hello go}, *interfaces.T)
	inter.M()       // hello go
}

func TestEmptyInterfacePrint(t *testing.T) {
	var i interface{}
	EmptyInterfacePrint(i) // (<nil>, <nil>)

	i = 42
	EmptyInterfacePrint(i) // (42, int)

	i = "hello"
	EmptyInterfacePrint(i) // (hello, string)
}

func TestTypeAssert(t *testing.T) {
	TypeAssert()
}

package function

import (
	"fmt"
	"testing"
)

func TestVertex_Abs(t *testing.T) {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

func TestFactorial(t *testing.T) {
	fmt.Println(Factorial(5))
}

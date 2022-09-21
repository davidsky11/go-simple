package errors

import (
	"fmt"
	"testing"
)

func TestErrorRun(t *testing.T) {
	if err := ErrorRun(); err != nil {
		fmt.Println(err)
	}
}

func TestErrorSqrt(t *testing.T) {
	fmt.Println(ErrorSqrt(2))
	fmt.Println(ErrorSqrt(-2))
}

func TestErrorNegativeSqrtRun(t *testing.T) {
	_, ens := ErrorNegativeSqrtRun(-2)
	fmt.Println(ens.Error())

	fmt.Println(ErrNegativeSqrt(-2).Error())
}

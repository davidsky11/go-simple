package grammar

import (
	"fmt"
	"testing"
)

func TestPrintIPAddr(t *testing.T) {
	PrintIPAddr()
}

func TestRoom_String(t *testing.T) {
	a := Room{"Dent", 42}
	z := Room{"Zaphod", 9001}
	fmt.Println(a, z)
}

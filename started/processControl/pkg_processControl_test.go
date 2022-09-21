package processControl

import (
	"github.com/katyusha/go-simple/started/function"
	"testing"
)

func TestTypeSwitch(t *testing.T) {
	TypeSwitch(21)
	TypeSwitch("hello")
	TypeSwitch(true)
	TypeSwitch(3.1415926)
	TypeSwitch(function.Vertex{3, 4})
}

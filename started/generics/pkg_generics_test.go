package generics

import "testing"

func TestNonGenericSum(t *testing.T) {
	NonGenericSum()
}

func TestGenericSum(t *testing.T) {
	GenericSum()
}

func TestGenericSumWithConstraint(t *testing.T) {
	GenericSumWithConstraint()
}

func TestGenericForReceiver(t *testing.T) {
	GenericForReceiver()
}

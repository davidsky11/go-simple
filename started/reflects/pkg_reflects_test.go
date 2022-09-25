package reflects

import "testing"

func TestReflect1(t *testing.T) {
	Reflect1()
}

func TestReflect2(t *testing.T) {
	Reflect2()
}

func TestReflect3(t *testing.T) {
	Reflect3()
}

func TestDynamicReflect(t *testing.T) {
	DynamicReflect()
}

func TestReflectValue1(t *testing.T) {
	ReflectValue1()
}

func TestReflectValue12(t *testing.T) {
	Reflectvalue2()
}

func TestCustomReflectFunc2DuckingTypeValue(t *testing.T) {
	CustomReflectFunc2DuckingTypeValue()
}

func TestMethodByNameCall(t *testing.T) {
	MethodByNameCall()
}

func TestReflectSetMapIndex(t *testing.T) {
	ReflectSetMapIndex()
}

func TestReflectChannelValue(t *testing.T) {
	ReflectChannelValue()
}

func TestReflectSelect(t *testing.T) {
	ReflectSelect()
}

func TestReflectZeroValue(t *testing.T) {
	ReflectZeroValue()
}

func TestReflectInterfaceValue(t *testing.T) {
	ReflectInterfaceValue()
}

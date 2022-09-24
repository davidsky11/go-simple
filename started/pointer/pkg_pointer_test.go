package pointer

import (
	"fmt"
	"strings"
	"testing"
)

func TestUnsafePointer(t *testing.T) {
	UnsafePointer()
}

func TestUnsafePointerDanger(t *testing.T) {
	UnsafePointerDanger()
}

func TestFoo(t *testing.T) {
	Foo()
}

func TestBar(t *testing.T) {
	Bar()
}

func TestPointer4SafePointer(t *testing.T) {
	Pointer4SafePointer()
}

func TestUnsafeLogPrint(t *testing.T) {
	UnsafeLogPrint()
}

func TestUnsafePointTransform(t *testing.T) {
	UnsafePointTransform()
}

func TestUnsafePointTransformDanger(t *testing.T) {
	UnsafePointTransformDanger()
}

func TestReflectStringHeaderPointer(t *testing.T) {
	ReflectStringHeaderPointer()
}

func TestReflectSliceHeaderPointer(t *testing.T) {
	ReflectSliceHeaderPointer()
}

func TestString2ByteSlice(t *testing.T) {
	//str := "Goland"
	// 对于官方标准编译器来说，上面这行将使str中的字节开辟在不可修改内存区。所以这里我们使用下面这行。
	str := strings.Join([]string{"Go", "land"}, "")
	s := String2ByteSlice(str)
	fmt.Printf("%s\n", s) // Goland
	s[5] = 'g'
	fmt.Println(str) // Golang

}

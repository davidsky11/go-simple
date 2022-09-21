package pointer

import (
	"fmt"
	"unsafe"
)

var X struct {
	a int64
	b bool
	c string
}

type Person struct {
	name string
	age  int
}

func UnsafePointer() {
	// 此函数用来取得一个值的尺寸（亦即此值的类型的尺寸）。
	const M, N = unsafe.Sizeof(X.c), unsafe.Sizeof(X)
	fmt.Println(M, N)

	// 此函数用来取得一个值在内存中的地址对齐保证（address alignment guarantee）。
	fmt.Println(unsafe.Alignof(X.a)) // 8
	fmt.Println(unsafe.Alignof(X.b)) // 1
	fmt.Println(unsafe.Alignof(X.c)) // 8

	//  此函数用来取得一个结构体值的某个字段的地址相对于此结构体值的地址的偏移。
	fmt.Println(unsafe.Offsetof(X.a)) // 0
	fmt.Println(unsafe.Offsetof(X.b)) // 8
	fmt.Println(unsafe.Offsetof(X.c)) // 16

	p := Person{name: "test", age: 1}
	fmt.Println(unsafe.Add(unsafe.Pointer(&p), 10))

	//unsafe.Slice()
}

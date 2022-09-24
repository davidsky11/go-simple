package pointer

// 一些Go的说明
// 		1. 非类型安全指针值是指针但uintptr值是整数
//		2. 不再被使用的内存块的回收时间点是不确定的
//		3. 一个值的地址在程序运行中可能改变
//		4. 一个值的生命范围可能并没有代码中看上去的大
//		5. *unsafe.Pointer是一个类型安全指针类型

// uintptr 可以用于运算（不安全）

import (
	"fmt"
	"reflect"
	"unsafe"
)

var X struct {
	a int64
	b bool
	c string
}

type T struct {
	a string
	b string
	c string
}

type S struct {
	b bool
}

var Y struct {
	a int64
	*S
	T
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

	fmt.Println("Y.a: ", unsafe.Offsetof(Y.a)) // 0
	fmt.Println("Y.S: ", unsafe.Offsetof(Y.S)) // 8
	fmt.Println("Y.T: ", unsafe.Offsetof(Y.T)) // 16

	// 此行编译不过，因为选择器Y.b中的隐含字段S为指针。
	//fmt.Println("Y.b: ", unsafe.Offsetof(Y.b)) // error

	// 此行可以编译过，但是它将打印出字段b在Y.S中的偏移量.
	fmt.Println("Y.S.b: ", unsafe.Offsetof(Y.S.b)) // 0

	// 此行可以编译过，因为选择器Y.c中的隐含字段T为非指针。
	fmt.Println("Y.c: ", unsafe.Offsetof(Y.c)) // 48

	// 打印字段c在Y.T中的偏移量
	fmt.Println("Y.T.c: ", unsafe.Offsetof(Y.T.c)) // 32

	p := Person{name: "test", age: 1}
	fmt.Println(unsafe.Add(unsafe.Pointer(&p), 10))

	//unsafe.Slice()
}

// Go 1.17引入的两个函数为：
//		func Add(ptr Pointer, len IntegerType) Pointer。
//			在一个（非安全）指针表示 的地址上添加一个偏移量，然后返回表示新地址的一个指针。
//		func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType。
//			从一个任意（安全）指针派生出一个指定长度的切片。

// ======> 这两个函数具有一定的危险性，需谨慎使用

func UnsafePointerDanger() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a) // 使a开辟在堆上
	// [0 0 0 3 0 0 0 0 0 9 0 11 0 0 0 0]
	elseSize := int(unsafe.Sizeof(a[0]))
	p9 := &a[9]
	up9 := unsafe.Pointer(p9)
	p3 := (*int)(unsafe.Add(up9, -6*elseSize))
	fmt.Println(*p3) // 3
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	// 下面是两个不正确的调用。因为它们的返回结果引用了未知的内存块。
	_ = unsafe.Add(up9, 7*elseSize)
	_ = unsafe.Slice(p9, 8)
}

// 假设此函数不会被内联
func createInt() *int {
	return new(int)
}

func Foo() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y) // 和y一样引用着同一个值
	var p2 = uintptr(unsafe.Pointer(z))

	// 此时，即使z指针值所引用的int值的地址仍旧存储
	// 在p2值中，但是此int值已经不再被使用了，所以垃圾
	// 回收器认为可以回收它所占据的内存块了。另一方面，
	// p0和p1各自所引用的int值仍旧将在下面被使用。

	// uintptr 值可以参与算术运算。
	p2 += 2
	p2--
	p2--

	*p0 = 1         // ok
	*(*int)(p1) = 2 // ok
	// 当*(*int)(unsafe.Pointer(p2)) = 3被执行的时候，此内存块有可能已经被回收了。
	*(*int)(unsafe.Pointer(p2)) = 3 // 危险操作！
}

type More struct {
	x int
	y *[1 << 23]byte
}

func Bar() {
	more := More{y: new([1 << 23]byte)}
	p := uintptr(unsafe.Pointer(&more.y[0]))

	// 值more仍旧在使用中并不能保证被值more.y所引用的值仍在被使用。

	// 一个聪明的编译器能够觉察到值more.y将不会再被用到，
	// 所以认为more.y值所占的内存块可以被回收了。

	*(*byte)(unsafe.Pointer(p)) = 1 // 危险操作
	println(more.x)                 // ok  继续使用值more，但只使用more.x字段

}

func Pointer4SafePointer() {
	x := 123                // 类型为 int
	p := unsafe.Pointer(&x) // 类型为 unsafe.Pointer
	pp := &p                // 类型为 *unsafe.Point
	p = unsafe.Pointer(pp)
	pp = (*unsafe.Pointer)(p)
}

// 如何正确地使用非类型安全指针？
//		模式一. 将类型*T1的一个值转换为非类型安全指针值，然后将此非类型安全指针值转换为类型*T2。

// math标准库包中的Float64bits函数
func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func float64frombits(b uint64) float64 {
	return *(*float64)(unsafe.Pointer(&b))
}

// 是将一个不再使用的字节切片转换为一个字符串（从而避免对底层字节序列的一次开辟和复制）
func byteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

//		模式二. 将一个非类型安全指针值转换为一个uintptr值，然后使用此uintptr值。

func UnsafeLogPrint() {
	type T struct{ a int }
	var t T
	fmt.Printf("%p\n", &t)                          // 0xc0000103d0
	println(&t)                                     // 0xc0000103d0
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // c0000103d0
}

//		模式三. 将一个非类型安全指针转换为一个uintptr值，然后此uintptr值参与各种算术运算，
//		再将算术运算的结果uintptr值转回非类型安全指针。

type Ti struct {
	x bool
	y [3]int16
}

const N = unsafe.Offsetof(Ti{}.y)
const M = unsafe.Sizeof(Ti{}.y[0])

func UnsafePointTransform() {
	t := Ti{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)

	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	// 对于这样地址加减运算，更推荐unsafe.Add函数来完成。
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2) // 789
}

// UnsafePointTransformDanger 不建议取地址分两行处理！！！！！
func UnsafePointTransformDanger() {
	t := Ti{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t)
	//ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))

	addr := uintptr(p) + N + M + M
	// 从这里到下一行代码执行之前，t值将不再被任何值引用，所以垃圾回收期认为它可以被回收了。
	// 一旦它真的被回收了，下面继续使用 t.y[2] 值的曾经的地址时非法和危险的！！！

	// 另一个危险的原因时t的地址在执行下一行之前可能改变。

	// 另一个潜在的危险是：如果再次期间发生了一些导致协程堆栈大小改变的情况，则记录在addr中
	// 的地址将失效。当然，此危险对于这个特定的例子并不存在。

	// ty2 取值变得不原子了（地址可能发生变化）  太容易掉坑了！！！
	ty2 := (*int16)(unsafe.Pointer(addr))
	fmt.Println(*ty2)
}

// 中间uintptr值可以参与&^清位运算来进行内存对齐计算，只要保证转换前后的非类型安全指针同时指
// 		向同一个内存块，整个转换就是合法安全的。

// 另一个需要注意的细节是最好不要将一个内存块的结尾边界地址存储在一个（安全或非安全）指针中。
//		这样做将导致紧随着此内存块的另一个内存块肯定不会被垃圾回收掉。

//		模式四. 将非类型安全指针值转换为uintptr值并传递给syscall.Syscall函数调用。

// syscall标准库包中的Syscall函数的原型为：
// 		func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
//	那么此函数是如何保证处于传递给它的地址参数值a1、a2和a3处的内存块在此函数执行过程中一定没有被回收和被移动呢？
//		- 此函数无法做出这样的保证。事实上，是编译器做出了这样的保证。 这是syscall.Syscall这样的函数的特权。其它自定义函数无法享受到这样的待遇。
//		- 我们可以认为编译器针对每个syscall.Syscall函数调用中的每个被转换为uintptr类型的非类型安
//			全指针实参添加了一些指令，从而保证此非类型安全指针所引用着的内存块在此调用返回之前不会被垃圾回收和移动。

//		模式五. 将reflect.Value.Pointer或者reflect.Value.UnsafeAddr方法的uintptr返回值立即转换为非类型安全指针。

//		模式六. 将一个reflect.SliceHeader或者reflect.StringHeader值的Data字段转换为非类型安全指针，以及其逆转换。

func ReflectStringHeaderPointer() {
	a := [...]byte{'G', 'o', 'l', 'a', 'n', 'g', ' '}
	s := "Java"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&a))
	hdr.Len = len(a)
	fmt.Println(s) // Goland

	// 现在，字符串s和切片a共享着底层的byte字节序列，从而使得此字符串中的字节变得可以修改。
	a[2], a[3], a[4], a[5] = 'o', 'g', 'l', 'e'
	a[6] = '!'
	fmt.Println(s) // Google
}

func ReflectSliceHeaderPointer() {
	a := [6]byte{'G', 'o', 'o', 'g', 'l', 'e'}
	bs := []byte("Golang!")
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	hdr.Data = uintptr(unsafe.Pointer(&a))

	hdr.Len = 2
	hdr.Cap = len(a)
	fmt.Printf("%s\n", bs) // Go
	bs = bs[:cap(bs)]
	fmt.Printf("%s\n", bs) // Google
}

// 我们只应该从一个已经存在的字符串值得到一个*reflect.StringHeader指针， 或者从
//		一个已经存在的切片值得到一个*reflect.SliceHeader指针， 而不应该从一个StringHeader值生
//		成一个字符串，或者从一个SliceHeader值生成一个切片。

// var hdr reflect.StringHeader
// hdr.Data = uintptr(unsafe.Pointer(new([5]byte)))
// 		在此时刻，上一行代码中刚开辟的数组内存块已经不再被任何值
// 		所引用，所以它可以被回收了。
// hdr.Len = 5
// s := *(*string)(unsafe.Pointer(&hdr)) // 危险！

func String2ByteSlice(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Cap = strHdr.Len
	sliceHdr.Len = strHdr.Len
	return
}

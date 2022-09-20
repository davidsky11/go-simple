package grammar

import "fmt"

// 基本数据类型（原生数据类型）：
//		整型、浮点型、复数型、布尔型、字符串、字符（byte、rune）
// 复合数据类型（派生数据类型）：
//		数组（array）、切片（slice）、映射（map）、函数（function）、结构体（struct）、通道（channel）、接口（interface）、指针（pointer）

// 整型分两大类。
//		有符号整型：int8、int16、int32、int64、int。
//		无符号整型：uint8、uint16、uint32、uint64、uint。
// 		uintptr：用于存放一个指针
//		其中uint8就是byte型，int16对应C语言的short型，int64对应C语言的long型。

// 浮点型：
//		float32：4字节，32位
//			最大值 math.MaxFloat32，大约是3.4×10E38
//			最小值 math.SmallestNonzeroFloat32，大约是1.4×10E-45
//		float64：8字节，64位
//			最大值 math.MaxFloat64，大约是1.8×10E308
//			最小值 math.SmallestNonzeroFloat64，大约是4.9×10E-324

// 复数型：复数型用于表示数学中的复数，如1+2j、1-2j、-1-2j等
// 		complex64：8字节，由float32类型的实部和虚部联合表示
//		complex128：16字节，由float64类型的实部和虚部联合表示

// 布尔型：布尔型的值只可以是常量true或者false

// 字符串：
// 定义多行字符串的方法如下:
// • 双引号书写字符串被称为字符串字面量（string literal），这种字面量不能跨行。
// • 多行字符串需要使用反引号“`”，多用于内嵌源码和内嵌数据。
// • 在反引号中的所有代码不会被编译器识别，而只是作为字符串的一部分。

func multiString() {
	var temp string
	temp = `
		x := 10
		y := 20
		z := 30
		fmt.PrintLn(x, "  ", y, "  ", z)
		x, y, z = y, z, x
		fmt.PrintLn(x, "  ", y, "  ", z)`
	fmt.Println(temp)
}

// 字符：字符串中的每一个元素叫作“字符”，定义字符时使用单引号
//		byte：1字节，表示 UTF-8 字符串的单个字节的值，unit8 的别名类型
//		rune：4字节，表示单个 unicode 字符，int32的别名类型

var char1 byte = 'a'
var char2 rune = '一'

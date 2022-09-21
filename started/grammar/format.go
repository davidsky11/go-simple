package grammar

import "fmt"

// 打印格式化通常使用fmt包

// 通用打印格式
//		%v：值的默认格式表示
//		%+v：类似 %v，但输出结构体时会添加字段名
//		%#v：值的Go语法表示
//		%T：值的类型的Go语法表示

func formatOut() {
	str99 := "steven"
	fmt.Printf("%T, %v \n", str99, str99)
	var char99 rune = '一'
	fmt.Printf("%T, %v \n", char99, char99)
	var char91 byte = 'b'
	fmt.Printf("%T, %v \n", char91, char91)
	var char98 int32 = 98
	fmt.Printf("%T, %v \n", char98, char98)
}

// 布尔型打印格式：%t

// 整型打印格式：
//		%b：二进制
//		%c：对应的unicode码值
//		%d：十进制
//		%8d：表示该整型长度是8，不足8则在数值前补空格；超过8，则以实际为准
//		%08d：表示该整型长度是8，不足8则在数值前补0；超过8，则以实际为准
//		%o：八进制
//		%q：该值对应的单引号括起来的Go语法字符字面值，必要时会采用安全的转义表示
//		%x：十六进制，使用 a~f
//		%X：十六进制，使用 A~F
//		%U：表示位unicode格式；U+1234，等价于 “U+%04X”

// 浮点型和复数型 打印
//		%b：无小数部分、二进制指数的科学计数法，如 -123456p-78
//		%e：(=%.6e)有6位小鼠部分的科学计数法，如 -1234.456e+78
//		%E：科学计数法，如 -1234.456E+78
//		%f：有6位小数部分，如 123.456123
//  	%F：等价于 %f
//		%g：根据实际情况采用 %e 或 %f 格式
//		%G：根据实际情况采用 %E 或 %F 格式

// 字符串与字节数组的打印
//		%s：直接输出字符串或者字节数组
//		%q：该值对应的双引号括起来的 GO 语法字符串字面值，必要时会采用安全的转义
//		%x：每个字节用两字符十六进制数表示，使用 a~f
//		%X：每个字节用两字符十六进制数表示，使用 A~F

type IPAddr [4]byte

// fmt 包中定义的 Stringer 是最普遍的接口之一。
//  type Stringer interface {
//    	String() string
//	}
//
// Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

func PrintIPAddr() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type Room struct {
	Name   string
	Number int
}

// 实现 Stringer interface 的 String() 方法，用于输出

func (p Room) String() string {
	return fmt.Sprintf("%v@(No.%v)", p.Name, p.Number)
}

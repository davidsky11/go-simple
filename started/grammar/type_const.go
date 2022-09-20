package grammar

import "fmt"

// 常量中的数据类型只可以是布尔型、数字型（整型、浮点型和复数型）和字符串

// 声明方式： const 标识符 [类型] = 值
// 可以省略类型说明符[type]，因为编译器可以根据变量的值来自动推断其类型

const B string = "Steven"
const C = "Steven"
const WIDTH, HEIGHT = 10, 20

// 常量用于枚举
// 常量组中如果不指定类型和初始值，则与上一行非空常量的值相同
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func constPrint() {
	fmt.Println(Unknown, Female, Male)
}

// iota，特殊常量值，是一个系统定义的可以被编译器修改的常量值。
// iota只能被用在常量的赋值中，在每一个const关键字出现时，被重置为0，然后每出现一个常量，iota所代表的数值会自动增加1。
// iota可以理解成常量组中常量的计数器，不论该常量的值是什么，只要有一个常量，那么iota就加1

const (
	con001 = iota
	con002 = iota
	con003 = iota
)

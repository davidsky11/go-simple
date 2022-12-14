package closure

import "fmt"

// 闭包（Closure）是词法闭包（Lexical Closure）的简称

// 闭包是由函数和与其相关的引用环境组合而成的实体。
// 在实现深约束时，需要创建一个能显式表示引用环境的东西，并将它与相关的子程序捆绑在一起，这样捆绑起来的整体被称为闭包。
// 函数 + 引用环境 = 闭包。

// 闭包只是在形式和表现上像函数，但实际上不是函数。

// 闭包在运行时可以有多个实例，不同的引用环境和相同的函数组合可以产生不同的实例。
// 闭包在某些编程语言中被称为Lambda表达式。

// 函数本身不存储任何信息，只有与引用环境结合后形成的闭包才具有“记忆性”。
// 函数是编译器静态的概念，而闭包是运行期动态的概念。

// 对象是附有行为的数据，而闭包是附有数据的行为。

// 闭包的优点：
// 		1. 加强模块化
//		2. 抽象。闭包是数据和行为的组合，这使得闭包具有较好的抽象能力。
//		3. 简化代码

// 一个编程语言需要以下特性来支持闭包。
// 		• 函数是一阶值（First-class value，一等公民），即函数可以作为另一个函数的返回值或参数，还可以作为一个变量的值。
//		• 函数可以嵌套定义，即在一个函数内部可以定义另一个函数。
//		• 允许定义匿名函数。
//		• 可以捕获引用环境，并把引用环境和函数代码组成一个可调用的实体。

// 闭包函数实现计数器
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum1=%d \t", sum)
		sum += x
		fmt.Printf("sum2=%d \t", sum)
		return sum
	}
}

func ClosureCounter() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(pos(i))
	}

	fmt.Println("--------------")
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(pos(i))
	}
}

// 由于闭包函数“捕获”了和它在同一作用域的其他常量和变量，
// 所以当闭包在任何地方被调用，闭包都可以使用这些常量或者变量。
// 它不关心这些变量是否已经超出作用域，只要闭包还在使用这些变量，这些变量就依然存在。

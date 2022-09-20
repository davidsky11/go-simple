package goroutine

import (
	"fmt"
	"time"
)

// 延迟函数调用并非绝对必要。两个需要的点：
//		1. 对于 恐慌/恢复特性是必要的
//		2. 可以邦族我们写出更整洁和更鲁棒的代码

// DeferTest defer关键字，形成延迟函数调用。
// 被延迟的函数调用的所有返回值必须全部被舍弃。
// 延迟调用堆栈
func DeferTest() {
	defer fmt.Println("The third line.")
	defer fmt.Println("The second line.")
	fmt.Println("The first line.")
}

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值  r 是返回值
	}()

	return n + n // 约等于  r = n + n; return
}

//  ----------->  延迟调用的实参的估值时刻

// 一个协程调用或者延迟调用的实参是在此调用发生时被估值的。
//		1. 对于一个延迟函数调用，它的实参是在此调用被推入延迟调用堆栈的时候被估值的。
//		2. 对于一个协程调用，它的实参是在此协程被创建的时候估值的。

//      一个匿名函数体内的表达式是在此函数被执行的时候才会被逐个估值的，不管此函数是被普通调用还是延迟/协程调用

// DeferMoreTest /**
// 输出：a: 2
//a: 1
//a: 0
//
//b: 3
//b: 3
//b: 3
//
// 为什么？
// 第一个循环中的i是在 fmt.Println函数调用被推入延迟调用堆栈的时候估的值
// 第二个循环中的i是在第二个匿名函数调用的退出阶段估的值（此时循环变量i的值已经变为3）
func DeferMoreTest() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			// 退出阶段估的值（此时循环变量i的值已经变为3）
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}

// DeferMoreTest1 正确示例1
func DeferMoreTest1() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			// 此 i 为形参i，非实参循环变量i
			defer func(i int) {
				fmt.Println("b:", i)
			}(i) // 这个 i 为for循环中的 实参 i
		}
	}()
}

// DeferMoreTest2 正确示例2
func DeferMoreTest2() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			i := i // 在下面的调用中，左 i 遮挡了右 i
			// <==>  var i = i
			defer func() {
				// 此 i 为上面的左 i，非循环变量 i
				fmt.Println("b:", i)
			}()
		}
	}()
}

// GoroutineActualVariable 协程的实参的估值时刻
// 实际打印  123 789
func GoroutineActualVariable() {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 执行时，a 是 789
	}(a) // 创建协程时，a 时 123

	a = 789

	time.Sleep(time.Second * 2)
}

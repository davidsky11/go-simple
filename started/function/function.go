package function

import (
	"fmt"
	"math"
	"strings"
)

// 普通函数需要先声明才能调用，一个函数的声明包括参数和函数名等

// 函数定义
// func 函数名 (参数列表) (返回参数列表) { }
// func funcName (pName type1, pName2 type2...) (output1 type1, output2 type2...) { }

// 在参数列表中，如果有多个参数变量，则以逗号分隔；如果相邻变量是同类型，则可以将类型省略。

func add(a, b int) {}

// 支持可变参数
func myFunc(arg ...int) {}

// 变量作用域
// 作用域是变量、常量、类型、函数的作用范围。
//    	在函数体内声明的变量称为局部变量，它们的作用域只在函数体内，生命周期同所在的函数。参数和返回值变量也是局部变量。
// 		在函数体内声明的变量称为局部变量，它们的作用域只在函数体内，生命周期同所在的函数。参数和返回值变量也是局部变量。
//
// 		全局变量可以在任何函数中使用。Go语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑

//		函数中定义的参数称为形式参数，形式参数会作为函数的局部变量来使用。

// 函数变量
func processCase(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

func stringToLower(str string, f func(string) string) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

// 直接将 func processCase 作为参数传入
func funcRunning() {
	result := stringToLower("AbdsDEDSGSEsdfsf", processCase)
	fmt.Println(result)
}

// 匿名函数
// Go语言支持匿名函数，即在需要使用函数时再定义函数。
// 匿名函数没有函数名，只有函数体，函数可以作为一种类型被赋值给变量，匿名函数也往往以变量方式被传递。
// 匿名函数经常被用于实现回调函数、闭包等。

func anonymousFunc() {
	// 匿名函数赋值给变量
	f := func(data string) {
		fmt.Println(data)
	}

	f("Go start...")
}

// 匿名函数用作回调函数
// 定义一个函数，遍历切片元素，对每个元素进行处理 （这个用法还挺优雅）
func visit(list []float64, f func(float64)) {
	for _, value := range list {
		f(value)
	}
}

func anonymousFuncForCallback() {
	arr := []float64{1, 9, 16, 25, 30}

	// 调用函数，对每个元素进行求平方根操作
	visit(arr, func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f \n", v)
	})

	// 调用函数，对每个元素进行求平方操作
	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.0f \n", v)
	})
}

// 可变参数
func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}

// 使用可变参数应注意如下细节。
// • 一个函数最多只能有一个可变参数。
// • 若参数列表中还有其他类型参数，则可变参数写在所有参数的最后。

// 递归函数
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func GetMultiple(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

// 使用递归需要注意如下事项。
//		• 递归函数的优点是定义简单，逻辑清晰。理论上，所有的递归函数都可以用循环的方式实现，但循环的逻辑不如递归清晰。
//		• 使用递归函数需要注意防止栈溢出。在计算机中，函数调用是通过栈（stack）这种数据结构实现的，每当进入一个函数调用，栈就会加一层，每当函数返回，栈就会减一层。由于栈的大小不是无限的，所以，递归调用的次数过多，会导致栈溢出。
//		• 使用递归函数的优点是逻辑简单清晰，缺点是过深的调用会导致栈溢出。

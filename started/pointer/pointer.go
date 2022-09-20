package pointer

import "fmt"

// 在Go语言中使用取地址符（＆）来获取变量的地址，一个变量前使用＆，会返回该变量的内存地址。

// 		Go语言指针的最大特点是：指针不能运算（不同于C语言）
// 		在Go语言中如果对指针进行运算会报错。

// nil指针也称为空指针。nil在概念上和其他语言的null、None、NULL一样，都指代零值或空值。

func ChangeValue() {
	b := 1808
	a := &b
	fmt.Println("b 的地址：", a)
	fmt.Println("*a 的值：", *a)
	*a++
	fmt.Println("b 的新值：", b)
}

// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到原内容数据。
// 严格来说Go语言只有值传递这一种传参方式，Go语言是没有引用传递的。

// Go语言中可以借助传指针来实现引用传递的效果。函数参数使用指针参数，传参的时候其实是复制一份指针参数，也就是复制了一份变量地址。

// 引用传递的作用如下。
// 		· 引用传递的作用如下。
//		· 传指针更轻量级（8 bytes），只需要传内存地址。如果参数是非指针参数，那么值传递的过程中，每次在复制上面就会花费相对较多的系统开销（内存和时间）。所以要传递大的结构体的时候，用指针是一个明智的选择。

// Go语言中slice、map、chan类型的实现机制都类似指针，所以可以直接传递，而不必取地址后传递指针。

func ChangeArrayPtr(a *[4]int) {
	fmt.Printf("-----ChangeArrayPtr 函数内：指针参数 a 的内存地址是： %p，值位：%v \n", &a, a)
	(*a)[1] = 188
}

// 函数传结构体
type Teacher struct {
	Name    string
	Age     int
	Married bool
	Sex     int8
}

func ChangeStructPtr(a *Teacher) {
	fmt.Printf("-----ChangeStructPtr 函数内：指针参数 a 的内存地址是： %p，值位：%v \n", &a, a)
	(*a).Name = "Katyusha"
	(*a).Age = 18
}

// Go语言中所有的传参都是值传递（传值），都是一个副本。
//		副本的内容有的是值类型（int、string、bool、array、struct属于值类型），这样在函数中就无法修改原内容数据；
//		有的是引用类型（pointer、 slice、map、chan属于引用类型），这样就可以修改原内容数据。

// 在实际编程时，应尽量使用函数来提高代码的复用性，
// 		对于占用内存较大的变量应尽量使用指针来减少资源的消耗。

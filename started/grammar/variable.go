package grammar

// 单个声明
var var11 int

// 批量声明
var (
	a int
	b string
	c []float32
	d func() bool
	e struct {
		x int
		y string
	}
)

// 未初始化变量的默认值有如下特点。
//• 整型和浮点型变量默认值：0。
//• 字符串默认值为空字符串。
//• 布尔型默认值为false。
//• 函数、指针变量、切片默认值为nil。

var m1 int = 10
var m2 = 10

var test int

func createVariable() {
	// 声明变量的首选形式，但是它只能被用在函数体内，
	// 而不可以用于全局变量的声明与赋值
	// 该变量名必须是没有定义过的变量，若定义过，将发生编译错误
	m3 := 10
	test = m3 + 10

	n1, n2 := 100, 200
	test = n1 + n2
}

// 多重赋值
func multiVariable() {
	a := 10
	b := 20
	b, a = a, b
}

func getData() (int, int) {
	return 10, 20
}

// 匿名变量 _ 既不占用命名空间，也不会分配内存
func test_getData() {
	a, _ := getData() // 舍弃第二个返回值
	_, b := getData() // 舍弃第一个返回值
	test = a + b
}

package collection

import "fmt"

func IterArray() {
	arr001 := [5]float64{67.6, 78.9, 123.05, 89, 99.99}
	arr002 := [...]int{2, 3, 5}
	for i := 0; i < len(arr001); i++ {
		fmt.Print(arr001[i], "\t")
	}

	fmt.Println()

	for _, value := range arr002 {
		fmt.Print(value, "\t")
	}
}

// Go语言中的数组并非引用类型，而是值类型。

// 当它们被分配给一个新变量时，会将原始数组复制出一份分配给新变量。因此对新变量进行更改，原始数组不会有反应。
// 将数组作为函数参数进行传递，它们将通过值传递，原始数组依然保持不变。
func nonChange() {
	a := [...]string{"CN", "EN", "IND"}
	b := a
	b[0] = "TH"
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
}

// 如果需要改变，就使用 指针传递

// ArrayCompareTest 映射map 和 切片类型都属于不可比较类型。
// 		所以任意两个映射值（或切片值）是不能相互比较的。
func ArrayCompareTest() {
	var a [16]byte
	var s []int
	var m map[string]int

	fmt.Println(a == a)                  // true
	fmt.Println(m == nil)                // true
	fmt.Println(s == nil)                //
	fmt.Println(nil == map[string]int{}) // false
	fmt.Println(nil == []int{})          // false
}

func ArrayLength() {
	var a [5]int
	fmt.Println(len(a), cap(a)) // 5 5
	var s []int
	fmt.Println(len(s), cap(s)) // 0 0
	s, s2 := []int{2, 3, 5}, []bool{}
	fmt.Println(len(s), cap(s), len(s2), cap(s2)) // 3 3 0 0
	var m map[int]bool
	fmt.Println(len(m)) // 0
	m, m2 := map[int]bool{1: true, 0: false}, map[int]int{}
	fmt.Println(len(m), len(m2)) // 2 0
}

// CollectionAssignment 容器赋值
// 		映射 or 切片，赋值时是共享底层的原色
//		数组，赋值时时直接赋值到目标数组
func CollectionAssignment() {
	m0 := map[int]int{0: 7, 1: 8, 2: 9}
	m1 := m0
	m1[0] = 2
	fmt.Println(m0, m1)
	// map[0:2 1:8 2:9]	 map[0:2 1:8 2:9]

	s0 := []int{7, 8, 9}
	s1 := s0
	s1[0] = 2
	fmt.Println(s0, s1)
	// [2 8 9]	[2 8 9]

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1)
	// [7 8 9]	[2 8 9]
}

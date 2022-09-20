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

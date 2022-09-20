package main

import (
	"fmt"
	"github.com/katyusha/go-simple/started/closure"
	"github.com/katyusha/go-simple/started/collection"
	"github.com/katyusha/go-simple/started/function"
	"github.com/katyusha/go-simple/started/pointer"
)

func main() {
	fmt.Println("Hello Go")

	// 调用函数必须大写开头
	closure.ClosureCounter()

	// 可变参数
	sum, avg, count := function.GetScore(90, 82.5, 77, 89)
	fmt.Printf("学生共有 %d 门成绩，总成绩为: %.2f， 平均成绩: %.2f \n", count, sum, avg)

	fac := 5
	fmt.Printf("%d 的阶乘是: %d \n", fac, function.GetMultiple(fac))

	pointer.ChangeValue()

	arr001 := [4]int{1, 2, 3, 4}
	pointer.ChangeArrayPtr(&arr001)
	fmt.Printf("after ChangeArrayPtr: %v \n", arr001)

	teacher001 := pointer.Teacher{Name: "Steven", Age: 26, Married: true, Sex: 1}
	fmt.Printf("变量 teacher001 的内存地址是: %p，值为：%v \n", &teacher001, teacher001)
	pointer.ChangeStructPtr(&teacher001)
	fmt.Printf("after ChangeStructPtr: %v \n", teacher001)

	collection.IterArray()
}

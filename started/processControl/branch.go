package processControl

import "fmt"

// if
// switch
// select：类似 switch，但是 select 会随机执行一个可运行的case。
//		如果没有case可运行，它将阻塞，直到有case可运行

func ifPlus() {
	// num的定义在if里，那么只能够在该if...else语句块中使用，否则编译器会报错
	if num := 10; num%2 == 0 {
		fmt.Println(num, "偶数")
	} else {
		fmt.Println(num, "奇数")
	}
}

// switch语句执行的过程自上而下，直到找到case匹配项，匹配项中无须使用break，
// 		因为Go语言中的switch默认给每个case自带break。
// 因此匹配成功后不会向下执行其他的 case 分支，而是跳出整个 switch。
//		可以添加fallthrough（中文含义是：贯穿），强制执行后面的case分支。
//		fallthrough必须放在case分支的最后一行。如果它出现在中间的某个地方，编译器就会报错

func switch1(int) {
	grade := ""
	score := 98.5
	switch { // switch省略不写，默认相当于：switch true
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的等级是： %s\n", grade)
}

func switch2() {
	year := 2022
	month := 2
	days := 0
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}
	fmt.Printf("%d 年 %d 月的天数为： %d\n", year, month, days)
}

// TypeSwitch 类型转换
func TypeSwitch(x interface{}) {
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型： %T", i)
	case int:
		fmt.Printf("x 的类型： int\n")
	case float64:
		fmt.Printf("x 的类型： float64\n")
	case func(int) float64:
		fmt.Printf("x 的类型： func(int) 型\n")
	case bool, string:
		fmt.Printf("x 的类型： bool 或 string 型\n")
	default:
		fmt.Printf("x 的类型：未知\n")
	}
}

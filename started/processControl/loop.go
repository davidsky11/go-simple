package processControl

import "fmt"

// for
// 循环嵌套
// break   continue    goto

func forLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	for ; ; i++ {
		if i > 20 {
			break
		}
		fmt.Printf("%d ", i)
	}

	//var i int
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	for {
		if i > 10 {
			break
		}
		fmt.Print(i)
		i++
	}
}

func forRange() {
	str := "123ABCagss火箭"
	for i, value := range str {
		fmt.Printf("第 %d 位的ASCII值 = %d， 字符是 %c \n", i, value, value)
	}
}

func forContinue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

// break与continue的区别如下。
// • break语句无条件跳出并结束当前的循环，然后执行循环体后的语句。
// • continue语句跳过当前的循环，而开始执行下一次循环。

func forGoto() {
	var C, c int
	C = 1

LOOP:
	for C < 50 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto LOOP // 若发现因子则不是素数
			}
		}
		fmt.Printf("%d \t", C)
	}
}

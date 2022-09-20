package goroutine

import (
	"fmt"
	"runtime"
)

func SetMaxProcessCount() {
	fmt.Println("逻辑CPU个数", runtime.NumCPU())

	runtime.GOMAXPROCS(8) // 设置逻辑CPU个数
}

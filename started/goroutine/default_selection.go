package goroutine

// 默认选择
// 		当 select 中的其它分支都没有准备好时，default 分支就会执行。

import (
	"fmt"
	"time"
)

func DefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

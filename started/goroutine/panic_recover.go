package goroutine

import (
	"fmt"
	"time"
)

// PanicCrashTest 在一个新协程中产生了一个恐慌，并且此协程在恐慌状况下退出，所以整个程序崩溃
func PanicCrashTest() {
	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}

}

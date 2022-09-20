package goroutine

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成
}

func WaitGroupTest() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2)
	go sayGreetings("hi!", 10)
	go sayGreetings("hello!", 10)
	wg.Wait() // 阻塞在这里，知道所有任务都已完成
}

// DeadLockTest 死锁示例
func DeadLockTest() {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		//wg.Done()		// 需要通知主协程，任务已经完成
		//wg.Wait()		// 阻塞在此
	}()
	wg.Wait() // 阻塞在此
}

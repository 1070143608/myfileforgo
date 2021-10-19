package test

import (
	"fmt"
	"github.com/bigwhite/functrace"
	"testing"
	"time"
)

// 巧用 defer, 在函数入口和出口时做一些操作
func TestDefer(t *testing.T) {
	defer functrace.Trace()()
	doSomething()
	b()
}

func doSomething() {
	defer functrace.Trace()()
	defer countTime("doSomething")()
	// 模拟耗时操作
	fmt.Println("run func2: ", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

// 统计某函数的运行时间
func countTime(msg string) func() {
	defer functrace.Trace()()

	start := time.Now()
	fmt.Println("run func countTime: ", time.Now())
	fmt.Printf("run func countTime: %s", msg)
	return func() {
		fmt.Printf("func name: %s run time: %f s \n", msg, time.Since(start).Seconds())
	}
}

func trace(s string) string {
	defer functrace.Trace()()
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	defer functrace.Trace()()
	fmt.Println("leaving:", s)
}

func a() {
	defer functrace.Trace()()
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer functrace.Trace()()
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

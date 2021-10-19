package main

import (
	"context"
	"example.com/study"
	"fmt"
	"github.com/bigwhite/functrace"
	"time"
)

func main() {
	defer

	// example 1
	functrace.Trace()()

	ctx := context.Background()
	config := &study.Config{
		MaxConn:        4,
		MaxIdle:        2,
		MaxWait:        2,
		MaxWaitTimeout: 3000,
	}
	conn := study.Prepare(ctx, config)
	for i := 0; i < 100; i++ {

		go func() {
			_, err := conn.New(ctx)
			if err != nil {
				fmt.Println(err.Error())
			}
		}()
		// 释放连接
		go conn.Release(ctx)
	}
	time.Sleep(time.Second * 60)

}

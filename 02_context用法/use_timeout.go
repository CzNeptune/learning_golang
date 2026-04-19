package main

import (
	"context"
	"fmt"
	"time"
)

// 场景： 如果你需要对一个用协程启动的函数做超时控制，可以用context来完成goroutine的控制

func useTimeout1() {
	// 设置一个用于超时控制的context ctx, ctx作为参数可以用来作为协程的超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ctx作为参数传递给需要做超时控制的函数

	go Monitor(ctx)

	time.Sleep(20 * time.Second)
}

func Monitor(ctx context.Context) {
	for {

		select {
		// 如果context 超时，ctx.Done()就会返回一个空接口 struct{}
		case <-ctx.Done():
			// 如果超时时间到了，就退出循环
			fmt.Println(ctx.Err())
			return
		// 如果没有超时，打印输出后继续循环
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("monitor")
		}

	}
}

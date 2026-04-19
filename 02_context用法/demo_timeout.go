package main

import (
	"context"
	"time"
)

func run() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go process(ctx)

	time.Sleep(5 * time.Second)

}

func process(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second): // 如果设置为1秒则能完成,如果设置为3秒则被取消
		println("process completed")
	case <-ctx.Done(): //任务超时取消时会走此逻辑
		println("process canceled", ctx.Err())
	}
}

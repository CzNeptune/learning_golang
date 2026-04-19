package main

import (
	"context"
	"fmt"
	"time"
)

/*
2.利用context，手动让2个goroutine同时结束[是不是更简单?]
*/
func useCancel1() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutin1 exit.")
				return
			default:
				fmt.Println("goroutin1 sleep 1s, keep going.")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutin2 exit.")
				return
			default:
				fmt.Println("goroutin2 sleep 1s, keep going.")
				time.Sleep(time.Second * 1)
			}
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("10s 时间到了，goroutine需要退出了.")

	// 利用context的方法，手动让2个goroutine同时结束
	cancel()

	time.Sleep(5 * time.Second)
}

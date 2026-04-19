package main

import (
	"context"
	"fmt"
	"time"
)

/*
3. 利用context传递key-value
*/
func useValue1() {

	// 为ctx设置一个key-value
	ctx := context.Background()
	ctx = context.WithValue(ctx, "hello", "world")
	x := ctx.Value("hello")
	fmt.Println("x=", x) // world

	// 将key-vluae值传递到goroutine
	go work(ctx)

	time.Sleep(3 * time.Second)

}
func work(ctx context.Context) {
	fmt.Println("do worker.")
	fmt.Println("hello=", ctx.Value("hello")) // world，利用context传递key-value
	// 继续传递到下层goroutine
	go subwork(ctx)
}

func subwork(ctx context.Context) {
	fmt.Println("do subwork.")
	fmt.Println("hello=", ctx.Value("hello")) // world，利用context传递key-value到更进一层
}

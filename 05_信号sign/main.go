package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	running := true
	// 创建信号通道,用于接收信号
	sigChan := make(chan os.Signal, 1)
	// signal.Notify注册信号,将信号通道注册为接收特定信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT:
			fmt.Println("处理SIGINT信号")
			running = false
		case syscall.SIGTERM:
			fmt.Println("处理SIGTERM信号")
			running = false
		}
	}()

	for running {
		time.Sleep(5 * time.Second)
		fmt.Println("正在运行")
	}

	// 将之前设置的所有信号取消注册并关闭通道
	signal.Stop(sigChan)
	fmt.Println("退出主函数")
}

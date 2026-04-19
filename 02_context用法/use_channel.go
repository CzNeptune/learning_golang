package main

import (
	"fmt"
	"time"
)

/*
1. 利用channel控制goroutine的停止
不使用context时,可以利用channel+select主动让goroutine退出
*/

func useChannel1() {

	stopChan := make(chan bool)

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("goroutin1 exit.")
				return
			default:
				fmt.Println("goroutin1 sleep 1s, keep going.")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("goroutin2 exit.")
				return
			default:
				fmt.Println("goroutin2 sleep 1s, keep going.")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("10s 时间到了，主进程需要退出了.")
	// 发送信号让goroute1结束
	stopChan <- true

	// 发送信号让goroute2结束
	stopChan <- true
	time.Sleep(5 * time.Second)
}

/*

2. 利用关闭channel的方法，让2个goroutine同时结束
close(ch) 让两个channel同时收到关闭信号,从而让两个goroutine同时结束
*/

func useChannel2() {

	stopChan := make(chan bool)

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("goroutin1 exit.")
				return
			default:
				fmt.Println("goroutin1 sleep 1s, keep going.")
				time.Sleep(time.Second * 2)
			}
		}
	}()

	go func() {
		for {
			select {
			case <-stopChan:
				fmt.Println("goroutin2 exit.")
				return
			default:
				fmt.Println("goroutin2 sleep 1s, keep going.")
				time.Sleep(time.Second * 3)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("10s 时间到了，主进程需要退出了.")
	// 利用关闭channel的方法，让2个goroutine同时结束
	close(stopChan)

	time.Sleep(5 * time.Second)
}

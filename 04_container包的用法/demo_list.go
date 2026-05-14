package main

import (
	"container/list"
	"fmt"
)

func dump_list(l *list.List) {
	l.Do(func(v any) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
}

func demo_list_1() {
	// 创建一个链表
	l := list.New()
}

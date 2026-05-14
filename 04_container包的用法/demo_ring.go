package main

import (
	"container/ring"
	"fmt"
)

// ring 是一个循环链表，每个节点包含一个值和指向下一个节点的指针
// 适合固定长度的缓冲区,轮询算法
// 环没有固定的起点与终点,遍历时要小心,否则会导致无限循环

func dump_ring(r *ring.Ring) {
	r.Do(func(v any) {
		fmt.Printf("%v ", v)
	})
	fmt.Println()
}

func demo_ring_1() {
	// 创建一个5个节点的环
	r := ring.New(5)
	r.Value = 0
	i := 1
	for p := r.Next(); p != r; p = p.Next() {
		p.Value = i
		i++
	}
	dump_ring(r)
	fmt.Println(r.Len())

	// 首节点往后移动2个位置
	r = r.Move(2)
	dump_ring(r)
}

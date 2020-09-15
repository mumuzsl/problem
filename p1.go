package main

import (
	"fmt"
	"math"
)

type Node struct {
	Val  int
	Next *Node
}

//func p1() int {
//	var n, k int
//	fmt.Scan(&n)
//	fmt.Scan(&k)
//
//	//link := make([]int, n)
//
//	var curr *Node
//	head := new(Node)
//	fmt.Scan(&head.Val)
//	curr = head
//	for i := 1; i < n; i++ {
//		node := new(Node)
//		fmt.Scan(&node.Val)
//		curr.Next = node
//	}
//
//	link = append(link[:k-1], link[k:]...)
//
//	for i := 0; i < n-1; i++ {
//		fmt.Println(link[i])
//	}
//
//	return link[0]
//}

func sum(k int64) int64 {
	var s int64
	for k != 0 {
		s += k % 10
		k = k / 10
	}
	return s
}

func compute(n int64) int64 {
	if n <= 18 {
		return n
	}
	var k int64
	var p int64
	for i := 2; i <= 13; i++ {
		p = int64(math.Pow10(i-1) - 1)
		if n >= p && float64(n) < math.Pow10(i)-1 {
			k = int64(9 * (i - 1))
			fmt.Println(k, n-p)
			return k + sum(n-p)
		}
	}
	return 0
}

func p2() {
	var T int
	fmt.Scan(&T)

	var n int64
	for i := 0; i < T; i++ {
		fmt.Scan(&n)
		fmt.Println(compute(n))
	}
}

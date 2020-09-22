package lc

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func All() {
	//r := myAtoi('  -42')
	//fmt.Println(r)
	//r := permutation('aab')
	//fmt.Println(r)
	//r := minDeletionSize({[}}string{'rrjk','furt','guzm'})
	//fmt.Println(r)
	//a := {[}}int{2, 4, 6, 2, 1, 3}
	//r := minSubsequence(a)
	//fmt.Println(r)
	//r := combinationSum({[}}int{2, 3, 5}, 8)
	//fmt.Println(r)
	//r := removeComments({[}}string{'/*Test program */', 'int main()', '{ ', '  // variable declaration ', 'int a, b, c;', '/* This is a test', '   multiline  ', '   comment for ', '   testing */', 'a = b + c;', '}'})
	//fmt.Println(r)
	//fmt.Println(0 | 1<<3 ^ 1<<3)
	//data := [][]byte{
	//	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	//	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	//	{'.', '9', '.', '.', '.', '.', '.', '6', '.'},
	//	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	//	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	//	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	//	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	//	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	//	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	//}
	//solveSudoku(data)
	//fmt.Println(data)
	r := swapNumbers([]int{1, 2})
	fmt.Println(r)
}

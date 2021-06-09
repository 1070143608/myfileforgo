package main

import (
	"log"
	_ "math"
	"sort"
	//"example/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	//names := []string{"hh", "mm", "dd"}
	//messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	//if err != nil {
	//	log.Fatal(err)
	//}

	// If no error was returned, print the returned message
	// to the console.
	//fmt.Println(messages)
	//fmt.Println(res)
}


func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, arrLen + 1)
	}
	dp[0][0] = 1
	for row := 1; row < steps+1; row++ {
		for col := 0; col < arrLen+1; col++ {
			var a int = 0
			var b int = 0
			var c int = 0
			b = dp[row-1][col]
			if col == 0 && row == 0 {
				continue
			}
			if col > 0 {
				a = dp[row-1][col-1]
			}
			if col + 1 <= arrLen {
				c = dp[row-1][col+1]
			}
			dp[row][col] = (a + b + c) % mod

		}
	}
	return dp[steps][0]

}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func isCousins(root *TreeNode, x int, y int) bool {
	var dp []*TreeNode
	dp = append(dp, root)
	for len(dp) != 0 {
		var temp [] *TreeNode
		var count int
		for _, node := range dp {
			if node.Left != nil && node.Right != nil{
				temp = append(temp, node.Left)
				temp = append(temp, node.Right)
				if (node.Left.Val == x && node.Right.Val == y) || (node.Left.Val == y && node.Right.Val == x) {
					return false
				}
			} else if node.Left != nil && node.Right == nil {
				temp = append(temp, node.Left)
			} else if node.Left == nil && node.Right != nil {
				temp = append(temp, node.Right)
			}
			if node.Left != nil && (node.Left.Val == x || node.Left.Val == y) {
				count += 1
			}
			if node.Right != nil && (node.Right.Val == x || node.Right.Val == y) {
				count += 1
			}
		}
		if count == 2 {
			return true
		}
		dp = temp
	}
	return false
}

func kthLargestValue(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])
	var result = make([]int, rows * cols)
	var dp = make([][]int, rows + 1)
	dp[0] = make([]int, cols + 1)
	for i, ints := range matrix {
		dp[i+1] = make([]int, cols + 1)
		for j, val := range ints {
			dp[i+1][j+1] = dp[i][j] ^ dp[i+1][j] ^ dp[i][j+1] ^ val
			result = append(result, dp[i+1][j+1])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result[k-1]
}


func reverseParentheses(s string) string {
	var stack []byte;
	for i := range s {
		if s[i] == ')' {
			temp := stack[len(stack)-1]
			temp_stack := []byte{}
			stack = stack[:len(stack)-1]
			for temp != '(' {
				temp_stack = append(temp_stack, temp)
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			for _, c := range temp_stack {
				stack = append(stack, c)
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}


func totalHammingDistance(nums []int) int {
	n := len(nums)
	var  res int
	var c int
	for i := 0; i < 30; i++ {
		c = 0
		for _, num := range nums {
			c += num >> i & 1
		}
		res += c * (n - c)
	}
	return res
}
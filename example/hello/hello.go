package main

import (
	"fmt"
	"log"
	_ "math"

	"example/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	names := []string{"hh", "mm", "dd"}
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(messages)
	res := numWays(2, 4)
	fmt.Println(res)
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
package greetings

import (
	"fmt"
	"github.com/bigwhite/functrace"
	"math"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	defer functrace.
		// Return a greeting that embeds the name in a message.
		Trace()()

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func change(amount int, coins []int) int {
	defer functrace.Trace()()
	dp := make([][]int, len(coins)+1)
	dp[0] = make([]int, amount+1)
	dp[0][0] = 1
	for i := 1; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
		for j := 1; j < amount+1; j++ {
			if coins[i-1] > j {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}

func numSquares(n int) int {
	defer functrace.Trace()()
	ns := int(math.Pow(float64(n), 0.5)) + 1
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		for j := 1; j < ns; j++ {
			if i >= j*j {
				if dp[i] == 0 {
					dp[i] = dp[i-j*j] + 1
				} else {
					dp[i] = min(dp[i], dp[i-j*j]+1)
				}
			} else {
				break
			}
		}
	}
	return dp[n]
}

func min(a int, b int) int {
	defer functrace.Trace()()
	if a > b {
		return b
	}
	return a
}

func peakIndexInMountainArray(arr []int) int {
	defer functrace.Trace()()
	left := 0
	right := len(arr) - 1
	var mid int
	for left < right {
		mid = int((left + right) / 2)
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return mid
}

func stoneGame(piles []int) bool {
	defer functrace.Trace()()
	left := 0
	right := len(piles) - 1
	var clik = 0
	var li = 0
	for left < right {
		if piles[left] > piles[right] {
			clik += piles[left]
			left += 1
		} else {
			clik += piles[right]
			right -= 1
		}
		if piles[left] > piles[right] {
			li += piles[left]
			left += 1
		} else {
			li += piles[right]
			right -= 1
		}
	}
	return clik > li
}

func IsNumber(s string) bool {
	defer functrace.Trace()()
	for _, c := range s {
		if !((c >= 65 && c <= 90) || (c >= 97 && c <= 122) || c == 46 || c == 43 || c == 45) {
			return false
		}

	}
	return true
}

type Interface interface {
	Tes1()
	Tes2()
	Tes3()
}

type DesFather struct {
	name int
}

type Des struct {
	DesFather
}

func (res *Des) Tes1() {
	defer functrace.Trace()()
	res.name = 111
	fmt.Println("test1")
}

func (res *Des) Tes2() {
	defer functrace.Trace()()
	fmt.Println("test2")
}

func (res *Des) Tes3() {
	defer functrace.Trace()()
	fmt.Println("test3")
}

func TestMain(ops ...Interface) {
	defer functrace.Trace()()
	fmt.Println("ddddddd")
	fmt.Println(len(ops))
	for _, op := range ops {
		op.Tes1()
	}
}

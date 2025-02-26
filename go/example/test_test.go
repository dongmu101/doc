package example

import (
	"fmt"
	"math"
	"testing"
)

// 青蛙跳台阶，总共5个台阶可以一次1步或者2步，跳上去有几种跳法
// 利用斐波那契算法
func Test青蛙跳台阶(t *testing.T) {
	n := 10
	ways := climbStairs(n)
	fmt.Printf("对于 %d 阶台阶，共有 %d 种上法\n", n, ways)
}
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs2(n-1) + climbStairs2(n-2)
}

func climbStairs2(n int) int {
	if n <= 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]

	}
	fmt.Println(dp)
	return dp[n]
}

// 给出一个数字在map数组中找到最相近的数字
func TestSolution(t *testing.T) {
	ranks := map[int]int{}
	ranks[1] = 93
	ranks[10] = 55
	ranks[15] = 30
	ranks[20] = 19
	ranks[23] = 11
	ranks[30] = 2
	result := solution(ranks, 16)
	fmt.Println("result:", ranks[result])
}

func solution(ranks map[int]int, ho int) int {
	gap := 0
	prevGap := 0
	result := 0
	for rank, honor := range ranks {
		gap = int(math.Abs(float64(honor - ho)))
		if gap <= prevGap {
			result = rank
		}
		prevGap = gap
	}
	return result
}

package main

import (
	"fmt"
	"math"
)

func main() {

}

//斐波拉契数列
func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

//回文数
func huiwen(x int) bool {
	if x == 0 {
		return true
	}
	var n = x
	revertedNumber := 0
	for {
		if n == 0 {
			break
		}
		revertedNumber = revertedNumber*10 + n%10
		n /= 10
		fmt.Println(n)
	}

	if x == revertedNumber {
		return true
	} else {
		return false
	}
}

//翻转数字
func reverseNumber(x int) int {
	var n int
	n = 0
	for {
		if x == 0 {
			break
		}
		n = n*10 + x%10
		x = x / 10
	}

	if n > math.MaxInt32 || n < math.MinInt32 {
		return 0
	}

	return n
}

//杨辉三角
//每一层左右都为1
func yanghui(x int) [][]int {
	ans := make([][]int, x)
	for i := range ans {
		ans[i] = make([]int, 0)
		//头和尾是1
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}




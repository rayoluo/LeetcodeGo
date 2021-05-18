package main

import "fmt"

/**
 * @Author: luo
 * @Author: hustly123@gmail.com
 * @Date: 2021/5/18 9:18
 * @Desc:
 */

var dp []int

type pair struct {
	m, n int
}

var m map[pair]int
var (
	v0, v1, v2, count int
	ok                bool
)

func countTriplets(arr []int) int {
	count = 0
	dp = make([]int, len(arr))
	m = make(map[pair]int)
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = dp[i-1] ^ arr[i]
	}
	for i := 0; i < len(arr)-1; i++ {
		for k := 1; k < len(arr); k++ {
			for j := i + 1; j <= k; j++ {
				if v1, ok = m[pair{i, j - 1}]; !ok {
					if i == 0 {
						m[pair{0, j - 1}] = dp[j-1]
					} else {
						m[pair{i, j - 1}] = dp[j-1] ^ dp[i-1]
					}
					v1 = m[pair{i, j - 1}]
				}
				if v2, ok = m[pair{j, k}]; !ok {
					m[pair{j, k}] = dp[k] ^ dp[j-1]
					v2 = m[pair{j, k}]
				}
				if v1 == v2 {
					count++
				}
			}
		}
	}
	return count
}

// 上面的方法能保证结果是正确的，但是对于长的测试用例超时 加一点剪枝能过
func countTripletsNew(arr []int) int {
	count = 0
	dp = make([]int, len(arr))
	m = make(map[pair]int)
	dp[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		dp[i] = dp[i-1] ^ arr[i]
	}
	for i := 0; i < len(arr)-1; i++ {
		for k := 1; k < len(arr); k++ {
			if v0, ok = m[pair{i, k}]; !ok {
				if i == 0 {
					m[pair{i, k}] = dp[k]
				} else {
					m[pair{i, k}] = dp[k] ^ dp[i-1]
				}
				v0 = m[pair{i, k}]
			}
			if v0 != 0 {
				continue
			}
			for j := i + 1; j <= k; j++ {
				if v1, ok = m[pair{i, j - 1}]; !ok {
					if i == 0 {
						m[pair{0, j - 1}] = dp[j-1]
					} else {
						m[pair{i, j - 1}] = dp[j-1] ^ dp[i-1]
					}
					v1 = m[pair{i, j - 1}]
				}
				if v2, ok = m[pair{j, k}]; !ok {
					m[pair{j, k}] = dp[k] ^ dp[j-1]
					v2 = m[pair{j, k}]
				}
				if v1 == v2 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	// fmt.Println(countTriplets([]int{2, 3, 1, 6, 7}))
	fmt.Println(countTriplets([]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(countTripletsNew([]int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}))

}
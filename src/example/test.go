package example

import "fmt"
import "framework"
import "math"
import "util"

type cc struct {
	obj framework.Object
}

var cc1 = cc{framework.Object{"CoinChange", "Example link: https://leetcode.com/problems/coin-change/"}}

func (cc1 *cc) Id() string {
	return cc1.obj.Id
}

func (cc1 *cc) Description() string {
	return cc1.obj.Description
}

func init() {
	framework.Register(&cc1)
}

func (cc1 *cc) Run() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 3))
	fmt.Println(coinChange(coins, 11))
	fmt.Println(coinChange(coins, -1))
}

// solution
func coinChange(coins []int, amount int) int {
	if amount <= 0 || len(coins) == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && dp[i-coins[j]] < math.MaxInt32 {
				dp[i] = util.Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	result := math.MaxInt32
	if dp[amount] != math.MaxInt32 {
		result = dp[amount]
	}
	return result
}

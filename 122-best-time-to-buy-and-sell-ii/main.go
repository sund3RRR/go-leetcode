// 122. Best Time to Buy and Sell Stock II
//
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/description/
//
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock. On each day, you can only buy one or sell one share of the stock.
// Return the maximum profit you can achieve.
package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}

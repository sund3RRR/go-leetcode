// 121. Best Time to Buy and Sell Stock
//
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	profit := 0

	for _, price := range prices {
		profit = max(profit, price-minPrice)
		minPrice = min(minPrice, price)
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

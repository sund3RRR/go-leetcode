// 714. Best Time to Buy and Sell Stock with Transaction Fee
//
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
//
// You are given an array prices where prices[i] is the price of a given stock on an integer day i.
// Every time you want to buy a stock, you pay the transaction fee.
// Return the maximum profit you can achieve.
// Note: You may complete as many transactions as you like, and you need to pay the transaction fee for each transaction.
package main

import "fmt"

func maxProfit(prices []int, fee int) int {
	hold := -prices[0]
	cash := 0

	for i := 1; i < len(prices); i++ {
		hold = max(hold, cash-prices[i])
		cash = max(cash, hold+prices[i]-fee)
	}

	return cash
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}

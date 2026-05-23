package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	return 0
}

func guessNumber(n int) int {
	var left = 0
	var right = n

	for {
		var mid = (left + right) / 2
		g := guess(mid)

		if g == -1 {
			right = mid - 1
		} else if g == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}
}

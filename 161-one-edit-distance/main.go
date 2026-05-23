// 161. One Edit Distance
//
// Даны две строки s и t. Нужно определить, находятся ли они ровно на расстоянии одной операции редактирования.
package main

import "fmt"

func oneEditDistance(s, t string) bool {
	ls, lt := len(s), len(t)
	if abs(ls-lt) > 1 {
		return false
	}

	// если строки равны — ноль правок, значит false
	if s == t {
		return false
	}

	for i := 0; i < min(ls, lt); i++ {
		if s[i] != t[i] {
			if ls == lt {
				return s[i+1:] == t[i+1:] // заменили символ
			} else if ls < lt {
				return s[i:] == t[i+1:] // вставили в s
			} else {
				return s[i+1:] == t[i:] // удалили из s
			}
		}
	}

	// Если дошли до конца и разница длины = 1 → true (например "ab" и "abc")
	return abs(ls-lt) == 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func main() {
	fmt.Println(oneEditDistance("abc", "abc"))
}

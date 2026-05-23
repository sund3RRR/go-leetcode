// 281. Zigzag Iterator
//
// https://leetcode.com/problems/zigzag-iterator
package main

import "fmt"

// Пара: какой массив и какой индекс
type pair struct {
	listIdx int
	elemIdx int
}

type ZigzagIterator struct {
	lists [][]int
	q     []pair
}

func Constructor(lists ...[]int) *ZigzagIterator {
	q := []pair{}
	for i, l := range lists {
		if len(l) > 0 {
			q = append(q, pair{i, 0})
		}
	}
	return &ZigzagIterator{lists: lists, q: q}
}

func (it *ZigzagIterator) HasNext() bool {
	return len(it.q) > 0
}

func (it *ZigzagIterator) Next() int {
	if !it.HasNext() {
		panic("no more elements")
	}

	// Берём первый элемент из очереди
	p := it.q[0]
	it.q = it.q[1:]

	val := it.lists[p.listIdx][p.elemIdx]

	// Если в этом списке есть следующий элемент — добавляем назад в очередь
	if p.elemIdx+1 < len(it.lists[p.listIdx]) {
		it.q = append(it.q, pair{p.listIdx, p.elemIdx + 1})
	}

	return val
}

func main() {
	v1 := []int{1, 2, 3}
	v2 := []int{4, 5, 6, 7}
	v3 := []int{8, 9}

	it := Constructor(v1, v2, v3)

	result := []int{}
	for it.HasNext() {
		result = append(result, it.Next())
	}

	fmt.Println(result) // [1 4 8 2 5 9 3 6 7]
}

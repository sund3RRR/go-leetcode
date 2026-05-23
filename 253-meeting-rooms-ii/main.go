// 253. Meeting Rooms II
//
// Дан массив интервалов [start, end], описывающих время встреч.
// Нужно вернуть минимальное количество переговорных комнат, необходимых для проведения всех встреч так, чтобы они не пересекались.
// sample = [ (1, 2), (1, 3), (2, 4), (2, 3), ]
package main

import (
	"fmt"
	"sort"
)

type Event struct {
	day   int
	delta int
}

func meetingRooms(sample [][2]int) int {
	events := make([]Event, 0, len(sample))

	for _, v := range sample {
		events = append(events, Event{
			day:   v[0],
			delta: 1,
		})
		events = append(events, Event{
			day:   v[1],
			delta: -1,
		})
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].day == events[j].day {
			return events[i].delta < events[j].delta
		}
		return events[i].day < events[j].day
	})

	var acc int
	var maxGuest int
	for _, v := range events {
		acc += v.delta
		maxGuest = max(maxGuest, acc)
	}

	return maxGuest
}

func main() {
	fmt.Println(meetingRooms([][2]int{{1, 2}, {1, 3}, {2, 4}, {2, 3}}))
}

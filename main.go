package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := generateData(20)
	fmt.Println(data)

	fmt.Println(selectionSort(a))
	fmt.Println(quickSort(a))
	fmt.Println(bubbleSort(a))
	fmt.Println(mergeSort(a))
}

// selectionSort

func selectionSort(array []int) []int {
	items := make([]int, len(array))
	copy(items, array)
	if len(items) < 2 {
		return items
	}
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	return items
}

// quickSort

func quickSort(array []int) []int {
	items := make([]int, len(array))
	copy(items, array)
	if len(items) < 2 {
		return items
	}
	if len(items) == 2 {
		if items[1] < items[0] {
			items[0], items[1] = items[1], items[0]
		}
		return items
	}
	rand.Seed(time.Now().UTC().UnixNano())
	randIndex := rand.Intn(len(items))

	var low, high, equals []int
	for _, v := range items {
		if v < items[randIndex] {
			low = append(low, v)
		} else if v > items[randIndex] {
			high = append(high, v)
		} else {
			equals = append(equals, v)
		}
	}

	low = quickSort(low)
	high = quickSort(high)

	items = append(low, equals...)
	items = append(items, high...)

	return items
}

// bubbleSort

func bubbleSort(array []int) []int {
	items := make([]int, len(array))
	copy(items, array)
	if len(items) < 2 {
		return items
	}
	run := true
	for run {
		run = false
		for i := range items {
			if i+1 < len(items) && items[i] > items[i+1] {
				items[i], items[i+1] = items[i+1], items[i]
				run = true
			}
		}
	}
	return items
}

// mergeSort

func mergeSort(array []int) []int {
	items := make([]int, len(array))
	copy(items, array)
	if len(items) < 2 {
		return items
	}
	if len(items) == 2 {
		if items[1] < items[0] {
			items[0], items[1] = items[1], items[0]
		}
		return items
	}
	var arrLow, arrHigh []int
	arrLow = items[:len(items)/2]
	arrHigh = items[len(items)/2:]

	arrLow = mergeSort(arrLow)
	arrHigh = mergeSort(arrHigh)

	return merge(arrLow, arrHigh)
}

func merge(left []int, right []int) []int {
	lst := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}
	if len(left) > 0 {
		lst = append(lst, left...)
	}
	if len(right) > 0 {
		lst = append(lst, right...)
	}

	return lst
}

// generate data

func generateData(count int) []int {
	var items []int
	for i := 0; i < count; i++ {
		items = append(items, rand.Intn(count))
	}

	return items
}

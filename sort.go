package main

import (
	"sort"
)

func linearSortDescNoBuffer(unsortedInts []int) {
	for i := 0; i < len(unsortedInts); i++ {
		for j := i + 1; j < len(unsortedInts); j++ {
			if unsortedInts[j] > unsortedInts[i] {
				temp := unsortedInts[i]
				unsortedInts[i] = unsortedInts[j]
				unsortedInts[j] = temp
			}
		}
	}
}

func linearSortNoBuffer(unsortedInts []int) (sortedInts []int) {
	pendingInts := unsortedInts
	for len(pendingInts) > 1 {
		maxIntIndex := 0
		maxInt := pendingInts[maxIntIndex]
		l := len(pendingInts)
		for i := 1; i < l; i++ {
			if pendingInts[i] > maxInt {
				maxIntIndex = i
				maxInt = pendingInts[maxIntIndex]
			}
		}
		sortedInts = append(sortedInts, maxInt)
		pendingInts = append(pendingInts[:maxIntIndex], pendingInts[maxIntIndex+1:]...)
	}
	sortedInts = append(sortedInts, pendingInts[0])
	return
}

func linearSortExplorationBuffer(unsortedInts []int) (sortedInts []int) {
	var indexExplored = make([]bool, len(unsortedInts))
	pendingIntCount := len(unsortedInts)
	for pendingIntCount > 1 {
		maxIntIndex := 0
		maxInt := unsortedInts[maxIntIndex]
		indexExplored[maxIntIndex] = true
		for i := 1; i < pendingIntCount; i++ {
			if indexExplored[i] == true {
				continue
			}
			if unsortedInts[i] > maxInt {
				maxIntIndex = i
				maxInt = unsortedInts[maxIntIndex]
			}
		}
		indexExplored[maxIntIndex] = true
		sortedInts = append(sortedInts, maxInt)
		pendingIntCount--
	}
	for i := 0; i < len(indexExplored); i++ {
		if indexExplored[i] == false {
			sortedInts = append(sortedInts, unsortedInts[i])
			break
		}
	}
	return
}

func quickSortDesc(unsortedInts []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(unsortedInts)))
}

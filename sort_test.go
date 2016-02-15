package main

import (
	"fmt"
	"testing"
)

func init() {
	unsortedInts = append(unsortedInts, unsortedInts...)
	unsortedInts = append(unsortedInts, unsortedInts...)
	unsortedInts = append(unsortedInts, unsortedInts...)
	unsortedInts = append(unsortedInts, unsortedInts...)
	fmt.Printf("Number of ints to be sorted: %d\n", len(unsortedInts))
}

func TestLinearSort(t *testing.T) {
	ints := unsortedInts
	linearSortDescNoBuffer(ints)
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			t.Fatalf("The array is not sorted in descending order")
		}
	}
	t.Logf("%v", ints)
}

func TestQuickSort(t *testing.T) {
	ints := unsortedInts
	quickSortDesc(ints)
	for i := 0; i < len(ints)-1; i++ {
		if ints[i] < ints[i+1] {
			t.Fatalf("The array is not sorted in descending order")
		}
	}
	t.Logf("%v", ints)
}

func TestLinearSortNoBuffer(t *testing.T) {
	sortedInts := linearSortNoBuffer(unsortedInts)
	for i := 0; i < len(sortedInts)-1; i++ {
		if sortedInts[i] < sortedInts[i+1] {
			t.Fatalf("The array is not sorted in descending order")
		}
	}
	t.Logf("%v", sortedInts)
}

func TestLinearSortExplorationBuffer(t *testing.T) {
	sortedInts := linearSortExplorationBuffer(unsortedInts)
	for i := 0; i < len(sortedInts)-1; i++ {
		if sortedInts[i] < sortedInts[i+1] {
			t.Fatalf("The array is not sorted in descending order")
		}
	}
	t.Logf("%v", sortedInts)
}

func BenchmarkLinearSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ints := unsortedInts
		linearSortDescNoBuffer(ints)
	}
}

func BenchmarkQuicksort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ints := unsortedInts
		quickSortDesc(ints)
	}
}

func BenchmarkLinearSortNoBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		linearSortNoBuffer(unsortedInts)
	}
}

func BenchmarkLinearSortExplorationBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		linearSortExplorationBuffer(unsortedInts)
	}
}

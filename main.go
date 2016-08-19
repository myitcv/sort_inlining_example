package main

import "fmt"

func main() {
	x := []string{"b", "a"}
	sortStrings(x)
	fmt.Println(x)
}

func dataLess(vs []string, i, j int) bool {
	return vs[i] < vs[j]
}

func sortStrings(vs []string) {
	n := len(vs)
	maxDepth := 0
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	quickSort(vs, 0, n, maxDepth)
}

func dataSwap(data []string, i, j int) {
	data[i], data[j] = data[j], data[i]
}

func quickSort(data []string, a, b, maxDepth int) {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			heapSort(data, a, b)
			return
		}
		maxDepth--
		mlo, mhi := doPivot(data, a, b)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			quickSort(data, a, mlo, maxDepth)
			a = mhi // i.e., quickSort(data, mhi, b)
		} else {
			quickSort(data, mhi, b, maxDepth)
			b = mlo // i.e., quickSort(data, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if dataLess(data, i, i-6) {
				dataSwap(data, i, i-6)
			}
		}
		insertionSort(data, a, b)
	}
}

func doPivot(data []string, lo, hi int) (midlo, midhi int) {
	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(data, lo, lo+s, lo+2*s)
		medianOfThree(data, m, m-s, m+s)
		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
	}
	medianOfThree(data, lo, m, hi-1)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo < i < a] < pivot
	//	data[a <= i < b] <= pivot
	//	data[b <= i < c] unexamined
	//	data[c <= i < hi-1] > pivot
	//	data[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && dataLess(data, a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !dataLess(data, pivot, b); b++ { // data[b] <= pivot
		}
		for ; b < c && dataLess(data, pivot, c-1); c-- { // data[c-1] > pivot
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] <= pivot
		dataSwap(data, b, c-1)
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !dataLess(data, pivot, hi-1) { // data[hi-1] = pivot
			dataSwap(data, c, hi-1)
			c++
			dups++
		}
		if !dataLess(data, b-1, pivot) { // data[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> data[m] <= pivot
		if !dataLess(data, m, pivot) { // data[m] = pivot
			dataSwap(data, m, b-1)
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	data[a <= i < b] unexamined
		//	data[b <= i < c] = pivot
		for {
			for ; a < b && !dataLess(data, b-1, pivot); b-- { // data[b] == pivot
			}
			for ; a < b && dataLess(data, a, pivot); a++ { // data[a] < pivot
			}
			if a >= b {
				break
			}
			// data[a] == pivot; data[b-1] < pivot
			dataSwap(data, a, b-1)
			a++
			b--
		}
	}
	// Swap pivot into middle
	dataSwap(data, pivot, b-1)
	return b - 1, c
}

// Insertion sort
func insertionSort(data []string, a, b int) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && dataLess(data, j, j-1); j-- {
			dataSwap(data, j, j-1)
		}
	}
}

// siftDown implements the heap property on data[lo, hi).
// first is an offset into the array where the root of the heap lies.
func siftDown(data []string, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && dataLess(data, first+child, first+child+1) {
			child++
		}
		if !dataLess(data, first+root, first+child) {
			return
		}
		dataSwap(data, first+root, first+child)
		root = child
	}
}

func heapSort(data []string, a, b int) {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		dataSwap(data, first, first+i)
		siftDown(data, lo, i, first)
	}
}

// Quicksort, loosely following Bentley and McIlroy,
// ``Engineering a Sort Function,'' SP&E November 1993.

// medianOfThree moves the median of the three values data[m0], data[m1], data[m2] into data[m1].
func medianOfThree(data []string, m1, m0, m2 int) {
	// sort 3 elements
	if dataLess(data, m1, m0) {
		dataSwap(data, m1, m0)
	}
	// data[m0] <= data[m1]
	if dataLess(data, m2, m1) {
		dataSwap(data, m2, m1)
		// data[m0] <= data[m2] && data[m1] < data[m2]
		if dataLess(data, m1, m0) {
			dataSwap(data, m1, m0)
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRange(data []string, a, b, n int) {
	for i := 0; i < n; i++ {
		dataSwap(data, a+i, b+i)
	}
}

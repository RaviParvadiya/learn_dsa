package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 1. Lomuto Partition Scheme
func quicksortLomuto(arr []int, low, high int) {
	if low < high {
		pi := partitionLomuto(arr, low, high)
		quicksortLomuto(arr, low, pi-1)
		quicksortLomuto(arr, pi+1, high)
	}
}

func partitionLomuto(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			if j != i { // Optional prevent self swap
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 2. Hoare Partition Scheme (Original - pivot at low)
func quicksortHoare(arr []int, low, high int) {
	if low < high {
		pi := partitionHoare(arr, low, high)
		quicksortHoare(arr, low, pi)
		quicksortHoare(arr, pi+1, high)
	}
}

func partitionHoare(arr []int, low, high int) int {
	pivot := arr[low]
	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 2b. Hoare Partition Scheme (Modified - pivot at high)
func quicksortHoareLastPivot(arr []int, low, high int) {
	if low < high {
		pi := partitionHoareLastPivot(arr, low, high)
		quicksortHoareLastPivot(arr, low, pi)
		quicksortHoareLastPivot(arr, pi+1, high)
	}
}

func partitionHoareLastPivot(arr []int, low, high int) int {
	pivot := arr[high] // Use last element as pivot
	i := low - 1
	j := high

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if j < low || arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 2c. Hoare Partition Scheme (Modified - pivot at mid)
func quicksortHoareMidPivot(arr []int, low, high int) {
	if low < high {
		pi := partitionHoareMidPivot(arr, low, high)
		quicksortHoareMidPivot(arr, low, pi)
		quicksortHoareMidPivot(arr, pi+1, high)
	}
}

func partitionHoareMidPivot(arr []int, low, high int) int {
	mid := low + (high-low)/2
	pivot := arr[mid]                       // Use mid element as pivot
	arr[mid], arr[low] = arr[low], arr[mid] // Move pivot to start

	i := low - 1
	j := high + 1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// 3. Three-Way (Dutch National Flag) Quicksort
func quicksortThreeWay(arr []int, low, high int) {
	if low < high {
		lt, gt := partitionThreeWay(arr, low, high)
		quicksortThreeWay(arr, low, lt-1)
		quicksortThreeWay(arr, gt+1, high)
	}
}

func partitionThreeWay(arr []int, low, high int) (int, int) {
	pivot := arr[low]
	lt := low
	gt := high
	i := low

	for i <= gt {
		if arr[i] < pivot {
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		} else if arr[i] > pivot {
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

// 4. Randomized Quicksort
func quicksortRandomized(arr []int, low, high int) {
	if low < high {
		pi := partitionRandomized(arr, low, high)
		quicksortRandomized(arr, low, pi-1)
		quicksortRandomized(arr, pi+1, high)
	}
}

func partitionRandomized(arr []int, low, high int) int {
	// Random pivot selection
	randomIndex := low + rand.Intn(high-low+1)
	arr[randomIndex], arr[high] = arr[high], arr[randomIndex]

	// Use Lomuto partition with random pivot
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// 5. Iterative Quicksort
func quicksortIterative(arr []int, low, high int) {
	stack := make([]int, high-low+1)
	top := -1

	// Push initial values
	top++
	stack[top] = low
	top++
	stack[top] = high

	for top >= 0 {
		// Pop high and low
		high = stack[top]
		top--
		low = stack[top]
		top--

		// Partition
		pi := partitionLomuto(arr, low, high)

		// Push left side to stack
		if pi-1 > low {
			top++
			stack[top] = low
			top++
			stack[top] = pi - 1
		}

		// Push right side to stack
		if pi+1 < high {
			top++
			stack[top] = pi + 1
			top++
			stack[top] = high
		}
	}
}

// 6. Hybrid Quicksort (with Insertion Sort)
func quicksortHybrid(arr []int, low, high int) {
	const threshold = 10

	if high-low < threshold {
		insertionSort(arr, low, high)
		return
	}

	if low < high {
		pi := partitionLomuto(arr, low, high)
		quicksortHybrid(arr, low, pi-1)
		quicksortHybrid(arr, pi+1, high)
	}
}

func insertionSort(arr []int, low, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 7. Dual-Pivot Quicksort
func quicksortDualPivot(arr []int, low, high int) {
	if low < high {
		lp, rp := partitionDualPivot(arr, low, high)
		quicksortDualPivot(arr, low, lp-1)
		quicksortDualPivot(arr, lp+1, rp-1)
		quicksortDualPivot(arr, rp+1, high)
	}
}

func partitionDualPivot(arr []int, low, high int) (int, int) {
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}

	p := arr[low]
	q := arr[high]

	lt := low + 1
	gt := high - 1
	i := low + 1

	for i <= gt {
		if arr[i] < p {
			arr[i], arr[lt] = arr[lt], arr[i]
			lt++
		} else if arr[i] >= q {
			for arr[gt] > q && i < gt {
				gt--
			}
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
			if arr[i] < p {
				arr[i], arr[lt] = arr[lt], arr[i]
				lt++
			}
		}
		i++
	}

	lt--
	gt++

	arr[low], arr[lt] = arr[lt], arr[low]
	arr[high], arr[gt] = arr[gt], arr[high]

	return lt, gt
}

// Helper function to copy array
func copyArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	original := []int{64, 34, 25, 12, 22, 11, 90, 88, 45, 50, 23, 36}

	// Test each variant
	fmt.Println("Original array:", original)
	fmt.Println()

	// 1. Lomuto
	arr1 := copyArray(original)
	quicksortLomuto(arr1, 0, len(arr1)-1)
	fmt.Println("1. Lomuto Partition:", arr1)

	// 2. Hoare
	arr2 := copyArray(original)
	quicksortHoare(arr2, 0, len(arr2)-1)
	fmt.Println("2. Hoare Partition:", arr2)

	// 2b. Hoare with last element pivot
	arr2b := copyArray(original)
	quicksortHoareLastPivot(arr2b, 0, len(arr2b)-1)
	fmt.Println("2b. Hoare Partition (last element pivot):", arr2b)

	// 2c. Hoare with mid element pivot
	arr2c := copyArray(original)
	quicksortHoareMidPivot(arr2c, 0, len(arr2c)-1)
	fmt.Println("2c. Hoare Partition (mid element pivot):", arr2c)

	// 3. Three-Way
	arr3 := copyArray(original)
	quicksortThreeWay(arr3, 0, len(arr3)-1)
	fmt.Println("3. Three-Way (Dutch National Flag):", arr3)

	// 4. Randomized
	arr4 := copyArray(original)
	quicksortRandomized(arr4, 0, len(arr4)-1)
	fmt.Println("4. Randomized Quicksort:", arr4)

	// 5. Iterative
	arr5 := copyArray(original)
	quicksortIterative(arr5, 0, len(arr5)-1)
	fmt.Println("5. Iterative Quicksort:", arr5)

	// 6. Hybrid
	arr6 := copyArray(original)
	quicksortHybrid(arr6, 0, len(arr6)-1)
	fmt.Println("6. Hybrid Quicksort:", arr6)

	// 7. Dual-Pivot
	arr7 := copyArray(original)
	quicksortDualPivot(arr7, 0, len(arr7)-1)
	fmt.Println("7. Dual-Pivot Quicksort:", arr7)
}

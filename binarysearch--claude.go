package main

import "fmt"

// 1. Basic Iterative Binary Search
func binarySearchIterative(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return -1
}

// 2. Recursive Binary Search
func binarySearchRecursive(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}
	
	mid := low + (high-low)/2
	
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, high)
	} else {
		return binarySearchRecursive(arr, target, low, mid-1)
	}
}

// 3. Binary Search - First Occurrence (leftmost)
func binarySearchFirst(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	result := -1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid] == target {
			result = mid
			high = mid - 1 // Continue searching left
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return result
}

// 4. Binary Search - Last Occurrence (rightmost)
func binarySearchLast(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	result := -1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid] == target {
			result = mid
			low = mid + 1 // Continue searching right
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return result
}

// 5. Binary Search - Count Occurrences
func binarySearchCount(arr []int, target int) int {
	first := binarySearchFirst(arr, target)
	if first == -1 {
		return 0
	}
	last := binarySearchLast(arr, target)
	return last - first + 1
}

// 6. Binary Search - Lower Bound (first element >= target)
func binarySearchLowerBound(arr []int, target int) int {
	low := 0
	high := len(arr)
	
	for low < high {
		mid := low + (high-low)/2
		
		if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	
	return low
}

// 7. Binary Search - Upper Bound (first element > target)
func binarySearchUpperBound(arr []int, target int) int {
	low := 0
	high := len(arr)
	
	for low < high {
		mid := low + (high-low)/2
		
		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	
	return low
}

// 8. Binary Search on Infinite Array (or unknown size)
func binarySearchInfinite(arr []int, target int) int {
	// Find range where target might exist
	low := 0
	high := 1
	
	// Exponentially increase high until we find a range
	for high < len(arr) && arr[high] < target {
		low = high
		high *= 2
	}
	
	// Ensure high doesn't exceed array bounds
	if high >= len(arr) {
		high = len(arr) - 1
	}
	
	// Standard binary search in this range
	return binarySearchRecursive(arr, target, low, high)
}

// 9. Binary Search on Rotated Sorted Array
func binarySearchRotated(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid] == target {
			return mid
		}
		
		// Check which half is sorted
		if arr[low] <= arr[mid] {
			// Left half is sorted
			if target >= arr[low] && target < arr[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// Right half is sorted
			if target > arr[mid] && target <= arr[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	
	return -1
}

// 10. Binary Search - Find Peak Element
func binarySearchPeak(arr []int) int {
	low := 0
	high := len(arr) - 1
	
	for low < high {
		mid := low + (high-low)/2
		
		if arr[mid] < arr[mid+1] {
			// Peak is on the right
			low = mid + 1
		} else {
			// Peak is on the left or mid is peak
			high = mid
		}
	}
	
	return low
}

// 11. Binary Search - Square Root (integer)
func binarySearchSqrt(n int) int {
	if n < 2 {
		return n
	}
	
	low := 1
	high := n / 2
	result := 0
	
	for low <= high {
		mid := low + (high-low)/2
		square := mid * mid
		
		if square == n {
			return mid
		} else if square < n {
			result = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return result
}

// 12. Ternary Search (divide into 3 parts)
func ternarySearch(arr []int, target, low, high int) int {
	if low > high {
		return -1
	}
	
	mid1 := low + (high-low)/3
	mid2 := high - (high-low)/3
	
	if arr[mid1] == target {
		return mid1
	}
	if arr[mid2] == target {
		return mid2
	}
	
	if target < arr[mid1] {
		return ternarySearch(arr, target, low, mid1-1)
	} else if target > arr[mid2] {
		return ternarySearch(arr, target, mid2+1, high)
	} else {
		return ternarySearch(arr, target, mid1+1, mid2-1)
	}
}

// 13. Exponential Search (for unbounded/infinite arrays)
func exponentialSearch(arr []int, target int) int {
	n := len(arr)
	
	// If target is at first position
	if arr[0] == target {
		return 0
	}
	
	// Find range for binary search by doubling
	i := 1
	for i < n && arr[i] <= target {
		i *= 2
	}
	
	// Binary search in found range
	low := i / 2
	high := min(i, n-1)
	
	return binarySearchRecursive(arr, target, low, high)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 14. Interpolation Search (better for uniformly distributed data)
func interpolationSearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	
	for low <= high && target >= arr[low] && target <= arr[high] {
		if low == high {
			if arr[low] == target {
				return low
			}
			return -1
		}
		
		// Probing position with interpolation
		pos := low + ((target-arr[low])*(high-low))/(arr[high]-arr[low])
		
		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}
	
	return -1
}

// 15. Binary Search in 2D Matrix (row-wise and column-wise sorted)
func binarySearch2D(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	
	rows := len(matrix)
	cols := len(matrix[0])
	
	// Start from top-right corner
	row := 0
	col := cols - 1
	
	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	
	return false
}

// 16. Binary Search - Find Closest Element
func binarySearchClosest(arr []int, target int) int {
	n := len(arr)
	
	// Edge cases
	if target <= arr[0] {
		return arr[0]
	}
	if target >= arr[n-1] {
		return arr[n-1]
	}
	
	low := 0
	high := n - 1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid] == target {
			return arr[mid]
		}
		
		if arr[mid] < target {
			// Check if target is closer to mid or mid+1
			if mid+1 <= high && target < arr[mid+1] {
				if target-arr[mid] < arr[mid+1]-target {
					return arr[mid]
				}
				return arr[mid+1]
			}
			low = mid + 1
		} else {
			// Check if target is closer to mid or mid-1
			if mid-1 >= low && target > arr[mid-1] {
				if arr[mid]-target < target-arr[mid-1] {
					return arr[mid]
				}
				return arr[mid-1]
			}
			high = mid - 1
		}
	}
	
	return arr[low]
}

// 17. Binary Search with Custom Comparator
type Comparator func(a, b int) int // returns -1, 0, 1

func binarySearchCustom(arr []int, target int, comp Comparator) int {
	low := 0
	high := len(arr) - 1
	
	for low <= high {
		mid := low + (high-low)/2
		cmp := comp(arr[mid], target)
		
		if cmp == 0 {
			return mid
		} else if cmp < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return -1
}

// 18. Binary Search for Objects
type Person struct {
	Name string
	Age  int
}

func binarySearchObjects(arr []Person, targetAge int) int {
	low := 0
	high := len(arr) - 1
	
	for low <= high {
		mid := low + (high-low)/2
		
		if arr[mid].Age == targetAge {
			return mid
		} else if arr[mid].Age < targetAge {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 13
	
	fmt.Println("Sorted array:", arr)
	fmt.Println("Target:", target)
	fmt.Println()
	
	// 1. Iterative
	fmt.Println("1. Iterative Binary Search:", binarySearchIterative(arr, target))
	
	// 2. Recursive
	fmt.Println("2. Recursive Binary Search:", binarySearchRecursive(arr, target, 0, len(arr)-1))
	
	// 3-5. Duplicates
	arrDup := []int{1, 2, 2, 2, 3, 4, 5, 5, 5, 6}
	fmt.Println("\nArray with duplicates:", arrDup)
	fmt.Println("3. First Occurrence of 2:", binarySearchFirst(arrDup, 2))
	fmt.Println("4. Last Occurrence of 2:", binarySearchLast(arrDup, 2))
	fmt.Println("5. Count of 5:", binarySearchCount(arrDup, 5))
	
	// 6-7. Bounds
	fmt.Println("\n6. Lower Bound of 8:", binarySearchLowerBound(arr, 8))
	fmt.Println("7. Upper Bound of 8:", binarySearchUpperBound(arr, 8))
	
	// 8. Infinite array
	fmt.Println("\n8. Binary Search Infinite Array:", binarySearchInfinite(arr, 15))
	
	// 9. Rotated array
	rotated := []int{15, 17, 19, 1, 3, 5, 7, 9, 11, 13}
	fmt.Println("\n9. Rotated Array:", rotated)
	fmt.Println("   Search 5:", binarySearchRotated(rotated, 5))
	
	// 10. Peak element
	peakArr := []int{1, 3, 20, 4, 1, 0}
	fmt.Println("\n10. Peak Element in:", peakArr)
	fmt.Println("    Peak at index:", binarySearchPeak(peakArr))
	
	// 11. Square root
	fmt.Println("\n11. Square Root of 25:", binarySearchSqrt(25))
	fmt.Println("    Square Root of 27:", binarySearchSqrt(27))
	
	// 12. Ternary search
	fmt.Println("\n12. Ternary Search for 13:", ternarySearch(arr, 13, 0, len(arr)-1))
	
	// 13. Exponential search
	fmt.Println("\n13. Exponential Search for 15:", exponentialSearch(arr, 15))
	
	// 14. Interpolation search
	uniform := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println("\n14. Interpolation Search in:", uniform)
	fmt.Println("    Search 50:", interpolationSearch(uniform, 50))
	
	// 15. 2D Matrix
	matrix := [][]int{
		{1, 4, 7, 11},
		{2, 5, 8, 12},
		{3, 6, 9, 16},
		{10, 13, 14, 17},
	}
	fmt.Println("\n15. 2D Matrix Search for 5:", binarySearch2D(matrix, 5))
	
	// 16. Closest element
	fmt.Println("\n16. Closest to 8:", binarySearchClosest(arr, 8))
	fmt.Println("    Closest to 10:", binarySearchClosest(arr, 10))
	
	// 17. Custom comparator
	fmt.Println("\n17. Custom Comparator:", 
		binarySearchCustom(arr, 13, func(a, b int) int {
			if a < b { return -1 }
			if a > b { return 1 }
			return 0
		}))
	
	// 18. Objects
	people := []Person{
		{"Alice", 20},
		{"Bob", 25},
		{"Charlie", 30},
		{"David", 35},
	}
	fmt.Println("\n18. Binary Search Objects (age 30):", binarySearchObjects(people, 30))
}
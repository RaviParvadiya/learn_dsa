package main

import "fmt"

// 1. Basic Linear Search (returns index)
func linearSearchBasic(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// 2. Linear Search with range-based for loop
func linearSearchRange(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

// 3. Linear Search (returns boolean)
func linearSearchBool(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

// 4. Linear Search from End (reverse direction)
func linearSearchReverse(arr []int, target int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

// 5. Recursive Linear Search
func linearSearchRecursive(arr []int, target, index int) int {
	// Base case: reached end
	if index >= len(arr) {
		return -1
	}
	
	// Found target
	if arr[index] == target {
		return index
	}
	
	// Recurse for next element
	return linearSearchRecursive(arr, target, index+1)
}

// 6. Linear Search - Find All Occurrences
func linearSearchAll(arr []int, target int) []int {
	indices := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			indices = append(indices, i)
		}
	}
	return indices
}

// 7. Linear Search - Count Occurrences
func linearSearchCount(arr []int, target int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			count++
		}
	}
	return count
}

// 8. Sentinel Linear Search (optimized - avoids bound check)
func linearSearchSentinel(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	
	// Save last element
	last := arr[n-1]
	
	// Put target as sentinel at end
	arr[n-1] = target
	
	i := 0
	// No need to check i < n because sentinel guarantees we'll find it
	for arr[i] != target {
		i++
	}
	
	// Restore last element
	arr[n-1] = last
	
	// Check if we found target before sentinel or at last position
	if i < n-1 || arr[n-1] == target {
		return i
	}
	
	return -1
}

// 9. Bidirectional Linear Search (search from both ends)
func linearSearchBidirectional(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	
	for left <= right {
		if arr[left] == target {
			return left
		}
		if arr[right] == target {
			return right
		}
		left++
		right--
	}
	
	return -1
}

// 10. Linear Search with Early Termination (for sorted arrays)
func linearSearchSorted(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
		// If current element is greater than target, target doesn't exist
		if arr[i] > target {
			return -1
		}
	}
	return -1
}

// 11. Linear Search with Custom Comparator
type Comparator func(a, b int) bool

func linearSearchCustom(arr []int, target int, comp Comparator) int {
	for i := 0; i < len(arr); i++ {
		if comp(arr[i], target) {
			return i
		}
	}
	return -1
}

// 12. Linear Search for Objects
type Person struct {
	Name string
	Age  int
}

func linearSearchObjects(arr []Person, targetAge int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].Age == targetAge {
			return i
		}
	}
	return -1
}

// 13. Linear Search with Predicate Function
type Predicate func(int) bool

func linearSearchPredicate(arr []int, pred Predicate) int {
	for i := 0; i < len(arr); i++ {
		if pred(arr[i]) {
			return i
		}
	}
	return -1
}

// 14. Transposition Linear Search (self-organizing - move found to front)
func linearSearchTransposition(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			// Move to front if not already there
			if i > 0 {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				return i - 1
			}
			return i
		}
	}
	return -1
}

// 15. Move-to-Front Linear Search (self-organizing - aggressive)
func linearSearchMoveToFront(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			// Move to front
			if i > 0 {
				val := arr[i]
				// Shift elements
				for j := i; j > 0; j-- {
					arr[j] = arr[j-1]
				}
				arr[0] = val
				return 0
			}
			return i
		}
	}
	return -1
}

// 16. Frequency Count Linear Search (self-organizing)
type FreqElement struct {
	value int
	freq  int
}

func linearSearchFrequency(arr []FreqElement, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i].value == target {
			arr[i].freq++
			
			// Bubble up based on frequency
			j := i
			for j > 0 && arr[j].freq > arr[j-1].freq {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				j--
			}
			
			return j
		}
	}
	return -1
}

func main() {
	arr := []int{10, 23, 45, 70, 11, 15, 23, 45}
	target := 23
	
	fmt.Println("Array:", arr)
	fmt.Println("Target:", target)
	fmt.Println()
	
	// 1. Basic
	fmt.Println("1. Basic Linear Search:", linearSearchBasic(arr, target))
	
	// 2. Range-based
	fmt.Println("2. Range-based Linear Search:", linearSearchRange(arr, target))
	
	// 3. Boolean return
	fmt.Println("3. Boolean Linear Search:", linearSearchBool(arr, target))
	
	// 4. Reverse
	fmt.Println("4. Reverse Linear Search:", linearSearchReverse(arr, target))
	
	// 5. Recursive
	fmt.Println("5. Recursive Linear Search:", linearSearchRecursive(arr, target, 0))
	
	// 6. Find all
	fmt.Println("6. Find All Occurrences:", linearSearchAll(arr, target))
	
	// 7. Count
	fmt.Println("7. Count Occurrences:", linearSearchCount(arr, target))
	
	// 8. Sentinel
	arr8 := []int{10, 23, 45, 70, 11, 15}
	fmt.Println("\n8. Sentinel Linear Search:", linearSearchSentinel(arr8, target))
	
	// 9. Bidirectional
	fmt.Println("9. Bidirectional Linear Search:", linearSearchBidirectional(arr, target))
	
	// 10. Sorted with early termination
	sortedArr := []int{5, 10, 15, 20, 25, 30}
	fmt.Println("\n10. Sorted Linear Search (early term):", linearSearchSorted(sortedArr, 15))
	fmt.Println("    Searching for 22 (doesn't exist):", linearSearchSorted(sortedArr, 22))
	
	// 11. Custom comparator
	fmt.Println("\n11. Custom Comparator (equals):", 
		linearSearchCustom(arr, target, func(a, b int) bool { return a == b }))
	
	// 12. Objects
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
	}
	fmt.Println("\n12. Linear Search Objects (age 25):", linearSearchObjects(people, 25))
	
	// 13. Predicate
	fmt.Println("\n13. Predicate (find even number):", 
		linearSearchPredicate(arr, func(x int) bool { return x%2 == 0 }))
	
	// 14. Transposition
	arr14 := []int{10, 23, 45, 70, 11, 15}
	fmt.Println("\n14. Transposition Search:")
	fmt.Println("    Before:", arr14)
	idx := linearSearchTransposition(arr14, 70)
	fmt.Println("    Found at:", idx)
	fmt.Println("    After:", arr14)
	
	// 15. Move-to-Front
	arr15 := []int{10, 23, 45, 70, 11, 15}
	fmt.Println("\n15. Move-to-Front Search:")
	fmt.Println("    Before:", arr15)
	idx2 := linearSearchMoveToFront(arr15, 70)
	fmt.Println("    Found at:", idx2)
	fmt.Println("    After:", arr15)
	
	// 16. Frequency Count
	arr16 := []FreqElement{
		{10, 0}, {23, 0}, {45, 0}, {70, 0},
	}
	fmt.Println("\n16. Frequency Count Search:")
	fmt.Println("    Initial:", arr16)
	linearSearchFrequency(arr16, 70)
	fmt.Println("    After 1 search:", arr16)
	linearSearchFrequency(arr16, 70)
	fmt.Println("    After 2 searches:", arr16)
}
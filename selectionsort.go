package main

import "fmt"

// 1. Basic Selection Sort (ascending order)
func selectionSortBasic(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find minimum element in unsorted portion
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		// Swap minimum with first unsorted element
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 2. Selection Sort (descending order)
func selectionSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find maximum element in unsorted portion
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		// Swap maximum with first unsorted element
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

// 3. Bidirectional Selection Sort (Cocktail Selection Sort)
func selectionSortBidirectional(arr []int) {
	n := len(arr)
	left := 0
	right := n - 1
	
	for left < right {
		// Find minimum in current range
		minIdx := left
		maxIdx := left
		
		for i := left; i <= right; i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
			if arr[i] > arr[maxIdx] {
				maxIdx = i
			}
		}
		
		// Swap minimum to left
		arr[left], arr[minIdx] = arr[minIdx], arr[left]
		
		// If max was at left position, it's now at minIdx
		if maxIdx == left {
			maxIdx = minIdx
		}
		
		// Swap maximum to right
		arr[right], arr[maxIdx] = arr[maxIdx], arr[right]
		
		left++
		right--
	}
}

// 4. Recursive Selection Sort
func selectionSortRecursive(arr []int, start int) {
	n := len(arr)
	
	// Base case
	if start >= n-1 {
		return
	}
	
	// Find minimum in remaining array
	minIdx := start
	for i := start + 1; i < n; i++ {
		if arr[i] < arr[minIdx] {
			minIdx = i
		}
	}
	
	// Swap
	arr[start], arr[minIdx] = arr[minIdx], arr[start]
	
	// Recurse for remaining array
	selectionSortRecursive(arr, start+1)
}

// 5. Stable Selection Sort (maintains relative order of equal elements)
func selectionSortStable(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find minimum
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		
		// Instead of swapping, shift elements to maintain stability
		key := arr[minIdx]
		for minIdx > i {
			arr[minIdx] = arr[minIdx-1]
			minIdx--
		}
		arr[i] = key
	}
}

// 6. Selection Sort with Early Termination
func selectionSortEarlyTermination(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		isSorted := true
		
		// Find minimum and check if array is sorted
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
				isSorted = false
			}
		}
		
		// If no smaller element found and rest is sorted, we're done
		if isSorted && minIdx == i {
			break
		}
		
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 7. Selection Sort for Linked List
type Node struct {
	data int
	next *Node
}

func selectionSortLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}
	
	current := head
	
	for current != nil {
		// Find minimum in remaining list
		minNode := current
		temp := current.next
		
		for temp != nil {
			if temp.data < minNode.data {
				minNode = temp
			}
			temp = temp.next
		}
		
		// Swap data
		current.data, minNode.data = minNode.data, current.data
		
		current = current.next
	}
	
	return head
}

// 8. Selection Sort with Custom Comparator
type Comparator func(a, b int) bool

func selectionSortCustom(arr []int, comp Comparator) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		selectedIdx := i
		for j := i + 1; j < n; j++ {
			if comp(arr[j], arr[selectedIdx]) {
				selectedIdx = j
			}
		}
		arr[i], arr[selectedIdx] = arr[selectedIdx], arr[i]
	}
}

// 9. Selection Sort for Objects
type Person struct {
	Name string
	Age  int
}

func selectionSortObjects(arr []Person) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j].Age < arr[minIdx].Age {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// 10. Optimized Selection Sort (skip if already in position)
func selectionSortOptimized(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		
		// Only swap if minimum is not already in position
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

// Helper functions
func copyArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func printLinkedList(head *Node) {
	current := head
	fmt.Print("[")
	for current != nil {
		fmt.Print(current.data)
		if current.next != nil {
			fmt.Print(" ")
		}
		current = current.next
	}
	fmt.Println("]")
}

func createLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{data: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		current.next = &Node{data: arr[i]}
		current = current.next
	}
	return head
}

func main() {
	original := []int{64, 25, 12, 22, 11, 90, 88, 34}
	
	fmt.Println("Original array:", original)
	fmt.Println()
	
	// 1. Basic
	arr1 := copyArray(original)
	selectionSortBasic(arr1)
	fmt.Println("1. Basic Selection Sort:", arr1)
	
	// 2. Descending
	arr2 := copyArray(original)
	selectionSortDescending(arr2)
	fmt.Println("2. Selection Sort (descending):", arr2)
	
	// 3. Bidirectional
	arr3 := copyArray(original)
	selectionSortBidirectional(arr3)
	fmt.Println("3. Bidirectional Selection Sort:", arr3)
	
	// 4. Recursive
	arr4 := copyArray(original)
	selectionSortRecursive(arr4, 0)
	fmt.Println("4. Recursive Selection Sort:", arr4)
	
	// 5. Stable
	arr5 := []int{4, 5, 3, 2, 4, 1}
	fmt.Println("\n5. Stable Selection Sort:")
	fmt.Println("   Original:", arr5)
	selectionSortStable(arr5)
	fmt.Println("   Sorted:", arr5)
	
	// 6. Early Termination
	arr6 := []int{11, 12, 22, 25, 34, 64, 88, 90}
	fmt.Println("\n6. Early Termination (already sorted):", arr6)
	selectionSortEarlyTermination(arr6)
	fmt.Println("   Result:", arr6)
	
	// 7. Linked List
	fmt.Print("\n7. Linked List Selection Sort: ")
	head := createLinkedList([]int{64, 25, 12, 22, 11})
	head = selectionSortLinkedList(head)
	printLinkedList(head)
	
	// 8. Custom Comparator (descending)
	arr8 := copyArray(original)
	selectionSortCustom(arr8, func(a, b int) bool { return a > b })
	fmt.Println("\n8. Custom Comparator (descending):", arr8)
	
	// 9. Objects
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
	}
	fmt.Println("\n9. Selection Sort for Objects:")
	fmt.Println("   Original:", people)
	selectionSortObjects(people)
	fmt.Println("   Sorted by Age:", people)
	
	// 10. Optimized
	arr10 := copyArray(original)
	selectionSortOptimized(arr10)
	fmt.Println("\n10. Optimized Selection Sort:", arr10)
}
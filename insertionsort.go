package main

import "fmt"

// 1. Basic Insertion Sort (with shifting)
func insertionSortBasic(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		
		// Shift elements greater than key to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 2. Insertion Sort with Swapping (instead of shifting)
func insertionSortSwapping(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i
		// Swap adjacent elements until correct position
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
}

// 3. Binary Insertion Sort (uses binary search to find position)
func insertionSortBinary(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		
		// Binary search to find insertion position
		pos := binarySearch(arr, key, 0, i-1)
		
		// Shift elements to make space
		for j := i - 1; j >= pos; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key
	}
}

func binarySearch(arr []int, key, low, high int) int {
	for low <= high {
		mid := low + (high-low)/2
		if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// 4. Recursive Insertion Sort
func insertionSortRecursive(arr []int, n int) {
	// Base case
	if n <= 1 {
		return
	}
	
	// Sort first n-1 elements
	insertionSortRecursive(arr, n-1)
	
	// Insert last element at correct position
	key := arr[n-1]
	j := n - 2
	
	for j >= 0 && arr[j] > key {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = key
}

// 5. Sentinel Insertion Sort (avoids boundary check)
func insertionSortSentinel(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}
	
	// Find minimum and place it at the beginning (sentinel)
	minIdx := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[minIdx] {
			minIdx = i
		}
	}
	arr[0], arr[minIdx] = arr[minIdx], arr[0]
	
	// Now we can skip the j >= 0 check
	for i := 2; i < n; i++ {
		key := arr[i]
		j := i - 1
		
		// No need to check j >= 0 because sentinel guarantees we'll stop
		for arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 6. Insertion Sort with Gap (Shell Sort building block)
func insertionSortWithGap(arr []int, gap int) {
	n := len(arr)
	for i := gap; i < n; i++ {
		key := arr[i]
		j := i - gap
		
		for j >= 0 && arr[j] > key {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = key
	}
}

// 7. Bidirectional Insertion Sort (Cocktail Insertion Sort)
func insertionSortBidirectional(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		
		// Check if we need to move left or right
		if key < arr[i-1] {
			// Move left (standard insertion)
			j := i - 1
			for j >= 0 && arr[j] > key {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = key
		}
	}
}

// 8. Insertion Sort for Linked List
type Node struct {
	data int
	next *Node
}

func insertionSortLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}
	
	sorted := &Node{} // Dummy node
	current := head
	
	for current != nil {
		next := current.next
		
		// Find position to insert
		prev := sorted
		for prev.next != nil && prev.next.data < current.data {
			prev = prev.next
		}
		
		// Insert current node
		current.next = prev.next
		prev.next = current
		
		current = next
	}
	
	return sorted.next
}

// 9. Insertion Sort with Early Termination (optimized)
func insertionSortEarlyTermination(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		// If element is already in correct position, skip
		if arr[i] >= arr[i-1] {
			continue
		}
		
		key := arr[i]
		j := i - 1
		
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 10. Stable Insertion Sort for Objects
type Student struct {
	Name  string
	Score int
}

func insertionSortObjects(arr []Student) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		
		// Sort by score (stable - maintains order for equal scores)
		for j >= 0 && arr[j].Score > key.Score {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
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
	original := []int{64, 34, 25, 12, 22, 11, 90, 88}
	
	fmt.Println("Original array:", original)
	fmt.Println()
	
	// 1. Basic (with shifting)
	arr1 := copyArray(original)
	insertionSortBasic(arr1)
	fmt.Println("1. Basic Insertion Sort (shifting):", arr1)
	
	// 2. Swapping
	arr2 := copyArray(original)
	insertionSortSwapping(arr2)
	fmt.Println("2. Insertion Sort (swapping):", arr2)
	
	// 3. Binary Insertion Sort
	arr3 := copyArray(original)
	insertionSortBinary(arr3)
	fmt.Println("3. Binary Insertion Sort:", arr3)
	
	// 4. Recursive
	arr4 := copyArray(original)
	insertionSortRecursive(arr4, len(arr4))
	fmt.Println("4. Recursive Insertion Sort:", arr4)
	
	// 5. Sentinel
	arr5 := copyArray(original)
	insertionSortSentinel(arr5)
	fmt.Println("5. Sentinel Insertion Sort:", arr5)
	
	// 6. With Gap (gap = 3)
	arr6 := copyArray(original)
	insertionSortWithGap(arr6, 3)
	insertionSortWithGap(arr6, 1) // Final pass with gap=1
	fmt.Println("6. Insertion Sort with Gap:", arr6)
	
	// 7. Bidirectional
	arr7 := copyArray(original)
	insertionSortBidirectional(arr7)
	fmt.Println("7. Bidirectional Insertion Sort:", arr7)
	
	// 8. Linked List
	fmt.Print("8. Linked List Insertion Sort: ")
	head := createLinkedList([]int{64, 34, 25, 12, 22})
	head = insertionSortLinkedList(head)
	printLinkedList(head)
	
	// 9. Early Termination
	arr9 := copyArray(original)
	insertionSortEarlyTermination(arr9)
	fmt.Println("9. Early Termination Insertion Sort:", arr9)
	
	// 10. Objects
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
		{"David", 92},
		{"Eve", 88},
	}
	fmt.Println("\n10. Insertion Sort for Objects:")
	fmt.Println("    Original:", students)
	insertionSortObjects(students)
	fmt.Println("    Sorted:", students)
}
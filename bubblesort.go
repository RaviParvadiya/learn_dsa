package main

import "fmt"

// 1. Basic Bubble Sort (no optimization)
func bubbleSortBasic(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 2. Optimized Bubble Sort (with swap flag)
func bubbleSortOptimized(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no swaps occurred, array is sorted
		if !swapped {
			break
		}
	}
}

// 3. Recursive Bubble Sort
func bubbleSortRecursive(arr []int, n int) {
	// Base case
	if n == 1 {
		return
	}
	
	// One pass of bubble sort - bubble largest to end
	swapped := false
	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
			swapped = true
		}
	}
	
	// Early termination if no swaps
	if !swapped {
		return
	}
	
	// Recurse for remaining array
	bubbleSortRecursive(arr, n-1)
}

// 4. Cocktail Shaker Sort (Bidirectional Bubble Sort)
func bubbleSortCocktail(arr []int) {
	n := len(arr)
	swapped := true
	start := 0
	end := n - 1
	
	for swapped {
		swapped = false
		
		// Forward pass (like normal bubble sort)
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		
		// If no swaps, array is sorted
		if !swapped {
			break
		}
		
		end--
		swapped = false
		
		// Backward pass
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		
		start++
	}
}

// 5. Odd-Even Sort (Parallel Bubble Sort variant)
func bubbleSortOddEven(arr []int) {
	n := len(arr)
	sorted := false
	
	for !sorted {
		sorted = true
		
		// Odd phase
		for i := 1; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		
		// Even phase
		for i := 0; i < n-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
}

// 6. Comb Sort (improved bubble sort with gap)
func bubbleSortComb(arr []int) {
	n := len(arr)
	gap := n
	shrink := 1.3
	swapped := true
	
	for gap > 1 || swapped {
		// Update gap
		gap = int(float64(gap) / shrink)
		if gap < 1 {
			gap = 1
		}
		
		swapped = false
		
		// Compare elements gap distance apart
		for i := 0; i+gap < n; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}
	}
}

// 7. Bubble Sort (Descending Order)
func bubbleSortDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] < arr[j+1] { // Changed comparison
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// 8. Bubble Sort for Linked List
type Node struct {
	data int
	next *Node
}

func bubbleSortLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}
	
	swapped := true
	
	for swapped {
		swapped = false
		current := head
		
		for current.next != nil {
			if current.data > current.next.data {
				// Swap data
				current.data, current.next.data = current.next.data, current.data
				swapped = true
			}
			current = current.next
		}
	}
	
	return head
}

// 9. Bubble Sort with Custom Comparator
type Comparator func(a, b int) bool

func bubbleSortCustom(arr []int, comp Comparator) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// 10. Bubble Sort for Objects
type Student struct {
	Name  string
	Score int
}

func bubbleSortObjects(arr []Student) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j].Score > arr[j+1].Score {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// 11. Bubble Sort with Last Swap Optimization
func bubbleSortLastSwap(arr []int) {
	n := len(arr)
	newN := n
	
	for n > 1 {
		newN = 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				newN = i // Remember last swap position
			}
		}
		n = newN // All elements after last swap are sorted
	}
}

// 12. Bubble Sort with Count (for educational purposes)
func bubbleSortWithCount(arr []int) (comparisons, swaps int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			comparisons++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return
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
	original := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Println("Original array:", original)
	fmt.Println()
	
	// 1. Basic
	arr1 := copyArray(original)
	bubbleSortBasic(arr1)
	fmt.Println("1. Basic Bubble Sort:", arr1)
	
	// 2. Optimized
	arr2 := copyArray(original)
	bubbleSortOptimized(arr2)
	fmt.Println("2. Optimized Bubble Sort:", arr2)
	
	// 3. Recursive
	arr3 := copyArray(original)
	bubbleSortRecursive(arr3, len(arr3))
	fmt.Println("3. Recursive Bubble Sort:", arr3)
	
	// 4. Cocktail Shaker
	arr4 := copyArray(original)
	bubbleSortCocktail(arr4)
	fmt.Println("4. Cocktail Shaker Sort:", arr4)
	
	// 5. Odd-Even
	arr5 := copyArray(original)
	bubbleSortOddEven(arr5)
	fmt.Println("5. Odd-Even Sort:", arr5)
	
	// 6. Comb Sort
	arr6 := copyArray(original)
	bubbleSortComb(arr6)
	fmt.Println("6. Comb Sort:", arr6)
	
	// 7. Descending
	arr7 := copyArray(original)
	bubbleSortDescending(arr7)
	fmt.Println("7. Bubble Sort (descending):", arr7)
	
	// 8. Linked List
	fmt.Print("8. Linked List Bubble Sort: ")
	head := createLinkedList([]int{64, 34, 25, 12, 22})
	head = bubbleSortLinkedList(head)
	printLinkedList(head)
	
	// 9. Custom Comparator
	arr9 := copyArray(original)
	bubbleSortCustom(arr9, func(a, b int) bool { return a > b })
	fmt.Println("9. Custom Comparator (desc):", arr9)
	
	// 10. Objects
	students := []Student{
		{"Alice", 85},
		{"Bob", 92},
		{"Charlie", 78},
		{"David", 95},
	}
	fmt.Println("\n10. Bubble Sort for Objects:")
	fmt.Println("    Original:", students)
	bubbleSortObjects(students)
	fmt.Println("    Sorted by Score:", students)
	
	// 11. Last Swap Optimization
	arr11 := copyArray(original)
	bubbleSortLastSwap(arr11)
	fmt.Println("\n11. Last Swap Optimization:", arr11)
	
	// 12. With Count
	arr12 := copyArray(original)
	comp, swaps := bubbleSortWithCount(arr12)
	fmt.Printf("\n12. With Statistics: %v\n", arr12)
	fmt.Printf("    Comparisons: %d, Swaps: %d\n", comp, swaps)
}
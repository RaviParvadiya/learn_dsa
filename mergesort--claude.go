package main

import "fmt"

// 1. Basic Recursive Merge Sort (Top-Down)
func mergeSortBasic(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSortBasic(arr[:mid])
	right := mergeSortBasic(arr[mid:])
	
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

// 2. In-Place Merge Sort (modifies original array)
func mergeSortInPlace(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		
		mergeSortInPlace(arr, left, mid)
		mergeSortInPlace(arr, mid+1, right)
		
		mergeInPlace(arr, left, mid, right)
	}
}

func mergeInPlace(arr []int, left, mid, right int) {
	// Create temp arrays
	n1 := mid - left + 1
	n2 := right - mid
	
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)
	
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	// Merge back into original array
	i, j, k := 0, 0, left
	
	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// 3. Iterative (Bottom-Up) Merge Sort
func mergeSortIterative(arr []int) {
	n := len(arr)
	
	// Start with merge subarrays of size 1, then 2, 4, 8...
	for size := 1; size < n; size *= 2 {
		// Pick starting index of left sub array
		for left := 0; left < n-1; left += 2 * size {
			mid := min(left+size-1, n-1)
			right := min(left+2*size-1, n-1)
			
			mergeInPlace(arr, left, mid, right)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 4. Three-Way Merge Sort (divides into 3 parts)
func mergeSortThreeWay(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	// Divide into three parts
	third := len(arr) / 3
	
	left := mergeSortThreeWay(arr[:third])
	mid := mergeSortThreeWay(arr[third : 2*third])
	right := mergeSortThreeWay(arr[2*third:])
	
	return mergeThree(left, mid, right)
}

func mergeThree(left, mid, right []int) []int {
	result := make([]int, 0, len(left)+len(mid)+len(right))
	i, j, k := 0, 0, 0
	
	// Merge all three
	for i < len(left) && j < len(mid) && k < len(right) {
		if left[i] <= mid[j] && left[i] <= right[k] {
			result = append(result, left[i])
			i++
		} else if mid[j] <= left[i] && mid[j] <= right[k] {
			result = append(result, mid[j])
			j++
		} else {
			result = append(result, right[k])
			k++
		}
	}
	
	// Merge remaining two
	for i < len(left) && j < len(mid) {
		if left[i] <= mid[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, mid[j])
			j++
		}
	}
	
	for i < len(left) && k < len(right) {
		if left[i] <= right[k] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[k])
			k++
		}
	}
	
	for j < len(mid) && k < len(right) {
		if mid[j] <= right[k] {
			result = append(result, mid[j])
			j++
		} else {
			result = append(result, right[k])
			k++
		}
	}
	
	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, mid[j:]...)
	result = append(result, right[k:]...)
	
	return result
}

// 5. Hybrid Merge Sort (with Insertion Sort for small arrays)
func mergeSortHybrid(arr []int, left, right int) {
	const threshold = 10
	
	if right-left <= threshold {
		insertionSort(arr, left, right)
		return
	}
	
	if left < right {
		mid := left + (right-left)/2
		mergeSortHybrid(arr, left, mid)
		mergeSortHybrid(arr, mid+1, right)
		mergeInPlace(arr, left, mid, right)
	}
}

func insertionSort(arr []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		key := arr[i]
		j := i - 1
		for j >= left && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 6. Natural Merge Sort (takes advantage of existing order)
func mergeSortNatural(arr []int) {
	n := len(arr)
	
	for {
		runs := findRuns(arr)
		
		// If only one run, array is sorted
		if len(runs) <= 1 {
			break
		}
		
		// Merge adjacent runs
		i := 0
		for i < len(runs)-1 {
			left := runs[i]
			mid := runs[i+1] - 1
			right := n - 1
			if i+2 < len(runs) {
				right = runs[i+2] - 1
			}
			
			mergeInPlace(arr, left, mid, right)
			i += 2
		}
	}
}

func findRuns(arr []int) []int {
	runs := []int{0}
	n := len(arr)
	
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			runs = append(runs, i)
		}
	}
	
	return runs
}

// 7. Parallel Merge Sort (conceptual - uses goroutines)
func mergeSortParallel(arr []int, depth int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	
	// Use goroutines for parallel execution up to certain depth
	if depth > 0 {
		leftChan := make(chan []int)
		rightChan := make(chan []int)
		
		go func() {
			leftChan <- mergeSortParallel(arr[:mid], depth-1)
		}()
		
		go func() {
			rightChan <- mergeSortParallel(arr[mid:], depth-1)
		}()
		
		left := <-leftChan
		right := <-rightChan
		
		return merge(left, right)
	}
	
	// Below threshold, use sequential
	left := mergeSortParallel(arr[:mid], 0)
	right := mergeSortParallel(arr[mid:], 0)
	
	return merge(left, right)
}

// 8. Merge Sort for Linked List
type Node struct {
	data int
	next *Node
}

func mergeSortLinkedList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}
	
	// Split the list into two halves
	mid := getMiddle(head)
	midNext := mid.next
	mid.next = nil
	
	// Recursively sort both halves
	left := mergeSortLinkedList(head)
	right := mergeSortLinkedList(midNext)
	
	// Merge sorted halves
	return mergeLists(left, right)
}

func getMiddle(head *Node) *Node {
	if head == nil {
		return head
	}
	
	slow := head
	fast := head.next
	
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	
	return slow
}

func mergeLists(left, right *Node) *Node {
	dummy := &Node{}
	current := dummy
	
	for left != nil && right != nil {
		if left.data <= right.data {
			current.next = left
			left = left.next
		} else {
			current.next = right
			right = right.next
		}
		current = current.next
	}
	
	if left != nil {
		current.next = left
	}
	if right != nil {
		current.next = right
	}
	
	return dummy.next
}

// 9. Merge Sort with Custom Comparator
type Comparator func(a, b int) bool

func mergeSortCustom(arr []int, comp Comparator) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSortCustom(arr[:mid], comp)
	right := mergeSortCustom(arr[mid:], comp)
	
	return mergeCustom(left, right, comp)
}

func mergeCustom(left, right []int, comp Comparator) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if comp(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

// 10. Merge Sort for Objects
type Person struct {
	Name string
	Age  int
}

func mergeSortObjects(arr []Person) []Person {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := mergeSortObjects(arr[:mid])
	right := mergeSortObjects(arr[mid:])
	
	return mergeObjects(left, right)
}

func mergeObjects(left, right []Person) []Person {
	result := make([]Person, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i].Age <= right[j].Age {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
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
	original := []int{38, 27, 43, 3, 9, 82, 10}
	
	fmt.Println("Original array:", original)
	fmt.Println()
	
	// 1. Basic Recursive
	arr1 := copyArray(original)
	result1 := mergeSortBasic(arr1)
	fmt.Println("1. Basic Recursive Merge Sort:", result1)
	
	// 2. In-Place
	arr2 := copyArray(original)
	mergeSortInPlace(arr2, 0, len(arr2)-1)
	fmt.Println("2. In-Place Merge Sort:", arr2)
	
	// 3. Iterative (Bottom-Up)
	arr3 := copyArray(original)
	mergeSortIterative(arr3)
	fmt.Println("3. Iterative Merge Sort:", arr3)
	
	// 4. Three-Way
	arr4 := copyArray(original)
	result4 := mergeSortThreeWay(arr4)
	fmt.Println("4. Three-Way Merge Sort:", result4)
	
	// 5. Hybrid
	arr5 := copyArray(original)
	mergeSortHybrid(arr5, 0, len(arr5)-1)
	fmt.Println("5. Hybrid Merge Sort:", arr5)
	
	// 6. Natural
	arr6 := []int{3, 9, 27, 38, 10, 43, 82} // Has some natural runs
	fmt.Println("\n6. Natural Merge Sort:")
	fmt.Println("   Original:", arr6)
	mergeSortNatural(arr6)
	fmt.Println("   Sorted:", arr6)
	
	// 7. Parallel
	arr7 := copyArray(original)
	result7 := mergeSortParallel(arr7, 2) // depth of 2 for parallelism
	fmt.Println("\n7. Parallel Merge Sort:", result7)
	
	// 8. Linked List
	fmt.Print("\n8. Linked List Merge Sort: ")
	head := createLinkedList([]int{38, 27, 43, 3, 9})
	head = mergeSortLinkedList(head)
	printLinkedList(head)
	
	// 9. Custom Comparator (descending)
	arr9 := copyArray(original)
	result9 := mergeSortCustom(arr9, func(a, b int) bool { return a > b })
	fmt.Println("\n9. Custom Comparator (desc):", result9)
	
	// 10. Objects
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
	}
	fmt.Println("\n10. Merge Sort for Objects:")
	fmt.Println("    Original:", people)
	sorted := mergeSortObjects(people)
	fmt.Println("    Sorted by Age:", sorted)
}
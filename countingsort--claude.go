package main

import "fmt"

// 1. Basic Counting Sort (for non-negative integers)
func countingSortBasic(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	
	// Find max element
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Create count array
	count := make([]int, max+1)
	
	// Count occurrences
	for _, v := range arr {
		count[v]++
	}
	
	// Build output array
	output := make([]int, 0, len(arr))
	for i := 0; i <= max; i++ {
		for j := 0; j < count[i]; j++ {
			output = append(output, i)
		}
	}
	
	return output
}

// 2. In-Place Counting Sort (modifies original array)
func countingSortInPlace(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	// Find max element
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Create count array
	count := make([]int, max+1)
	
	// Count occurrences
	for _, v := range arr {
		count[v]++
	}
	
	// Overwrite original array
	idx := 0
	for i := 0; i <= max; i++ {
		for j := 0; j < count[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
}

// 3. Stable Counting Sort (preserves relative order of equal elements)
func countingSortStable(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	
	// Find max element
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Create count array
	count := make([]int, max+1)
	
	// Count occurrences
	for _, v := range arr {
		count[v]++
	}
	
	// Convert to cumulative count (prefix sum)
	for i := 1; i <= max; i++ {
		count[i] += count[i-1]
	}
	
	// Build output array (traverse from right to maintain stability)
	output := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		output[count[val]-1] = val
		count[val]--
	}
	
	return output
}

// 4. Counting Sort with Negative Numbers
func countingSortWithNegatives(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	
	// Find min and max
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	
	// Create count array with offset
	rangeSize := max - min + 1
	count := make([]int, rangeSize)
	
	// Count occurrences (with offset)
	for _, v := range arr {
		count[v-min]++
	}
	
	// Build output array
	output := make([]int, 0, len(arr))
	for i := 0; i < rangeSize; i++ {
		for j := 0; j < count[i]; j++ {
			output = append(output, i+min)
		}
	}
	
	return output
}

// 5. Counting Sort for Limited Range (with min and max parameters)
func countingSortLimitedRange(arr []int, min, max int) []int {
	if len(arr) == 0 {
		return arr
	}
	
	// Create count array
	rangeSize := max - min + 1
	count := make([]int, rangeSize)
	
	// Count occurrences
	for _, v := range arr {
		if v < min || v > max {
			panic("Value out of specified range")
		}
		count[v-min]++
	}
	
	// Build output array
	output := make([]int, 0, len(arr))
	for i := 0; i < rangeSize; i++ {
		for j := 0; j < count[i]; j++ {
			output = append(output, i+min)
		}
	}
	
	return output
}

// 6. Counting Sort for Objects (sorting by key)
type Person struct {
	Name string
	Age  int
}

func countingSortObjects(arr []Person) []Person {
	if len(arr) == 0 {
		return arr
	}
	
	// Find max age
	maxAge := arr[0].Age
	for _, p := range arr {
		if p.Age > maxAge {
			maxAge = p.Age
		}
	}
	
	// Create count array
	count := make([]int, maxAge+1)
	
	// Count occurrences
	for _, p := range arr {
		count[p.Age]++
	}
	
	// Convert to cumulative count
	for i := 1; i <= maxAge; i++ {
		count[i] += count[i-1]
	}
	
	// Build output array (stable sort)
	output := make([]Person, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		age := arr[i].Age
		output[count[age]-1] = arr[i]
		count[age]--
	}
	
	return output
}

// 7. Optimized Counting Sort (dynamic range)
func countingSortOptimized(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	
	// Find min and max in single pass
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	
	// If all elements are same
	if min == max {
		return arr
	}
	
	// Create count array with optimal size
	rangeSize := max - min + 1
	count := make([]int, rangeSize)
	
	// Count occurrences
	for _, v := range arr {
		count[v-min]++
	}
	
	// Build output array
	output := make([]int, len(arr))
	idx := 0
	for i := 0; i < rangeSize; i++ {
		for j := 0; j < count[i]; j++ {
			output[idx] = i + min
			idx++
		}
	}
	
	return output
}

// 8. Counting Sort as Subroutine (for Radix Sort)
func countingSortRadix(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10) // For decimal digits 0-9
	
	// Count occurrences of digits
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}
	
	// Convert to cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	
	// Build output array (stable)
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
	
	// Copy output to arr
	copy(arr, output)
}

// Helper function to copy array
func copyArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func main() {
	// Test basic counting sort
	arr1 := []int{4, 2, 2, 8, 3, 3, 1}
	fmt.Println("Original array:", arr1)
	fmt.Println("1. Basic Counting Sort:", countingSortBasic(copyArray(arr1)))
	
	// Test in-place
	arr2 := copyArray(arr1)
	countingSortInPlace(arr2)
	fmt.Println("2. In-Place Counting Sort:", arr2)
	
	// Test stable
	arr3 := copyArray(arr1)
	fmt.Println("3. Stable Counting Sort:", countingSortStable(arr3))
	
	// Test with negatives
	arr4 := []int{-5, -10, 0, -3, 8, 5, -1, 10}
	fmt.Println("\n4. With Negatives:", arr4)
	fmt.Println("   Sorted:", countingSortWithNegatives(arr4))
	
	// Test limited range
	arr5 := []int{5, 7, 6, 8, 5, 9, 7}
	fmt.Println("\n5. Limited Range (5-9):", arr5)
	fmt.Println("   Sorted:", countingSortLimitedRange(arr5, 5, 9))
	
	// Test with objects
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 25},
		{"David", 20},
		{"Eve", 30},
	}
	fmt.Println("\n6. Sorting Objects by Age:")
	fmt.Println("   Original:", people)
	fmt.Println("   Sorted:", countingSortObjects(people))
	
	// Test optimized
	arr7 := []int{100, 105, 102, 108, 103}
	fmt.Println("\n7. Optimized (dynamic range):", arr7)
	fmt.Println("   Sorted:", countingSortOptimized(arr7))
	
	// Test radix sort helper
	arr8 := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("\n8. Radix Sort (using counting sort):", arr8)
	// Find max for number of digits
	max := arr8[0]
	for _, v := range arr8 {
		if v > max {
			max = v
		}
	}
	// Apply counting sort for each digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortRadix(arr8, exp)
	}
	fmt.Println("   Sorted:", arr8)
}
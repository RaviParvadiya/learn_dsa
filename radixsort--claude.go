package main

import "fmt"

// ==================== RADIX SORT VARIANTS ====================

// 1. LSD Radix Sort (Least Significant Digit) - Most Common
func radixSortLSD(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	// Find maximum to know number of digits
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Do counting sort for every digit
	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp)
	}
}

func countingSortByDigit(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 10)
	
	// Count occurrences
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}
	
	// Cumulative count
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	
	// Build output (traverse from right for stability)
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
	
	copy(arr, output)
}

// 2. MSD Radix Sort (Most Significant Digit)
func radixSortMSD(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Find the most significant digit position
	exp := 1
	for max/exp >= 10 {
		exp *= 10
	}
	
	radixSortMSDHelper(arr, exp)
}

func radixSortMSDHelper(arr []int, exp int) {
	if len(arr) <= 1 || exp < 1 {
		return
	}
	
	// Count sort by current digit
	buckets := make([][]int, 10)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}
	
	for _, v := range arr {
		digit := (v / exp) % 10
		buckets[digit] = append(buckets[digit], v)
	}
	
	// Recursively sort each bucket
	for i := range buckets {
		if len(buckets[i]) > 1 {
			radixSortMSDHelper(buckets[i], exp/10)
		}
	}
	
	// Concatenate buckets
	idx := 0
	for i := range buckets {
		for j := range buckets[i] {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}

// 3. Binary Radix Sort (sorts by bits instead of digits)
func radixSortBinary(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Process each bit
	for bit := 0; (1 << bit) <= max; bit++ {
		countingSortByBit(arr, bit)
	}
}

func countingSortByBit(arr []int, bit int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 2)
	
	// Count 0s and 1s
	for i := 0; i < n; i++ {
		bitValue := (arr[i] >> bit) & 1
		count[bitValue]++
	}
	
	// Cumulative count
	count[1] += count[0]
	
	// Build output
	for i := n - 1; i >= 0; i-- {
		bitValue := (arr[i] >> bit) & 1
		output[count[bitValue]-1] = arr[i]
		count[bitValue]--
	}
	
	copy(arr, output)
}

// 4. Radix Sort for Negative Numbers
func radixSortWithNegatives(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	// Separate positive and negative numbers
	positive := []int{}
	negative := []int{}
	
	for _, v := range arr {
		if v >= 0 {
			positive = append(positive, v)
		} else {
			negative = append(negative, -v) // Make positive
		}
	}
	
	// Sort both arrays
	radixSortLSD(positive)
	radixSortLSD(negative)
	
	// Combine: negatives (reversed) + positives
	idx := 0
	for i := len(negative) - 1; i >= 0; i-- {
		arr[idx] = -negative[i]
		idx++
	}
	for i := 0; i < len(positive); i++ {
		arr[idx] = positive[i]
		idx++
	}
}

// 5. Radix Sort with Different Base (Base-16 for hex)
func radixSortBase16(arr []int) {
	if len(arr) == 0 {
		return
	}
	
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	
	// Process in base 16 (hexadecimal)
	for exp := 1; max/exp > 0; exp *= 16 {
		countingSortByDigitBase16(arr, exp)
	}
}

func countingSortByDigitBase16(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n)
	count := make([]int, 16) // 0-F in hex
	
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 16
		count[digit]++
	}
	
	for i := 1; i < 16; i++ {
		count[i] += count[i-1]
	}
	
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 16
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
	
	copy(arr, output)
}

// 6. In-Place MSD Radix Sort (space-efficient)
func radixSortMSDInPlace(arr []int, low, high, exp int) {
	if low >= high || exp < 1 {
		return
	}
	
	// Count array for digits 0-9
	count := make([]int, 11) // 10 digits + 1 for offset
	
	// Count occurrences
	for i := low; i <= high; i++ {
		digit := (arr[i] / exp) % 10
		count[digit+1]++
	}
	
	// Convert to starting indices
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}
	
	// Rearrange elements
	temp := make([]int, high-low+1)
	for i := low; i <= high; i++ {
		digit := (arr[i] / exp) % 10
		temp[count[digit]] = arr[i]
		count[digit]++
	}
	
	// Copy back
	for i := low; i <= high; i++ {
		arr[i] = temp[i-low]
	}
	
	// Recursively sort each bucket
	for i := 0; i < 10; i++ {
		bucketStart := low + count[i]
		bucketEnd := low + count[i+1] - 1
		if bucketStart < bucketEnd {
			radixSortMSDInPlace(arr, bucketStart, bucketEnd, exp/10)
		}
	}
}

// 7. String Radix Sort (LSD for fixed-length strings)
func radixSortStrings(arr []string) {
	if len(arr) == 0 {
		return
	}
	
	// Find max length
	maxLen := len(arr[0])
	for _, s := range arr {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	
	// Sort from rightmost character to leftmost
	for pos := maxLen - 1; pos >= 0; pos-- {
		countingSortByChar(arr, pos, maxLen)
	}
}

func countingSortByChar(arr []string, pos, maxLen int) {
	n := len(arr)
	output := make([]string, n)
	count := make([]int, 256) // ASCII characters
	
	// Count occurrences
	for i := 0; i < n; i++ {
		charIdx := 0
		if pos < len(arr[i]) {
			charIdx = int(arr[i][pos])
		}
		count[charIdx]++
	}
	
	// Cumulative count
	for i := 1; i < 256; i++ {
		count[i] += count[i-1]
	}
	
	// Build output
	for i := n - 1; i >= 0; i-- {
		charIdx := 0
		if pos < len(arr[i]) {
			charIdx = int(arr[i][pos])
		}
		output[count[charIdx]-1] = arr[i]
		count[charIdx]--
	}
	
	copy(arr, output)
}

// ==================== BUCKET SORT VARIANTS ====================

// 1. Basic Bucket Sort (for uniformly distributed data)
func bucketSortBasic(arr []float64) {
	if len(arr) == 0 {
		return
	}
	
	n := len(arr)
	
	// Create n empty buckets
	buckets := make([][]float64, n)
	for i := range buckets {
		buckets[i] = make([]float64, 0)
	}
	
	// Put elements into buckets
	for _, v := range arr {
		idx := int(v * float64(n))
		if idx == n {
			idx = n - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}
	
	// Sort individual buckets using insertion sort
	for i := range buckets {
		insertionSortFloat(buckets[i])
	}
	
	// Concatenate all buckets
	idx := 0
	for i := range buckets {
		for j := range buckets[i] {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}

func insertionSortFloat(arr []float64) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 2. Bucket Sort for Integers (with range)
func bucketSortIntegers(arr []int) {
	if len(arr) == 0 {
		return
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
	
	// Create buckets
	bucketCount := len(arr)
	bucketRange := (max - min) / bucketCount + 1
	buckets := make([][]int, bucketCount)
	
	// Distribute into buckets
	for _, v := range arr {
		idx := (v - min) / bucketRange
		if idx >= bucketCount {
			idx = bucketCount - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}
	
	// Sort each bucket
	for i := range buckets {
		insertionSortInt(buckets[i])
	}
	
	// Concatenate
	idx := 0
	for i := range buckets {
		for j := range buckets[i] {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}

func insertionSortInt(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// 3. Bucket Sort with Fixed Number of Buckets
func bucketSortFixedBuckets(arr []int, numBuckets int) {
	if len(arr) == 0 {
		return
	}
	
	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	
	// Create buckets
	bucketRange := float64(max-min+1) / float64(numBuckets)
	buckets := make([][]int, numBuckets)
	
	// Distribute
	for _, v := range arr {
		idx := int(float64(v-min) / bucketRange)
		if idx >= numBuckets {
			idx = numBuckets - 1
		}
		buckets[idx] = append(buckets[idx], v)
	}
	
	// Sort and concatenate
	idx := 0
	for i := range buckets {
		insertionSortInt(buckets[i])
		for j := range buckets[i] {
			arr[idx] = buckets[i][j]
			idx++
		}
	}
}

// Helper functions
func copyIntArray(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func copyFloatArray(arr []float64) []float64 {
	result := make([]float64, len(arr))
	copy(result, arr)
	return result
}

func main() {
	fmt.Println("============ RADIX SORT ============")
	
	// Test data
	intArr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	
	// 1. LSD Radix Sort
	arr1 := copyIntArray(intArr)
	fmt.Println("\n1. LSD Radix Sort:")
	fmt.Println("   Original:", arr1)
	radixSortLSD(arr1)
	fmt.Println("   Sorted:", arr1)
	
	// 2. MSD Radix Sort
	arr2 := copyIntArray(intArr)
	fmt.Println("\n2. MSD Radix Sort:")
	fmt.Println("   Original:", arr2)
	radixSortMSD(arr2)
	fmt.Println("   Sorted:", arr2)
	
	// 3. Binary Radix Sort
	arr3 := copyIntArray(intArr)
	fmt.Println("\n3. Binary Radix Sort:")
	fmt.Println("   Original:", arr3)
	radixSortBinary(arr3)
	fmt.Println("   Sorted:", arr3)
	
	// 4. With Negatives
	arr4 := []int{170, -45, 75, -90, 24, -2, 66}
	fmt.Println("\n4. Radix Sort with Negatives:")
	fmt.Println("   Original:", arr4)
	radixSortWithNegatives(arr4)
	fmt.Println("   Sorted:", arr4)
	
	// 5. Base-16 Radix Sort
	arr5 := copyIntArray(intArr)
	fmt.Println("\n5. Radix Sort Base-16 (Hexadecimal):")
	fmt.Println("   Original:", arr5)
	radixSortBase16(arr5)
	fmt.Println("   Sorted:", arr5)
	
	// 6. In-Place MSD Radix Sort
	arr6 := copyIntArray(intArr)
	fmt.Println("\n6. In-Place MSD Radix Sort:")
	fmt.Println("   Original:", arr6)
	max := arr6[0]
	for _, v := range arr6 {
		if v > max {
			max = v
		}
	}
	exp := 1
	for max/exp >= 10 {
		exp *= 10
	}
	radixSortMSDInPlace(arr6, 0, len(arr6)-1, exp)
	fmt.Println("   Sorted:", arr6)
	
	// 7. String Radix Sort
	strArr := []string{"abc", "aaa", "bcd", "bbb", "xyz", "aab"}
	fmt.Println("\n7. String Radix Sort:")
	fmt.Println("   Original:", strArr)
	radixSortStrings(strArr)
	fmt.Println("   Sorted:", strArr)
	
	fmt.Println("\n============ BUCKET SORT ============")
	
	// 1. Basic Bucket Sort (floats 0.0 to 1.0)
	floatArr := []float64{0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434}
	fmt.Println("\n1. Basic Bucket Sort (floats):")
	fmt.Println("   Original:", floatArr)
	bucketSortBasic(copyFloatArray(floatArr))
	fmt.Println("   Sorted:", floatArr)
	
	// 2. Bucket Sort for Integers
	arr5 := []int{29, 25, 3, 49, 9, 37, 21, 43}
	fmt.Println("\n2. Bucket Sort (integers):")
	fmt.Println("   Original:", arr5)
	bucketSortIntegers(arr5)
	fmt.Println("   Sorted:", arr5)
	
	// 3. Fixed Buckets
	arr6 := copyIntArray(intArr)
	fmt.Println("\n3. Bucket Sort (5 fixed buckets):")
	fmt.Println("   Original:", arr6)
	bucketSortFixedBuckets(arr6, 5)
	fmt.Println("   Sorted:", arr6)
	
	fmt.Println("\n========================================")
	fmt.Println("Note: Radix sort works best with counting")
	fmt.Println("sort as subroutine for O(d*n) complexity!")
}
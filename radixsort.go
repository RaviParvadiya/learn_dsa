package main

import "fmt"

func RadixSort(arr []int) {
	// Find max element
	maxValue := 0
	for _, v := range arr {
		if maxValue < v {
			maxValue = v
		}
	}

	// Preallocate buckets
	buckets := make([][]int, 10)
	for i := range buckets {
		buckets[i] = make([]int, 0, len(arr)/10+1)
	}

	div := 1
	for maxValue/div > 0 {
		radix(arr, buckets, div)
		div *= 10
	}

	fmt.Println(arr)
}

func radix(arr []int, buckets [][]int, div int) {
	// Distribute elements into buckets based on current digit
	for _, v := range arr {
		radixIdx := (v / div) % 10
		buckets[radixIdx] = append(buckets[radixIdx], v)
	}

	// Collect elements back into main array
	writeIdx := 0
	for i := range buckets {
		length := copy(arr[writeIdx:], buckets[i])
		writeIdx += length
		buckets[i] = buckets[i][:0]
	}
}

func main() {
	arr := []int{33, 45, 40, 25, 17, 24}

	RadixSort(arr)
}

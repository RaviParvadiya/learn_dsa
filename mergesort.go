package main

import "fmt"

func mergeSort1(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2
	mergeSort1(arr, low, mid)
	mergeSort1(arr, mid+1, high)

	merge1(arr, low, mid, high)
}

func merge1(arr []int, low, mid, high int) {
    temp := make([]int, 0, high-low+1)

    i := low
    j := mid + 1

    for i <= mid && j <= high {
        if arr[i] <= arr[j] {
            temp = append(temp, arr[i])
            i++
        } else {
            temp = append(temp, arr[j])
            j++
        }
    }

    // copy remaining elements
    for i <= mid {
        temp = append(temp, arr[i])
        i++
    }
    for j <= high {
        temp = append(temp, arr[j])
        j++
    }

    // copy back to original array
    for k := 0; k < len(temp); k++ {
        arr[low+k] = temp[k]
    }
}

func mergeSort2(arr []int, low, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2
	mergeSort2(arr, low, mid)
	mergeSort2(arr, mid+1, high)

	merge2(arr, low, mid, high)
}


func merge2(arr []int, low, mid, high int) {
	temp := make([]int, high-low+1)

	i := low      // left pointer
	j := mid + 1  // right pointer
	k := 0        // temp pointer

	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}

	// remaining left part
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}

	// remaining right part
	for j <= high {
		temp[k] = arr[j]
		j++
		k++
	}

	// copy back
	for x := 0; x < len(temp); x++ {
		arr[low+x] = temp[x]
	}
}


func main() {
	arr1 := []int{33, 45, 40, 25, 17, 24}
	mergeSort1(arr1, 0, len(arr1)-1)
	
	arr2 := []int{33, 45, 40, 25, 17, 24}
	mergeSort2(arr2, 0, len(arr2)-1)
}

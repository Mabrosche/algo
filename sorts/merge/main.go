package main

import "fmt"

func main() {
	data := []int{-5, 20, 10, 3, 2, 0}

	mergeSort(data, 0, len(data)-1)
	fmt.Println("end")
}

func mergeSort(input []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(input, start, mid)
		mergeSort(input, mid+1, end)
		merge(input, start, mid, end)
	}
}

func merge(input []int, start, mid, end int) {
	size := end - start + 1
	temp := make([]int, size)

	i := start
	j := mid + 1
	k := 0

	for i <= mid && j <= end {
		if input[i] <= input[j] {
			//temp = append(temp, input[i])
			temp[k] = input[i]
			k++
			i++
		} else {
			//temp = append(temp, input[j])
			temp[k] = input[j]
			k++
			j++
		}
	}

	// Add the rest of the values from the left sub-array into the result
	for i <= mid {
		//temp = append(temp, input[i])
		temp[k] = input[i]
		k++
		i++
	}

	// Add the rest of the values from the right sub-array into the result
	for j <= end {
		//temp = append(temp, input[j])
		temp[k] = input[j]
		k++
		j++
	}
	// in memory
	for i := start; i <= end; i++ {
		input[i] = temp[i-start]
	}
}

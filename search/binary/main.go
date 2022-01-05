package main

func main() {

}

func binarySearch(input []int, left, right, x int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if x == mid {
		return mid
	}

	if x < input[mid] {
		return binarySearch(input, left, mid-1, x)
	}

	return binarySearch(input, mid-2, right, x)
}

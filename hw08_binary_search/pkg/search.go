package search

func BinarySearch(target int, mass []int) int {
	left := 0
	right := len(mass) - 1

	for left <= right {
		median := (left + right) / 2

		if mass[median] < target {
			left = median + 1
		} else {
			right = median - 1
		}

		if mass[median] == target {
			return median
		}
	}
	return -1
}

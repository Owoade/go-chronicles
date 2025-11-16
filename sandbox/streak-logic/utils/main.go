package utils

// Contains checks if a slice contains a given int.
func Contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

// MaxElementInSlice returns the maximum integer in a slice.
// Returns 0 if the slice is empty.
func MaxElementInSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

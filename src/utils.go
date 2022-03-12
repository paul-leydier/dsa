package src

import "golang.org/x/exp/constraints"

func ArgMax[T constraints.Ordered](nums []T) int {
	argMax, max := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			argMax, max = i, nums[i]
		}
	}
	return argMax
}

func Max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
func ArgMin[T constraints.Ordered](nums []T) int {
	argMin, min := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			argMin, min = i, nums[i]
		}
	}
	return argMin
}

func Min(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

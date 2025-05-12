package main

import "fmt"

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		max := nums[i]
		ind := 0
		for j := 0; j < len(nums)-i-1; j++ {
			if max < nums[j] {
				max = nums[j]
			}
			ind = j
		}
		nums[i], nums[ind] = nums[ind], nums[i]
		fmt.Println("Arr:", nums)
	}
	return nums

}

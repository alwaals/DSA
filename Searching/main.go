package main

import "fmt"

func main() {
	fmt.Println("Started Service!")
	nums := []int{3, 7, 10, 13, 15}
	target := 10
	fmt.Println(linearSearch(nums, target))
	fmt.Println(binarySearch(nums, target))
	fmt.Println(binaryRecursive(nums, target, 0, len(nums)-1))
}

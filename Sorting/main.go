package main

import "fmt"

func main() {
	fmt.Println("Sorting started!")
	nums := []int{25, 10, 1, 3, 7, 12}
	fmt.Println("After bubble sorting:", bubbleSort(nums))

	fmt.Println("After selection sorting:",selectionSort(nums))
}

package main

import "fmt"

func main() {
	input := []int{2, 1, 2, 5, 6, 7, 5, 6, 7}
	result := singleNumber(input)
	fmt.Println(result)

}

func singleNumber(nums []int) int {
	count := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			count = 1
			if nums[i] == nums[j] && i != j {
				count = 2

				break
			}
		}
		if count == 1 {
			//fmt.Println(nums[i])
			result = nums[i]
		}
	}
	return result
}

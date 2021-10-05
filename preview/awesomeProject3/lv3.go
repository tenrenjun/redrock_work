package main

import (
	"fmt"
)
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
func main(){
	var slice1 []int
	var v int
	fmt.Println("input:")
	for v !=-1 {
		fmt.Scan(&v)
		slice1=append(slice1,v)
	}
	slice1=slice1[0:len(slice1)-2]
	bubbleSort(slice1)
	fmt.Println("output:",slice1)
}

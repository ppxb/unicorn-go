package test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(randomArray(arr))
}

func randomArray(nums []int) (random int) {
	randomIndex := rand.Intn(len(nums))
	random = nums[randomIndex]
	return
}

func find(nums []int, target int) (index int) {
	index = -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return
}

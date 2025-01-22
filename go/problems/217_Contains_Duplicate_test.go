package main

import (
	"fmt"
	"leet-it-go/util"
	"slices"
	"testing"
)

func containsDuplicate(nums []int) bool {
	arr := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if slices.Contains(arr, nums[i]) {
			return true
		}
		arr = append(arr, nums[i])
	}

	return false
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []util.TestCase[[]int, bool]{
		{
			Input:    []int{1, 2, 3, 1},
			Expected: true,
		},
		{
			Input:    []int{1, 2, 3, 4},
			Expected: false,
		},
		{
			Input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			Expected: true,
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test#%d", idx+1), func(t *testing.T) {
			output := containsDuplicate(testCase.Input)
			err := util.Test(testCase.Expected, output)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

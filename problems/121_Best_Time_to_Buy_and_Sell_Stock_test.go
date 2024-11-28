package main

import (
	"fmt"
	"leet-it-go/util"
	"testing"
)

func maxProfit(prices []int) int {
	maxProfit := 0
	entry := prices[0]

	for _, price := range prices {
		if price-entry > maxProfit {
			maxProfit = price - entry
		}
		if price < entry {
			entry = price
		}
	}
	return maxProfit
}

func TestMaxProfit(t *testing.T) {
	testCases := []util.TestCase[[]int, int]{
		{
			Input:    []int{10, 1, 5, 6, 7, 1},
			Expected: 6,
		},
		{
			Input:    []int{10, 8, 7, 5, 2},
			Expected: 0,
		},
		{
			Input:    []int{7, 1, 5, 3, 6, 4},
			Expected: 5,
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test#%d", idx+1), func(t *testing.T) {
			output := maxProfit(testCase.Input)
			err := util.Test(testCase.Expected, output)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

package main

import (
	"fmt"
	"leet-it-go/util"
	"regexp"
	"strings"
	"testing"
)

func isPalindrome(s string) bool {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	clean := regex.ReplaceAllString(strings.ToLower(s), "")
	n := len(clean)
	for i := 0; i < n/2; i++ {
		if clean[i] != clean[n-1-i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	testCases := []util.TestCase[string, bool]{
		{
			Input:    "A man, a plan, a canal: Panama",
			Expected: true,
		},
		{
			Input:    "race a car",
			Expected: false,
		},
		{
			Input:    " ",
			Expected: true,
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test#%d", idx+1), func(t *testing.T) {
			output := isPalindrome(testCase.Input)
			err := util.Test(testCase.Expected, output)
			if err != nil {
				t.Error(err)
			}
		})
	}
}

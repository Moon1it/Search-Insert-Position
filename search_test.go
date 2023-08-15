package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {

	testTable := []struct {
		nums     []int
		target   int
		expected int
	}{
		{
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 5, 6},
			target:   6,
			expected: 3,
		},
		{
			nums:     []int{1, 3, 5, 6, 12, 14},
			target:   0,
			expected: 0,
		},
		{
			nums:     []int{1, 3, 5, 6, 10, 12, 14, 43, 54},
			target:   15,
			expected: 7,
		},
		{
			nums:     []int{3, 5, 6, 14, 43, 54, 142, 155},
			target:   80,
			expected: 6,
		},
		{
			nums:     []int{3, 5, 6, 7, 14, 16, 43, 54, 67, 142, 155},
			target:   46,
			expected: 7,
		},
	}

	for _, testCase := range testTable {

		result := SearchInsert(testCase.nums, testCase.target)
		t.Logf("Calling SearchInsert(%d, %d), result: %d\n", testCase.nums, testCase.target, testCase.expected)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expect %d, got %d\n\n", testCase.expected, result))
	}

}

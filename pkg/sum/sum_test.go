package sum_test

import (
	"testing"

	"github.com/tirathawat/softQ/pkg/sum"
)

func TestSum(t *testing.T) {
	testcases := []struct {
		testName string
		input    []int
		want     int
	}{
		{
			testName: "should return 3 when 1 and 2",
			input:    []int{1, 2},
			want:     3,
		},
		{
			testName: "should return 1 when 1 and 0",
			input:    []int{1, 0},
			want:     1,
		},
		{
			testName: "should return -3 when -2 and -1",
			input:    []int{-2, -1},
			want:     -3,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			got := sum.Sum(tc.input...)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

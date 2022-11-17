package ticket_test

import (
	"testing"

	"github.com/tirathawat/softQ/pkg/ticket"
)

func TestTicketPrice(t *testing.T) {
	testcases := []struct {
		testName string
		age      uint
		want     float64
	}{
		{
			testName: "should return 0 when age is 0",
			age:      0,
			want:     0,
		},
		{
			testName: "should free ticket when age under 3",
			age:      1,
			want:     0,
		},
		{
			testName: "should free ticket when age is 3",
			age:      3,
			want:     0,
		},
		{
			testName: "should tiket $15 when age is over 3",
			age:      4,
			want:     15,
		},
		{
			testName: "should tiket $15 when age is 15",
			age:      15,
			want:     15,
		},
		{
			testName: "should tiket $30 when age over 15",
			age:      16,
			want:     30,
		},
		{
			testName: "should tiket $30 when age is 50",
			age:      50,
			want:     30,
		},
		{
			testName: "should tiket $30 when age is 51",
			age:      51,
			want:     5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.testName, func(t *testing.T) {
			got := ticket.GetPrice(tc.age)

			if got != tc.want {
				t.Errorf("want %v, got %v", tc.want, got)
			}
		})
	}
}

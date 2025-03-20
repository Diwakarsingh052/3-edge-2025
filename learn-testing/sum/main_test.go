package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// add _test at the end of the fileName to make it a test file
// go test ./... // run all the tests for the project

// function names must start with word Test to signal it is a test
func TestSumInt(t *testing.T) {
	// Figure out two things
	// What are inputs (parameters)
	// Expected output
	input := []int{1, 2, 3, 4, 5}
	want := 15 // Expected output
	got := SumInt(input)
	if got != want {
		// test would continue on if test case fail
		t.Errorf("sum of 1 to 5 should be %v; got %v", want, got)

		// Uncomment next line to stop the test if it fails at this point.
		//t.Fatalf("sum of 1 to 5 should be %v; got %v", want, got)
	}

	want = 0
	got = SumInt(nil)

	if got != want {
		t.Errorf("sum of nil should be %v; got %v", want, got)
	}

}

// Table test
func TestTableTestSumInt(t *testing.T) {

	tt := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "one to five numbers",
			args: []int{1, 2, 3, 4, 5},
			want: 15,
		},
		{
			name: "empty slice",
			args: nil,
			want: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := SumInt(tc.args)
			require.Equal(t, tc.want, got)
		})

	}

}

package common

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMap(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int) int
		output []int
	}{
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i int) int { return i + 1 },
			output: []int{2, 3, 4},
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i int) int { return i + 1 },
			output: []int{},
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			output := Map(test.input, test.fn)
			if !reflect.DeepEqual(output, test.output) {
				t.Errorf("Map(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int) bool
		output []int
	}{
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i int) bool { return i > 1 },
			output: []int{2, 3},
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i int) bool { return i > 1 },
			output: []int{},
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			output := Filter(test.input, test.fn)
			if !reflect.DeepEqual(output, test.output) {
				t.Errorf("Filter(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int, int) int
		output int
	}{
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i, j int) int { return i + j },
			output: 6,
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i, j int) int { return i + j },
			output: 0,
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			output := Reduce(test.input, test.fn, 0)
			if output != test.output {
				t.Errorf("Reduce(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int) []int
		output []int
	}{
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i int) []int { return []int{i + 1} },
			output: []int{2, 3, 4},
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i int) []int { return []int{i + 1} },
			output: []int{},
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			output := FlatMap(test.input, test.fn)
			if !cmp.Equal(output, test.output) {
				t.Errorf("FlatMap(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

func TestSome(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int) bool
		output bool
	}{
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i int) bool { return i > 1 },
			output: true,
		},
		{
			name:   "ok-false",
			input:  []int{1, 2, 3},
			fn:     func(i int) bool { return i > 3 },
			output: false,
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i int) bool { return i > 1 },
			output: false,
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			output := Some(test.input, test.fn)
			if output != test.output {
				t.Errorf("Some(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

func TestAll(t *testing.T) {
	table := []struct {
		name   string
		input  []int
		fn     func(int) bool
		output bool
	}{
		{
			name:   "ok-false",
			input:  []int{1, 2, 3},
			fn:     func(i int) bool { return i > 1 },
			output: false,
		},
		{
			name:   "ok",
			input:  []int{1, 2, 3},
			fn:     func(i int) bool { return i >= 1 },
			output: true,
		},
		{
			name:   "empty",
			input:  []int{},
			fn:     func(i int) bool { return i > 1 },
			output: false,
		},
	}
	for _, test := range table {
		test := test
		t.Run(test.name, func(t *testing.T) {
			output := All(test.input, test.fn)
			if output != test.output {
				t.Errorf("All(%v, %v) = %v, want %v", test.input, reflect.ValueOf(test.fn).Type(), output, test.output)
			}
		})
	}
}

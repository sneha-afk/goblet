package set

import (
	"testing"
)

func Test_NewSet(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "takes in single value",
			input: []int{1},
		},
		{
			name:  "takes mulitple values",
			input: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewSet(tt.input...)

			for _, e := range tt.input {
				if !result.Contains(e) {
					t.Errorf("could not find element within constructed set: %d", e)
				}
			}
		})
	}
}

func Test_Union(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected []int
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: []int{},
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: []int{1},
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.Union(twoset)
			if result.Count() != len(tt.expected) {
				t.Errorf("different number of elements: got %d, expected %d", result.Count(), len(tt.expected))
			}
			for _, e := range tt.expected {
				if !result.Contains(e) {
					t.Errorf("did not find element: wanted %d in %v", e, result)
				}
			}
		})
	}
}

func Test_Intersection(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected []int
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: []int{},
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: []int{3},
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: []int{},
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: []int{},
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.Intersection(twoset)
			if result.Count() != len(tt.expected) {
				t.Errorf("different number of elements: got %d, expected %d", result.Count(), len(tt.expected))
			}
			for _, e := range tt.expected {
				if !result.Contains(e) {
					t.Errorf("did not find element: wanted %d in %v", e, result)
				}
			}
		})
	}
}

func Test_Difference(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected []int
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: []int{1, 2, 3},
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: []int{1, 2},
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: []int{},
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: []int{1},
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.Difference(twoset)
			if result.Count() != len(tt.expected) {
				t.Errorf("different number of elements: got %d, expected %d", result.Count(), len(tt.expected))
			}
			for _, e := range tt.expected {
				if !result.Contains(e) {
					t.Errorf("did not find element: wanted %d in %v", e, result)
				}
			}
		})
	}
}

func Test_SymmetricDifference(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected []int
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: []int{},
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: []int{1},
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.SymmetricDifference(twoset)
			if result.Count() != len(tt.expected) {
				t.Errorf("different number of elements: got %d, expected %d", result.Count(), len(tt.expected))
			}
			for _, e := range tt.expected {
				if !result.Contains(e) {
					t.Errorf("did not find element: wanted %d in %v", e, result)
				}
			}
		})
	}
}

func Test_IsSubset(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected bool
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: false,
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: true,
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: false,
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: true,
		},
		{
			name:     "left is big",
			one:      []int{1, 2, 3, 4, 5, 6},
			two:      []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "right is big",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3, 4, 5, 6},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.IsSubset(twoset)
			if result != tt.expected {
				t.Errorf("onesubset is %s and twoset is %s: got %t, expected %t", oneset, twoset, result, tt.expected)
			}
		})
	}
}

func Test_IsSuperset(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected bool
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: false,
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: true,
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: true,
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: false,
		},
		{
			name:     "left is big",
			one:      []int{1, 2, 3, 4, 5, 6},
			two:      []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "right is big",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3, 4, 5, 6},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.IsSuperset(twoset)
			if result != tt.expected {
				t.Errorf("onesubset is %s and twoset is %s: got %t, expected %t", oneset, twoset, result, tt.expected)
			}
		})
	}
}

func Test_IsEqual(t *testing.T) {
	tests := []struct {
		name     string
		one      []int
		two      []int
		expected bool
	}{
		{
			name:     "equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "non-equivalent sets",
			one:      []int{1, 2, 3},
			two:      []int{4, 5, 6},
			expected: false,
		},
		{
			name:     "overlap",
			one:      []int{1, 2, 3},
			two:      []int{3, 4, 5},
			expected: false,
		},
		{
			name:     "both empty",
			one:      []int{},
			two:      []int{},
			expected: true,
		},
		{
			name:     "left empty",
			one:      []int{1},
			two:      []int{},
			expected: false,
		},
		{
			name:     "right empty",
			one:      []int{},
			two:      []int{1},
			expected: false,
		},
		{
			name:     "left is big",
			one:      []int{1, 2, 3, 4, 5, 6},
			two:      []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "right is big",
			one:      []int{1, 2, 3},
			two:      []int{1, 2, 3, 4, 5, 6},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oneset := NewSet(tt.one...)
			twoset := NewSet(tt.two...)

			result := oneset.IsEqual(twoset)
			if result != tt.expected {
				t.Errorf("onesubset is %s and twoset is %s: got %t, expected %t", oneset, twoset, result, tt.expected)
			}
		})
	}
}

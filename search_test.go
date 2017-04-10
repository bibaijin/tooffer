package tooffer

import (
	"testing"
)

func TestSearch(t *testing.T) {
	cases := []struct {
		inArray  [][]int
		inNumber int
		want     bool
	}{
		{
			[][]int{
				[]int{1, 2, 8, 9},
				[]int{2, 4, 9, 12},
				[]int{4, 7, 10, 13},
				[]int{6, 8, 11, 15},
			},
			7,
			true,
		},
		{
			[][]int{
				[]int{1, 2, 8, 9},
				[]int{2, 4, 9, 12},
				[]int{4, 7, 10, 13},
				[]int{6, 8, 11, 15},
			},
			5,
			false,
		},
	}

	for _, c := range cases {
		got := Search(c.inArray, c.inNumber)
		if got != c.want {
			t.Errorf("Search failed, want: %s, got: %s.", c.want, got)
		}
	}
}

func TestSpaceEscape(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{
			"We are happy.",
			"We%20are%20happy.",
		},
	}

	for _, c := range cases {
		got := SpaceEscape(c.in)
		if got != c.want {
			t.Errorf("SpaceEscape failed, want: %s, got: %s.", c.want, got)
		}
	}
}

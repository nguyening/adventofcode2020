package day3

import "testing"

func equal(a, b *Coord) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil {
		return false
	} else if b == nil {
		return false
	}

	return a.Row == b.Row && a.Column == b.Column
}

func TestNextPosition(t *testing.T) {
	cases := []struct {
		dim  Coord
		in   Coord
		sled Slope
		want *Coord
	}{
		{Coord{11, 11}, Coord{0, 0}, Slope{3, 1}, &Coord{1, 3}},
		{Coord{11, 11}, Coord{1, 3}, Slope{3, 1}, &Coord{2, 6}},
		{Coord{11, 11}, Coord{2, 6}, Slope{3, 1}, &Coord{3, 9}},
		{Coord{11, 11}, Coord{3, 9}, Slope{3, 1}, &Coord{4, 1}},
		{Coord{11, 11}, Coord{4, 1}, Slope{3, 1}, &Coord{5, 4}},
		{Coord{11, 11}, Coord{5, 4}, Slope{3, 1}, &Coord{6, 7}},
		{Coord{11, 11}, Coord{6, 7}, Slope{3, 1}, &Coord{7, 10}},
		{Coord{11, 11}, Coord{7, 10}, Slope{3, 1}, &Coord{8, 2}},
		{Coord{11, 11}, Coord{8, 2}, Slope{3, 1}, &Coord{9, 5}},
		{Coord{11, 11}, Coord{9, 5}, Slope{3, 1}, &Coord{10, 8}},
		{Coord{11, 11}, Coord{10, 8}, Slope{3, 1}, nil},
	}
	for i, tc := range cases {
		got := nextPosition(tc.dim, tc.in, tc.sled)
		if equal(got, tc.want) == false {
			t.Errorf("#%d got %v want %v", i, got, tc.want)
		}
	}
}

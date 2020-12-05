package main

import "testing"

func TestSeat(t *testing.T) {
	cases := []struct{
		in Seat
		want int
	}{
		{ Seat{70, 7}, 567 },
		{ Seat{14, 7}, 119 },
		{ Seat{102, 4}, 820 },
	}
	for i, tc := range cases {
		if got := tc.in.ID(); got != tc.want {
			t.Errorf("#%d got %d want %d", i, got, tc.want)
		}
	}
}

func TestStringToSpec(t *testing.T) {
	equal := func (a, b []Op) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	cases := []struct{
		in string
		want []Op
	}{
		{ "BFFFBBFRRR", []Op{OpUpper,OpLower,OpLower,OpLower,OpUpper,OpUpper,OpLower,OpUpper,OpUpper,OpUpper} },
		{ "FFFBBBFRRR", []Op{OpLower,OpLower,OpLower,OpUpper,OpUpper,OpUpper,OpLower,OpUpper,OpUpper,OpUpper} },
		{ "BBFFBBFRLL", []Op{OpUpper,OpUpper,OpLower,OpLower,OpUpper,OpUpper,OpLower,OpUpper,OpLower,OpLower} },
	}
	for i, tc := range cases {
		if got := StringToSpec(tc.in); equal(got, tc.want) == false {
			t.Errorf("%d got %v want %v", i, got, tc.want)
		}
	}
}

func TestFindSeat(t *testing.T) {
	cases := []struct{
		in string
		want Seat
	}{
		{ "FBFBBFFRLR", Seat{44, 5} },
		{ "BFFFBBFRRR", Seat{70, 7} },
		{ "FFFBBBFRRR", Seat{14, 7} },
		{ "BBFFBBFRLL", Seat{102, 4} },
	}
	for i, tc := range cases {
		if got := FindSeat(StringToSpec(tc.in)); got != tc.want {
			t.Errorf("%d got %v want %v", i, got, tc.want)
		}
	}
}

package day1

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct{
		in []int
		want int
	}{
		{ in: []int{ 2020, 0 }, want: 0 },
		{ in: []int{ 2019, 1 }, want: 2019 },
		{ in: []int{ 2020, 1, 2, 3, 500, 0 }, want: 0 },
		{
			in: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			want: 514579,
		},
	}
	for _, tc := range cases {
		got := Solve(tc.in)
		if got != tc.want {
			t.Errorf("got %d want %d", got, tc.want)
		}
	}
}

func TestSolvePartTwo(t *testing.T) {
	cases := []struct{
		in []int
		want int
	}{
		{ in: []int{ 2020, 0, 0 }, want: 0 },
		{ in: []int{ 2018, 1, 1 }, want: 2018 },
		{ in: []int{ 2019, 1, 2, 3, 500, 0 }, want: 0 },
		{
			in: []int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			want: 241861950,
		},
	}
	for _, tc := range cases {
		got := SolvePartTwo(tc.in)
		if got != tc.want {
			t.Errorf("got %d want %d", got, tc.want)
		}
	}
}

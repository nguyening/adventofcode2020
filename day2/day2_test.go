package day2

import "testing"

func TestValid(t *testing.T) {
	cases := []struct{
		in string
		policy Policy
		want bool
	}{
		{ "", Policy{1, 10, 'a'}, false },
		{ "", Policy{0, 10, 'a'}, true },
		{ "abcde", Policy{1, 3, 'a'}, true },
		{ "cdefg", Policy{1, 3, 'b'}, false },
		{ "ccccccccc", Policy{2, 9, 'c'}, true },
	}
	for i, tc := range cases {
		if got := Valid(tc.policy, tc.in); got != tc.want {
			t.Errorf("#%d got %t want %t", i, got, tc.want)
		}
	}
}

func TestValidPartTwo(t *testing.T) {
	cases := []struct{
		in string
		policy PolicyPartTwo
		want bool
	}{
		{ "", PolicyPartTwo{1, 10, 'a'}, false },
		{ "abcda", PolicyPartTwo{1, 5, 'a'}, false },
		{ "abcde", PolicyPartTwo{1, 5, 'a'}, true },
		{ "abcde", PolicyPartTwo{1, 3, 'a'}, true },
		{ "cdefg", PolicyPartTwo{1, 3, 'b'}, false },
		{ "ccccccccc", PolicyPartTwo{2, 9, 'c'}, false },
	}
	for i, tc := range cases {
		if got := ValidPartTwo(tc.policy, tc.in); got != tc.want {
			t.Errorf("#%d got %t want %t", i, got, tc.want)
		}
	}
}



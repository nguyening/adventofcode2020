package main

import "strings"
import "testing"

func TestCountValid(t *testing.T) {
	cases := []struct{
		in string
		want int
	}{
		{
			in: `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm`,
			want: 1,
		},
		{
			in: `iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929`,
			want: 0,
		},
		{
			in: `hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm`,
			want: 1,
		},
		{
			in: `hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`,
			want: 0,
		},
	}
	for i, tc := range cases {
		got := CountValid(strings.NewReader(tc.in))
		if got != tc.want {
			t.Errorf("#%d got %d want %d", i, got, tc.want)
		}
	}
}

func TestField(t *testing.T) {
	t.Run("Validate", func(t *testing.T) {
		cases := []struct{
			field Field
			in string
			want bool
		}{
			{ FieldBirthYear, "2002", true },
			{ FieldBirthYear, "2003", false },
			{ FieldHeight, "60in", true },
			{ FieldHeight, "190cm", true },
			{ FieldHeight, "190in", false },
			{ FieldHeight, "190", false },
			{ FieldHairColor, "#123abc", true },
			{ FieldHairColor, "#123abz", false },
			{ FieldHairColor, "123abc", false },
			{ FieldEyeColor, "brn", true },
			{ FieldEyeColor, "wat", false },
			{ FieldPassportID, "000000001", true },
			{ FieldPassportID, "0123456789", false },
		}
		for i, tc := range cases {
			if got := tc.field.Validate(tc.in); got != tc.want {
				t.Errorf("#%d got %t want %t", i, got, tc.want)
			}
		}
	})
}

func TestCountValidWithValues(t *testing.T) {
	cases := []struct{
		in string
		want int
	}{
		{
			in: `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926`,
			want: 0,
		},
		{
			in: `iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946`,
			want: 0,
		},
		{
			in: `hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277`,
			want: 0,
		},
		{
			in: `hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`,
			want: 0,
		},
		{
		in: `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f`,
		want: 1,
		},
		{
			in: `eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm`,
			want: 1,
		},
		{
			in: `hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022`,
			want: 1,
		},
		{
			in: `iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`,
			want: 1,
		},
	}
	for i, tc := range cases {
		got := CountValidWithValues(strings.NewReader(tc.in))
		if got != tc.want {
			t.Errorf("#%d got %d want %d", i, got, tc.want)
		}
	}
}

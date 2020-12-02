package main

import "bufio"
import "flag"
import "fmt"
import "os"
import "strings"
import "strconv"

import "github.com/nguyening/adventofcode2020/day2"

var parttwo = flag.Bool("parttwo", false, "interpret values as policy described in part two")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Print("Count number of valid passwords.")
		flag.PrintDefaults()
	}
	flag.Parse()

	var count int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		var ret bool
		if *parttwo {
			ret = two(s.Text())
		} else {
			ret = one(s.Text())
		}

		if ret {
			count++
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", count)
}

func one(s string) bool {
	var err error
	tokens := strings.Split(s, " ")
	if len(tokens) != 3 {
		panic(fmt.Sprintf("unexpected string format: %v", tokens))
	}

	runes, passwd := []rune(tokens[1]), tokens[2]
	if len(runes) != 2 {
		panic(fmt.Sprintf("unexpected rune: %s", tokens[1]))
	}

	var policy day2.Policy
	policy.Rune = runes[0]

	tokens = strings.Split(tokens[0], "-")
	if len(tokens) != 2 {
		panic(fmt.Sprintf("unexpected policy format: %v", tokens))
	}

	policy.Min, err = strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}

	policy.Max, err = strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	if day2.Valid(policy, passwd) {
		return true
	}
	return false
}

func two(s string) bool {
	var err error
	tokens := strings.Split(s, " ")
	if len(tokens) != 3 {
		panic(fmt.Sprintf("unexpected string format: %v", tokens))
	}

	runes, passwd := []rune(tokens[1]), tokens[2]
	if len(runes) != 2 {
		panic(fmt.Sprintf("unexpected rune: %s", tokens[1]))
	}

	var policy day2.PolicyPartTwo
	policy.Rune = runes[0]

	tokens = strings.Split(tokens[0], "-")
	if len(tokens) != 2 {
		panic(fmt.Sprintf("unexpected policy format: %v", tokens))
	}

	policy.PositionA, err = strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}

	policy.PositionB, err = strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	if day2.ValidPartTwo(policy, passwd) {
		return true
	}
	return false
}

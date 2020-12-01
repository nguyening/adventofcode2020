package main

import "bufio"
import "flag"
import "fmt"
import "os"
import "strconv"

import "github.com/nguyening/adventofcode2020/day1"

var triple = flag.Bool("triple", false, "find product of triple that sums to 2020")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Print("Prints product of tuple that sums to 2020.\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	var report []int
	for s.Scan() {
		n, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			panic(err)
		}

		report = append(report, int(n))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	if *triple == true {
		fmt.Printf("%d\n", day1.SolvePartTwo(report))
	} else {
		fmt.Printf("%d\n", day1.Solve(report))
	}
}

package main

import "bufio"
import "flag"
import "fmt"
import "os"

import "github.com/nguyening/adventofcode2020/day3"

var parttwo = flag.Bool("parttwo", false, "interpret the grid as described in part two")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Print("Count the number of trees found along grid slope.")
		flag.PrintDefaults()
	}
	flag.Parse()

	var grid [][]rune

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	if *parttwo {
		product := 1
		for _, sled := range []day3.Slope {
			{ Right: 1, Down: 1 },
			{ Right: 3, Down: 1 },
			{ Right: 5, Down: 1 },
			{ Right: 7, Down: 1 },
			{ Right: 1, Down: 2 },
		} {
			product *= day3.CountTrees(grid, sled)
		}

		fmt.Printf("%d\n", product)
	} else {
		count := day3.CountTrees(grid, day3.Slope{ Right: 3, Down: 1 })
		fmt.Printf("%d\n", count)
	}
}

package main

import "flag"
import "fmt"
import "bufio"
import "os"

type Seat struct {
	Row int
	Column int
}
func (s Seat) ID() int { return s.Row * 8 + s.Column }

type Op int
const (
	OpUpper Op = iota
	OpLower
)

func FindIndex(lower, upper int, op Op) (int, int) {
	switch op {
	case OpUpper:
		return lower + (upper-lower)/2 + 1, upper
	case OpLower:
		return lower, lower + (upper-lower)/2
	}

	panic("unhandled op")
}

func FindSeat(spec []Op) Seat {
	ri, rj := 0, 127
	ci, cj := 0, 7
	for i, op := range spec {
		if i < 7 {
			ri, rj = FindIndex(ri, rj, op)
		} else {
			ci, cj = FindIndex(ci, cj, op)
		}
	}
	if ri != rj {
		panic(fmt.Errorf("didn't get a seat: %d, %d", ri, rj))
	}
	if ci != cj {
		panic(fmt.Errorf("didn't get a seat: %d, %d", ci, cj))
	}
	return Seat{ Row: ri, Column: ci }
}

func StringToSpec(s string) (ops []Op) {
	for _, rune_ := range s {
		switch rune_ {
			case 'F':
				ops = append(ops, OpLower)
			case 'B':
				ops = append(ops, OpUpper)
			case 'L':
				ops = append(ops, OpLower)
			case 'R':
				ops = append(ops, OpUpper)
			default:
				panic("unhandled rune")
		}
	}
	return
}

var parttwo = flag.Bool("parttwo", false, "interpret the file as described in part two")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Print("Get the highest seat ID.")
		flag.PrintDefaults()
	}
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	var tickets []int
	for s.Scan() {
		seat := FindSeat(StringToSpec(s.Text()))
		tickets = append(tickets, seat.ID())
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	if *parttwo {
		var occ [8*128]bool
		for _, sid := range tickets {
			occ[sid] = true
		}

		// front seats [0, 7]
		// back seats (8 * 127) + [0, 7]
		for sid := 8; sid < (8 * 127); sid++ {
			if occ[sid] == false && occ[sid-1] == true && occ[sid+1] == true {
				fmt.Println(sid)
				break
			}
		}
	} else {
		var max int
		for _, sid := range tickets {
			if sid > max {
				max = sid
			}
		}
		fmt.Println(max)
	}
}

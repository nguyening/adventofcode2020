// --- Day 11: Seating System ---
// 
// Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!
// 
// By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).
// 
// The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:
// 
// L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
// 
// Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:
// 
//     If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
//     If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
//     Otherwise, the seat's state does not change.
// 
// Floor (.) never changes; seats don't move, and nobody sits on the floor.
// 
// After one round of these rules, every seat in the example layout becomes occupied:
// 
// #.##.##.##
// #######.##
// #.#.#..#..
// ####.##.##
// #.##.##.##
// #.#####.##
// ..#.#.....
// ##########
// #.######.#
// #.#####.##
// 
// After a second round, the seats with four or more occupied adjacent seats become empty again:
// 
// #.LL.L#.##
// #LLLLLL.L#
// L.L.L..L..
// #LLL.LL.L#
// #.LL.LL.LL
// #.LLLL#.##
// ..L.L.....
// #LLLLLLLL#
// #.LLLLLL.L
// #.#LLLL.##
// 
// This process continues for three more rounds:
// 
// #.##.L#.##
// #L###LL.L#
// L.#.#..#..
// #L##.##.L#
// #.##.LL.LL
// #.###L#.##
// ..#.#.....
// #L######L#
// #.LL###L.L
// #.#L###.##
// 
// #.#L.L#.##
// #LLL#LL.L#
// L.L.L..#..
// #LLL.##.L#
// #.LL.LL.LL
// #.LL#L#.##
// ..L.L.....
// #L#LLLL#L#
// #.LLLLLL.L
// #.#L#L#.##
// 
// #.#L.L#.##
// #LLL#LL.L#
// L.#.L..#..
// #L##.##.L#
// #.#L.LL.LL
// #.#L#L#.##
// ..L.L.....
// #L#L##L#L#
// #.LLLLLL.L
// #.#L#L#.##
// 
// At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.
// 
// Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?
// 
// --- Part Two ---
// 
// As soon as people start to arrive, you realize your mistake. People don't just care about adjacent seats - they care about the first seat they can see in each of those eight directions!
// 
// Now, instead of considering just the eight immediately adjacent seats, consider the first seat in each of those eight directions. For example, the empty seat below would see eight occupied seats:
// 
// .......#.
// ...#.....
// .#.......
// .........
// ..#L....#
// ....#....
// .........
// #........
// ...#.....
// 
// The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:
// 
// .............
// .L.L.#.#.#.#.
// .............
// 
// The empty seat below would see no occupied seats:
// 
// .##.##.
// #.#.#.#
// ##...##
// ...L...
// ##...##
// #.#.#.#
// .##.##.
// 
// Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an occupied seat to become empty (rather than four or more from the previous rules). The other rules still apply: empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.
// 
// Given the same starting layout as above, these new rules cause the seating area to shift around as follows:
// 
// L.LL.LL.LL
// LLLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLLL
// L.LLLLLL.L
// L.LLLLL.LL
// 
// #.##.##.##
// #######.##
// #.#.#..#..
// ####.##.##
// #.##.##.##
// #.#####.##
// ..#.#.....
// ##########
// #.######.#
// #.#####.##
// 
// #.LL.LL.L#
// #LLLLLL.LL
// L.L.L..L..
// LLLL.LL.LL
// L.LL.LL.LL
// L.LLLLL.LL
// ..L.L.....
// LLLLLLLLL#
// #.LLLLLL.L
// #.LLLLL.L#
// 
// #.L#.##.L#
// #L#####.LL
// L.#.#..#..
// ##L#.##.##
// #.##.#L.##
// #.#####.#L
// ..#.#.....
// LLL####LL#
// #.L#####.L
// #.L####.L#
// 
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##LL.LL.L#
// L.LL.LL.L#
// #.LLLLL.LL
// ..L.L.....
// LLLLLLLLL#
// #.LLLLL#.L
// #.L#LL#.L#
// 
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##L#.#L.L#
// L.L#.#L.L#
// #.L####.LL
// ..#.#.....
// LLL###LLL#
// #.LLLLL#.L
// #.L#LL#.L#
// 
// #.L#.L#.L#
// #LLLLLL.LL
// L.L.L..#..
// ##L#.#L.L#
// L.L#.LL.L#
// #.LLLL#.LL
// ..#.L.....
// LLL###LLL#
// #.LLLLL#.L
// #.L#LL#.L#
// 
// Again, at this point, people stop shifting around and the seating area reaches equilibrium. Once this occurs, you count 26 occupied seats.
// 
// Given the new visibility method and the rule change for occupied seats becoming empty, once equilibrium is reached, how many seats end up occupied?
// 
package main

import "bufio"
import "flag"
import "fmt"
import "os"

const (
	FlrPos rune = '.'
	EmpPos rune = 'L'
	OccPos rune = '#'
)

func VisibleOccupiedPartTwo(grid [][]rune, r, c int) (occ int) {
	R := len(grid)
	C := len(grid[0])

	for _, dir := range [][2]int{
		// cardinal
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},

		// diags
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	} {

		Loop:
		for d := 1; d < R || d < C; d++ {
			dr, dc := dir[0]*d, dir[1]*d
			if r+dr > R-1 || r+dr < 0 || c+dc > C-1 || c+dc < 0 {
				break
			}

			switch grid[r+dr][c+dc] {
			case FlrPos:
				continue
			case EmpPos:
				break Loop
			case OccPos:
				occ++
				break Loop
			}
		}
	}
	return
}

func Neighbors(r, c, R, C int) (nbr [][2]int) {
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			if (dr == 0 && dr == dc) || (r + dr < 0) || (r + dr > R-1) || (c + dc < 0) || (c + dc > C-1) {
				continue
			}

			nbr = append(nbr, [2]int{r+dr, c+dc})
		}
	}
	return nbr
}

type Rule int
const (
	PartOneRule Rule = iota
	PartTwoRule
)

func Simulate(grid [][]rune, rule Rule) (newgrid [][]rune, changed bool) {
	R := len(grid)
	C := len(grid[0])

	newgrid = make([][]rune, R)
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			pos := grid[r][c]
			newgrid[r] = append(newgrid[r], pos)

			switch pos {
			case FlrPos:
			case EmpPos:
				var free bool

				switch rule {
				case PartOneRule:
					free = true

					for _, nbr := range Neighbors(r, c, R, C) {
						switch grid[nbr[0]][nbr[1]] {
						case EmpPos:
						case FlrPos:
						case OccPos:
							free = false
						}
					}
				case PartTwoRule:
					free = VisibleOccupiedPartTwo(grid, r, c) == 0
				}

				if free {
					newgrid[r][c] = OccPos
					changed = true
				}
			case OccPos:
				var occ int
				switch rule {
				case PartOneRule:
					for _, nbr := range Neighbors(r, c, R, C) {
						switch grid[nbr[0]][nbr[1]] {
						case EmpPos:
						case FlrPos:
						case OccPos:
							occ++
						}
					}
					if occ > 3 {
						newgrid[r][c] = EmpPos
						changed = true
					}
				case PartTwoRule:
					occ = VisibleOccupiedPartTwo(grid, r, c)
					if occ > 4 {
						newgrid[r][c] = EmpPos
						changed = true
					}
				}
			}
		}
	}

	return
}

var p2 = flag.Bool("parttwo", false, "interpret the file as described in part two")
func main() {
	flag.Parse()

	s := bufio.NewScanner(os.Stdin)
	var grid [][]rune
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	var rule Rule
	if *p2 {
		rule = PartTwoRule
		// fmt.Println(parttwo())
	} else {
		rule = PartOneRule
	}

	for {
		var changed bool
		grid, changed = Simulate(grid, rule)
		if !changed {
			break
		}
	}

	var total int
	R := len(grid)
	C := len(grid[0])
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if grid[r][c] == OccPos {
				total++
			}
		}
	}
	fmt.Println(total)
}

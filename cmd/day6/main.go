// --- Day 6: Custom Customs ---
// 
// As your flight approaches the regional airport where you'll switch to a much larger plane, customs declaration forms are distributed to the passengers.
// 
// The form asks a series of 26 yes-or-no questions marked a through z. All you need to do is identify the questions for which anyone in your group answers "yes". Since your group is just you, this doesn't take very long.
// 
// However, the person sitting next to you seems to be experiencing a language barrier and asks if you can help. For each of the people in their group, you write down the questions for which they answer "yes", one per line. For example:
// 
// abcx
// abcy
// abcz
// 
// In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z. (Duplicate answers to the same question don't count extra; each question counts at most once.)
// 
// Another group asks for your help, then another, and eventually you've collected answers from every group on the plane (your puzzle input). Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:
// 
// abc
// 
// a
// b
// c
// 
// ab
// ac
// 
// a
// a
// a
// a
// 
// b
// 
// This list represents answers from five groups:
// 
//     The first group contains one person who answered "yes" to 3 questions: a, b, and c.
//     The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
//     The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
//     The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
//     The last group contains one person who answered "yes" to only 1 question, b.
// 
// In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.
// 
// For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?
// 
// --- Part Two ---
// 
// As you finish the last group's customs declaration, you notice that you misread one word in the instructions:
// 
// You don't need to identify the questions to which anyone answered "yes"; you need to identify the questions to which everyone answered "yes"!
// 
// Using the same example as above:
// 
// abc
// 
// a
// b
// c
// 
// ab
// ac
// 
// a
// a
// a
// a
// 
// b
// 
// This list represents answers from five groups:
// 
//     In the first group, everyone (all 1 person) answered "yes" to 3 questions: a, b, and c.
//     In the second group, there is no question to which everyone answered "yes".
//     In the third group, everyone answered yes to only 1 question, a. Since some people did not answer "yes" to b or c, they don't count.
//     In the fourth group, everyone answered yes to only 1 question, a.
//     In the fifth group, everyone (all 1 person) answered "yes" to 1 question, b.
// 
// In this example, the sum of these counts is 3 + 0 + 1 + 1 + 1 = 6.
// 
// For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?
// 
package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "os"

var p2 = flag.Bool("parttwo", false, "interpret the file as described in part two")

func main() {
	flag.Parse()

	var count int
	if *p2 {
		count = parttwo(os.Stdin)
	} else {
		count = partone(os.Stdin)
	}
	fmt.Println(count)
}

type Bitmask uint32
func Set(b Bitmask, slot int) Bitmask { return b | (1 << slot) }
func Clear(b Bitmask, slot int) Bitmask { return b &^ (1 << slot) }
func Flip(b Bitmask, slot int) Bitmask { return b ^ (1 << slot) }
func IsSet(b Bitmask, slot int) bool { return (b & (1 << slot)) != 0 }
func And(a, b Bitmask) Bitmask { return a & b }
func Count(b Bitmask) (count int) { 
	for i := 0; i < 32; i++ {
		if IsSet(b, i) {
			count++
		}
	}
	return
}
func (b Bitmask) String() string { return fmt.Sprintf("%#b", b) }
const AllOnes = Bitmask(1<<32 - 1)

func parttwo(r io.Reader) (count int) {
	s := bufio.NewScanner(r)
	var open bool
	for {
		group := AllOnes

		for {
			var bits Bitmask
			open = s.Scan()
			if s.Text() == "" {
				break
			} else {
				for _, rune_ := range s.Text() {
					i := byte(rune_) - byte('a')
					if i > 26 {
						panic(fmt.Sprintf("unsupported rune: %s", string(rune_)))
					}

					bits = Set(bits, int(i))
				}
				group = And(group, bits)
			}
			if !open {
				break
			}
		}

		count += Count(group)
		if !open {
			break
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return count
}

func partone(r io.Reader) (count int) {
	s := bufio.NewScanner(r)
	var open bool
	for {
		var ok [26]bool

		for {
			open = s.Scan()
			if s.Text() == "" {
				break
			} else {
				for _, rune_ := range s.Text() {
					i := byte(rune_) - byte('a')
					if t := ok[i]; !t {
						ok[i] = true
						count++
					}
				}
			}
			if !open {
				break
			}
		}

		if !open {
			break
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return count
}

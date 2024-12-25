package day21

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) [][]byte {
	scanner := bufio.NewScanner(in)
	lines := [][]byte{}
	for scanner.Scan() {
		b := scanner.Bytes()
		lines = append(lines, make([]byte, len(b)))
		copy(lines[len(lines)-1], b)
	}
	return lines
}

type Pad struct {
	buttons map[byte]int
	banned  int
}

var numpad = Pad{map[byte]int{'7': 0, '8': 1, '9': 2, '4': 3, '5': 4, '6': 5, '1': 6, '2': 7, '3': 8, '0': 10, 'A': 11}, 9}
var dirpad = Pad{map[byte]int{'^': 1, 'A': 2, '<': 3, 'v': 4, '>': 5}, 0}

func hash(start int, end int, lastPress int) int {
	return (start * 12 * 12) + (end * 12) + lastPress
}

const BAD = 999999999999999

// Count how many presses it would take to use a dirpad to move from the start index to the end index, given the last button pushed on this pad
func numberOfMoves(memos *[][]int, pad Pad, start int, end int, lastPress int, depth int) int {
	if depth == 0 {
		return 1
	}
	if start == pad.banned {
		return BAD
	}
	h := hash(start, end, lastPress)
	if (*memos)[depth][h] != 0 {
		return (*memos)[depth][h]
	}
	if start == end {
		// We are already positioned over the correct button, just press A
		return numberOfMoves(memos, dirpad, lastPress, dirpad.buttons['A'], dirpad.buttons['A'], depth-1)
	}
	sx := start % 3
	sy := start / 3
	ex := end % 3
	ey := end / 3
	dx := ex - sx
	dy := ey - sy

	xMoves := BAD
	if dx > 0 {
		// >
		xMoves = numberOfMoves(memos, dirpad, lastPress, dirpad.buttons['>'], dirpad.buttons['A'], depth-1) +
			numberOfMoves(memos, pad, start+1, end, dirpad.buttons['>'], depth)
	} else if dx < 0 {
		// <
		xMoves = numberOfMoves(memos, dirpad, lastPress, dirpad.buttons['<'], dirpad.buttons['A'], depth-1) +
			numberOfMoves(memos, pad, start-1, end, dirpad.buttons['<'], depth)
	}
	yMoves := BAD
	if dy > 0 {
		// v
		yMoves = numberOfMoves(memos, dirpad, lastPress, dirpad.buttons['v'], dirpad.buttons['A'], depth-1) +
			numberOfMoves(memos, pad, start+3, end, dirpad.buttons['v'], depth)
	} else if dy < 0 {
		// ^
		yMoves = numberOfMoves(memos, dirpad, lastPress, dirpad.buttons['^'], dirpad.buttons['A'], depth-1) +
			numberOfMoves(memos, pad, start-3, end, dirpad.buttons['^'], depth)
	}
	moves := xMoves
	if xMoves == BAD || yMoves < xMoves {
		moves = yMoves
	}
	(*memos)[depth][h] = moves
	return moves
}

func Level1(in io.Reader) string {
	lines := parse(in)
	depth := 3
	memos := make([][]int, depth+1)
	for i := range memos {
		memos[i] = make([]int, 12*12*12)
	}
	total := 0
	for _, line := range lines {
		numMoves := numberOfMoves(&memos, numpad, numpad.buttons['A'], numpad.buttons[line[0]], dirpad.buttons['A'], depth)
		for i := 0; i < len(line)-1; i += 1 {
			newMoves := numberOfMoves(&memos, numpad, numpad.buttons[line[i]], numpad.buttons[line[i+1]], dirpad.buttons['A'], depth)
			numMoves += newMoves
		}
		total += numMoves * util.ToInt(string(line[:3]))
	}
	return fmt.Sprint(total)
	// <v<A>>^AvA^A <vA<AA>>^AAvA<^A>AAvA^     A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A
	//    <   A > A   v <<   AA >  ^ AA >      A v   AA ^ A   < v  AAA >  ^ A
	//        ^   A          <<      ^^        A     >>   A        VVV      A
	//            3                            7          9                 A
	// v<<A>>^AvA^A v<<A>>^AAv<A<A>>^AAvAA^<A> Av<A>^AA<A>Av<A<A>>^AAAvA^<A>A
	//    <   A > A    <   AA  v <   AA >>  ^  A  v  AA ^ A  v <   AAA >  ^ A
	//        ^   A        ^^        <<        A     >>   A        vvv      A
	//            3                            7          9                 A
}

func Level2(in io.Reader) string {
	lines := parse(in)
	depth := 26
	memos := make([][]int, depth+1)
	for i := range memos {
		memos[i] = make([]int, 12*12*12)
	}
	total := 0
	for _, line := range lines {
		numMoves := numberOfMoves(&memos, numpad, numpad.buttons['A'], numpad.buttons[line[0]], dirpad.buttons['A'], depth)
		for i := 0; i < len(line)-1; i += 1 {
			newMoves := numberOfMoves(&memos, numpad, numpad.buttons[line[i]], numpad.buttons[line[i+1]], dirpad.buttons['A'], depth)
			numMoves += newMoves
		}
		total += numMoves * util.ToInt(string(line[:3]))
	}
	return fmt.Sprint(total)
}

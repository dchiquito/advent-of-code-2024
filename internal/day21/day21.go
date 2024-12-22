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

type Pad = map[byte]int

var numpad = Pad{'7': 0, '8': 1, '9': 2, '4': 3, '5': 4, '6': 5, '1': 6, '2': 7, '3': 8, '0': 10, 'A': 11}
var dirpad = Pad{'^': 1, 'A': 2, '<': 3, 'v': 4, '>': 5}

func moves(pad Pad, start byte, end byte) []byte {
	moves := []byte{}
	si := pad[start]
	sx := si % 3
	sy := si / 3

	ei := pad[end]
	ex := ei % 3
	ey := ei / 3

	dx := ex - sx
	dy := ey - sy

	// The order matters here
	// We are dodging a hole in either the top left or the top right:
	// X O O
	// O O O
	// X O O
	// So we first move right, then up/down, then left
	if dx > 0 {
		for i := 0; i < dx; i += 1 {
			moves = append(moves, '>')
		}
	}
	if dy < 0 {
		for i := dy; i < 0; i += 1 {
			moves = append(moves, '^')
		}
	}
	if dy > 0 {
		for i := 0; i < dy; i += 1 {
			moves = append(moves, 'v')
		}
	}
	if dx < 0 {
		for i := dx; i < 0; i += 1 {
			moves = append(moves, '<')
		}
	}
	moves = append(moves, 'A')
	return moves
}

func allMoves(pad Pad, arr []byte) []byte {
	var prev byte = 'A'
	ms := []byte{}
	for _, next := range arr {
		ms = append(ms, moves(pad, prev, next)...)
		prev = next
	}
	return ms
}
func moveNumpad(memo *[]int, si int, ei int, depth int) int {
	// si := numpad[start]
	// ei := numpad[end]
	if (*memo)[si*12+ei] != 0 {
		return (*memo)[si*12+ei]
	}

	sx := si % 3
	sy := si / 3

	ex := ei % 3
	ey := ei / 3

	dx := ex - sx
	dy := ey - sy

	if dx == 0 {
		if dy > 0 {
			numMoves := moveNumpad(memo, si-3, ei, depth) + moveDirpad('^', depth-1)
			(*memo)[si*12+ei] = numMoves
			return numMoves
		}
	}
	return 0
}

func moveDirpad(c byte, depth int) int {
	if depth == 1 {

	}
	return 0
}

func Level1(in io.Reader) string {
	lines := parse(in)
	total := 0
	for _, line := range lines {
		fmt.Println(string(line))
		m1 := allMoves(numpad, line)
		fmt.Println(string(m1))
		m2 := allMoves(dirpad, m1)
		fmt.Println(string(m2))
		m3 := allMoves(dirpad, m2)
		fmt.Println(string(m3))
		fmt.Println(len(m3), "*", util.ToInt(string(line[:len(line)-1])))
		total += len(m3) * util.ToInt(string(line[:len(line)-1]))
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
	// The plan:
	// moveNumpad: calculates the best path for a move on the numpad
	//   attempts moving closer in both X and Y, picks the shorter length
	//   memoized with start, end
	// moveDirpad: calculates the best path for a move on the dirpad, given a depth
	//   attempts moving closer in both X and Y, picks the shorter length
	//   memoized with start, end, depth
}

func Level2(in io.Reader) string {
	return ""
}

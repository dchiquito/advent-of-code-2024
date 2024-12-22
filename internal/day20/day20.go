package day20

import (
	"bufio"
	"fmt"
	"io"
)

// const Size int = 15
const Size int = 141

func parse(in io.Reader) ([]bool, int, int, int, int) {
	scanner := bufio.NewScanner(in)
	walls := make([]bool, Size*Size)
	y := 0
	sx := 0
	sy := 0
	ex := 0
	ey := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		for x, c := range line {
			if c == '#' {
				walls[y*Size+x] = true
			} else if c == 'S' {
				sx = x
				sy = y
			} else if c == 'E' {
				ex = x
				ey = y
			}
		}
		y += 1
	}
	return walls, sx, sy, ex, ey
}

func walk(walls []bool, sx int, sy int) []int {
	weights := make([]int, Size*Size)
	q := []int{sy*Size + sx}
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		x := i % Size
		y := i / Size
		w := weights[i]
		if x > 0 && !walls[i-1] && weights[i-1] == 0 {
			weights[i-1] = w + 1
			q = append(q, i-1)
		}
		if x < Size-1 && !walls[i+1] && weights[i+1] == 0 {
			weights[i+1] = w + 1
			q = append(q, i+1)
		}
		if y > 0 && !walls[i-Size] && weights[i-Size] == 0 {
			weights[i-Size] = w + 1
			q = append(q, i-Size)
		}
		if y < Size-1 && !walls[i+Size] && weights[i+Size] == 0 {
			weights[i+Size] = w + 1
			q = append(q, i+Size)
		}
	}
	weights[sy*Size+sx] = 0
	return weights
}

func Level1(in io.Reader) string {
	walls, sx, sy, ex, ey := parse(in)
	startToEnd := walk(walls, sx, sy)
	endToStart := walk(walls, ex, ey)
	dist := startToEnd[ey*Size+ex]
	// savings := 1 + 2
	savings := 100 + 2
	canSave := func(a int, b int) bool {
		return startToEnd[b]+endToStart[a] <= dist-savings || startToEnd[a]+endToStart[b] <= dist-savings
	}
	shortcuts := 0
	for y := 1; y < Size-1; y += 1 {
		for x := 1; x < Size-1; x += 1 {
			i := y*Size + x
			if walls[i] {
				l := i - 1
				r := i + 1
				u := i - Size
				d := i + Size
				wl := !walls[l]
				wr := !walls[r]
				wu := !walls[u]
				wd := !walls[d]
				if wl && wr && canSave(l, r) {
					shortcuts += 1
				}
				if wl && wu && canSave(l, u) {
					shortcuts += 1
				}
				if wl && wd && canSave(l, d) {
					shortcuts += 1
				}
				if wr && wu && canSave(r, u) {
					shortcuts += 1
				}
				if wr && wd && canSave(r, d) {
					shortcuts += 1
				}
				if wu && wd && canSave(u, d) {
					shortcuts += 1
				}
			}
		}
	}
	return fmt.Sprint(shortcuts)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Level2(in io.Reader) string {
	walls, sx, sy, ex, ey := parse(in)
	startToEnd := walk(walls, sx, sy)
	endToStart := walk(walls, ex, ey)
	dist := startToEnd[ey*Size+ex]
	blink := 20
	// savings := 50
	savings := 100
	canSave := func(a int, b int, d int) bool {
		return startToEnd[a]+endToStart[b]+d <= dist-savings
	}
	shortcuts := 0
	for ay := 1; ay < Size-1; ay += 1 {
		for ax := 1; ax < Size-1; ax += 1 {
			ai := ay*Size + ax
			if !walls[ai] {
				for dy := -blink; dy <= blink; dy += 1 {
					by := ay + dy
					if by < 1 || by >= Size-1 {
						continue
					}
					for dx := abs(dy) - blink; dx <= blink-abs(dy); dx += 1 {
						bx := ax + dx
						if bx < 1 || bx >= Size-1 {
							continue
						}
						bi := by*Size + bx
						if !walls[bi] && canSave(ai, bi, abs(dy)+abs(dx)) {
							shortcuts += 1
						}
					}
				}
			}
		}
	}
	return fmt.Sprint(shortcuts)
}

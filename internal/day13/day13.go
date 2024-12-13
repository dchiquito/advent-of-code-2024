package day13

import (
	"bufio"
	"fmt"
	"io"
)

func ToInt(s []byte) int {
	total := 0
	for i := 0; i < len(s); i += 1 {
		total *= 10
		total += int(s[i]) - 48
	}
	return total
}

type Machine struct {
	ax int
	ay int
	bx int
	by int
	px int
	py int
}

func parse(in io.Reader) []Machine {
	scanner := bufio.NewScanner(in)
	machines := make([]Machine, 0, 400)
	for scanner.Scan() {
		lineA := scanner.Bytes()
		ax := int((lineA[12]-48)*10 + lineA[13] - 48)
		ay := int((lineA[18]-48)*10 + lineA[19] - 48)
		scanner.Scan()
		lineB := scanner.Bytes()
		bx := int((lineB[12]-48)*10 + lineB[13] - 48)
		by := int((lineB[18]-48)*10 + lineB[19] - 48)
		scanner.Scan()
		prizeLine := scanner.Bytes()
		var comma int
		for comma = 10; prizeLine[comma] != 44; comma += 1 {
		}
		px := ToInt(prizeLine[9:comma])
		py := ToInt(prizeLine[comma+4:])
		scanner.Scan()
		machines = append(machines, Machine{ax, ay, bx, by, px, py})
	}
	return machines
}

func Level1(in io.Reader) string {
	machines := parse(in)
	total := 0
	for _, m := range machines {
		if m.ax*m.by == m.bx*m.ay {
			fmt.Println("hmmm", m)
		}
		// i*ax + j*bx == px
		// i*ay + j*by == py
		// i == (py - j*by)/ay
		// (ax/ay)*(py-j*by) + j*bx == px
		// py*(ax/ay) - j*by*(ax/ay) + j*bx == px
		// py*(ax/ay) + j(bx - by*(ax/ay)) == px
		// j == (px - py*(ax/ay)) / (bx - by*(ax/ay))
		// very cool, but ax/ay is not necessarily an integer
		// j == (px*ay/ay - py*ax/ay) / (bx*ay/ay - by*ax/ay)
		// j == (px*ay - py*ax) / (bx*ay - by*ax)
		iNum := m.px*m.by - m.py*m.bx
		iDenom := m.ax*m.by - m.ay*m.bx
		jNum := (m.px * m.ay) - (m.py * m.ax)
		jDenom := (m.bx * m.ay) - (m.by * m.ax)
		if iNum%iDenom == 0 && jNum%jDenom == 0 {
			iNum := m.px*m.by - m.py*m.bx
			iDenom := m.ax*m.by - m.ay*m.bx
			aPresses := iNum / iDenom
			bPresses := jNum / jDenom
			total += 3*aPresses + bPresses
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	machines := parse(in)
	total := 0
	for _, m := range machines {
		//      9223372036854775807
		m.px += 10000000000000
		m.py += 10000000000000
		iNum := m.px*m.by - m.py*m.bx
		iDenom := m.ax*m.by - m.ay*m.bx
		jNum := (m.px * m.ay) - (m.py * m.ax)
		jDenom := (m.bx * m.ay) - (m.by * m.ax)
		if iNum%iDenom == 0 && jNum%jDenom == 0 {
			aPresses := iNum / iDenom
			bPresses := jNum / jDenom
			total += 3*aPresses + bPresses
		}
	}
	return fmt.Sprint(total)
}

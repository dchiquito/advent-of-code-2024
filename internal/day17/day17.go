package day17

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

type CPU struct {
	a       int
	b       int
	c       int
	pc      int
	program []int
	out     []int
}

func parse(in io.Reader) CPU {
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	_, a := util.ChompInt(scanner.Bytes(), 11)
	scanner.Scan()
	_, b := util.ChompInt(scanner.Bytes(), 11)
	scanner.Scan()
	_, c := util.ChompInt(scanner.Bytes(), 11)
	scanner.Scan()
	scanner.Scan()

	programLine := scanner.Bytes()[9:]
	i := 0
	var p int
	program := make([]int, 0, 20)
	for i < len(programLine) {
		i, p = util.ChompInt(programLine, i)
		program = append(program, p)
	}
	return CPU{a, b, c, 0, program, make([]int, 0, 100)}
}

func (cpu *CPU) Step() bool {
	if cpu.pc >= len(cpu.program) {
		return false
	}

	instr := cpu.program[cpu.pc]
	literalOp := cpu.program[cpu.pc+1]
	cpu.pc += 2
	var op int = -1
	switch literalOp {
	case 0:
		op = 0
	case 1:
		op = 1
	case 2:
		op = 2
	case 3:
		op = 3
	case 4:
		op = cpu.a
	case 5:
		op = cpu.b
	case 6:
		op = cpu.c
	}
	switch instr {
	case 0: // adv (division)
		num := cpu.a
		denom := 1 << op
		cpu.a = num / denom
	case 1: // bxl (bitwise xor)
		cpu.b = cpu.b ^ literalOp
	case 2: //bst (modulo 8)
		cpu.b = op % 8
	case 3: //jnz (jump not zero)
		if cpu.a != 0 {
			cpu.pc = literalOp
		}
	case 4: // bxc (bitwise xor)
		cpu.b = cpu.b ^ cpu.c
	case 5: //out
		cpu.out = append(cpu.out, op%8)
	case 6: // bdv
		num := cpu.a
		denom := 1 << op
		cpu.b = num / denom
	case 7: // cdv
		num := cpu.a
		denom := 1 << op
		cpu.c = num / denom
	}
	return true
}

func Level1(in io.Reader) string {
	cpu := parse(in)
	for cpu.Step() {
	}
	answer := fmt.Sprint(cpu.out[0])
	for i := 1; i < len(cpu.out); i += 1 {
		answer = answer + "," + fmt.Sprint(cpu.out[i])
	}
	return answer
}

func solve(cpuInit CPU, a int, place int) int {
	if place < 0 {
		return a
	}
	for j := 0; j < 8; j += 1 {
		cpu := cpuInit
		trialA := a + (j << (place * 3))
		cpu.a = trialA
		for cpu.Step() {
		}
		if len(cpu.out) > place && cpu.out[place] == cpu.program[place] {
			solution := solve(cpuInit, trialA, place-1)
			if solution > -1 {
				return solution
			}
		}
	}
	return -1
}

func Level2(in io.Reader) string {
	//2,4 b = a%8
	//1,1 b ^= 1
	//7,5 c = a/1<<b
	//1,5 b = b^0b101
	//4,0 b = b^c
	//0,3 a = a/8
	//5,5 print b%8
	//3,0 loop if a!=0
	// Observations:
	// The program is a simple loop.
	// It computes a value to print based on A, then divides A by 8 and repeats until A is 0.
	// Therefore, the last digit is only dependent on the 3 most significant bits of A, the second to last digit on the next 3 bits, and so on.
	// We can simply work our way backwards through the input to solve for A.

	cpuInit := parse(in)
	a := 0
	place := len(cpuInit.program) - 1
	solution := solve(cpuInit, a, place)
	return fmt.Sprint(solution)
}

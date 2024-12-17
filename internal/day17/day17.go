package day17

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

type CPU struct {
	a       int
	b       int
	c       int
	pc      int
	program []int
	out     []string
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
	return CPU{a, b, c, 0, program, make([]string, 0, 100)}
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
	fmt.Println("exec", instr, literalOp, op)
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
		cpu.out = append(cpu.out, fmt.Sprint(op%8))
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
	return strings.Join(cpu.out, ",")
}

func Level2(in io.Reader) string {
	return ""
}

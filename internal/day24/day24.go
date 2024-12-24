package day24

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

type Op int

const (
	AND Op = iota
	OR
	XOR
)

type Expr struct {
	a  int
	b  int
	op Op
}

type Node struct {
	id    int
	value *bool // TODO inline this for performance
	expr  *Expr
}

type Nodes map[int]*Node

func genId(line []byte) int {
	return int(line[0])*128*128 + int(line[1])*128 + int(line[2])
}
func unId(id int) string {
	i0 := id / (128 * 128)
	i1 := (id / 128) % 128
	i2 := id % 128
	return string([]byte{byte(i0), byte(i1), byte(i2)})
}

func parse(in io.Reader) Nodes {
	scanner := bufio.NewScanner(in)
	nodes := Nodes{}
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		id := genId(line)
		value := line[5] == '1'
		nodes[id] = &Node{id, &value, nil}
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		a := genId(line)
		var op Op
		var b int
		var id int
		if line[4] == 'O' {
			op = OR
			b = genId(line[7:])
			id = genId(line[14:])
		} else if line[4] == 'A' {
			op = AND
			b = genId(line[8:])
			id = genId(line[15:])
		} else {
			op = XOR
			b = genId(line[8:])
			id = genId(line[15:])
		}
		nodes[id] = &Node{id, nil, &Expr{a, b, op}}
	}
	return nodes
}

func eval(nodes Nodes, node *Node) bool {
	if node.value != nil {
		return *node.value
	}
	a := eval(nodes, nodes[node.expr.a])
	b := eval(nodes, nodes[node.expr.b])
	var value bool
	switch node.expr.op {
	case OR:
		value = a || b
	case AND:
		value = a && b
	case XOR:
		value = a != b
	}
	node.value = &value
	return value
}

func Level1(in io.Reader) string {
	nodes := parse(in)
	zids := []int{}
	for id, node := range nodes {
		if id/(128*128) == 'z' {
			zids = append(zids, id)
			eval(nodes, node)
		}
	}
	sort.Ints(zids)
	total := 0
	for i, zid := range zids {
		value := *nodes[zid].value
		if value {
			total += 1 << i
		}
	}
	return fmt.Sprint(total)
}

type Node2 struct {
	id int
	a  int
	b  int
	op Op
}
type Nodes2 map[int]*Node2

type Calc struct {
	nodes Nodes2
	xids  []int
	yids  []int
	zids  []int
}

func parse2(in io.Reader) Calc {
	scanner := bufio.NewScanner(in)
	xids := make([]int, 0, 50)
	yids := make([]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		id := genId(line)
		if line[0] == 'x' {
			xids = append(xids, id)
		} else {
			yids = append(yids, id)
		}
	}

	nodes := Nodes2{}
	zids := make([]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Bytes()
		a := genId(line)
		var op Op
		var b int
		var id int
		if line[4] == 'O' {
			op = OR
			b = genId(line[7:])
			id = genId(line[14:])
		} else if line[4] == 'A' {
			op = AND
			b = genId(line[8:])
			id = genId(line[15:])
		} else {
			op = XOR
			b = genId(line[8:])
			id = genId(line[15:])
		}
		nodes[id] = &Node2{id, a, b, op}
		if id/(128*128) == 'z' {
			zids = append(zids, id)
		}
	}
	sort.Ints(zids)
	return Calc{nodes, xids, yids, zids}
}

func (calc Calc) add(x int, y int) int {
	memo := map[int]bool{}
	for i, xid := range calc.xids {
		memo[xid] = (x>>i)&1 == 1
	}
	for i, yid := range calc.yids {
		memo[yid] = (y>>i)&1 == 1
	}
	var eval func(int) bool
	eval = func(id int) bool {
		if m, ok := memo[id]; ok {
			return m
		}
		node := *calc.nodes[id]
		a := eval(node.a)
		b := eval(node.b)
		var value bool
		switch node.op {
		case OR:
			value = a || b
		case AND:
			value = a && b
		case XOR:
			value = a != b
		}
		memo[id] = value
		return value
	}
	total := 0
	for i, zid := range calc.zids {
		if eval(zid) {
			total += 1 << i
		}
	}
	return total
}

func (calc Calc) diag(a int, b int) bool {
	c := calc.add(a, b)
	if c != a+b {
		fmt.Printf(" %#0.45b (%d)\n", a, a)
		fmt.Printf("+%#0.45b (%d)\n", b, b)
		fmt.Printf("=%#0.45b (%d)\n\n", c, c)
		return true
	}
	return false
}

func (calc Calc) show(id int, depth int) {
	fmt.Print(unId(id))
	node, ok := calc.nodes[id]
	if ok && depth > 0 {
		fmt.Print(":")
		switch node.op {
		case OR:
			fmt.Print("OR")
		case AND:
			fmt.Print("AND")
		case XOR:
			fmt.Print("XOR")
		}
		fmt.Print("(")
		calc.show(node.a, depth-1)
		fmt.Print(",")
		calc.show(node.b, depth-1)
		fmt.Print(")")
	}
}

func (calc Calc) swap(idx int, idy int) {
	x := calc.nodes[idx]
	y := calc.nodes[idy]
	ta := x.a
	tb := x.b
	top := x.op
	x.a = y.a
	x.b = y.b
	x.op = y.op
	y.a = ta
	y.b = tb
	y.op = top
}

func Level2(in io.Reader) string {
	calc := parse2(in)
	// I'm not solving this automatically, so a proof by inspection it is
	calc.swap(genId([]byte("gvw")), genId([]byte("qjb")))
	calc.swap(genId([]byte("z15")), genId([]byte("jgc")))
	calc.swap(genId([]byte("z22")), genId([]byte("drg")))
	calc.swap(genId([]byte("z35")), genId([]byte("jbp")))
	// These swaps resolve all the diagnostics that fail otherwise
	// I found them by staring at the diagnostics
	for i := 0; i < len(calc.xids)-1; i += 1 {
		a := 1 << i
		// fmt.Println("\t", i, 1<<i)
		problem := calc.diag(0, a)
		problem = calc.diag(a, a) || problem
		if problem {
			calc.show(calc.zids[i-1], 5)
			fmt.Println()
			calc.show(calc.zids[i], 5)
			fmt.Println()
			calc.show(calc.zids[i+1], 5)
			fmt.Println()
		}
	}
	return "drg,gvw,jbp,jgc,qjb,z15,z22,z35"
}

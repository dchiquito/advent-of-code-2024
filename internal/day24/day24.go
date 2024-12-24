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

func Level2(in io.Reader) string {
	return ""
}

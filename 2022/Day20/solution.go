package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

type circNode struct {
	val   int
	moved bool
	next  *circNode
	prec  *circNode
}

type List struct {
	head *circNode
}

func newNode(val int) *circNode {
	return &circNode{val, false, nil, nil}
}

func addNewNode(l *List, val int) {
	NewNode := newNode(val)
	if l.head == nil {
		NewNode.prec = NewNode
		NewNode.next = NewNode
		l.head = NewNode
		return
	}
	current := l.head
	for current.next != l.head {
		current = current.next
	}
	current.next, NewNode.prec = NewNode, current
	NewNode.next, l.head.prec = l.head, NewNode
}

func printList(l List) {
	p := l.head
	Print(p.val, " ")
	p = p.next
	for p != l.head {
		Print(p.val, " ")
		p = p.next
	}
	Println()
}

func sposta(p *circNode, shift int) {
	if shift == 0 {
		return
	}
	p.prec.next = p.next
	p.next.prec = p.prec

	if shift > 0 {
		target := p.next
		for i := 1; i < shift; i++ {
			target = target.next
		}
		p.next = target.next
		p.prec = target
		target.next.prec = p
		target.next = p
	} else {
		target := p.prec
		for i := 1; i < -shift; i++ {
			target = target.prec
		}
		p.prec = target.prec
		p.next = target
		target.prec.next = p
		target.prec = p
	}
}

func part2(lista List, slice []int) {
	nodes := make([]*circNode, 0)
	p := lista.head
	for {
		nodes = append(nodes, p)
		p = p.next
		if p == lista.head {
			break
		}
	}
	length := len(nodes)
	for round := 0; round < 10; round++ {
		for _, val := range slice {
			for _, node := range nodes {
				if node.val == val && !node.moved {
					sposta(node, node.val%(length-1))
					node.moved = true
					break
				}
			}
		}
		for _, node := range nodes {
			node.moved = false
		}
	}
	p = lista.head
	for p.val != 0 {
		p = p.next
	}
	sum := 0
	for i := 0; i < 3000; i++ {
		p = p.next
		if i == 999 || i == 1999 || i == 2999 {
			sum += p.val
		}
	}
	Println(sum)
}


func part1(lista List, slice []int) {
	nodes := make([]*circNode, 0)
	p := lista.head
	for {
		nodes = append(nodes, p)
		p = p.next
		if p == lista.head {
			break
		}
	}

	length := len(nodes)
	for _, val := range slice {
		for _, node := range nodes {
			if node.val == val && !node.moved {
				sposta(node, node.val % (length - 1))
				node.moved = true
				break
			}
		}
	}
	p = lista.head
	for p.val != 0 {
		p = p.next
	}

	sum := 0
	for i := 0; i < 3000; i++ {
		p = p.next
		if i == 999 || i == 1999 || i == 2999 {
			sum += p.val
		}
	}
	Println(sum)
}


func input(filename string, part2 bool) (List, []int) {
	file, _ := os.Open(filename + ".txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	lista := List{}
	slice := make([]int, 0)

	for sc.Scan() {
		temp, _ := strconv.Atoi(sc.Text())
		if part2 {
			temp *= 811589153
		}
		slice = append(slice, temp)
		addNewNode(&lista, temp)
	}
	return lista, slice
}

func main() {
	lista, slice := input(os.Args[1], false)
	part1(lista, slice)
	lista, slice = input(os.Args[1], true)
	part2(lista, slice)
}

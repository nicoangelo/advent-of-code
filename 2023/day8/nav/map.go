package nav

import "fmt"

type Map struct {
	directions []int // 0 = left; 1 = right
	nodes      map[string]Node
}

type Node struct {
	next [2]string
}

func (m *Map) FromLines(lines []string) {
	m.nodes = make(map[string]Node, len(lines)-2)
	for i, l := range lines {
		if i == 0 {
			m.directionsFromLine(l)
		}
		if i >= 2 {
			m.nodes[l[0:3]] = Node{next: [2]string{l[7:10], l[12:15]}}
		}
	}
}

func (m *Map) directionsFromLine(l string) {
	m.directions = make([]int, len(l))
	for i, r := range l {
		if r == 'R' {
			m.directions[i] = 1
		}
	}
}

func (m *Map) Navigate(from string, to string) int {
	startNode, ok := m.nodes[from]
	if !ok {
		panic("Start node does not exist")
	}
	endNode, ok := m.nodes[to]
	if !ok {
		panic("End node does not exist")
	}

	nextNode := startNode
	stepCount := 0
	for {
		nextNodeKey := nextNode.next[m.directions[stepCount%len(m.directions)]]
		nextNode = m.nodes[nextNodeKey]
		stepCount++
		if nextNode == endNode {
			break
		}
		if stepCount >= 1000000 {
			fmt.Println("Stopped navigating after 1M attempts to find the end node")
			break
		}
	}
	return stepCount
}

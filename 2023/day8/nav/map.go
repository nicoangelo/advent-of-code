package nav

import "fmt"

type Node struct {
	Name string
	next [2]string
}

type Map struct {
	directions []int // 0 = left; 1 = right
	nodes      map[string]Node
}

func (m *Map) FromLines(lines []string) {
	m.nodes = make(map[string]Node, len(lines)-2)
	for i, l := range lines {
		if i == 0 {
			m.directionsFromLine(l)
		}
		if i >= 2 {
			m.nodes[l[0:3]] = Node{
				Name: l[0:3],
				next: [2]string{l[7:10], l[12:15]}}
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

func (m *Map) GetNodeKeys() []string {
	keys := make([]string, 0, len(m.nodes))
	for k := range m.nodes {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map) NewNavigationContext(from string, endCondition func(*Node) bool) *NavigationContext {
	startNode, ok := m.nodes[from]
	if !ok {
		panic("Start node does not exist")
	}
	return &NavigationContext{
		startNode:    &startNode,
		endCondition: endCondition,
		maap:         m,
	}
}

type NavigationContext struct {
	stepCount    int
	startNode    *Node
	endCondition func(*Node) bool
	maap         *Map
	currentNode  *Node
}

func (ctx *NavigationContext) Navigate() int {
	ctx.StartNavigation()
	for {
		ctx.NavigateStep()
		if ctx.endCondition(ctx.currentNode) {
			break
		}
		if ctx.stepCount >= 1000000 {
			fmt.Println("Stopped navigating after 1M attempts to find the end node")
			break
		}
	}
	return ctx.stepCount
}

func (ctx *NavigationContext) StartNavigation() {
	ctx.currentNode = ctx.startNode
	ctx.stepCount = 0
}

func (ctx *NavigationContext) NavigateStep() *Node {
	nextNodeKey := ctx.currentNode.next[ctx.maap.directions[ctx.stepCount%len(ctx.maap.directions)]]
	nn := ctx.maap.nodes[nextNodeKey]
	ctx.currentNode = &nn
	ctx.stepCount++
	return ctx.currentNode
}

func (ctx *NavigationContext) GetCurrentNode() *Node {
	return ctx.currentNode
}

func (ctx *NavigationContext) GetCurrentStepCount() int {
	return ctx.stepCount
}

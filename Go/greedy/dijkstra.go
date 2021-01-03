package greedy

import (
	"fmt"
	"sort"
)

const infinity = int(^uint(0) >> 1)

//Node : Holds the information for each vertex
type Node struct {
	Name string
}

// Edge holds information for each edge in the Graph
type Edge struct {
	source *Node
	dest   *Node
	cost   int
}

// Graph holds the graph
type Graph struct {
	Edges []*Edge
	Nodes map[*Node]bool
}

// AddNode : adds node in the graph
func (g *Graph) AddNode(node *Node) {
	if g.Nodes == nil {
		g.Nodes = make(map[*Node]bool)
	}

	if _, ok := g.Nodes[node]; !ok {
		g.Nodes[node] = true
	}
}

// AddEdge : Adds edge in the graph
func (g *Graph) AddEdge(src, dest *Node, cost int) {
	e := &Edge{
		source: src,
		dest:   dest,
		cost:   cost,
	}
	g.Edges = append(g.Edges, e)
	g.AddNode(src)
	g.AddNode(dest)
}

// DijkstraGraphCreation : Starting point of Dijkstra Algo - Creates graph
func DijkstraGraphCreation() map[*Node]int {
	a := &Node{"a"}
	b := &Node{"b"}
	c := &Node{"c"}
	d := &Node{"d"}
	e := &Node{"e"}
	f := &Node{"f"}
	g := &Node{"g"}

	graph := Graph{}
	graph.AddEdge(c, a, 4)
	graph.AddEdge(c, f, 6)
	graph.AddEdge(a, f, 5)
	graph.AddEdge(a, d, 7)
	graph.AddEdge(f, d, 2)
	graph.AddEdge(a, b, 2)
	graph.AddEdge(d, b, 6)
	graph.AddEdge(d, g, 6)
	graph.AddEdge(f, g, 6)
	graph.AddEdge(b, g, 8)
	graph.AddEdge(b, e, 3)
	graph.AddEdge(g, e, 7)

	costTable := graph.Dijkstra(a)
	for node, cost := range costTable {
		fmt.Printf("To reach destination[%v] from source[%v] = Cost[%d]", node, a, cost)
	}

	return costTable

}

// Dijkstra : Dijkstra Algorithm Implementation
func (g *Graph) Dijkstra(source *Node) map[*Node]int {
	costTable := createCostTable(g.Nodes, source)
	var visited []*Node

	for len(visited) != len(g.Nodes) {
		node := g.getClosestNonVistedNode(costTable, visited)
		visited = append(visited, node)

		for _, edges := range g.Edges {
			if edges.source == node {
				neighborCost := costTable[node] + edges.cost
				costTable[edges.dest] = min(neighborCost, costTable[edges.dest])
			} else if edges.dest == node {
				neighborCost := costTable[node] + edges.cost
				costTable[edges.source] = min(neighborCost, costTable[edges.source])
			}
		}
		fmt.Printf("\nNode[%v] costTable:", node.Name)
		showCostTable(costTable)
	}
	return costTable
}

func showCostTable(costTable map[*Node]int) {
	for key, val := range costTable {
		fmt.Printf("\t Node[%v]-{%v} ", key.Name, val)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// getClosestNonVistedNode : creates a sorted cost table based on the non-vistied nodes and then returns the lowest cost node
func (g *Graph) getClosestNonVistedNode(costTable map[*Node]int, visited []*Node) *Node {
	type costTableNode struct {
		node *Node
		cost int
	}

	var costTableToSort []costTableNode

	for node, cost := range costTable {
		isVisited := false
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}

		if !isVisited {
			costTableToSort = append(costTableToSort, costTableNode{node, cost})
		}
	}

	sort.Slice(costTableToSort, func(i, j int) bool {
		return costTableToSort[i].cost < costTableToSort[j].cost
	})

	return costTableToSort[0].node
}

func createCostTable(nodes map[*Node]bool, source *Node) map[*Node]int {
	costTable := make(map[*Node]int)
	costTable[source] = 0

	for key := range nodes {
		if key != source {
			costTable[key] = infinity
		}
	}
	return costTable
}

package graph

import (
	"fmt"
	"testing"
)

func TestUndirectedGraph(t *testing.T) {
	g := NewUndirected()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		g.AddVertex(v)
	}

	if len(g.edges) != 10 {
		fmt.Println(g)
		t.Error()
	}

	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}

	if g.IsEdge(0, 8) == false || g.IsEdge(0, 9) == true || g.CheckVertex(2) != true {
		fmt.Println(g)
		t.Error()
	}

	// AddEdge should fail for already existing Edge
	err := g.AddEdge(0, 2, 1)
	if err == nil {
		fmt.Println(g)
		t.Error()
	}

	// AddVertex should fail for already existing vertex
	err = g.AddVertex(0)
	if err == nil {
		fmt.Println(g)
		t.Error()
	}

	g.RemoveVertex(VertexId(9))

	if g.IsVertex(VertexId(9)) {
		fmt.Println(g.edges[9] == nil)
		t.Error()
	}

	// RemoveVertex should fail for unknown vertex
	err = g.RemoveVertex(VertexId(9))

	if err == nil {
		fmt.Println(g.edges[9] == nil)
		t.Error()
	}

	g.RemoveEdge(0, 8)

	if g.IsEdge(VertexId(0), VertexId(8)) == true || g.edgesCount != 7 {
		fmt.Println(g.IsEdge(VertexId(0), VertexId(8)), g.edgesCount)
		t.Error()
	}

	// RemoveEdge should fail for unknown egde
	err = g.RemoveEdge(0, 8)

	if err == nil {
		fmt.Println(g)
		t.Error()
	}

	c := g.EdgesIter()

	countEdge := 0
	for _ = range c {
		countEdge++
	}

	if g.EdgesCount() != countEdge {
		t.Error()
	}

	d := g.VerticesIter()
	verticesCount := g.Order()
	countVertices := 0

	for _ = range d {
		countVertices++
	}

	if countVertices != verticesCount {
		fmt.Println(countVertices, g.edges)
		t.Error()
	}

	g.TouchVertex(9)
	if _, ok := g.edges[9]; !ok {
		t.Error()
	}
}

func TestDirectedGraph(t *testing.T) {
	g := NewDirected()

	for i := 0; i < 10; i++ {
		g.AddVertex(VertexId(i))
	}

	if len(g.edges) != 10 {
		fmt.Println(g)
		t.Error()
	}

	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2), 1)
	}

	r := g.Reverse()
	if !r.IsEdge(1, 7) || r.IsEdge(8, 0) {
		fmt.Println(r.edges)
		t.Error()
	}

	successors := r.GetSuccessors(1).VerticesIter()
	for successor := range successors {
		if successor != 3 &&
			successor != 5 &&
			successor != 7 &&
			successor != 9 {
			t.Error()
		}
	}

	predecessors := g.GetPredecessors(1).VerticesIter()
	for predecessor := range predecessors {
		if predecessor != 3 &&
			predecessor != 5 &&
			predecessor != 7 &&
			predecessor != 9 {
			t.Error()
		}
	}
}

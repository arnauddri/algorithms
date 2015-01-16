package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		g.AddVertex(v)
	}

	if len(g.edges) != 10 {
		fmt.Println(g)
		t.Error()
	}

	for i := 0; i < 10; i++ {
		g.AddEdge(VertexId(i), VertexId(i%2))
	}

	if g.isEdge(0, 8) == false || g.isEdge(0, 9) == true || g.CheckVertex(2) != true {
		fmt.Println(g)
		t.Error()
	}

	// AddEdge should fail for already existing Edge
	err := g.AddEdge(0, 2)
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

	if g.isVertex(VertexId(9)) {
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

	if g.isEdge(VertexId(0), VertexId(8)) || g.edgesCount != 7 {
		fmt.Println(g.isEdge(VertexId(0), VertexId(8)), g.edgesCount)
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
		fmt.Println(countEdge, g.edges)
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

	g.touchVertex(9)
	if _, ok := g.edges[9]; !ok {
		t.Error()
	}

}

func TestBfs(t *testing.T) {
	h := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(VertexId(i), VertexId(i+1))
	}

	bfsMap := bfs(h, VertexId(2))

	for i := 0; i < len(bfsMap); i++ {
		if _, ok := bfsMap[VertexId(i)]; !ok {
			t.Error()
		}
	}
}

func TestBfsShortestPath(t *testing.T) {
	h := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(VertexId(i), VertexId(i+1))
	}

	distance := bfsShortestPath(h, VertexId(0))

	for i := 0; i < len(distance); i++ {
		if distance[VertexId(i)] != i {
			t.Error()
		}

		if bfsDist(h, VertexId(0), VertexId(i)) != i {
			t.Error()
		}
	}
}

func TestDfs(t *testing.T) {
	h := NewGraph()

	for i := 0; i < 10; i++ {
		v := VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(VertexId(i), VertexId(i+1))
	}

	dfsMap := dfs(h, VertexId(2))

	for i := 0; i < len(dfsMap); i++ {
		if _, ok := dfsMap[VertexId(i)]; !ok {
			t.Error()
		}
	}
}

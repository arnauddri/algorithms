package yen

import (
	"reflect"
	"testing"

	"github.com/arnauddri/algorithms/data-structures/graph"
)

func v(ids ...int) []graph.VertexId {
	out := make([]graph.VertexId, len(ids))
	for i, id := range ids {
		out[i] = graph.VertexId(id)
	}
	return out
}

// The example from the Wikipedia article, with C..H mapped to 0..5.
func wikipediaGraph() *graph.DirGraph {
	g := graph.NewDirected()
	for i := 0; i < 6; i++ {
		g.AddVertex(graph.VertexId(i))
	}
	edges := [][3]int{
		{0, 1, 3}, // C -> D
		{0, 2, 2}, // C -> E
		{1, 3, 4}, // D -> F
		{2, 1, 1}, // E -> D
		{2, 3, 2}, // E -> F
		{2, 4, 3}, // E -> G
		{3, 4, 2}, // F -> G
		{3, 5, 1}, // F -> H
		{4, 5, 2}, // G -> H
	}
	for _, e := range edges {
		g.AddEdge(graph.VertexId(e[0]), graph.VertexId(e[1]), e[2])
	}
	return g
}

func TestKShortestPathsDirected(t *testing.T) {
	g := wikipediaGraph()

	paths := KShortestPaths(g, 0, 5, 4)

	expected := []Path{
		{Vertices: v(0, 2, 3, 5), Cost: 5},
		{Vertices: v(0, 2, 4, 5), Cost: 7},
		{Vertices: v(0, 1, 3, 5), Cost: 8},
		{Vertices: v(0, 2, 1, 3, 5), Cost: 8},
	}

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("got %v, expected %v", paths, expected)
	}
}

func TestKShortestPathsUndirected(t *testing.T) {
	g := graph.NewUndirected()
	for i := 0; i < 4; i++ {
		g.AddVertex(graph.VertexId(i))
	}
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(0, 2, 2)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 2, 1)

	paths := KShortestPaths(g, 0, 3, 10)

	expected := []Path{
		{Vertices: v(0, 1, 3), Cost: 2},
		{Vertices: v(0, 2, 3), Cost: 4}, // same cost as the next two: fewer vertices wins
		{Vertices: v(0, 1, 2, 3), Cost: 4},
		{Vertices: v(0, 2, 1, 3), Cost: 4},
	}

	if !reflect.DeepEqual(paths, expected) {
		t.Errorf("got %v, expected %v", paths, expected)
	}

	// Every returned path must be loopless.
	for _, p := range paths {
		seen := map[graph.VertexId]bool{}
		for _, x := range p.Vertices {
			if seen[x] {
				t.Errorf("path %v visits %d twice", p.Vertices, x)
			}
			seen[x] = true
		}
	}
}

func TestKShortestPathsEdgeCases(t *testing.T) {
	g := wikipediaGraph()

	if paths := KShortestPaths(g, 5, 0, 3); len(paths) != 0 {
		t.Errorf("unreachable target must return no path, got %v", paths)
	}

	if paths := KShortestPaths(g, 0, 5, 0); len(paths) != 0 {
		t.Errorf("k = 0 must return no path, got %v", paths)
	}

	if paths := KShortestPaths(g, 0, 42, 3); len(paths) != 0 {
		t.Errorf("unknown vertex must return no path, got %v", paths)
	}

	paths := KShortestPaths(g, 0, 0, 3)
	if len(paths) != 1 || paths[0].Cost != 0 || !reflect.DeepEqual(paths[0].Vertices, v(0)) {
		t.Errorf("source == target must return the empty path, got %v", paths)
	}
}

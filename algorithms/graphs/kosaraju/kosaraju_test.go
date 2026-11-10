package kosaraju

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestScc(t *testing.T) {
	h := graph.NewDirected()
	addEdge := func(h *graph.DirGraph, i, j int) {
		h.AddEdge(graph.VertexId(i), graph.VertexId(j), 1)
	}

	for i := 1; i < 10; i++ {
		h.AddVertex(graph.VertexId(i))
	}

	addEdge(h, 1, 7)
	addEdge(h, 7, 4)
	addEdge(h, 4, 1)

	addEdge(h, 7, 9)

	addEdge(h, 9, 6)
	addEdge(h, 6, 3)
	addEdge(h, 3, 9)

	addEdge(h, 6, 8)

	addEdge(h, 5, 8)
	addEdge(h, 2, 5)
	addEdge(h, 8, 2)

	// Three cycles of three vertices: {1,7,4}, {9,6,3}, {8,5,2}.
	sccs := Scc(h)
	if len(sccs) != 3 {
		t.Fatalf("expected 3 SCCs, got %d", len(sccs))
	}
	for _, scc := range sccs {
		if scc.Len() != 3 {
			t.Errorf("expected an SCC of 3 vertices, got %d", scc.Len())
		}
	}
}

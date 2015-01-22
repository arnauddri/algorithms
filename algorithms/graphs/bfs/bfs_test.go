package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestBfs(t *testing.T) {
	h := graph.NewGraph()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1))
	}

	bfsMap := make(map[graph.VertexId]bool)
	checkVertices := func(v graph.VertexId) {
		bfsMap[v] = true
	}

	Bfs(h, graph.VertexId(2), checkVertices)

	for i := 0; i < len(bfsMap); i++ {
		if _, ok := bfsMap[graph.VertexId(i)]; !ok {
			t.Error()
		}
	}
}

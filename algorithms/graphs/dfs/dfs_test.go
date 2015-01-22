package dfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestDfs(t *testing.T) {
	h := graph.NewGraph()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1))
	}

	dfsMap := make(map[graph.VertexId]bool)
	checkVertices := func(v graph.VertexId) {
		dfsMap[v] = true
	}

	dfs(h, graph.VertexId(2), checkVertices)

	for i := 0; i < len(dfsMap); i++ {
		if _, ok := dfsMap[graph.VertexId(i)]; !ok {
			t.Error()
		}
	}
}

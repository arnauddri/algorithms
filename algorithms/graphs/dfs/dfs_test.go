package dfs

import (
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestUndirectedDfs(t *testing.T) {
	h := graph.NewUndirected()

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

	undirectedDfs(h, graph.VertexId(2), checkVertices)

	for i := 0; i < len(dfsMap); i++ {
		if _, ok := dfsMap[graph.VertexId(i)]; !ok {
			fmt.Println(dfsMap)
			t.Error()
		}
	}
}

func TestDirectedDfs(t *testing.T) {
	h := graph.NewDirected()

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

	directedDfs(h, graph.VertexId(3), checkVertices)

	for i := 3; i < len(dfsMap); i++ {
		if _, ok := dfsMap[graph.VertexId(i)]; !ok {
			fmt.Println(dfsMap)
			t.Error()
		}
	}
}

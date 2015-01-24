package kosaraju

import (
	"fmt"
	"github.com/arnauddri/algorithms/algorithms/graphs/dfs"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
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

	dfs.DirectedDfs(h, graph.VertexId(9), func(v graph.VertexId) {
		//fmt.Println(v)
	})

	s := Scc(h)
	fmt.Println(s)

}

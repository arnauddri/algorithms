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
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	counter := 0
	UndirectedDfs(h, graph.VertexId(3), func(v graph.VertexId) {
		counter += int(v)
	})

	if counter != 45 {
		fmt.Println(counter)
		t.Error()
	}
}

func TestDirectedDfs(t *testing.T) {
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	counter := 0
	DirectedDfs(h, graph.VertexId(3), func(v graph.VertexId) {
		counter += int(v)
	})

	if counter != 42 {
		fmt.Println(counter)
		t.Error()
	}
}

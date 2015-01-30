package dijkstra

import (
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestDijkstra(t *testing.T) {
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 20*i)
	}

	counter := 0
	prev := ShortestPath(h, graph.VertexId(1))
	fmt.Println(prev)

	if counter != 0 {
		fmt.Println(counter)
		t.Error()
	}
}

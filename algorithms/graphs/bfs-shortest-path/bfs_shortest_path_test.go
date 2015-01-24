package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"testing"
)

func TestBfsShortestPath(t *testing.T) {
	h := graph.NewDirected()

	for i := 0; i < 10; i++ {
		v := graph.VertexId(i)
		h.AddVertex(v)
	}

	for i := 0; i < 9; i++ {
		h.AddEdge(graph.VertexId(i), graph.VertexId(i+1), 1)
	}

	distance := ShortestPath(h, graph.VertexId(0))

	for i := 0; i < len(distance); i++ {
		if distance[graph.VertexId(i)] != i {
			t.Error()
		}

		if GetDist(h, graph.VertexId(0), graph.VertexId(i)) != i {
			t.Error()
		}
	}
}

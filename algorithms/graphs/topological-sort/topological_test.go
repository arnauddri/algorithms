package topological

import (
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
	//"github.com/arnauddri/algorithms/data-structures/stack"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	h := graph.NewDirected()

	h.AddVertex(graph.VertexId(2))
	h.AddVertex(graph.VertexId(3))
	h.AddVertex(graph.VertexId(5))
	h.AddVertex(graph.VertexId(7))
	h.AddVertex(graph.VertexId(8))
	h.AddVertex(graph.VertexId(9))
	h.AddVertex(graph.VertexId(10))
	h.AddVertex(graph.VertexId(11))

	h.AddEdge(graph.VertexId(7), graph.VertexId(8))
	h.AddEdge(graph.VertexId(7), graph.VertexId(11))
	h.AddEdge(graph.VertexId(5), graph.VertexId(11))
	h.AddEdge(graph.VertexId(3), graph.VertexId(8))
	h.AddEdge(graph.VertexId(3), graph.VertexId(10))
	h.AddEdge(graph.VertexId(8), graph.VertexId(9))
	h.AddEdge(graph.VertexId(11), graph.VertexId(10))
	h.AddEdge(graph.VertexId(11), graph.VertexId(2))
	h.AddEdge(graph.VertexId(11), graph.VertexId(9))

	s := Sort(h)

	first := int(s.Pop().(graph.VertexId))
	if first != 7 &&
		first != 5 &&
		first != 3 {
		fmt.Print(first, first != 7)
		t.Error()
	}
}

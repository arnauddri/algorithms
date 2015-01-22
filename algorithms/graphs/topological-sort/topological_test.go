package topological

import (
	//"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
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

	//for i := 0; i < 10; i++ {
	//h.AddVertex(graph.VertexId(i))
	//}

	//for i := 0; i < 9; i++ {
	//h.AddEdge(graph.VertexId(i), graph.VertexId(i+1))
	//}

	Sort(h)
	//for i := 3; i < len(dfsMap); i++ {
	//if _, ok := dfsMap[graph.VertexId(i)]; !ok {
	//fmt.Println(dfsMap)
	//t.Error()
	//}
	//}
}

package dijkstra

import (
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
	//"github.com/arnauddri/algorithms/data-structures/queue"
)

func ShortestPath(g *graph.DirGraph, source graph.VertexId) {
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	Q := make([]graph.VertexId, 0)
	vertices := g.VerticesIter()

	for vertex := range vertices {
		if source != vertex {
			dist[vertex] = -1
			prev[vertex] = 0
		}
		Q = append(Q, vertex)
	}

	for len(Q) > 0 {
		min := -1
		var u graph.VertexId
		for vertex := range Q {
			if int(dist[graph.VertexId(vertex)]) > 0 && int(dist[graph.VertexId(vertex)]) < min {
				min = dist[graph.VertexId(vertex)]
				u = graph.VertexId(vertex)
			}
		}

		for successor := range g.GetSuccessors(u).VerticesIter() {
			fmt.Println(successor)
			alt := dist[u] + g.GetEdge(u, successor)

			if alt < dist[successor] {
				dist[successor] = alt
				prev[successor] = u
			}
		}
	}
}

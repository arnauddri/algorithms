// Dijkstra algorithm for undirected weighed graph
// Pseudo-code: http://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Using_a_priority_queue
package dijkstra

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/priority-queue"
)

func ShortestPath(g *graph.UnGraph, source graph.VertexId) map[graph.VertexId]graph.VertexId {
	visited := make(map[graph.VertexId]bool, g.VerticesCount())
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	Q := pq.NewMin()
	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if source != vertex {
			dist[vertex] = 1000
			prev[vertex] = 0
		}
		Q.Insert(*pq.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(graph.VertexId)
		visited[u] = true

		for neighbour := range g.GetNeighbours(u).VerticesIter() {
			if !visited[neighbour] {
				alt := dist[u] + g.GetEdge(u, neighbour)

				if alt < dist[neighbour] {
					dist[neighbour] = alt
					prev[neighbour] = u
					Q.ChangePriority(neighbour, alt)
				}
			}
		}
	}
	return prev
}

package dijkstra

import (
	"fmt"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/priority-queue"
)

func ShortestPath(g *graph.DirGraph, source graph.VertexId) map[graph.VertexId]int {
	visited := make(map[graph.VertexId]bool, g.VerticesCount())
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	Q := pq.New()
	vertices := g.VerticesIter()

	dist[source] = 0
	for vertex := range vertices {
		if source != vertex {
			dist[vertex] = -1
			prev[vertex] = 0
		}
		Q.Insert(*pq.NewItem(vertex, dist[vertex]))
	}

	for Q.Len() > 0 {
		u := Q.Extract().Value.(graph.VertexId)
		visited[u] = true

		for successor := range g.GetSuccessors(u).VerticesIter() {
			if !visited[successor] {
				alt := dist[u] + g.GetEdge(u, successor)
				fmt.Println(u, dist[u], successor, alt, g.GetEdge(u, successor))

				if alt < dist[successor] || dist[successor] == -1 {
					dist[successor] = alt
					prev[successor] = u
					fmt.Println("alt", successor, alt, dist)
					Q.ChangePriority(successor, alt)
				}
			}
		}
	}
	return dist
}

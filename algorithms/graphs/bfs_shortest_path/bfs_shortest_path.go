package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
)

func bfsShortestPath(g *graph.Graph, start graph.VertexId) (dist map[graph.VertexId]int) {
	queue := []graph.VertexId{start}
	visited := make(map[graph.VertexId]bool)
	dist = make(map[graph.VertexId]int)
	var next []graph.VertexId

	for len(queue) > 0 {
		next = []graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()

			for neighbour := range neighbours {

				ok, _ := visited[neighbour]
				if !ok {
					dist[neighbour] = dist[vertex] + 1
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
	return
}

func bfsDist(g *graph.Graph, from graph.VertexId, to graph.VertexId) int {
	return bfsShortestPath(g, from)[to]
}

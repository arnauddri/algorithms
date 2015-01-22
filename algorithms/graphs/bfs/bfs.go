package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
)

func Bfs(g *graph.DirGraph, start graph.VertexId, fn func(graph.VertexId)) {
	queue := []graph.VertexId{start}
	visited := make(map[graph.VertexId]bool)
	var next []graph.VertexId

	for len(queue) > 0 {
		next = []graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()
			fn(vertex)

			for neighbour := range neighbours {

				_, ok := visited[neighbour]
				if !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
}

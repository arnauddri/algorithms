package dfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
)

func dfs(g *graph.Graph, v graph.VertexId) map[graph.VertexId]bool {
	stack := []graph.VertexId{v}
	visited := make(map[graph.VertexId]bool)

	for len(stack) > 0 {
		l := len(stack) - 1
		v = stack[l]
		stack = stack[:l]

		if _, ok := visited[v]; !ok {
			visited[v] = true
			neighbours := g.GetNeighbours(v).VerticesIter()
			for neighbour := range neighbours {
				stack = append(stack, neighbour)
			}
		}
	}
	return visited
}

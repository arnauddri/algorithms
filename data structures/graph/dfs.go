package graph

import ()

func dfs(g *Graph, v VertexId) map[VertexId]bool {
	stack := []VertexId{v}
	visited := make(map[VertexId]bool)

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

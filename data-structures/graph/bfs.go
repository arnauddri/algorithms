package graph

import ()

func bfs(g *Graph, start VertexId) map[VertexId]bool {
	queue := []VertexId{start}
	visited := make(map[VertexId]bool)
	var next []VertexId

	for len(queue) > 0 {
		next = []VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()

			for neighbour := range neighbours {

				_, ok := visited[neighbour]
				if !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
	return visited
}

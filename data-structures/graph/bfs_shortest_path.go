package graph

import ()

func bfsShortestPath(g *Graph, start VertexId) (dist map[VertexId]int) {
	queue := []VertexId{start}
	visited := make(map[VertexId]bool)
	dist = make(map[VertexId]int)
	var next []VertexId

	for len(queue) > 0 {
		next = []VertexId{}
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

func bfsDist(g *Graph, from VertexId, to VertexId) int {
	return bfsShortestPath(g, from)[to]
}

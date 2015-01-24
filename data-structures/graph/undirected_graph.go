package graph

type UnGraph struct {
	graph
}

func NewUndirected() *UnGraph {
	return &UnGraph{
		graph{
			edgesCount: 0,
			edges:      make(map[VertexId]map[VertexId]int),
			isDirected: false,
		},
	}
}

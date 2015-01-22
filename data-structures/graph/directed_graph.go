package graph

type DirGraph struct {
	graph
}

func NewDirected() *DirGraph {
	return &DirGraph{
		graph{
			edgesCount: 0,
			edges:      make(map[VertexId]map[VertexId]bool),
			isDirected: true,
		},
	}
}

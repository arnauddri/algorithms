package graph

type UnGraph struct {
	graph
}

func NewUndirected() *UnGraph {
	return &UnGraph{
		graph{
			edgesCount: 0,
			edges:      make(map[VertexId]map[VertexId]bool),
			isDirected: false,
		},
	}
}

func (g *UnGraph) GetNeighbours(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			if connected, ok := g.edges[vertex]; ok {
				for VertexId, _ := range connected {
					ch <- VertexId
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&vertexIterableHelper{iterFunc: iterator})
}

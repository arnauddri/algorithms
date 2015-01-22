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

func (g *DirGraph) GetPredecessors(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			for i := 0; i < g.Order(); i++ {
				if g.edges[VertexId(i)][vertex] {
					ch <- VertexId(i)
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&vertexIterableHelper{iterFunc: iterator})
}

func (g *DirGraph) GetSuccessors(vertex VertexId) VerticesIterable {
	iterator := func() <-chan VertexId {
		ch := make(chan VertexId)
		go func() {
			for i := 0; i < g.Order(); i++ {
				if g.edges[vertex][VertexId(i)] {
					ch <- VertexId(i)
				}
			}
			close(ch)
		}()
		return ch
	}

	return VerticesIterable(&vertexIterableHelper{iterFunc: iterator})
}

func (g *DirGraph) Reverse() *DirGraph {
	r := NewDirected()
	edges := g.EdgesIter()

	for edge := range edges {
		r.AddEdge(edge.To, edge.From)
	}

	return r
}

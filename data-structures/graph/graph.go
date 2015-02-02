package graph

import (
	"errors"
)

type VertexId uint

type Vertices []VertexId

type Edge struct {
	From VertexId
	To   VertexId
}

type graph struct {
	edges      map[VertexId]map[VertexId]int
	edgesCount int
	isDirected bool
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VerticesIterable interface {
	VerticesIter() <-chan VertexId
}

func (g *graph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if g.isDirected {
					ch <- Edge{from, to}
				} else {
					if from < to {
						ch <- Edge{from, to}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *graph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *graph) CheckVertex(vertex VertexId) bool {
	_, exists := g.edges[vertex]

	return exists
}

func (g *graph) TouchVertex(vertex VertexId) {
	if _, ok := g.edges[vertex]; !ok {
		g.edges[vertex] = make(map[VertexId]int)
	}
}

func (g *graph) AddVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i != nil {
		return errors.New("Vertex already exists")
	}

	g.edges[vertex] = make(map[VertexId]int)

	return nil
}

func (g *graph) RemoveVertex(vertex VertexId) error {
	if !g.IsVertex(vertex) {
		return errors.New("Unknown vertex")
	}

	delete(g.edges, vertex)

	for _, connectedVertices := range g.edges {
		delete(connectedVertices, vertex)
	}

	return nil
}

func (g *graph) IsVertex(vertex VertexId) (exist bool) {
	_, exist = g.edges[vertex]

	return
}

func (g *graph) VerticesCount() int {
	return len(g.edges)
}

func (g *graph) AddEdge(from, to VertexId, weight int) error {
	if from == to {
		return errors.New("Cannot add self loop")
	}

	if !g.CheckVertex(from) || !g.CheckVertex(to) {
		return errors.New("Vertices don't exist")
	}

	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i > 0 || j > 0 {
		return errors.New("Edge already defined")
	}

	g.TouchVertex(from)
	g.TouchVertex(to)

	g.edges[from][to] = weight

	if !g.isDirected {
		g.edges[to][from] = weight
	}

	g.edgesCount++

	return nil
}

func (g *graph) RemoveEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == -1 || j == -1 {
		return errors.New("Edge doesn't exist")
	}

	g.edges[from][to] = -1

	if !g.isDirected {
		g.edges[to][from] = -1
	}

	g.edgesCount--

	return nil
}

func (g *graph) IsEdge(from, to VertexId) bool {
	connected, ok := g.edges[from]

	if !ok {
		return false
	}

	weight := connected[to]
	return weight > 0
}

func (g *graph) Order() int {
	return len(g.edges)
}

func (g *graph) EdgesCount() int {
	return g.edgesCount
}

func (g *graph) GetEdge(from, to VertexId) int {
	return g.edges[from][to]
}

func (g *graph) GetNeighbours(vertex VertexId) VerticesIterable {
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

type vertexIterableHelper struct {
	iterFunc func() <-chan VertexId
}

func (helper *vertexIterableHelper) VerticesIter() <-chan VertexId {
	return helper.iterFunc()
}

package graph

import (
	"errors"
)

type VertexId uint

type Vertices []VertexId

type Edge struct {
	Tail VertexId
	Head VertexId
}

type EdgesIterable interface {
	EdgesIter() <-chan Edge
}

type VerticesIterable interface {
	VerticesIter() <-chan VertexId
}

type Graph interface {
	EdgesIter()
	VerticesIter()
	CheckVertex(vertex VertexId) bool
	TouchVertex(vertex VertexId)
	AddVertex(vertex VertexId) error
	RemoveVertex(vertex VertexId) error
	IsVertex(vertex VertexId) bool
	AddEdge(from, to VertexId) error
	RemoveEdge(from, to VertexId) error
	IsEdge(from, to VertexId) bool
	Order() int
	EdgesCount() int
}

type graph struct {
	edges      map[VertexId]map[VertexId]bool
	edgesCount int
	isDirected bool
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
		g.edges[vertex] = make(map[VertexId]bool)
	}
}

func (g *graph) AddVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i != nil {
		return errors.New("Vertex already exists")
	}

	g.edges[vertex] = make(map[VertexId]bool)

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

func (g *graph) AddEdge(from, to VertexId) error {
	if from == to {
		return errors.New("Cannot add self loop")
	}

	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == true || j == true {
		return errors.New("Edge already defined")
	}

	g.TouchVertex(from)
	g.TouchVertex(to)

	g.edges[from][to] = true

	if !g.isDirected {
		g.edges[to][from] = true
	}

	g.edgesCount++

	return nil
}

func (g *graph) RemoveEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == false || j == false {
		return errors.New("Edge doesn't exist")
	}

	g.edges[from][to] = false

	if !g.isDirected {
		g.edges[to][from] = false
	}

	g.edgesCount--

	return nil
}

func (g *graph) IsEdge(from, to VertexId) (exist bool) {
	connected, ok := g.edges[from]

	if !ok {
		panic("Vertex doesn't exit")
	}

	exist, _ = connected[to]
	return
}

func (g *graph) Order() int {
	return len(g.edges)
}

func (g *graph) EdgesCount() int {
	return g.edgesCount
}

type vertexIterableHelper struct {
	iterFunc func() <-chan VertexId
}

func (helper *vertexIterableHelper) VerticesIter() <-chan VertexId {
	return helper.iterFunc()
}

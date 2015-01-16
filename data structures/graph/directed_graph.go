package graph

import (
	"errors"
)

type DirGraph struct {
	edges      map[VertexId]map[VertexId]bool
	edgesCount int
}

func NewDirGraph() *DirGraph {
	g := new(DirGraph)
	g.edges = make(map[VertexId]map[VertexId]bool)
	g.edgesCount = 0

	return g
}

func (g *DirGraph) EdgesIter() <-chan Edge {
	ch := make(chan Edge)
	go func() {
		for from, connectedVertices := range g.edges {
			for to, _ := range connectedVertices {
				if from < to {
					ch <- Edge{from, to}
				}
			}
		}
		close(ch)
	}()
	return ch
}

func (g *DirGraph) VerticesIter() <-chan VertexId {
	ch := make(chan VertexId)
	go func() {
		for vertex, _ := range g.edges {
			ch <- vertex
		}
		close(ch)
	}()
	return ch
}

func (g *DirGraph) CheckVertex(vertex VertexId) bool {
	_, exists := g.edges[vertex]

	return exists
}

func (g *DirGraph) touchVertex(vertex VertexId) {
	if _, ok := g.edges[vertex]; !ok {
		g.edges[vertex] = make(map[VertexId]bool)
	}
}

func (g *DirGraph) AddVertex(vertex VertexId) error {
	i, _ := g.edges[vertex]
	if i != nil {
		return errors.New("Vertex already exists")
	}

	g.edges[vertex] = make(map[VertexId]bool)

	return nil
}

func (g *DirGraph) RemoveVertex(vertex VertexId) error {
	if !g.isVertex(vertex) {
		return errors.New("Unknown vertex")
	}

	delete(g.edges, vertex)

	for _, connectedVertices := range g.edges {
		delete(connectedVertices, vertex)
	}

	return nil
}

func (g *DirGraph) AddEdge(from, to VertexId) error {
	if from == to {
		return errors.New("Cannot add self loop")
	}

	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == true || j == true {
		return errors.New("Edge already defined")
	}

	g.touchVertex(from)
	g.touchVertex(to)

	g.edges[from][to] = true

	g.edgesCount++

	return nil
}

func (g *DirGraph) RemoveEdge(from, to VertexId) error {
	i, _ := g.edges[from][to]
	j, _ := g.edges[to][from]

	if i == false || j == false {
		return errors.New("Edge doesn't exist")
	}

	g.edges[from][to] = false

	g.edgesCount--

	return nil
}

func (g *DirGraph) Order() int {
	return len(g.edges)
}

func (g *DirGraph) EdgesCount() int {
	return g.edgesCount
}

func (g *DirGraph) GetNeighbours(vertex VertexId) VerticesIterable {
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

	return VerticesIterable(&_vertexIterableHelper{iterFunc: iterator})
}

func (g *DirGraph) isVertex(vertex VertexId) (exist bool) {
	_, exist = g.edges[vertex]

	return
}

func (g *DirGraph) isEdge(from, to VertexId) (exist bool) {
	connected, ok := g.edges[from]

	if !ok {
		panic("Vertex doesn't exit")
	}

	exist, _ = connected[to]
	return
}

package topological

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/stack"
)

func Sort(g *graph.DirGraph) *stack.Stack {
	var visit func(v graph.VertexId)
	sorted := stack.New()

	visited := make(map[graph.VertexId]bool)

	mark := make(map[graph.VertexId]bool)

	visit = func(v graph.VertexId) {
		if mark[v] {
			return
		}

		mark[v] = true
		successors := g.GetSuccessors(v).VerticesIter()

		for successor := range successors {
			if !mark[successor] && !visited[successor] {
				visit(successor)
			}
		}

		visited[v] = true
		mark[v] = false
		sorted.Push(v)
	}

	vertices := g.VerticesIter()
	for vertex := range vertices {
		if !visited[vertex] {
			visit(vertex)
		}
	}

	return sorted
}

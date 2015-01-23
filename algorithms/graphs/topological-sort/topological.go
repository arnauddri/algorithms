package topological

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/stack"
)

func Sort(g *graph.DirGraph) *stack.Stack {
	var visit func(v graph.VertexId)
	sorted := stack.New()
	visited := make(map[graph.VertexId]bool)
	marked := make(map[graph.VertexId]bool)

	visit = func(v graph.VertexId) {
		if marked[v] {
			panic("Not a Directed Acyclic Graph (DAG)")
		}

		marked[v] = true
		successors := g.GetSuccessors(v).VerticesIter()

		for successor := range successors {
			if !marked[successor] && !visited[successor] {
				visit(successor)
			}
		}

		visited[v] = true
		marked[v] = false
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

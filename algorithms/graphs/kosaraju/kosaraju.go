package kosaraju

import (
	"github.com/arnauddri/algorithms/algorithms/graphs/dfs"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/stack"
)

func Scc(g *graph.DirGraph) []stack.Stack {
	s := stack.New()
	n := g.Order()
	scc := make([]stack.Stack, 0)
	visitedOnce := make(map[graph.VertexId]bool)
	vertices := g.VerticesIter()

	for s.Len() != n {
		for vertex := range vertices {
			dfs.DirectedDfs(g, vertex, func(v graph.VertexId) {
				if visitedOnce[v] == false {
					s.Push(v)
					visitedOnce[v] = true
				}
			})
		}
	}

	r := g.Reverse()
	visitedTwice := make(map[graph.VertexId]bool)
	index := 0
	for s.Len() > 0 {
		vertex := s.Pop()
		scc = append(scc, *stack.New())

		if visitedTwice[vertex.(graph.VertexId)] == false {
			dfs.DirectedDfs(r, vertex.(graph.VertexId), func(v graph.VertexId) {
				if visitedTwice[graph.VertexId(v)] == false {
					visitedTwice[graph.VertexId(v)] = true
					scc[len(scc)-1].Push(v)
				}
			})
		}
		index++
	}
	return scc
}

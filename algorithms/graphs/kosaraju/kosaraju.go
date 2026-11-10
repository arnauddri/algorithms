package kosaraju

import (
	"github.com/arnauddri/algorithms/algorithms/graphs/dfs"
	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/stack"
)

// Scc returns the strongly connected components with more than one vertex.
func Scc(g *graph.DirGraph) []*stack.Stack {
	s := stack.New()
	n := g.Order()
	SCCs := make([]*stack.Stack, 0)
	visited := make(map[graph.VertexId]bool)
	vertices := g.VerticesIter()

	for s.Len() != n {
		vertex := <-vertices
		dfs.DirectedDfs(g, vertex, func(v graph.VertexId) {
			if visited[v] == false {
				s.Push(v)
				visited[v] = true
			}
		})
	}

	r := g.Reverse()
	visited = make(map[graph.VertexId]bool)
	for s.Len() > 0 {
		vertex := s.Pop().(graph.VertexId)
		scc := stack.New()

		if visited[vertex] == false {
			dfs.DirectedDfs(r, vertex, func(v graph.VertexId) {
				if visited[graph.VertexId(v)] == false {
					visited[graph.VertexId(v)] = true
					scc.Push(v)
				}
			})
		}
		if scc.Len() > 1 {
			SCCs = append(SCCs, scc)
		}
	}
	return SCCs
}

// Yen's algorithm: the K loopless shortest paths between two vertices of a
// weighted graph, found by repeatedly running Dijkstra on "spur" paths that
// deviate from the already found paths.
// Pseudo-code: https://en.wikipedia.org/wiki/Yen%27s_algorithm
package yen

import (
	"math"
	"sort"

	"github.com/arnauddri/algorithms/data-structures/graph"
	"github.com/arnauddri/algorithms/data-structures/priority-queue"
)

// Graph is satisfied by both *graph.UnGraph and *graph.DirGraph.
type Graph interface {
	VerticesIter() <-chan graph.VertexId
	GetNeighbours(graph.VertexId) graph.VerticesIterable
	GetEdge(from, to graph.VertexId) int
	IsEdge(from, to graph.VertexId) bool
}

// Path is a sequence of vertices from source to target and its total weight.
type Path struct {
	Vertices []graph.VertexId
	Cost     int
}

type edge struct {
	from, to graph.VertexId
}

// KShortestPaths returns up to k loopless paths from source to target, from
// cheapest to most expensive. Ties are broken by number of vertices, then by
// vertex ids. Fewer than k paths are returned when the graph has fewer.
func KShortestPaths(g Graph, source, target graph.VertexId, k int) []Path {
	var found []Path
	if k <= 0 {
		return found
	}

	first, ok := shortestPath(g, source, target, nil, nil)
	if !ok {
		return found
	}
	found = append(found, first)

	var candidates []Path
	seen := map[string]bool{key(first.Vertices): true}

	for len(found) < k {
		last := found[len(found)-1].Vertices

		for i := 0; i < len(last)-1; i++ {
			spur := last[i]
			root := last[:i+1]

			// Block the edges that already lead out of this root in a found path.
			blockedEdges := make(map[edge]bool)
			for _, p := range found {
				if len(p.Vertices) > i && samePrefix(p.Vertices, root) {
					blockedEdges[edge{p.Vertices[i], p.Vertices[i+1]}] = true
				}
			}

			// Block the root vertices (except the spur) to keep paths loopless.
			blockedVertices := make(map[graph.VertexId]bool)
			for _, v := range root[:i] {
				blockedVertices[v] = true
			}

			spurPath, ok := shortestPath(g, spur, target, blockedEdges, blockedVertices)
			if !ok {
				continue
			}

			vertices := append(append([]graph.VertexId{}, root...), spurPath.Vertices[1:]...)
			k := key(vertices)
			if seen[k] {
				continue
			}
			seen[k] = true

			candidates = append(candidates, Path{
				Vertices: vertices,
				Cost:     pathCost(g, root) + spurPath.Cost,
			})
		}

		if len(candidates) == 0 {
			break
		}

		sort.Slice(candidates, func(a, b int) bool { return less(candidates[a], candidates[b]) })
		found = append(found, candidates[0])
		candidates = candidates[1:]
	}

	return found
}

// shortestPath is Dijkstra from source to target that ignores the given edges
// and vertices. Already settled vertices are skipped when popped again.
func shortestPath(g Graph, source, target graph.VertexId, blockedEdges map[edge]bool, blockedVertices map[graph.VertexId]bool) (Path, bool) {
	dist := make(map[graph.VertexId]int)
	prev := make(map[graph.VertexId]graph.VertexId)
	visited := make(map[graph.VertexId]bool)

	for v := range g.VerticesIter() {
		dist[v] = math.MaxInt
	}
	if _, ok := dist[source]; !ok || blockedVertices[source] {
		return Path{}, false
	}
	dist[source] = 0

	Q := pq.NewMin()
	Q.Insert(*pq.NewItem(source, 0))

	for Q.Len() > 0 {
		u := Q.Extract().Value.(graph.VertexId)
		if visited[u] {
			continue
		}
		visited[u] = true
		if u == target {
			break
		}

		for n := range g.GetNeighbours(u).VerticesIter() {
			if visited[n] || blockedVertices[n] || blockedEdges[edge{u, n}] || !g.IsEdge(u, n) {
				continue
			}
			alt := dist[u] + g.GetEdge(u, n)
			if alt < dist[n] {
				dist[n] = alt
				prev[n] = u
				Q.Insert(*pq.NewItem(n, alt))
			}
		}
	}

	if !visited[target] {
		return Path{}, false
	}

	var vertices []graph.VertexId
	for v := target; ; v = prev[v] {
		vertices = append([]graph.VertexId{v}, vertices...)
		if v == source {
			break
		}
	}

	return Path{Vertices: vertices, Cost: dist[target]}, true
}

func pathCost(g Graph, vertices []graph.VertexId) int {
	cost := 0
	for i := 0; i < len(vertices)-1; i++ {
		cost += g.GetEdge(vertices[i], vertices[i+1])
	}
	return cost
}

func samePrefix(vertices, prefix []graph.VertexId) bool {
	if len(vertices) < len(prefix) {
		return false
	}
	for i, v := range prefix {
		if vertices[i] != v {
			return false
		}
	}
	return true
}

func less(a, b Path) bool {
	if a.Cost != b.Cost {
		return a.Cost < b.Cost
	}
	if len(a.Vertices) != len(b.Vertices) {
		return len(a.Vertices) < len(b.Vertices)
	}
	for i := range a.Vertices {
		if a.Vertices[i] != b.Vertices[i] {
			return a.Vertices[i] < b.Vertices[i]
		}
	}
	return false
}

func key(vertices []graph.VertexId) string {
	b := make([]byte, 0, len(vertices)*4)
	for _, v := range vertices {
		b = append(b, byte(v>>24), byte(v>>16), byte(v>>8), byte(v), ',')
	}
	return string(b)
}

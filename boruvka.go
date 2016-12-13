package boruvkamst

import (
	"fmt"
	"sort"

	"github.com/parpat/boruvkamst/quickFind"
)

//Input: A connected graph G whose edges have distinct weights
//1   Initialize a forest T to be a set of one-vertex trees, one for each vertex of the graph.
//2   While T has more than one component:
// 3     For each component C of T:
// 4       Begin with an empty set of edges S
// 5       For each vertex v in C:
// 6         Find the cheapest edge from v to a vertex outside of C, and add it to S
// 7       Add the cheapest edge in S to T
// 8
//   Combine trees connected by edges to form bigger components
// 9   Output: T is the minimum spanning tree of G.

func boruvka(g Graph) EdgeList {
	//mstEdges := *new(EdgeList)
	//1   Initialize a forest T to be a set of one-vertex trees, one for each vertex of the graph.
	forest := quickFind.InitializeQF(len(g.Vertices))
	T := *new(EdgeList)
	//2   While T has more than one component:
	for comps := forest.GetComponents(); len(comps) > 1; comps = forest.GetComponents() {
		//fmt.Printf("Components: %v\n", comps)
		// 3 For each component C of T:
		for _, c := range comps {
			// 4 Begin with an empty set of edges S
			S := *new(EdgeList)
			// 5 For each vertex v in C:
			fmt.Printf("component: %v\n", c)
			for v := 0; v < len(g.Vertices); v++ {
				if forest.ID[v] == c {
					// 6 Find the cheapest edge from v to a vertex outside of C, and add it to S
					for _, cedge := range g.Vertices[v].AdjEdges {
						if forest.Find(cedge.Start) != forest.Find(cedge.End) {
							S = append(S, cedge)
							break
						}
					}
				}
			}
			// 7 Add the cheapest edge in S to T
			if S != nil {
				sort.Sort(S)
				fmt.Printf("S: %v\n", S)
				T = append(T, S[0])
			}
		}
		// 8 Combine trees connected by edges to form bigger components
		for _, combEdge := range T {
			if forest.Find(combEdge.Start) != forest.Find(combEdge.End) {
				fmt.Printf("Combined vertices %v %v\n", combEdge.Start, combEdge.End)
				forest.Union(combEdge.Start, combEdge.End)
			}
		}
	}

	removeDuplicates(&T)
	sort.Sort(T)
	return T
}

func removeDuplicates(dup *EdgeList) {
	found := make(map[int]bool)
	j := 0
	for i, de := range *dup {
		if !found[de.Weight] {
			found[de.Weight] = true
			(*dup)[j] = (*dup)[i]
			j++
		}
	}
	*dup = (*dup)[:j]
}

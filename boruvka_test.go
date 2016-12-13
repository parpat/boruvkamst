package boruvkamst

import (
	"fmt"
	"testing"
)

var (
	testEdges = EdgeList{
		NewEdge(0, 1, 13),
		NewEdge(0, 2, 6),
		NewEdge(1, 2, 7),
		NewEdge(1, 3, 1),
		NewEdge(2, 3, 14), //4
		NewEdge(2, 4, 8),
		NewEdge(2, 7, 20),
		NewEdge(3, 4, 9),
		NewEdge(3, 5, 3),
		NewEdge(4, 5, 2), //9
		NewEdge(4, 9, 18),
		NewEdge(6, 7, 15),
		NewEdge(6, 8, 5),
		NewEdge(6, 9, 19),
		NewEdge(6, 10, 10), //14
		NewEdge(7, 9, 17),
		NewEdge(8, 10, 11),
		NewEdge(9, 10, 16),
		NewEdge(9, 11, 4),
		NewEdge(10, 11, 12), //19
	}

	expectedMST = EdgeList{
		testEdges[3],
		testEdges[9],
		testEdges[8],
		testEdges[18],
		testEdges[12],
		testEdges[1],
		testEdges[2],
		testEdges[14],
		testEdges[19],
		testEdges[11],
		testEdges[10],
	}
)

func TestBoruvka(t *testing.T) {
	g := NewGraph(12)
	for _, edge := range testEdges {
		g.AddEdge(edge)
	}

	mstEdges := boruvka(g)

	for i, mstEdge := range mstEdges {
		if mstEdge.Weight != expectedMST[i].Weight {
			t.Fail()
			t.Logf("Got MST edge: W%d -- Expected MST edge: W%d", mstEdge.Weight, expectedMST[i].Weight)
		}
		fmt.Printf(" MST edge: %d\n", mstEdge.Weight)
	}
}

// var (
// 	testEdges = EdgeList{
// 		NewEdge(4, 5, 35),
// 		NewEdge(4, 7, 37),
// 		NewEdge(5, 7, 28),
// 		NewEdge(0, 7, 16),
// 		NewEdge(1, 5, 32),
// 		NewEdge(0, 4, 38),
// 		NewEdge(2, 3, 17),
// 		NewEdge(1, 7, 19),
// 		NewEdge(0, 2, 26),
// 		NewEdge(1, 2, 36),
// 		NewEdge(1, 3, 29),
// 		NewEdge(2, 7, 34),
// 		NewEdge(6, 2, 40),
// 		NewEdge(3, 6, 52),
// 		NewEdge(6, 0, 58),
// 		NewEdge(6, 4, 93),
// 	}
//
// 	expectedMST = EdgeList{
// 		testEdges[3],
// 		testEdges[6],
// 		testEdges[7],
// 		testEdges[8],
// 		testEdges[2],
// 		testEdges[0],
// 		testEdges[12],
// 	}
// )

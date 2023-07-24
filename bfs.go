package main

import (
	"fmt"
	"math/rand"
)

// the max size of the graph.
const maxNodesInGraph = 1_00_000

// the max edges per node.
const maxEdges = 10

type Node struct {
	id       int
	edges    []Node
	distance int
}

func main() {
	/*
	 This program uses a BFS to find the distance between to given node.
	 It creates a random graph on the fly while loading edges (loadEdges)
	 and calculate the distance to node 21 from root node 1.
	 The loadEdges-method can be replaced via DB or API call.
	*/
	fmt.Println("Follow 21!")
	var root Node
	root.id = 1
	root.distance = 0

	distance := BFS(root, 21)
	fmt.Printf("Distance: %d\n", distance)
}

var numNodes = 0

func BFS(start Node, goalNodeId int) int {
	queue := make([]Node, 0)
	queue = append(queue, start)
	for len(queue) > 0 {
		if numNodes%1000 == 0 {
			fmt.Printf("Nodes: %d\n", len(nodes))
		}
		x := queue[0]
		queue = queue[1:]
		if x.id == goalNodeId {
			return x.distance
		}
		for _, edge := range loadEdges(x) {
			if edge.distance == 0 {
				edge.distance = x.distance + 1
				queue = append(queue, edge)
			}
		}
	}
	return 0
}

var nodes = make([]Node, 0)

func loadEdges(source Node) []Node {
	if contains(nodes, source) {
		return source.edges
	} else {
		numNodes++
		nodes = append(nodes, source)
	}
	numEdgesToCreate := rand.Intn(maxEdges)
	channels := make(chan Node, numEdgesToCreate)
	edges := make([]Node, numEdgesToCreate)
	for i := 1; i < numEdgesToCreate; i++ {
		go func(ch chan Node) {
			var edge Node
			edge.id = rand.Intn(maxNodesInGraph)
			edge.distance = 0
			ch <- getOrElse(nodes, edge)
		}(channels)
	}

	for i := 1; i < numEdgesToCreate; i++ {
		edges = append(edges, getOrElse(edges, <-channels))
	}

	source.edges = edges
	return edges
}

func contains(s []Node, e Node) bool {
	for _, a := range s {
		if a.id == e.id {
			return true
		}
	}
	return false
}

func getOrElse(s []Node, e Node) Node {
	for _, a := range s {
		if a.id == e.id {
			return a
		}
	}
	return e
}

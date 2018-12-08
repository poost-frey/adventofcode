package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// expected answer = 138
	input := "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

	tree := parseInput(input)

	node := parseNode(tree)

	sum := sumMetadata(node)

	fmt.Printf("%v", tree)
}

func parseNode(tree []int) node {
	numC := tree[0]
	numM := tree[1]
	c, m := getChildrenAndMetadata(numM, tree[2:])
	return node{
		numChildren: numC,
		numMetadata: numM,
	}
}

func getChildrenAndMetadata(metadataCount int, tree []int) ([]node, []int) {
	metadata := tree[len(tree)-1-metadataCount:]

	return nil, metadata
}

func sumMetadata(n node) int {
	sum := 0
	for _, c := range n.children {
		sum += sumMetadata(c)
	}
	for _, m := range n.metadata {
		sum += m
	}
	return sum
}

func parseInput(str string) (output []int) {
	split := strings.Split(str, " ")
	for _, r := range split {
		output = append(output, atoi(string(r)))
	}
	return
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

type node struct {
	numChildren int
	numMetadata int
	children    []node
	metadata    []int
}

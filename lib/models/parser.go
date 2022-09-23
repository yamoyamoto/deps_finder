package models

type Parser interface {
	FindDependingNodes(file File, allNodes []Node) []Node
}

package models

type Node struct {
	Id   int64
	Name string
	Path string
}

type Link struct {
	From     *Node
	To       *Node
	Strength float64
}

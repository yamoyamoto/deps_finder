package models

import "reflect"

type Node struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Size int64  `json:"size,omitempty"`
}

func (node Node) In(nodes []Node) bool {
	return contain(node, nodes)
}

func (node Node) Equal(node2 Node) bool {
	return node.Name == node2.Name
}

type Nodes []Node

func (nodes Nodes) Add(node Node) Nodes {
	for i, n := range nodes {
		if n.Name == node.Name {
			nodes[i].Size++
			return nodes
		}
	}
	return append(nodes, node)
}

type Link struct {
	From     *Node `json:"from,omitempty"`
	To       *Node `json:"to,omitempty"`
	Strength int64 `json:"strength,omitempty"`
}

func (link Link) In(links []Link) bool {
	return contain(link, links)
}

type Links []Link

func (links Links) Add(link Link) Links {
	for i, l := range links {
		if l.From.Equal(*link.From) && l.To.Equal(*link.To) {
			links[i].Strength++
			return links
		}
	}
	return append(links, link)
}

func contain[T any](elem T, list []T) bool {
	for _, e := range list {
		if reflect.DeepEqual(elem, e) {
			return true
		}
	}
	return false
}

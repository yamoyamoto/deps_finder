package finders

import "depsfinder/lib/models"

type Dependencies struct {
	Nodes models.Nodes
	Links models.Links
}

func (deps *Dependencies) Merge(mergedDeps *Dependencies) *Dependencies {
	if mergedDeps == nil {
		return deps
	}

	for _, node := range mergedDeps.Nodes {
		deps.Nodes = deps.Nodes.Add(node)
	}
	for _, link := range mergedDeps.Links {
		deps.Links = deps.Links.Add(link)
	}
	return deps
}

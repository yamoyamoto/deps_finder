package lib

import (
	"depsfinder/lib/finders"
	"log"
)

type DepsFinder interface {
	Find(dirPath string) (*finders.Dependencies, error)
}

type DependenciesDTO struct {
	Nodes []NodeDataDTO `json:"nodes"`
	Links []LinkDataDTO `json:"links"`
}

type NodeDataDTO struct {
	Name string `json:"name,omitempty"`
	Size int64  `json:"size,omitempty"`
}

type LinkDataDTO struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Strength int64  `json:"strength,omitempty"`
}

func FindDeps(depsFinder DepsFinder, dirPath string) (*DependenciesDTO, error) {
	dependencies, err := depsFinder.Find(dirPath)
	if err != nil {
		return nil, err
	}

	for _, link := range dependencies.Links {
		log.Printf("found dependencies: %s->%s, strength: %d", link.From.Name, link.To.Name, link.Strength)
	}

	nodeDataDTO := []NodeDataDTO{}
	for _, node := range dependencies.Nodes {
		nodeDataDTO = append(nodeDataDTO, NodeDataDTO{
			Name: node.Name,
			Size: node.Size,
		})
	}

	linkDataDTO := []LinkDataDTO{}
	for _, link := range dependencies.Links {
		linkDataDTO = append(linkDataDTO, LinkDataDTO{
			From:     link.From.Name,
			To:       link.To.Name,
			Strength: link.Strength,
		})
	}

	return &DependenciesDTO{
		Nodes: nodeDataDTO,
		Links: linkDataDTO,
	}, nil
}

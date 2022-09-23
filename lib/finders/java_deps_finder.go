package finders

import (
	"depsfinder/lib/models"
	"depsfinder/lib/services"
)

type JavaDepsFinder struct {
	Parser    JavaParser
	DirWalker services.DirWalker
}

type JavaParser interface {
	FindDependenciesFromFile(file models.File) *Dependencies
}

func NewJavaDepsFinder(parser JavaParser, walker services.DirWalker) *JavaDepsFinder {
	return &JavaDepsFinder{
		Parser:    parser,
		DirWalker: walker,
	}
}

func (finder *JavaDepsFinder) Find(dirPath string) (*Dependencies, error) {
	files, err := finder.DirWalker.ListFileRecursive(dirPath)
	if err != nil {
		return nil, err
	}

	dependencies := &Dependencies{
		Nodes: models.Nodes{},
		Links: models.Links{},
	}
	for _, file := range files {
		if file.FileType != models.Java {
			continue
		}
		deps := finder.Parser.FindDependenciesFromFile(file)
		dependencies = dependencies.Merge(deps)
	}

	return dependencies, nil
}

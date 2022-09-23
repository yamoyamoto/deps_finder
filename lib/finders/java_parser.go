package finders

import (
	"depsfinder/lib/models"
	"regexp"
)

type JavaParserImpl struct {
}

func NewJavaParser() *JavaParserImpl {
	return &JavaParserImpl{}
}

const (
	IMPORT_KEYWORD_REG  = "import\\s+(.+);"
	PACKAGE_KEYWORD_REG = "package\\s+(.+);"
)

var importKeywordReg = regexp.MustCompile(IMPORT_KEYWORD_REG)
var packageKeywordReg = regexp.MustCompile(PACKAGE_KEYWORD_REG)

func (parser *JavaParserImpl) FindDependenciesFromFile(file models.File) *Dependencies {
	content, err := file.GetContent()
	if err != nil {
		return nil
	}

	importMatches := importKeywordReg.FindAllStringSubmatch(content, -1)
	if len(importMatches) == 0 {
		return nil
	}
	packageMatches := packageKeywordReg.FindStringSubmatch(content)
	if len(packageMatches) == 0 {
		return nil
	}
	fromNode := models.Node{
		Name: packageMatches[1],
		Path: packageMatches[1],
	}

	nodes := models.Nodes{}
	links := models.Links{}
	for _, match := range importMatches[1:] {
		toNode := models.Node{
			Name: match[1],
			Path: match[1],
			Size: 1,
		}
		nodes = nodes.Add(toNode)

		links = links.Add(models.Link{
			From:     &fromNode,
			To:       &toNode,
			Strength: 1,
		})
	}

	return &Dependencies{
		Nodes: nodes,
		Links: links,
	}
}

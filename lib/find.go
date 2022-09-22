package lib

import (
	"depsfinder/lib/models"
	"log"
)

type DepsFinder interface {
	Find() (*Dependencies, error)
}

type Dependencies struct {
	Nodes []models.Node
	Links []models.Link
}

func FindDeps(depsFinder DepsFinder) error {
	dependencies, err := depsFinder.Find()
	if err != nil {
		return err
	}

	log.Printf("found dependencies: %#v", dependencies)

	return nil
}

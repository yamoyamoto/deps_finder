package lib

import (
	"depsfinder/lib/finders"
	"log"
)

type DepsFinder interface {
	Find(dirPath string) (*finders.Dependencies, error)
}

func FindDeps(depsFinder DepsFinder, dirPath string) error {
	dependencies, err := depsFinder.Find(dirPath)
	if err != nil {
		return err
	}

	for _, link := range dependencies.Links {
		log.Printf("found dependencies: %s->%s, strength: %d", link.From.Name, link.To.Name, link.Strength)
	}

	return nil
}

package services

import (
	"depsfinder/lib/models"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type DirWalker interface {
	ListFileRecursive(dirPath string) ([]models.File, error)
}

type DirWalkerImpl struct {
}

func NewDirWalker() *DirWalkerImpl {
	return &DirWalkerImpl{}
}

func (walk *DirWalkerImpl) ListFileRecursive(dirPath string) ([]models.File, error) {
	paths, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var files []models.File
	for _, path := range paths {
		if path.IsDir() {
			f, err := walk.ListFileRecursive(filepath.Join(dirPath, path.Name()))
			if err != nil {
				return nil, fmt.Errorf("dirwalk %s: %w", filepath.Join(dirPath, path.Name()), err)
			}
			files = append(files, f...)
			continue
		}
		file, err := models.NewFile(filepath.Join(dirPath, path.Name()))
		if err != nil {
			return nil, fmt.Errorf("dirwalk %s: %w", filepath.Join(dirPath, path.Name()), err)
		}
		files = append(files, *file)
	}

	return files, nil
}

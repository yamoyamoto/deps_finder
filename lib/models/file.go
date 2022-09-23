package models

import (
	"fmt"
	"regexp"
)

type File struct {
	Path     string
	FileType FileType
}

func NewFile(filePath string) (*File, error) {
	fileType, err := getFileTypeFromFilePath(filePath)
	if err != nil {
		return nil, err
	}
	return &File{
		Path:     filePath,
		FileType: fileType,
	}, nil
}

type FileType int64

const (
	Unknown FileType = iota
	PHP
	Go
	TypeScript
	Java
)

const (
	PHPExt        = ".php"
	GoExt         = ".go"
	TypeScriptExt = ".ts"
	JavaExt       = ".java"
)

var extMap = map[FileType]string{
	PHP:        PHPExt,
	Go:         GoExt,
	TypeScript: TypeScriptExt,
	Java:       JavaExt,
}

func getFileTypeFromFilePath(filePath string) (FileType, error) {
	for fileType, ext := range extMap {
		match, err := regexp.MatchString(fmt.Sprintf("^.+%s$", ext), filePath)
		if err != nil {
			return Unknown, err
		}
		if match {
			return fileType, nil
		}
	}
	return Unknown, nil
}

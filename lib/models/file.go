package models

import (
	"fmt"
	"regexp"
)

type File struct {
	Path string
}

type FileType int64

const (
	Unknown FileType = iota
	PHP
	Go
	TypeScript
)

const (
	PHPExt        = ".php"
	GoExt         = ".go"
	TypeScriptExt = ".ts"
)

var extMap = map[FileType]string{
	PHP:        PHPExt,
	Go:         GoExt,
	TypeScript: TypeScriptExt,
}

func (file *File) GetFileType() (FileType, error) {
	for fileType, ext := range extMap {
		match, err := regexp.MatchString(fmt.Sprintf("^.+%s$", ext), file.Path)
		if err != nil {
			return Unknown, err
		}
		if match {
			return fileType, nil
		}
	}
	return Unknown, nil
}

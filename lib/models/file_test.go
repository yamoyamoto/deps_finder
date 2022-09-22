package models_test

import (
	"depsfinder/lib/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileType(t *testing.T) {
	t.Run("拡張子がマッチしていればそのFileTypeを返す", func(t *testing.T) {
		testCases := []struct {
			name     string
			filePath string
			expected models.FileType
		}{
			{
				"PHP",
				"/path/to/hoge.php",
				models.PHP,
			},
			{
				"TypeScript",
				"/path/to/hello.ts",
				models.TypeScript,
			},
			{
				"GO",
				"/path/to/hugahuga.go",
				models.Go,
			},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				file := models.File{
					Path: tt.filePath,
				}
				fileType, err := file.GetFileType()

				assert.Nil(t, err)
				assert.Equal(t, tt.expected, fileType)
			})
		}
	})

	t.Run("不明な拡張子の場合Unknownを返す", func(t *testing.T) {
		unknownTypeFile := models.File{
			Path: "path/to/unknown-file.txt",
		}
		fileType, err := unknownTypeFile.GetFileType()

		assert.Nil(t, err)
		assert.Equal(t, models.Unknown, fileType)
	})
}

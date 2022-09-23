package models_test

import (
	"depsfinder/lib/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFile(t *testing.T) {
	t.Run("拡張子がマッチしていればそのFileTypeが入る", func(t *testing.T) {
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
				"Go",
				"/path/to/hugahuga.go",
				models.Go,
			},
			{
				"Java",
				"/path/to/aiueo.java",
				models.Java,
			},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				file, err := models.NewFile(tt.filePath)

				assert.Nil(t, err)
				assert.Equal(t, tt.filePath, file.Path)
				assert.Equal(t, tt.expected, file.FileType)
			})
		}
	})

	t.Run("不明な拡張子の場合Unknownが入る", func(t *testing.T) {
		file, err := models.NewFile("path/to/unknown-file.txt")

		assert.Nil(t, err)
		assert.Equal(t, "path/to/unknown-file.txt", file.Path)
		assert.Equal(t, models.Unknown, file.FileType)
	})
}

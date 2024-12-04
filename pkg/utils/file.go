package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CreateUploadDir() error {
	paths := []string{
		"uploads",
		"uploads/products",
		"uploads/avatars",
	}

	for _, path := range paths {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func GenerateFileName(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	filename := strings.TrimSuffix(file.Filename, ext)
	return fmt.Sprintf("%s_%d%s",
		strings.ToLower(strings.ReplaceAll(filename, " ", "_")),
		time.Now().UnixNano(),
		ext,
	)
}

func ValidateFileType(file *multipart.FileHeader) bool {
	allowedType := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	return allowedTypes[file.Header.Get("Content-Type")]
}

func ValidateFileSize(file *multipart.FileHeader) bool {
	maxSize := int64(5 * 1024 * 1024) // 5MB
	return file.Size <= maxSize
}

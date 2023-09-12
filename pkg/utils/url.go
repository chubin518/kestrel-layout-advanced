package utils

import (
	"net/url"
	"path/filepath"
)

// GetFileName
func GetFileName(path string) (string, error) {
	u, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	return filepath.Base(u.Path), nil
}

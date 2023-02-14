//go:build feature_test

package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type mockClient struct {
	baseDir string
}

func (c *mockClient) Create(ctx context.Context, distPath string, content []byte) error {
	path := fmt.Sprintf("%s/%s", c.baseDir, distPath)
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(path, content, os.ModePerm)
}

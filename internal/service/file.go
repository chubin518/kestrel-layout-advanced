package service

import (
	"context"
	"os"
	"path/filepath"

	"github.com/chubin518/kestrel-layout-advanced/internal/model"
	"github.com/chubin518/kestrel-layout-advanced/pkg/logging"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

// List
func (s *FileService) List(ctx context.Context, dir string) ([]*model.FileRecord, error) {
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	list := make([]*model.FileRecord, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		fi, err := fileInfo.Info()
		if err != nil {
			return list, err
		}
		list = append(list, &model.FileRecord{
			Name:     fi.Name(),
			Path:     filepath.Join(dir, fi.Name()),
			Modified: fi.ModTime(),
		})
	}
	return list, nil
}

// TreeList
func (s *FileService) TreeList(ctx context.Context, dir string) ([]*model.TreeNode, error) {
	logging.InfoContext(ctx, "TreeList: %v", dir)
	fileInfos, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	list := make([]*model.TreeNode, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		path := filepath.Join(dir, fileInfo.Name())
		node := &model.TreeNode{
			Label: fileInfo.Name(),
			Value: path,
		}
		if fileInfo.IsDir() {
			children, err := s.TreeList(ctx, path)
			if err != nil {
				return nil, err
			}
			node.Children = children
		}
		list = append(list, node)
	}
	return list, nil
}

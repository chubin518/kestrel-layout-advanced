package utils

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"

	"github.com/ulikunitz/xz"
	"github.com/yeka/zip"
)

// DecompressTarXz decompresses a tar.xz archive
func DecompressTarXz(srcPath, dstDir string) error {
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()
	xr, err := xz.NewReader(file)
	if err != nil {
		return err
	}
	tr := tar.NewReader(xr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
		dstPath := filepath.Join(dstDir, header.Name)
		os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		if header.Typeflag == tar.TypeReg {
			dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err := io.Copy(dst, tr); err != nil {
				return err
			}
		}
	}
	return nil
}

// DecompressZipWithPassword
func DecompressZipWithPassword(srcPath, dstDir, password string) error {
	reader, err := zip.OpenReader(srcPath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.IsEncrypted() {
			file.SetPassword(password)
		}

		dstPath := filepath.Join(dstDir, file.Name)
		if file.FileInfo().IsDir() {
			continue
		}
		os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		dst, err := os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer dst.Close()

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}

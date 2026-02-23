package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func ToTar(dir string) error {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}

		relativePath, _ := filepath.Rel(dir, path)
		header.Name = relativePath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.Mode().IsRegular() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tw, file)
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	tarFile, err := os.Create(fmt.Sprintf("tar-%d.tar", time.Now().UnixMilli()))
	if err != nil {
		return err
	}

	if _, err := io.Copy(tarFile, buf); err != nil {
		return err
	}

	return nil
}

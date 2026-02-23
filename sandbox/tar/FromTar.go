package main

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

func FromTar(path string) error {
	tarFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tr := tar.NewReader(tarFile)

	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if hdr.Typeflag == tar.TypeDir {
			os.MkdirAll(filepath.Join("extract", hdr.Name), hdr.FileInfo().Mode())
		}
		if hdr.Typeflag == tar.TypeReg {
			f, _ := os.Create(filepath.Join("extract", hdr.Name))
			io.Copy(f, tr)
			f.Close()
		}
	}

	return nil
}

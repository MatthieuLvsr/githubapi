package zipmaker

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipDirectory(source, target string) error{
	f, err := os.Create(target)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// header.Name = filepath.Join(source, path)
		header.Name = relPath
		if info.IsDir(){
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir(){
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
	
			_, err = io.Copy(headerWriter, f)
			return err
		}
		return err
	})
}
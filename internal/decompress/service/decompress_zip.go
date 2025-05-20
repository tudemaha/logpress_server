package service

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func DecompressZIP(filename string) (err error) {
	zipReader, err := zip.OpenReader("./dump/compressed/" + filename + ".sql.zip")
	if err != nil {
		return err
	}
	defer zipReader.Close()

	uncompressedPath := "./dump/uncompressed/" + filename + ".sql"

	for _, f := range zipReader.File {
		if strings.Contains(f.Name, "/") {
			continue
		}

		dst, err := os.Create(uncompressedPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		fileInArc, err := f.Open()
		if err != nil {
			return err
		}
		defer fileInArc.Close()

		if _, err := io.Copy(dst, fileInArc); err != nil {
			return nil
		}
	}

	return nil
}

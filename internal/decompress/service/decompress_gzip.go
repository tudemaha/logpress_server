package service

import (
	"compress/gzip"
	"io"
	"os"
)

func DecompressGZIP(filename string) (err error) {
	gzipped, err := os.Open("./dump/compressed/" + filename + ".sql.gz")
	if err != nil {
		return err
	}
	defer gzipped.Close()

	gzipReader, err := gzip.NewReader(gzipped)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	uncompressedPath := "./dump/uncompressed/" + filename + ".sql"
	uncompressed, err := os.Create(uncompressedPath)
	if err != nil {
		return err
	}
	defer uncompressed.Close()

	_, err = io.Copy(uncompressed, gzipReader)
	if err != nil {
		return err
	}

	return nil
}

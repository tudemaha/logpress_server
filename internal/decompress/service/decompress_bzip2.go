package service

import (
	"compress/bzip2"
	"io"
	"os"
)

func DecompressBZIP2(filename string) (err error) {
	bzipped, err := os.Open("./dump/compressed/" + filename + ".sql.bz2")
	if err != nil {
		return err
	}
	defer bzipped.Close()

	bzipReader := bzip2.NewReader(bzipped)

	uncompressedPath := "./dump/uncompressed/" + filename + ".sql"
	uncompressed, err := os.Create(uncompressedPath)
	if err != nil {
		return err
	}
	defer uncompressed.Close()

	_, err = io.Copy(uncompressed, bzipReader)
	if err != nil {
		return err
	}

	return nil
}

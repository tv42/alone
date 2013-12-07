package main

import (
	"compress/gzip"
	"github.com/surma/gocpio"
	"io"
	"log"
	"os"
)

func getSize(r io.Seeker) (int64, error) {
	var size int64
	var err error

	size, err = r.Seek(0, os.SEEK_END)
	if err != nil {
		return 0, err
	}
	_, err = r.Seek(0, os.SEEK_SET)
	if err != nil {
		return size, err
	}
	return size, nil
}

func writeCpio(w io.Writer, init io.ReadSeeker) error {
	var size int64
	var err error
	size, err = getSize(init)
	if err != nil {
		return err
	}

	archive := cpio.NewWriter(w)
	err = archive.WriteHeader(&cpio.Header{
		Mode: 0755,
		Size: size,
		Type: cpio.TYPE_REG,
		Name: "init",
	})
	if err != nil {
		return err
	}
	_, err = io.Copy(archive, init)
	if err != nil {
		return err
	}
	return archive.Close()
}

func main() {
	w := gzip.NewWriter(os.Stdout)
	w.Header.Name = "initramfs.cpio"
	var err error
	err = writeCpio(w, os.Stdin)
	if err != nil {
		log.Fatalf("cpio failed: %v", err)
	}
	err = w.Close()
	if err != nil {
		log.Fatalf("gzip failed: %v", err)
	}
}

package main

import (
	"compress/gzip"
	"debug/elf"
	"io"
	"log"
	"os"

	"github.com/surma/gocpio"
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

func isStatic(r io.ReaderAt) (bool, error) {
	f, err := elf.NewFile(r)
	if err != nil {
		return false, err
	}
	libs, err := f.ImportedLibraries()
	if err != nil {
		return false, err
	}
	if len(libs) > 0 {
		return false, nil
	}
	return true, nil
}

func main() {
	static, err := isStatic(os.Stdin)
	if err != nil {
		log.Fatalf("cannot decode ELF header: %v", err)
	}
	if !static {
		log.Fatal("input binary must be statically linked")
	}

	w := gzip.NewWriter(os.Stdout)
	w.Header.Name = "initramfs.cpio"
	err = writeCpio(w, os.Stdin)
	if err != nil {
		log.Fatalf("cpio failed: %v", err)
	}
	err = w.Close()
	if err != nil {
		log.Fatalf("gzip failed: %v", err)
	}
}

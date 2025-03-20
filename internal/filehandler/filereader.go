package filehandler

import "bufio"
import "os"

type Reader struct {
    Filename string
    Filereader *bufio.Reader
    File *os.File
}

func (r *Reader) Read(b []byte) (int, error) {
    return r.Filereader.Read(b)
}

func NewFilereader(filename string) (*Reader, error) {
    file, err := os.Open(filename)
    if err != nil {
	return nil, err
    }

    return &Reader{File:file, Filename:filename, Filereader:bufio.NewReader(file)}, nil
}

func (r *Reader) Getsize() (uint32, error) {
    n, err := r.File.Stat()
    if err != nil {
	panic(err)
    }
    return uint32(n.Size()), nil
}


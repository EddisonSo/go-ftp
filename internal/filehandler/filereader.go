package filehandler

import "bufio"
import "os"
import "log/slog"

type Reader struct {
    Filename string
    Filereader *bufio.Reader
    File *os.File
    Logger *slog.Logger
}

func (r *Reader) Read(b []byte) (int, error) {
    n, err := r.Filereader.Read(b)
    r.Logger.Info("Read from file: " + r.Filename)
    return n, err
}

func NewFilereader(filename string, logger *slog.Logger) (*Reader, error) {
    file, err := os.Open(filename)
    if err != nil {
	return nil, err
    }

    return &Reader{File:file, Filename:filename, Filereader:bufio.NewReader(file), Logger:logger}, nil
}

func (r *Reader) Getsize() (uint32, error) {
    n, err := r.File.Stat()
    if err != nil {
	panic(err)
    }
    return uint32(n.Size()), nil
}


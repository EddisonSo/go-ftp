package filehandler

import "bufio"
import "os"
import "fmt"

type Writer struct {
    Filename string
    Filereader *bufio.Writer
}

func (w *Writer) Write(b []byte) (int, error) {
    fmt.Println("Writing to file: ", w.Filename)
    fmt.Println("Data: ", string(b))
    n, err := w.Filereader.Write(b)
    w.Filereader.Flush()
    return n, err
}

func NewFilewriter(filename string) (*Writer, error) {
    file, err := os.Create(filename)
    if err != nil {
	return nil, err
    }

    return &Writer{Filename:filename, Filereader:bufio.NewWriter(file)}, nil
}

package filehandler

import "bufio"
import "os"
import "fmt"
import "log/slog"

type Writer struct {
    Filename string
    Filereader *bufio.Writer
    Logger *slog.Logger
}

func (w *Writer) Write(b []byte) (int, error) {
    w.Logger.Info("Writing to file: " + w.Filename)
    w.Logger.Info("Data: " + fmt.Sprintf("%s", b))
    n, err := w.Filereader.Write(b)
    w.Filereader.Flush()
    return n, err
}

func NewFilewriter(filename string, logger *slog.Logger) (*Writer, error) {
    file, err := os.Create(filename)
    if err != nil {
	return nil, err
    }

    return &Writer{Filename:filename, Filereader:bufio.NewWriter(file), Logger:logger}, nil
}

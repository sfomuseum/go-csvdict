package csvdict

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// type Writer implements a `encoding/csv` style writer for CSV documents with named columns.
type Writer struct {
	Writer     *csv.Writer
	Fieldnames []string
}


func NewWriter(wr io.Writer, fieldnames []string) (*Writer, error) {

	writer := csv.NewWriter(wr)

	dw := Writer{Writer: writer, Fieldnames: fieldnames}
	return &dw, nil
}

func NewWriterFromPath(path string, fieldnames []string) (*Writer, error) {

	fh, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for writing, %w", path, err)
	}

	return NewWriter(fh, fieldnames)
}

func (dw Writer) WriteHeader() error {

	return dw.Writer.Write(dw.Fieldnames)
}

// to do - check flags for whether or not to be liberal when missing keys
// (20160516/thisisaaronland)

func (dw Writer) WriteRow(row map[string]string) error {

	out := make([]string, 0)

	for _, k := range dw.Fieldnames {

		v, ok := row[k]

		if !ok {
			v = ""
		}

		out = append(out, v)
	}

	return dw.Writer.Write(out)
}

// Flush writes any buffered data to the underlying writer. To check if an error occurred during the Flush, call Error.
func (dw Writer) Flush() error {
	dw.Writer.Flush()
	return nil
}

// Error reports any error that has occurred during a previous Write or Flush.
func (dw Writer) Error() error {
	return dw.Writer.Error()
}

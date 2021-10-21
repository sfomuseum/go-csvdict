package csvdict

import (
	"bufio"
	"bytes"
	_ "fmt"
	"io"
	"os"
	"sort"
	"testing"
)

func TestWriter(t *testing.T) {

	path := "fixtures/test.csv"

	fh, err := os.Open(path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path, err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		t.Fatalf("Failed to read body for %s, %v", path, err)
	}

	_, err = fh.Seek(0, 0)

	if err != nil {
		t.Fatalf("Failed to seek file to 0, %v", err)
	}

	csv_r, err := NewReader(fh)

	if err != nil {
		t.Fatalf("Failed to create reader, %v", err)
	}

	rows := make([]map[string]string, 0)

	for {
		row, err := csv_r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("Failed to read row, %v", err)
		}

		rows = append(rows, row)
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	fieldnames := make([]string, 0)

	for k, _ := range rows[0] {
		fieldnames = append(fieldnames, k)
	}

	sort.Strings(fieldnames)

	csv_wr, err := NewWriter(wr, fieldnames)

	if err != nil {
		t.Fatalf("Failed to create new writer, %v", err)
	}

	err = csv_wr.WriteHeader()

	if err != nil {
		t.Fatalf("Failed to write CSV header, %v", err)
	}

	for i, row := range rows {

		err := csv_wr.WriteRow(row)

		if err != nil {
			t.Fatalf("Failed to write row (%d), %v", i, err)
		}

		csv_wr.Flush()
	}

	if !bytes.Equal(body, buf.Bytes()) {
		t.Fatalf("Unexpected output")
	}
}

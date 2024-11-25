package csvdict

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {

	path := "fixtures/test.csv"

	r, err := os.Open(path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path, err)
	}

	defer r.Close()

	scanner := bufio.NewScanner(r)
	count_lines := 0

	for scanner.Scan() {
		count_lines += 1
	}

	err = scanner.Err()

	if err != nil {
		t.Fatalf("Scanner reported an error, %v", err)
	}

	_, err = r.Seek(0, 0)

	if err != nil {
		t.Fatalf("Failed to seek file to 0, %v", err)
	}

	csv_r, err := NewReader(r)

	if err != nil {
		t.Fatalf("Failed to create reader, %v", err)
	}

	// Test the Read method

	count_rows := 0

	for {

		row, err := csv_r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("Failed to read row, %v", err)
		}

		_, ok := row["label"]

		if !ok {
			t.Fatalf("Row is missing 'label' column")
		}

		count_rows += 1
	}

	if count_rows != count_lines-1 {
		t.Fatalf("Expected %d rows, but got %d", count_lines-1, count_rows)
	}

	// Test the Iterator method

	_, err = r.Seek(0, 0)

	if err != nil {
		t.Fatalf("Failed to seek file to 0, %v", err)
	}

	csv_r, err = NewReader(r)

	if err != nil {
		t.Fatalf("Failed to create reader, %v", err)
	}

	count_rows = 0

	for row, err := range csv_r.Iterate() {

		if err != nil {
			t.Fatalf("Failed to iterate row, %v", err)
		}

		_, ok := row["label"]

		if !ok {
			t.Fatalf("Row is missing 'label' column")
		}

		count_rows += 1
	}

	if count_rows != count_lines-1 {
		t.Fatalf("Expected %d rows, but got %d", count_lines-1, count_rows)
	}

}

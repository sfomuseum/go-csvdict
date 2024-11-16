package csvdict

import (
	"bufio"
	"os"
	"testing"
)

func TestReader(t *testing.T) {

	path := "fixtures/test.csv"

	fh, err := os.Open(path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", path, err)
	}

	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	count_lines := 0

	for scanner.Scan() {
		count_lines += 1
	}

	err = scanner.Err()

	if err != nil {
		t.Fatalf("Scanner reported an error, %v", err)
	}

	_, err = fh.Seek(0, 0)

	if err != nil {
		t.Fatalf("Failed to seek file to 0, %v", err)
	}

	csv_r, err := NewReader(fh)

	if err != nil {
		t.Fatalf("Failed to create reader, %v", err)
	}

	count_rows := 0

	for row, err := range csv_r.Read() {

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
}

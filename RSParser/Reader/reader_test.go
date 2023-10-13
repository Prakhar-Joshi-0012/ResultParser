package reader

import (
	"testing"
)

func TestLineByLine(t *testing.T) {
	rdr := NewReader()
	filename := "test.txt"
	rdr.ReadFile(filename)
	students := rdr.parseStudents()
	for _, line := range students {
		t.Errorf("%v", line)
	}

}

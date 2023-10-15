package reader

import (
	"testing"
)

func TestLineByLine(t *testing.T) {
	rdr := NewReader(12)
	filename := "test.txt"
	rdr.ReadFile(filename)
	rdr.ParseStudents()
	for _, stream := range rdr.Streams {
		t.Errorf("%v", stream.Students)
	}
}

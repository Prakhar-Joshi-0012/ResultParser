package writer

import (
	. "ResultParser/RSParser/Reader"
	"testing"
)

func Test(t *testing.T) {
	rdr := NewReader(12)
	filename := "test"
	rdr.ReadFile("test.txt")
	rdr.ParseStudents()
	wr := NewWriter(filename, 12)
	err := wr.WriteStudents(rdr.Streams)
	if err != nil {
		t.Errorf("%v", err)
	}

}

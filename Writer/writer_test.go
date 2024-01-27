package writer

import (
	. "ResultParser/Reader"
	"testing"
)

func Test(t *testing.T) {
	rdr := NewReader(10)
	filename := "test2"
	rdr.ReadFile("test2.TXT")
	rdr.ParseStudents()
	wr := NewWriter(filename, 10)
	err := wr.WriteStudents(rdr.Streams)
	if err != nil {
		t.Errorf("%v", err)
	}

}

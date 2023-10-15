package parser

import (
	reader "ResultParser/RSParser/Reader"
	writer "ResultParser/RSParser/Writer"
	"io"
)

func Parse(std int, file io.ByteReader) error {
	rdr := reader.NewReader(std)
	rdr.ReadFile("test.txt")
	rdr.ParseStudents()
	wr := writer.NewWriter("test", 12)
	wr.WriteStudents(rdr.Streams)
	return nil
}

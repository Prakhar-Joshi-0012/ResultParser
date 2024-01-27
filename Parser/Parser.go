package parser

import (
	reader "ResultParser/Reader"
	writer "ResultParser/Writer"
	"fmt"
)

func Parse(std int, filepath string) error {
	rdr := reader.NewReader(std)
	rdr.ReadFile(filepath)
	fmt.Printf("%v", (rdr.Streams))
	rdr.ParseStudents()
	wr := writer.NewWriter("test", std)
	wr.WriteStudents(rdr.Streams)
	return nil
}

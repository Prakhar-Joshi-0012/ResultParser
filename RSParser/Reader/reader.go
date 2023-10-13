package reader

import (
	student "ResultParser/RSParser/Student"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FileReader struct {
	data [][]string
}

func (fr *FileReader) ReadFile(f string) {
	linebyline(fr, f)
}

func NewReader() FileReader {
	return FileReader{data: [][]string{}}
}

func ErrorEncountered(err error, msg ...string) error {
	e := fmt.Sprintf("Error: %v\n", err)
	for _, msg := range msg {
		e += fmt.Sprintf("%v\n", msg)
	}
	return errors.New(e)
}

func linebyline(rdr *FileReader, f string) error {
	fr, e := os.Open(f)
	if e != nil {
		return ErrorEncountered(e)
	}
	defer fr.Close()
	reader := bufio.NewReader(fr)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return ErrorEncountered(e)
		}
		line = strings.Replace(strings.Replace(strings.Replace(
			strings.Replace(line, `-`, " ", -1), "_", " ", -1), `:`, " ", -1), `*`, " ", -1)
		re := regexp.MustCompile(`\S+`)
		linesplit := re.FindAllString(line, -1)
		if len(linesplit) != 0 {
			rdr.data = append(rdr.data, linesplit)
		}
	}
	return nil
}

func (fr *FileReader) parseStudents() []student.Student {
	var index int = 0
	var size int = len(fr.data)
	StudentData := []student.Student{}

	for index < size {
		line := fr.data[index]
		_, e := strconv.Atoi(line[0])
		if e == nil {
			name, rno, gender, subCode := parseStudent(fr.data[index])
			index++
			marks := parseMarks(fr.data[index])
			StudentData = append(StudentData, student.NewStudent(name, rno, gender, subCode, marks))
		}
		index++
	}
	return StudentData
}

func parseStudent(s []string) (string, string, string, []string) {
	end := len(s)
	rno := s[0]
	gender := s[1]
	var pos int
	var name string
	_, e := strconv.Atoi(s[3])
	if e != nil {
		pos = 4
		name = strings.Join(s[2:4], " ")
	} else {
		pos = 3
		name = s[2]
	}
	subCode := []string{}
	for pos < end {
		_, e := strconv.Atoi(s[pos])
		if e == nil {
			subCode = append(subCode, s[pos])
		}
		pos++
	}
	return name, rno, gender, subCode
}

func parseMarks(s []string) []string {
	marks := []string{}
	for _, m := range s {
		_, e := strconv.Atoi(m)
		if e == nil {
			marks = append(marks, m)
		}
	}
	return marks
}

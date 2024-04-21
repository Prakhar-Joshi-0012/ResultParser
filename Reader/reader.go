package reader

import (
	. "ResultParser/Errors"
	student "ResultParser/Student"
	. "ResultParser/Subjects"
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type FileReader struct {
	data    [][]string
	std     int
	Streams []Streams
}

func (fr *FileReader) ReadFile(f string) {
	linebyline(fr, f)
}

func NewReader(std int) *FileReader {
	if std == 12 {
		return &FileReader{
			data: [][]string{},
			std:  std,
			Streams: []Streams{
				{Name: "Arts", SubCodes: map[string]bool{}, Students: []student.Student{}},
				{Name: "Commerce", SubCodes: map[string]bool{}, Students: []student.Student{}},
				{Name: "Science", SubCodes: map[string]bool{}, Students: []student.Student{}},
			},
		}
	} else {
		return &FileReader{
			data: [][]string{},
			std:  std,
			Streams: []Streams{
				{Name: "Marks", SubCodes: map[string]bool{}, Students: []student.Student{}},
			},
		}
	}
}

func linebyline(rdr *FileReader, f string) error {
	fr, e := os.Open(f)
	if e != nil {
		return ErrorEncountered(e, "File not Found")
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

func (fr *FileReader) ParseStudents() {
	var index int = 0
	var size int = len(fr.data)

	for index < size {
		line := fr.data[index]
		_, e := strconv.Atoi(line[0])
		if e == nil {
			name, rno, gender, subCode := parseStudent(fr.data[index])
			_, streamIn := ReturnStream(fr.std, subCode)
			status := fr.data[index][len(fr.data[index])-1]
			var marks []string
			index++
			if status == "ABST" {
				index--
				for i := 0; i < len(subCode); i++ {
					marks = append(marks, "0")
				}
			} else {
				marks = parseMarks(fr.data[index])
			}
			for _, k := range subCode {
				if _, e := fr.Streams[streamIn].SubCodes[k]; !e {
					fr.Streams[streamIn].SubCodes[k] = true
				}
			}
			fr.Streams[streamIn].Students = append(fr.Streams[streamIn].Students, student.NewStudent(name, rno, gender, subCode, marks))
		}
		index++
	}
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
		if s[pos] == "PASS" || s[pos] == "ABST" || s[pos] == "COMP" {
			break
		}
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

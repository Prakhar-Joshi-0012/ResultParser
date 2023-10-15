package writer

import (
	. "ResultParser/RSParser/Errors"
	. "ResultParser/RSParser/Subjects"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Writer struct {
	file     excelize.File
	std      int
	filename string
	SubCodes *Subcodes
}

type Subcodes struct {
	data map[string]int
}

func NewWriter(s string, std int) *Writer {
	return &Writer{
		file:     *excelize.NewFile(),
		std:      std,
		filename: s,
	}
}

func Generate(sbj map[string]bool) *Subcodes {
	var offset int = 3 // Rollnumber, Name, Subject list
	Subcodes := Subcodes{data: map[string]int{}}
	ind := 0
	for k, _ := range sbj {
		Subcodes.data[k] = ind + offset
		ind++
	}
	return &Subcodes
}

func (wr *Writer) WriteStudents(streams []Streams) error {
	for _, stream := range streams {
		index, err := wr.file.NewSheet(stream.Name)
		if err != nil {
			return ErrorEncountered(err, "Error encountered in creating sheet")
		}
		wr.file.SetActiveSheet(index)
		wr.SubCodes = Generate(stream.SubCodes)
		rno := 1
		wr.file.SetCellValue(stream.Name, "A1", "RollNumber")
		wr.file.SetCellValue(stream.Name, "B1", "Name")
		for k, collno := range wr.SubCodes.data {
			cellColName, err := excelize.ColumnNumberToName(collno)
			if err != nil {
				return err
			}
			wr.file.SetCellValue(stream.Name, strings.Join([]string{cellColName, strconv.Itoa(rno)}, ""), Subject_Code_12[k])
		}
		rno = 2
		for _, student := range stream.Students {
			wr.file.SetCellValue(stream.Name, strings.Join([]string{"A", strconv.Itoa(rno)}, ""), student.RollNo)
			wr.file.SetCellValue(stream.Name, strings.Join([]string{"B", strconv.Itoa(rno)}, ""), student.Name)
			for k, collno := range wr.SubCodes.data {
				cellColName, err := excelize.ColumnNumberToName(collno)
				if err != nil {
					return err
				}
				marks, f := student.Marks[k]
				if f {
					wr.file.SetCellValue(stream.Name, strings.Join([]string{cellColName, strconv.Itoa(rno)}, ""), marks)
				}
			}
			rno++
		}
	}
	wr.file.DeleteSheet("Sheet1")
	if err := wr.file.SaveAs(strings.Join([]string{wr.filename, ".xlsx"}, "")); err != nil {
		return ErrorEncountered(err)
	}
	return nil
}

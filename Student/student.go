package student

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name   string
	RollNo string
	Gender string
	Marks  map[string]int
}

func NewStudent(name string, rollno string, gender string, SubCode []string, marks []string) Student {
	return Student{
		Name:   name,
		RollNo: rollno,
		Gender: gender,
		Marks:  bind(SubCode, marks),
	}
}

func bind(SubCode []string, marks []string) map[string]int {
	fmt.Printf("Marks %v\n", marks)
	if len(SubCode) != len(marks) {
		fmt.Printf("Parsing : SubCodes doesn't match with marks")
		fmt.Printf("SubCode : %v\n", SubCode)
		fmt.Printf("Marks : %v\n", marks)
		panic("Parsing Error: SubCodes doesn't match with marks")
	}
	store := map[string]int{}
	for i := 0; i < len(marks); i++ {
		val, e := strconv.Atoi(marks[i])
		if e != nil {
			panic("type(marks) != integer")
		}
		store[SubCode[i]] = val
	}
	return store
}

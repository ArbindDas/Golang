package main

import (
	"fmt"
)

type Student struct {
	Name       string
	Subjects   map[string]float64
	TotalMarks float64
}

func createStudent(name string) *Student {

	subject := map[string]float64{
		"Math": 70,
		"eng":  70,
		"com":  70,
		"pop":  70,
		"nep":  70,
	}
	return &Student{Name: name, Subjects: subject, TotalMarks: 500}
}

func (s *Student) markSheet() {
	var obtainedMarks float64

	for _, mark := range s.Subjects {
		obtainedMarks += mark
	}

	percentage := obtainedMarks / float64(len(s.Subjects))

	fmt.Printf("Marksheet for %s:\n", s.Name)
	fmt.Println("Total Marks:", s.TotalMarks)
	fmt.Println("Obtained Marks:", obtainedMarks)
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Println("Grade:", s.calculateGrade(percentage))
}

func (s *Student) calculateGrade(percentage float64) string {

	switch {
	case percentage >= 90:
		return "A+"
	case percentage >= 80:
		return "A"
	case percentage >= 70:
		return "B+"
	case percentage >= 60:
		return "B"
	case percentage >= 50:
		return "C+"
	case percentage >= 40:
		return "C"
	default:
		return "Fail"
	}

}

func main() {

	student := createStudent("Arbind")

	student.markSheet()

	fmt.Println("jai shree ram")
}

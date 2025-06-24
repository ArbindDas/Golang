package main

import "fmt"

type StudentMarks struct {
	s1, s2, s3, s4, s5 float64
}

func createAStudent() *StudentMarks {
	return &StudentMarks{s1: 70, s2: 80, s3: 81, s4: 89, s5: 100}
}

func (s *StudentMarks) marksSheet() {
	Total := s.s1 + s.s2 + s.s3 + s.s4 + s.s5
	fmt.Println("total marks is : 500")
	fmt.Println("the total marks obtain  by student : ", Total)

	percentage := (Total) / 5

	fmt.Println("Percenatge ", percentage)

	if percentage <= 90 || percentage >= 100 {
		fmt.Println("A+")
		return
	}
	if percentage <= 80 || percentage >= 90 {
		fmt.Println("A")
		return
	}
	if percentage <= 70 || percentage >= 80 {
		fmt.Println("B+")
		return
	}
	if percentage <= 60 || percentage >= 70 {
		fmt.Println("B")
		return
	}
	if percentage <= 50 || percentage >= 60 {
		fmt.Println("C+")
		return
	}

	if percentage <= 40 || percentage >= 50 {
		fmt.Println("C")
		return
	}

	if percentage <= 30 || percentage >= 40 {
		fmt.Println("just fail")
		return
	}

}

func main() {

	// Arbind := StudentMarks{
	// 	s1: 90,
	// 	s2: 80,
	// 	s3: 80,
	// 	s4: 78,
	// 	s5: 98,
	// }

	Suraj := createAStudent()
	Suraj.marksSheet()

	// Arbind.marksSheet()

	fmt.Println("jai shree ram prabhu")
}

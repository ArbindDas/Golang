package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Roll  int
	Name  string
	Class string
}

// creating a constructor
func initializeStudent(roll int, name string, class string) *Student {
	return &Student{

		Roll:  roll,
		Name:  name,
		Class: class,

		//  Private fields	Unexported (lowercase) fields
		// Public fields	Exported (uppercase) fields
	}
}

func (s *Student) getRoll() int {
	fmt.Println("getting roll ")
	return s.Roll
}

// this func show my current student details
func (s *Student) showingStudent() {

	fmt.Println(s.Roll)
	fmt.Println(s.Name)
	fmt.Println(s.Class)
}

func main() {

	jai := initializeStudent(1, "jai", "33")

	fmt.Println(jai)

	fmt.Println("jai shree ram")

	Arbind := Student{
		Roll:  1,
		Name:  "Arbind",
		Class: "Ten",
	}

	Arbind.getRoll()

	data, _ := json.MarshalIndent(Arbind, "", " ")

	fmt.Println(string(data))

	Arbind.showingStudent()
	// fmt.Println(Arbind)
	// The _ is used to ignore the error returned by MarshalIndent.

	// Normally, you should check the error, like:

	// data, err := json.MarshalIndent(Arbind, "", " ")
	// if err != nil {
	// 	fmt.Println("Error:", err)
}
